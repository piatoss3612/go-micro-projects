// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

contract Todo {
    address owner;
    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier validId(uint _id) {
        require(_id < tasks.length);
        _;
    }

    function add(string memory _content) public onlyOwner {
        tasks.push(Task(_content, false));
    }

    function get(uint _id) public view validId(_id) returns(Task memory) {
        return tasks[_id];
    }

    function list() public view returns(Task[] memory) {
        return tasks;
    }

    function update(uint _id, string memory _content) public onlyOwner validId(_id) {
        tasks[_id].content = _content;
    }

    function toggle(uint _id) public onlyOwner validId(_id) {
        tasks[_id].status = !tasks[_id].status;
    }

    function remove(uint _id) public onlyOwner validId(_id) {
        for (uint i = _id; i < tasks.length - 1;) {
            tasks[i] = tasks[i+1];
            unchecked {
                i += 1;
            }
        }
        tasks.pop();
    }
}