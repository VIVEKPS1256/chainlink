// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {OffRamp} from "../../offRamp/OffRamp.sol";
import {IgnoreContractSize} from "./IgnoreContractSize.sol";

contract OffRampHelper is OffRamp, IgnoreContractSize {
  mapping(uint64 sourceChainSelector => uint256 overrideTimestamp) private s_sourceChainVerificationOverride;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    SourceChainConfigArgs[] memory sourceChainConfigs
  ) OffRamp(staticConfig, dynamicConfig, sourceChainConfigs) {}

  function setExecutionStateHelper(
    uint64 sourceChainSelector,
    uint64 sequenceNumber,
    Internal.MessageExecutionState state
  ) public {
    _setExecutionState(sourceChainSelector, sequenceNumber, state);
  }

  function getExecutionStateBitMap(uint64 sourceChainSelector, uint64 bitmapIndex) public view returns (uint256) {
    return s_executionStates[sourceChainSelector][bitmapIndex];
  }

  function releaseOrMintSingleToken(
    Internal.RampTokenAmount calldata sourceTokenAmount,
    bytes calldata originalSender,
    address receiver,
    uint64 sourceChainSelector,
    bytes calldata offchainTokenData
  ) external returns (Client.EVMTokenAmount memory) {
    return
      _releaseOrMintSingleToken(sourceTokenAmount, originalSender, receiver, sourceChainSelector, offchainTokenData);
  }

  function releaseOrMintTokens(
    Internal.RampTokenAmount[] calldata sourceTokenAmounts,
    bytes calldata originalSender,
    address receiver,
    uint64 sourceChainSelector,
    bytes[] calldata offchainTokenData
  ) external returns (Client.EVMTokenAmount[] memory) {
    return _releaseOrMintTokens(sourceTokenAmounts, originalSender, receiver, sourceChainSelector, offchainTokenData);
  }

  function trialExecute(
    Internal.Any2EVMRampMessage memory message,
    bytes[] memory offchainTokenData
  ) external returns (Internal.MessageExecutionState, bytes memory) {
    return _trialExecute(message, offchainTokenData);
  }

  function executeSingleReport(
    Internal.ExecutionReportSingleChain memory rep,
    uint256[] memory manualExecGasLimits
  ) external {
    _executeSingleReport(rep, manualExecGasLimits);
  }

  function batchExecute(
    Internal.ExecutionReportSingleChain[] memory reports,
    uint256[][] memory manualExecGasLimits
  ) external {
    _batchExecute(reports, manualExecGasLimits);
  }

  function verify(
    uint64 sourceChainSelector,
    bytes32[] memory hashedLeaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) external view returns (uint256 timestamp) {
    return super._verify(sourceChainSelector, hashedLeaves, proofs, proofFlagBits);
  }

  function _verify(
    uint64 sourceChainSelector,
    bytes32[] memory hashedLeaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) internal view override returns (uint256 timestamp) {
    uint256 overrideTimestamp = s_sourceChainVerificationOverride[sourceChainSelector];

    return overrideTimestamp == 0
      ? super._verify(sourceChainSelector, hashedLeaves, proofs, proofFlagBits)
      : overrideTimestamp;
  }

  /// @dev Test helper to override _verify result for easier exec testing
  function setVerifyOverrideResult(uint64 sourceChainSelector, uint256 overrideTimestamp) external {
    s_sourceChainVerificationOverride[sourceChainSelector] = overrideTimestamp;
  }

  /// @dev Test helper to directly set a root's timestamp
  function setRootTimestamp(uint64 sourceChainSelector, bytes32 root, uint256 timestamp) external {
    s_roots[sourceChainSelector][root] = timestamp;
  }
}
