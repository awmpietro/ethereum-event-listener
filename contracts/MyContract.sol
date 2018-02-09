pragma solidity ^0.4.9;

contract MyContract {
  string public name;
  uint256 public count;

  event _MyEvent(string name, uint256 count);

  function greet(string _name) public {
    name = _name;
    count += 1;
    _MyEvent(_name, count);
  }
}
