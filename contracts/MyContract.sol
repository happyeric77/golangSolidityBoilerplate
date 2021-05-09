pragma solidity ^0.8.0;

contract MyContract {
    address public owner;
    string public item;
    constructor () {
        owner = msg.sender;
    }
    function setItem(string memory data) external {
        item = data;
    }
}