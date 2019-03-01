pragma solidity 0.4.24;

/**
 * @title The LINK exchange contract
 */
contract LinkEx {

  uint256 private historicRate;
  uint256 private rate;
  uint256 private rateHeight;

  function currentRate() view public returns (uint256) {
    if (rateHeight != 0 && block.number == rateHeight) {
      return rate;
    }
    return historicRate;
  }

  function update(uint256 _rate) public {
    rateHeight = block.number;
    rate = _rate;
  }
}
