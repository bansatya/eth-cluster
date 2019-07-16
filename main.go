package main

import (
   "geth-dapp-demo/quiz"
    "context"
    "log"
    "fmt"
    "os"
    "strings"
    "bufio"
    "math/big"
    //"errors"
    //"crypto/Keccak256"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/joho/godotenv"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
)   
var myenv map[string]string

const envLoc = ".env"
// ErrTransactionWait should be returned/printed when we encounter an error that may be a result of the transaction not being confirmed yet.
const ErrTransactionWait = "if you've just started the application, wait a while for the network to confirm your transaction."


func loadEnv() {
    var err error
    if myenv, err = godotenv.Read(envLoc); err != nil {
        log.Printf("could not load env from %s: %v", envLoc, err)
    }
}

func main(){
    loadEnv()

    ctx := context.Background()

  //   client, err := ethclient.Dial(os.Getenv("GATEWAY"))
    client, err := ethclient.Dial("/home/ubuntu/eth_node/node1/datadir/geth.ipc")
    if err != nil {
        log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
    }
    defer client.Close()

    //accountAddress := common.HexToAddress("0xF6cc1cba1BE28d9fEB2be11534ECf7FEE8bFE1CE")
    //balance, _ := client.BalanceAt(ctx, accountAddress, nil)
    //fmt.Printf("Balance: %d\n",balance)

   session := NewSession(ctx)  



// Load or Deploy contract, and update session with contract instance
    if myenv["CONTRACTADDR"] == "" {
        session = NewContract(session, client, "name", "satya")
    }

    // If we have an existing contract, load it; if we've deployed a new contract, attempt to load it.
    if myenv["CONTRACTADDR"] != "" {
        session = LoadContract(session, client)
    }



// Loop to implement simple CLI
    for {
       fmt.Printf(
        "Pick an option:\n" + "" +
            "1. Show question.\n" +
            "2. Send answer.\n" +
            "3. Check if you answered correctly.\n" +
            "4. Exit.\n",
    )

    // Reads a single UTF-8 character (rune)
    // from STDIN and switches to case.
    runes := []rune(readStringStdin())
    switch runes[0] {
    case '1':
        readQuestion(session)
        break
    case '2':
        fmt.Println("Type in your answer")
        sendAnswer(session, readStringStdin())
        break
    case '3':
        checkCorrect(session)
        break
    case '4':
        fmt.Println("Bye!")
        return
    default:
        fmt.Println("Invalid option. Please try again.")
        break
    }
}

}

func NewSession(ctx context.Context) (session quiz.QuizSession) {
    loadEnv()
    keystore, err := os.Open(myenv["KEYSTORE"])
    if err != nil {
        log.Printf(
            "could not load keystore from location %s: %v\n",
            myenv["KEYSTORE"],
            err,
        )
    }
    defer keystore.Close()

    keystorepass := myenv["KEYSTOREPASS"]
    auth, err := bind.NewTransactor(keystore, keystorepass)
   // *auth.GasLimit  = big.NewInt(3141592);
    if err != nil {
        log.Printf("%s\n", err)
    }

    // Return session without contract instance
    return quiz.QuizSession{
       // TransactOpts: *auth,
        TransactOpts: bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 4700000,
                GasPrice: big.NewInt(1000000),
	},
        CallOpts: bind.CallOpts{
            From:    auth.From,
            Context: ctx,
        },
    }
}

// NewContract deploys a contract if no existing contract exists
func NewContract(session quiz.QuizSession, client *ethclient.Client, question string, answer string) (quiz.QuizSession) {
    loadEnv()

    // Hash answer before sending it over Ethereum network.
    contractAddress, tx, instance, err := quiz.DeployQuiz(&session.TransactOpts, client, question, stringToKeccak256(answer))
    if err != nil {
        log.Fatalf("could not deploy contract: %v\n", err)
    }
    fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

    session.Contract = instance
    updateEnvFile("CONTRACTADDR", contractAddress.Hex())
    return session
}

// LoadContract loads a contract if one exists
func LoadContract(session quiz.QuizSession, client *ethclient.Client) quiz.QuizSession {
    loadEnv()

    addr := common.HexToAddress(myenv["CONTRACTADDR"])
    instance, err := quiz.NewQuiz(addr, client)
    if err != nil {
        log.Fatalf("could not load contract: %v\n", err)
       log.Println(ErrTransactionWait)
    }
    session.Contract = instance
    return session
}

// Utility functions

// stringToKeccak256 converts a string to a keccak256 hash of type [32]byte
func stringToKeccak256(s string) [32]byte {
    var output [32]byte
    copy(output[:], crypto.Keccak256([]byte(s))[:])
    return output
}

// updateEnvFile updates our env file with a key-value pair
func updateEnvFile(k string, val string) {
    myenv[k] = val
    err := godotenv.Write(myenv, envLoc)
    if err != nil {
        log.Printf("failed to update %s: %v\n", envLoc, err)
    }
}


// readQuestion prints out question stored in contract.
func readQuestion(session quiz.QuizSession) {
    qn, err := session.Question()
    if err != nil {
        log.Printf("could not read question from contract: %v\n", err)
        log.Println(ErrTransactionWait)
        return
    }
    fmt.Printf("Question: %s\n", qn)
    return
}

// sendAnswer sends answer to contract as a keccak256 hash.
func sendAnswer(session quiz.QuizSession, ans string) {
    // Send answer
    txSendAnswer, err := session.SendAnswer(stringToKeccak256(ans))
    if err != nil {
        log.Printf("could not send answer to contract: %v\n", err)
        return
    }
    fmt.Printf("Answer sent! Please wait for tx %s to be confirmed.\n", txSendAnswer.Hash().Hex())
    return
}

// checkCorrect makes a contract message call to check if
// the current account owner has answered the question correctly.
func checkCorrect(session quiz.QuizSession) {
    win, err := session.CheckBoard()
    if err != nil {
        log.Printf("could not check leaderboard: %v\n", err)
        log.Println(ErrTransactionWait)
        return
    }
    fmt.Printf("Were you correct?: %v\n", win)
    return
}

// readStringStdin reads a string from STDIN and strips any trailing \n characters from it.
func readStringStdin() string {
    reader := bufio.NewReader(os.Stdin)
    inputVal, err := reader.ReadString('\n')
    if err != nil {
        log.Printf("invalid option: %v\n", err)
        return ""
    }

    output := strings.TrimSuffix(inputVal, "\n") // Important!
    return output
}
