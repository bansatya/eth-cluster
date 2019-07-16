// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quiz

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// QuizABI is the input ABI used to generate the binding from.
const QuizABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_ans\",\"type\":\"bytes32\"}],\"name\":\"sendAnswer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"question\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkBoard\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"answer\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"leaderBoard\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"name\":\"updateLeaderBoard\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_qn\",\"type\":\"string\"},{\"name\":\"_ans\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// QuizBin is the compiled bytecode used for deploying new contracts.
var QuizBin = "0x608060405234801561001057600080fd5b506040516105453803806105458339818101604052604081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8281019050602081018481111561006157600080fd5b815185600182028301116401000000008211171561007e57600080fd5b50509291906020018051906020019092919050505081600090805190602001906100a99291906100b8565b5080600181905550505061015d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f957805160ff1916838001178555610127565b82800160010185558215610127579182015b8281111561012657825182559160200191906001019061010b565b5b5090506101349190610138565b5090565b61015a91905b8082111561015657600081600090555060010161013e565b5090565b90565b6103d98061016c6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806317d1653c146100675780633fad9ae0146100ad57806377f46bff1461013057806385bb7d6914610152578063a54a928814610170578063e0e390cd146101cc575b600080fd5b6100936004803603602081101561007d57600080fd5b8101908080359060200190929190505050610214565b604051808215151515815260200191505060405180910390f35b6100b561022a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f55780820151818401526020810190506100da565b50505050905090810190601f1680156101225780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101386102c8565b604051808215151515815260200191505060405180910390f35b61015a61031c565b6040518082815260200191505060405180910390f35b6101b26004803603602081101561018657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610322565b604051808215151515815260200191505060405180910390f35b6101fa600480360360208110156101e257600080fd5b81019080803515159060200190929190505050610342565b604051808215151515815260200191505060405180910390f35b60006102238260015414610342565b9050919050565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102c05780601f10610295576101008083540402835291602001916102c0565b820191906000526020600020905b8154815290600101906020018083116102a357829003601f168201915b505050505081565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905090565b60015481565b60026020528060005260406000206000915054906101000a900460ff1681565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905091905056fea265627a7a72305820767bcd73d6371d94f8e876de48f03536546331b2401fd980329e068b9ceed71564736f6c634300050a0032"

// DeployQuiz deploys a new Ethereum contract, binding an instance of Quiz to it.
func DeployQuiz(auth *bind.TransactOpts, backend bind.ContractBackend, _qn string, _ans [32]byte) (common.Address, *types.Transaction, *Quiz, error) {
	parsed, err := abi.JSON(strings.NewReader(QuizABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QuizBin), backend, _qn, _ans)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// Quiz is an auto generated Go binding around an Ethereum contract.
type Quiz struct {
	QuizCaller     // Read-only binding to the contract
	QuizTransactor // Write-only binding to the contract
	QuizFilterer   // Log filterer for contract events
}

// QuizCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuizCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuizTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuizFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuizSession struct {
	Contract     *Quiz             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuizCallerSession struct {
	Contract *QuizCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuizTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuizTransactorSession struct {
	Contract     *QuizTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuizRaw struct {
	Contract *Quiz // Generic contract binding to access the raw methods on
}

// QuizCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuizCallerRaw struct {
	Contract *QuizCaller // Generic read-only contract binding to access the raw methods on
}

// QuizTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuizTransactorRaw struct {
	Contract *QuizTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuiz creates a new instance of Quiz, bound to a specific deployed contract.
func NewQuiz(address common.Address, backend bind.ContractBackend) (*Quiz, error) {
	contract, err := bindQuiz(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// NewQuizCaller creates a new read-only instance of Quiz, bound to a specific deployed contract.
func NewQuizCaller(address common.Address, caller bind.ContractCaller) (*QuizCaller, error) {
	contract, err := bindQuiz(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuizCaller{contract: contract}, nil
}

// NewQuizTransactor creates a new write-only instance of Quiz, bound to a specific deployed contract.
func NewQuizTransactor(address common.Address, transactor bind.ContractTransactor) (*QuizTransactor, error) {
	contract, err := bindQuiz(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuizTransactor{contract: contract}, nil
}

// NewQuizFilterer creates a new log filterer instance of Quiz, bound to a specific deployed contract.
func NewQuizFilterer(address common.Address, filterer bind.ContractFilterer) (*QuizFilterer, error) {
	contract, err := bindQuiz(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuizFilterer{contract: contract}, nil
}

// bindQuiz binds a generic wrapper to an already deployed contract.
func bindQuiz(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuizABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.QuizCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transact(opts, method, params...)
}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() constant returns(bytes32)
func (_Quiz *QuizCaller) Answer(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Quiz.contract.Call(opts, out, "answer")
	return *ret0, err
}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() constant returns(bytes32)
func (_Quiz *QuizSession) Answer() ([32]byte, error) {
	return _Quiz.Contract.Answer(&_Quiz.CallOpts)
}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() constant returns(bytes32)
func (_Quiz *QuizCallerSession) Answer() ([32]byte, error) {
	return _Quiz.Contract.Answer(&_Quiz.CallOpts)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() constant returns(bool)
func (_Quiz *QuizCaller) CheckBoard(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Quiz.contract.Call(opts, out, "checkBoard")
	return *ret0, err
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() constant returns(bool)
func (_Quiz *QuizSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() constant returns(bool)
func (_Quiz *QuizCallerSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// LeaderBoard is a free data retrieval call binding the contract method 0xa54a9288.
//
// Solidity: function leaderBoard(address ) constant returns(bool)
func (_Quiz *QuizCaller) LeaderBoard(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Quiz.contract.Call(opts, out, "leaderBoard", arg0)
	return *ret0, err
}

// LeaderBoard is a free data retrieval call binding the contract method 0xa54a9288.
//
// Solidity: function leaderBoard(address ) constant returns(bool)
func (_Quiz *QuizSession) LeaderBoard(arg0 common.Address) (bool, error) {
	return _Quiz.Contract.LeaderBoard(&_Quiz.CallOpts, arg0)
}

// LeaderBoard is a free data retrieval call binding the contract method 0xa54a9288.
//
// Solidity: function leaderBoard(address ) constant returns(bool)
func (_Quiz *QuizCallerSession) LeaderBoard(arg0 common.Address) (bool, error) {
	return _Quiz.Contract.LeaderBoard(&_Quiz.CallOpts, arg0)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() constant returns(string)
func (_Quiz *QuizCaller) Question(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Quiz.contract.Call(opts, out, "question")
	return *ret0, err
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() constant returns(string)
func (_Quiz *QuizSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() constant returns(string)
func (_Quiz *QuizCallerSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactor) SendAnswer(opts *bind.TransactOpts, _ans [32]byte) (*types.Transaction, error) {
	return _Quiz.contract.Transact(opts, "sendAnswer", _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactorSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}

// UpdateLeaderBoard is a paid mutator transaction binding the contract method 0xe0e390cd.
//
// Solidity: function updateLeaderBoard(bool ok) returns(bool)
func (_Quiz *QuizTransactor) UpdateLeaderBoard(opts *bind.TransactOpts, ok bool) (*types.Transaction, error) {
	return _Quiz.contract.Transact(opts, "updateLeaderBoard", ok)
}

// UpdateLeaderBoard is a paid mutator transaction binding the contract method 0xe0e390cd.
//
// Solidity: function updateLeaderBoard(bool ok) returns(bool)
func (_Quiz *QuizSession) UpdateLeaderBoard(ok bool) (*types.Transaction, error) {
	return _Quiz.Contract.UpdateLeaderBoard(&_Quiz.TransactOpts, ok)
}

// UpdateLeaderBoard is a paid mutator transaction binding the contract method 0xe0e390cd.
//
// Solidity: function updateLeaderBoard(bool ok) returns(bool)
func (_Quiz *QuizTransactorSession) UpdateLeaderBoard(ok bool) (*types.Transaction, error) {
	return _Quiz.Contract.UpdateLeaderBoard(&_Quiz.TransactOpts, ok)
}
