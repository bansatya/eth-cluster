pragma solidity >=0.5.2 <0.6.0;

contract Quiz {
    string public question;
    bytes32 public answer;
    mapping (address => bool) public leaderBoard;

    constructor(string memory _qn, bytes32 _ans) public {
        question = _qn;
        answer = _ans;
    }

    function sendAnswer(bytes32 _ans) public returns (bool){
        return updateLeaderBoard(answer == _ans);
    }

    function updateLeaderBoard(bool ok) public returns (bool){
        leaderBoard[msg.sender] = ok;
        return true;
    }

    function checkBoard() public view returns (bool){
        return leaderBoard[msg.sender];
    }
}
