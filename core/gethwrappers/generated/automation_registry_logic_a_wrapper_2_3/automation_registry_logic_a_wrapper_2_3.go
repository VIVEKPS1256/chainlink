// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package automation_registry_logic_a_wrapper_2_3

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type AutomationRegistryBase23BillingConfig struct {
	GasFeePPB         uint32
	FlatFeeMilliCents *big.Int
	PriceFeed         common.Address
	FallbackPrice     *big.Int
	MinSpend          *big.Int
}

var AutomationRegistryLogicAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractAutomationRegistryLogicB2_3\",\"name\":\"logicB\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ArrayHasNoEntries\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CheckDataExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateEntry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasLimitCanOnlyIncrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasLimitOutsideRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfFaultyOracles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBillingToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTrigger\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTriggerType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MigrationNotPermitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSettleOffchain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSettleOnchain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyActiveSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyActiveTransmitters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByLINKToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrRegistrar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByPayee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByProposedAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByProposedPayee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByUpkeepPrivilegeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyFinanceAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPausedUpkeep\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySimulatedBackend\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyUnpausedUpkeep\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParameterLengthError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepeatedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepeatedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"TargetCheckReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyOracles\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TranscoderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepNotCanceled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpkeepNotNeeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"privilegeConfig\",\"type\":\"bytes\"}],\"name\":\"AdminPrivilegeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"BillingConfigOverridden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"BillingConfigOverrideRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"gasFeePPB\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"flatFeeMilliCents\",\"type\":\"uint24\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fallbackPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"minSpend\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structAutomationRegistryBase2_3.BillingConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"BillingConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trigger\",\"type\":\"bytes\"}],\"name\":\"CancelledUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newModule\",\"type\":\"address\"}],\"name\":\"ChainSpecificModuleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dedupKey\",\"type\":\"bytes32\"}],\"name\":\"DedupKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"FundsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trigger\",\"type\":\"bytes\"}],\"name\":\"InsufficientFundsUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payments\",\"type\":\"uint256[]\"}],\"name\":\"NOPsSettledOffchain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"PayeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"PaymentWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trigger\",\"type\":\"bytes\"}],\"name\":\"ReorgedUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trigger\",\"type\":\"bytes\"}],\"name\":\"StaleUpkeepReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"UpkeepAdminTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"UpkeepAdminTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"atBlockHeight\",\"type\":\"uint64\"}],\"name\":\"UpkeepCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newCheckData\",\"type\":\"bytes\"}],\"name\":\"UpkeepCheckDataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"gasLimit\",\"type\":\"uint96\"}],\"name\":\"UpkeepGasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"UpkeepMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"UpkeepOffchainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"totalPayment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"trigger\",\"type\":\"bytes\"}],\"name\":\"UpkeepPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"privilegeConfig\",\"type\":\"bytes\"}],\"name\":\"UpkeepPrivilegeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"importedFrom\",\"type\":\"address\"}],\"name\":\"UpkeepReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"performGas\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"UpkeepRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"triggerConfig\",\"type\":\"bytes\"}],\"name\":\"UpkeepTriggerConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpkeepUnpaused\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancelUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"values\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"checkCallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"},{\"internalType\":\"enumAutomationRegistryBase2_3.UpkeepFailureReason\",\"name\":\"upkeepFailureReason\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"triggerData\",\"type\":\"bytes\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"},{\"internalType\":\"enumAutomationRegistryBase2_3.UpkeepFailureReason\",\"name\":\"upkeepFailureReason\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastGasWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkUSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"},{\"internalType\":\"enumAutomationRegistryBase2_3.UpkeepFailureReason\",\"name\":\"upkeepFailureReason\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fastGasWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkUSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"executeCallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"},{\"internalType\":\"enumAutomationRegistryBase2_3.UpkeepFailureReason\",\"name\":\"upkeepFailureReason\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fallbackTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"migrateUpkeeps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedUpkeeps\",\"type\":\"bytes\"}],\"name\":\"receiveUpkeeps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"enumAutomationRegistryBase2_3.Trigger\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"billingToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"triggerConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"registerUpkeep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b506040516200643f3803806200643f83398101604081905262000035916200062f565b80816001600160a01b031663ca30e6036040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000075573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200009b91906200062f565b826001600160a01b031663226cf83c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010091906200062f565b836001600160a01b031663614486af6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200013f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200016591906200062f565b846001600160a01b0316636709d0e56040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001a4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001ca91906200062f565b856001600160a01b0316635425d8ac6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000209573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200022f91906200062f565b866001600160a01b031663a08714c06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200026e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200029491906200062f565b876001600160a01b031663c5b964e06040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002d3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002f9919062000656565b886001600160a01b031663ac4dc59a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000338573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200035e91906200062f565b3380600081620003b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620003e857620003e8816200056b565b5050506001600160a01b0380891660805287811660a05286811660c05285811660e052848116610100528316610120526025805483919060ff19166001838181111562000439576200043962000679565b0217905550806001600160a01b0316610140816001600160a01b03168152505060c0516001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200049a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620004c091906200068f565b60ff1660a0516001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000504573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200052a91906200068f565b60ff16146200054c576040516301f86e1760e41b815260040160405180910390fd5b5050506001600160a01b039095166101605250620006b4945050505050565b336001600160a01b03821603620005c55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620003ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b03811681146200062c57600080fd5b50565b6000602082840312156200064257600080fd5b81516200064f8162000616565b9392505050565b6000602082840312156200066957600080fd5b8151600281106200064f57600080fd5b634e487b7160e01b600052602160045260246000fd5b600060208284031215620006a257600080fd5b815160ff811681146200064f57600080fd5b60805160a05160c05160e05161010051610120516101405161016051615d1b620007246000396000818160dc01526101720152600050506000612914015260008181611d10015261202901526000612a4801526000613bd901526000612b2c01526000611b520152615d1b6000f3fe608060405260043610620000da5760003560e01c80638e86139b116200007f578063ce7dc5b41162000055578063ce7dc5b414620002da578063f2fde38b14620002ff578063f7d334ba146200032457620000da565b80638e86139b146200025c578063c62cf6841462000281578063c804802214620002b557620000da565b806379ba509711620000b557806379ba509714620001f257806385c1b0ba146200020a5780638da5cb5b146200022f57620000da565b806329c5efad1462000122578063349e8cca146200016257806371791aa014620001b8575b7f00000000000000000000000000000000000000000000000000000000000000003660008037600080366000845af43d6000803e8080156200011b573d6000f35b3d6000fd5b005b3480156200012f57600080fd5b50620001476200014136600462004403565b62000349565b6040516200015994939291906200452b565b60405180910390f35b3480156200016f57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200162000159565b348015620001c557600080fd5b50620001dd620001d736600462004403565b6200063a565b60405162000159979695949392919062004568565b348015620001ff57600080fd5b506200012062000dc3565b3480156200021757600080fd5b506200012062000229366004620045ef565b62000ec6565b3480156200023c57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff1662000192565b3480156200026957600080fd5b50620001206200027b366004620046c8565b62001bd3565b3480156200028e57600080fd5b50620002a6620002a036600462004733565b62001f5b565b60405190815260200162000159565b348015620002c257600080fd5b5062000120620002d436600462004828565b62002308565b348015620002e757600080fd5b5062000147620002f9366004620048fe565b620027e0565b3480156200030c57600080fd5b50620001206200031e36600462004975565b620028a6565b3480156200033157600080fd5b50620001dd6200034336600462004828565b620028be565b600060606000806200035a620028fc565b600086815260046020908152604091829020825161012081018452815460ff8082161515835261010080830490911615159483019490945263ffffffff620100008204811695830195909552660100000000000081048516606083015273ffffffffffffffffffffffffffffffffffffffff6a01000000000000000000009091048116608083015260018301546fffffffffffffffffffffffffffffffff811660a08401526bffffffffffffffffffffffff70010000000000000000000000000000000082041660c08401527c0100000000000000000000000000000000000000000000000000000000900490941660e0820152600290910154909216908201525a9150600080826080015173ffffffffffffffffffffffffffffffffffffffff1663f00e6a2a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620004b1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620004d79190620049a2565b73ffffffffffffffffffffffffffffffffffffffff16601660000160149054906101000a900463ffffffff1663ffffffff1689604051620005199190620049c2565b60006040518083038160008787f1925050503d806000811462000559576040519150601f19603f3d011682016040523d82523d6000602084013e6200055e565b606091505b50915091505a62000570908562004a0f565b9350816200059b57600060405180602001604052806000815250600796509650965050505062000631565b80806020019051810190620005b1919062004a80565b909750955086620005df57600060405180602001604052806000815250600496509650965050505062000631565b60185486517401000000000000000000000000000000000000000090910463ffffffff1610156200062d57600060405180602001604052806000815250600596509650965050505062000631565b5050505b92959194509250565b60006060600080600080600062000650620028fc565b60006200065d8a6200296e565b905060006014604051806101200160405290816000820160009054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff166bffffffffffffffffffffffff16815260200160008201600c9054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160109054906101000a900462ffffff1662ffffff1662ffffff1681526020016000820160139054906101000a900461ffff1661ffff1661ffff1681526020016000820160159054906101000a900460ff1660ff1660ff1681526020016000820160169054906101000a900460ff161515151581526020016000820160179054906101000a900460ff161515151581526020016000820160189054906101000a900460ff161515151581526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090506000600460008d8152602001908152602001600020604051806101200160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900460ff161515151581526020016000820160029054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160069054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200160008201600a9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1681526020016001820160109054906101000a90046bffffffffffffffffffffffff166bffffffffffffffffffffffff166bffffffffffffffffffffffff16815260200160018201601c9054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505090506000808360a001511562000a225750506040805160208101825260008082529290910151919a5098506009975089965063ffffffff16945085935083925062000db7915050565b606083015163ffffffff9081161462000a6e5750506040805160208101825260008082529290910151919a5098506001975089965063ffffffff16945085935083925062000db7915050565b82511562000aaf5750506040805160208101825260008082529290910151919a5098506002975089965063ffffffff16945085935083925062000db7915050565b62000aba8462002a24565b80945081985082995050505062000ae18e858786604001518b8b888a610100015162002c2c565b9050806bffffffffffffffffffffffff168360c001516bffffffffffffffffffffffff16101562000b455750506040805160208101825260008082529290910151919a5098506006975089965063ffffffff16945085935083925062000db7915050565b5050600062000b568d858e62002fae565b90505a9750600080836080015173ffffffffffffffffffffffffffffffffffffffff1663f00e6a2a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000bae573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000bd49190620049a2565b73ffffffffffffffffffffffffffffffffffffffff16601660000160149054906101000a900463ffffffff1663ffffffff168460405162000c169190620049c2565b60006040518083038160008787f1925050503d806000811462000c56576040519150601f19603f3d011682016040523d82523d6000602084013e62000c5b565b606091505b50915091505a62000c6d908b62004a0f565b99508162000cf7576018548151780100000000000000000000000000000000000000000000000090910463ffffffff16101562000cd55750506040805160208101825260008082529390910151929b509950600898505063ffffffff16945062000db7915050565b604090930151929a50600399505063ffffffff909116955062000db792505050565b8080602001905181019062000d0d919062004a80565b909d509b508c62000d495750506040805160208101825260008082529390910151929b509950600498505063ffffffff16945062000db7915050565b6018548c517401000000000000000000000000000000000000000090910463ffffffff16101562000da55750506040805160208101825260008082529390910151929b509950600598505063ffffffff16945062000db7915050565b5050506040015163ffffffff16945050505b92959891949750929550565b60015473ffffffffffffffffffffffffffffffffffffffff16331462000e4a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600173ffffffffffffffffffffffffffffffffffffffff82166000908152601d602052604090205460ff16600381111562000f055762000f05620044c0565b1415801562000f515750600373ffffffffffffffffffffffffffffffffffffffff82166000908152601d602052604090205460ff16600381111562000f4e5762000f4e620044c0565b14155b1562000f89576040517f0ebeec3c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60165473ffffffffffffffffffffffffffffffffffffffff1662000fd9576040517fd12d7d8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082900362001015576040517f2c2fc94100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805161012081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526000808567ffffffffffffffff8111156200107c576200107c62004289565b604051908082528060200260200182016040528015620010a6578160200160208202803683370190505b50905060008667ffffffffffffffff811115620010c757620010c762004289565b6040519080825280602002602001820160405280156200115e57816020015b604080516101208101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c0820181905260e0820181905261010082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909201910181620010e65790505b50905060008767ffffffffffffffff8111156200117f576200117f62004289565b604051908082528060200260200182016040528015620011b457816020015b60608152602001906001900390816200119e5790505b50905060008867ffffffffffffffff811115620011d557620011d562004289565b6040519080825280602002602001820160405280156200120a57816020015b6060815260200190600190039081620011f45790505b50905060008967ffffffffffffffff8111156200122b576200122b62004289565b6040519080825280602002602001820160405280156200126057816020015b60608152602001906001900390816200124a5790505b50905060005b8a81101562001889578b8b8281811062001284576200128462004acd565b60209081029290920135600081815260048452604090819020815161012081018352815460ff8082161515835261010080830490911615159783019790975263ffffffff620100008204811694830194909452660100000000000081048416606083015273ffffffffffffffffffffffffffffffffffffffff6a01000000000000000000009091048116608083015260018301546fffffffffffffffffffffffffffffffff811660a08401526bffffffffffffffffffffffff70010000000000000000000000000000000082041660c08401527c0100000000000000000000000000000000000000000000000000000000900490931660e0820152600290910154909116938101939093529a50909850620013a19050896200319e565b60808801516040517f1a5da6c800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c8116600483015290911690631a5da6c890602401600060405180830381600087803b1580156200141157600080fd5b505af115801562001426573d6000803e3d6000fd5b505050508785828151811062001440576200144062004acd565b6020026020010181905250600560008a815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686828151811062001494576200149462004acd565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910182015260008a81526007909152604090208054620014d39062004afc565b80601f0160208091040260200160405190810160405280929190818152602001828054620015019062004afc565b8015620015525780601f10620015265761010080835404028352916020019162001552565b820191906000526020600020905b8154815290600101906020018083116200153457829003601f168201915b50505050508482815181106200156c576200156c62004acd565b6020026020010181905250601e60008a81526020019081526020016000208054620015979062004afc565b80601f0160208091040260200160405190810160405280929190818152602001828054620015c59062004afc565b8015620016165780601f10620015ea5761010080835404028352916020019162001616565b820191906000526020600020905b815481529060010190602001808311620015f857829003601f168201915b505050505083828151811062001630576200163062004acd565b6020026020010181905250601f60008a815260200190815260200160002080546200165b9062004afc565b80601f0160208091040260200160405190810160405280929190818152602001828054620016899062004afc565b8015620016da5780601f10620016ae57610100808354040283529160200191620016da565b820191906000526020600020905b815481529060010190602001808311620016bc57829003601f168201915b5050505050828281518110620016f457620016f462004acd565b60200260200101819052508760c001516bffffffffffffffffffffffff16876200171f919062004b51565b60008a815260046020908152604080832080547fffff0000000000000000000000000000000000000000000000000000000000001681556001810184905560020180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055600790915281209198506200179c919062004225565b6000898152601e60205260408120620017b59162004225565b6000898152601f60205260408120620017ce9162004225565b600089815260066020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556200180f60028a62003255565b5060c0880151604080516bffffffffffffffffffffffff909216825273ffffffffffffffffffffffffffffffffffffffff8c1660208301528a917fb38647142fbb1ea4c000fc4569b37a4e9a9f6313317b84ee3e5326c1a6cd06ff910160405180910390a280620018808162004b67565b91505062001266565b5061010087015173ffffffffffffffffffffffffffffffffffffffff166000908152601c6020526040902054620018c290879062004a0f565b61010088015173ffffffffffffffffffffffffffffffffffffffff166000908152601c60205260408120919091558b8b868167ffffffffffffffff8111156200190f576200190f62004289565b60405190808252806020026020018201604052801562001939578160200160208202803683370190505b50898888886040516020016200195798979695949392919062004d27565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282526016547faab9edd6000000000000000000000000000000000000000000000000000000008452915190935073ffffffffffffffffffffffffffffffffffffffff808e1693638e86139b939091169163c71249ab91600391869163aab9edd69160048083019260209291908290030181865afa15801562001a09573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001a2f919062004e06565b866040518463ffffffff1660e01b815260040162001a509392919062004e2b565b600060405180830381865afa15801562001a6e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405262001ab6919081019062004e52565b6040518263ffffffff1660e01b815260040162001ad4919062004e8b565b600060405180830381600087803b15801562001aef57600080fd5b505af115801562001b04573d6000803e3d6000fd5b50506040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d81166004830152602482018b90527f000000000000000000000000000000000000000000000000000000000000000016925063a9059cbb91506044016020604051808303816000875af115801562001b9e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001bc4919062004ea0565b50505050505050505050505050565b6002336000908152601d602052604090205460ff16600381111562001bfc5762001bfc620044c0565b1415801562001c3257506003336000908152601d602052604090205460ff16600381111562001c2f5762001c2f620044c0565b14155b1562001c6a576040517f0ebeec3c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080808080808062001c80888a018a62005107565b965096509650965096509650965060005b875181101562001f4f57600073ffffffffffffffffffffffffffffffffffffffff1687828151811062001cc85762001cc862004acd565b60200260200101516080015173ffffffffffffffffffffffffffffffffffffffff160362001ddc5785818151811062001d055762001d0562004acd565b6020026020010151307f000000000000000000000000000000000000000000000000000000000000000060405162001d3d9062004264565b73ffffffffffffffffffffffffffffffffffffffff938416815291831660208301529091166040820152606001604051809103906000f08015801562001d87573d6000803e3d6000fd5b5087828151811062001d9d5762001d9d62004acd565b60200260200101516080019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b62001e9488828151811062001df55762001df562004acd565b602002602001015188838151811062001e125762001e1262004acd565b602002602001015187848151811062001e2f5762001e2f62004acd565b602002602001015187858151811062001e4c5762001e4c62004acd565b602002602001015187868151811062001e695762001e6962004acd565b602002602001015187878151811062001e865762001e8662004acd565b60200260200101516200326c565b87818151811062001ea95762001ea962004acd565b60200260200101517f74931a144e43a50694897f241d973aecb5024c0e910f9bb80a163ea3c1cf5a7188838151811062001ee75762001ee762004acd565b602002602001015160c001513360405162001f329291906bffffffffffffffffffffffff92909216825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60405180910390a28062001f468162004b67565b91505062001c91565b50505050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16331480159062001f8f575062001f8d60093362003789565b155b1562001fc7576040517fd48b678b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a163b62002016576040517f09ee12d500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200202187620037b9565b905060008a307f0000000000000000000000000000000000000000000000000000000000000000604051620020569062004264565b73ffffffffffffffffffffffffffffffffffffffff938416815291831660208301529091166040820152606001604051809103906000f080158015620020a0573d6000803e3d6000fd5b50905062002191826040518061012001604052806000151581526020016000151581526020018d63ffffffff16815260200163ffffffff801681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006bffffffffffffffffffffffff168152602001600063ffffffff1681526020018a73ffffffffffffffffffffffffffffffffffffffff168152508b89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91506200326c9050565b601680547c0100000000000000000000000000000000000000000000000000000000900463ffffffff1690601c620021c98362005238565b91906101000a81548163ffffffff021916908363ffffffff16021790555050817fbae366358c023f887e791d7a62f2e4316f1026bd77f6fb49501a917b3bc5d0128b8b6040516200224292919063ffffffff92909216825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60405180910390a2817fcba2d5723b2ee59e53a8e8a82a4a7caf4fdfe70e9f7c582950bf7e7a5c24e83d87876040516200227e929190620052a7565b60405180910390a2817f2b72ac786c97e68dbab71023ed6f2bdbfc80ad9bb7808941929229d71b7d566485604051620022b8919062004e8b565b60405180910390a2817f3e8740446213c8a77d40e08f79136ce3f347d13ed270a6ebdf57159e0faf485084604051620022f2919062004e8b565b60405180910390a2509998505050505050505050565b6000818152600460209081526040808320815161012081018352815460ff8082161515835261010080830490911615159583019590955263ffffffff620100008204811694830194909452660100000000000081048416606083015273ffffffffffffffffffffffffffffffffffffffff6a01000000000000000000009091048116608083015260018301546fffffffffffffffffffffffffffffffff811660a08401526bffffffffffffffffffffffff70010000000000000000000000000000000082041660c08401527c0100000000000000000000000000000000000000000000000000000000900490931660e082015260029091015490911691810191909152906200242c60005473ffffffffffffffffffffffffffffffffffffffff1690565b61010083015173ffffffffffffffffffffffffffffffffffffffff90811660009081526022602090815260408083206002015460155482517f57e871e70000000000000000000000000000000000000000000000000000000081529251968616331497506bffffffffffffffffffffffff90911695939416926357e871e7926004808401939192918290030181865afa158015620024ce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620024f49190620052bd565b9050836060015163ffffffff166000036200253b576040517ffbc0357800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606084015163ffffffff9081161462002580576040517f9c0083a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82158015620025b3575060008581526005602052604090205473ffffffffffffffffffffffffffffffffffffffff163314155b15620025eb576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826200260157620025fe60328262004b51565b90505b6000858152600460205260409020805463ffffffff8084166601000000000000027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff909216919091179091556200265e9060029087906200325516565b506000826bffffffffffffffffffffffff168560a001516fffffffffffffffffffffffffffffffff161015620026d15760a08501516200269f9084620052d7565b90508460c001516bffffffffffffffffffffffff16816bffffffffffffffffffffffff161115620026d1575060c08401515b808560c00151620026e39190620052d7565b600087815260046020908152604080832060010180547fffffffff000000000000000000000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000006bffffffffffffffffffffffff9687160217905561010089015173ffffffffffffffffffffffffffffffffffffffff168352601c909152902054620027759183169062004a0f565b61010086015173ffffffffffffffffffffffffffffffffffffffff166000908152601c602052604080822092909255905167ffffffffffffffff84169188917f91cb3bb75cfbd718bbfccc56b7f53d92d7048ef4ca39a3b7b7c6d4af1f7911819190a3505050505050565b600060606000806000634b56a42e60e01b8888886040516024016200280893929190620052ff565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152905062002893898262000349565b929c919b50995090975095505050505050565b620028b062003a58565b620028bb8162003adb565b50565b600060606000806000806000620028e588604051806020016040528060008152506200063a565b959e949d50929b5090995097509550909350915050565b3273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146200296c576040517fb60ac5db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000818160045b600f81101562002a03577fff000000000000000000000000000000000000000000000000000000000000008216838260208110620029b757620029b762004acd565b1a60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614620029ee57506000949350505050565b80620029fa8162004b67565b91505062002975565b5081600f1a600181111562002a1c5762002a1c620044c0565b949350505050565b600080600080846040015162ffffff1690506000808263ffffffff161190506000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801562002ab2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062002ad891906200534e565b509450909250505060008113158062002af057508142105b8062002b15575082801562002b15575062002b0c824262004a0f565b8463ffffffff16105b1562002b2657601954965062002b2a565b8096505b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801562002b96573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062002bbc91906200534e565b509450909250505060008113158062002bd457508142105b8062002bf9575082801562002bf9575062002bf0824262004a0f565b8463ffffffff16105b1562002c0a57601a54955062002c0e565b8095505b868662002c1b8a62003bd2565b965096509650505050509193909250565b600080808089600181111562002c465762002c46620044c0565b0362002c56575061ea6062002cb0565b600189600181111562002c6d5762002c6d620044c0565b0362002c7e575062014c0862002cb0565b6040517ff2b2d41200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008a60800151600162002cc59190620053a3565b62002cd59060ff166040620053bf565b60185462002d05906103a49074010000000000000000000000000000000000000000900463ffffffff1662004b51565b62002d11919062004b51565b601554604080517fde9ee35e0000000000000000000000000000000000000000000000000000000081528151939450600093849373ffffffffffffffffffffffffffffffffffffffff169263de9ee35e92600480820193918290030181865afa15801562002d83573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062002da99190620053d9565b9092509050818362002dbd83601862004b51565b62002dc99190620053bf565b60808f015162002ddb906001620053a3565b62002dec9060ff166115e0620053bf565b62002df8919062004b51565b62002e04919062004b51565b62002e10908562004b51565b6101008e01516040517f125441400000000000000000000000000000000000000000000000000000000081526004810186905291955073ffffffffffffffffffffffffffffffffffffffff1690631254414090602401602060405180830381865afa15801562002e84573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062002eaa9190620052bd565b8d6060015161ffff1662002ebf9190620053bf565b9450505050600062002ed28b8662003ccc565b60008d815260046020526040902054909150610100900460ff161562002f385760008c81526023602090815260409182902082518084019093525463ffffffff811680845262ffffff640100000000909204821693830193845284529151909116908201525b600062002f888c6040518061010001604052808d63ffffffff1681526020018681526020018781526020018c81526020018b81526020018a81526020018581526020016000151581525062003e27565b6020810151815191925062002f9d91620053fe565b9d9c50505050505050505050505050565b6060600083600181111562002fc75762002fc7620044c0565b0362003094576000848152600760205260409081902090517f6e04ff0d00000000000000000000000000000000000000000000000000000000916200300f91602401620054c9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152905062003197565b6001836001811115620030ab57620030ab620044c0565b0362002c7e57600082806020019051810190620030c9919062005540565b6000868152600760205260409081902090519192507f40691db400000000000000000000000000000000000000000000000000000000916200311091849160240162005654565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091529150620031979050565b9392505050565b60008181526005602052604090205473ffffffffffffffffffffffffffffffffffffffff163314620031fc576040517fa47c170600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000818152600460205260409020546601000000000000900463ffffffff90811614620028bb576040517f9c0083a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006200326383836200401d565b90505b92915050565b601454760100000000000000000000000000000000000000000000900460ff1615620032c4576040517f24522f3400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60175483517c010000000000000000000000000000000000000000000000000000000090910463ffffffff16101562003329576040517fae7235df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108fc856040015163ffffffff1610806200336f5750601654604086015163ffffffff780100000000000000000000000000000000000000000000000090920482169116115b15620033a7576040517f14c237fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000868152600460205260409020546a0100000000000000000000900473ffffffffffffffffffffffffffffffffffffffff161562003412576040517f6e3b930b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61010085015173ffffffffffffffffffffffffffffffffffffffff90811660009081526022602052604090205467010000000000000090041662003482576040517f1183afea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846004600088815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548160ff02191690831515021790555060408201518160000160026101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600a6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060c08201518160010160106101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555060e082015181600101601c6101000a81548163ffffffff021916908363ffffffff1602179055506101008201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050836005600088815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600760008881526020019081526020016000209081620036c791906200576f565b5060c085015161010086015173ffffffffffffffffffffffffffffffffffffffff166000908152601c602052604090205462003712916bffffffffffffffffffffffff169062004b51565b61010086015173ffffffffffffffffffffffffffffffffffffffff166000908152601c6020908152604080832093909355888252601e9052206200375783826200576f565b506000868152601f602052604090206200377282826200576f565b506200378060028762004121565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151562003263565b601554604080517f57e871e70000000000000000000000000000000000000000000000000000000081529051600092839273ffffffffffffffffffffffffffffffffffffffff90911691839183916385df51fd9160019184916357e871e79160048083019260209291908290030181865afa1580156200383d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620038639190620052bd565b6200386f919062004a0f565b6040518263ffffffff1660e01b81526004016200388e91815260200190565b602060405180830381865afa158015620038ac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620038d29190620052bd565b60165460408051602081019390935230908301527c0100000000000000000000000000000000000000000000000000000000900463ffffffff166060820152608001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083015201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060045b600f811015620039e65783828281518110620039a257620039a262004acd565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080620039dd8162004b67565b91505062003982565b50846001811115620039fc57620039fc620044c0565b60f81b81600f8151811062003a155762003a1562004acd565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535062003a4f8162005896565b95945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146200296c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640162000e41565b3373ffffffffffffffffffffffffffffffffffffffff82160362003b5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000e41565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801562003c43573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062003c6991906200534e565b5093505092505060008213158062003c8057508042105b8062003cb457506000846040015162ffffff1611801562003cb4575062003ca8814262004a0f565b846040015162ffffff16105b1562003cc5575050601b5492915050565b5092915050565b604080516060810182526000808252602080830182815283850183905273ffffffffffffffffffffffffffffffffffffffff868116845260229092528483208054640100000000810462ffffff1690925263ffffffff8216855285517ffeaf968c00000000000000000000000000000000000000000000000000000000815295519495909484936701000000000000009093049092169163feaf968c9160048082019260a0929091908290030181865afa15801562003d8f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062003db591906200534e565b5093505092505060008213158062003dcc57508042105b8062003e0057506000866040015162ffffff1611801562003e00575062003df4814262004a0f565b866040015162ffffff16105b1562003e16576001830154604085015262003e1e565b604084018290525b50505092915050565b6040805160808101825260008082526020820181905291810182905260608101919091526000836060015161ffff16836060015162003e679190620053bf565b90508260e00151801562003e7a5750803a105b1562003e8357503a5b60008360a0015184604001518560200151866000015162003ea5919062004b51565b62003eb19085620053bf565b62003ebd919062004b51565b62003ec99190620053bf565b905062003eeb8460c00151604001518262003ee59190620058d9565b6200412f565b6bffffffffffffffffffffffff168352608084015162003f119062003ee59083620058d9565b6bffffffffffffffffffffffff16604084015260c08401516020015160009062003f4a9062ffffff16683635c9adc5dea00000620053bf565b9050600081633b9aca008760a001518860c001516000015163ffffffff1689604001518a600001518962003f7f9190620053bf565b62003f8b919062004b51565b62003f979190620053bf565b62003fa39190620053bf565b62003faf9190620058d9565b62003fbb919062004b51565b905062003fd78660c00151604001518262003ee59190620058d9565b6bffffffffffffffffffffffff1660208601526080860151620040009062003ee59083620058d9565b6bffffffffffffffffffffffff1660608601525050505092915050565b60008181526001830160205260408120548015620041165760006200404460018362004a0f565b85549091506000906200405a9060019062004a0f565b9050818114620040c65760008660000182815481106200407e576200407e62004acd565b9060005260206000200154905080876000018481548110620040a457620040a462004acd565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620040da57620040da62005915565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062003266565b600091505062003266565b6000620032638383620041d3565b60006bffffffffffffffffffffffff821115620041cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201527f3620626974730000000000000000000000000000000000000000000000000000606482015260840162000e41565b5090565b60008181526001830160205260408120546200421c5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562003266565b50600062003266565b508054620042339062004afc565b6000825580601f1062004244575050565b601f016020900490600052602060002090810190620028bb919062004272565b6103ca806200594583390190565b5b80821115620041cf576000815560010162004273565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715620042df57620042df62004289565b60405290565b604051610100810167ffffffffffffffff81118282101715620042df57620042df62004289565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171562004356576200435662004289565b604052919050565b600067ffffffffffffffff8211156200437b576200437b62004289565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112620043b957600080fd5b8135620043d0620043ca826200435e565b6200430c565b818152846020838601011115620043e657600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156200441757600080fd5b82359150602083013567ffffffffffffffff8111156200443657600080fd5b6200444485828601620043a7565b9150509250929050565b60005b838110156200446b57818101518382015260200162004451565b50506000910152565b600081518084526200448e8160208601602086016200444e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600a811062004527577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b841515815260806020820152600062004548608083018662004474565b9050620045596040830185620044ef565b82606083015295945050505050565b871515815260e0602082015260006200458560e083018962004474565b9050620045966040830188620044ef565b8560608301528460808301528360a08301528260c083015298975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114620028bb57600080fd5b8035620045ea81620045ba565b919050565b6000806000604084860312156200460557600080fd5b833567ffffffffffffffff808211156200461e57600080fd5b818601915086601f8301126200463357600080fd5b8135818111156200464357600080fd5b8760208260051b85010111156200465957600080fd5b602092830195509350508401356200467181620045ba565b809150509250925092565b60008083601f8401126200468f57600080fd5b50813567ffffffffffffffff811115620046a857600080fd5b602083019150836020828501011115620046c157600080fd5b9250929050565b60008060208385031215620046dc57600080fd5b823567ffffffffffffffff811115620046f457600080fd5b62004702858286016200467c565b90969095509350505050565b803563ffffffff81168114620045ea57600080fd5b803560028110620045ea57600080fd5b60008060008060008060008060006101008a8c0312156200475357600080fd5b6200475e8a620045dd565b98506200476e60208b016200470e565b97506200477e60408b01620045dd565b96506200478e60608b0162004723565b95506200479e60808b01620045dd565b945060a08a013567ffffffffffffffff80821115620047bc57600080fd5b620047ca8d838e016200467c565b909650945060c08c0135915080821115620047e457600080fd5b620047f28d838e01620043a7565b935060e08c01359150808211156200480957600080fd5b50620048188c828d01620043a7565b9150509295985092959850929598565b6000602082840312156200483b57600080fd5b5035919050565b600067ffffffffffffffff8211156200485f576200485f62004289565b5060051b60200190565b600082601f8301126200487b57600080fd5b813560206200488e620043ca8362004842565b82815260059290921b84018101918181019086841115620048ae57600080fd5b8286015b84811015620048f357803567ffffffffffffffff811115620048d45760008081fd5b620048e48986838b0101620043a7565b845250918301918301620048b2565b509695505050505050565b600080600080606085870312156200491557600080fd5b84359350602085013567ffffffffffffffff808211156200493557600080fd5b620049438883890162004869565b945060408701359150808211156200495a57600080fd5b5062004969878288016200467c565b95989497509550505050565b6000602082840312156200498857600080fd5b81356200319781620045ba565b8051620045ea81620045ba565b600060208284031215620049b557600080fd5b81516200319781620045ba565b60008251620049d68184602087016200444e565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115620032665762003266620049e0565b8015158114620028bb57600080fd5b600082601f83011262004a4657600080fd5b815162004a57620043ca826200435e565b81815284602083860101111562004a6d57600080fd5b62002a1c8260208301602087016200444e565b6000806040838503121562004a9457600080fd5b825162004aa18162004a25565b602084015190925067ffffffffffffffff81111562004abf57600080fd5b620044448582860162004a34565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c9082168062004b1157607f821691505b60208210810362004b4b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b80820180821115620032665762003266620049e0565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362004b9b5762004b9b620049e0565b5060010190565b600081518084526020808501945080840160005b8381101562004c7b5781518051151588528381015115158489015260408082015163ffffffff908116918a01919091526060808301518216908a015260808083015173ffffffffffffffffffffffffffffffffffffffff908116918b019190915260a0808401516fffffffffffffffffffffffffffffffff16908b015260c0808401516bffffffffffffffffffffffff16908b015260e080840151909216918a0191909152610100918201511690880152610120909601959082019060010162004bb6565b509495945050505050565b600081518084526020808501945080840160005b8381101562004c7b57815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010162004c9a565b600081518084526020808501808196508360051b8101915082860160005b8581101562004d1a57828403895262004d0784835162004474565b9885019893509084019060010162004cec565b5091979650505050505050565b60e081528760e082015260006101007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8a111562004d6457600080fd5b8960051b808c8386013783018381038201602085015262004d888282018b62004ba2565b915050828103604084015262004d9f818962004c86565b9050828103606084015262004db5818862004c86565b9050828103608084015262004dcb818762004cce565b905082810360a084015262004de1818662004cce565b905082810360c084015262004df7818562004cce565b9b9a5050505050505050505050565b60006020828403121562004e1957600080fd5b815160ff811681146200319757600080fd5b60ff8416815260ff8316602082015260606040820152600062003a4f606083018462004474565b60006020828403121562004e6557600080fd5b815167ffffffffffffffff81111562004e7d57600080fd5b62002a1c8482850162004a34565b60208152600062003263602083018462004474565b60006020828403121562004eb357600080fd5b8151620031978162004a25565b600082601f83011262004ed257600080fd5b8135602062004ee5620043ca8362004842565b82815260059290921b8401810191818101908684111562004f0557600080fd5b8286015b84811015620048f3578035835291830191830162004f09565b8035620045ea8162004a25565b80356fffffffffffffffffffffffffffffffff81168114620045ea57600080fd5b80356bffffffffffffffffffffffff81168114620045ea57600080fd5b600082601f83011262004f7f57600080fd5b8135602062004f92620043ca8362004842565b828152610120928302850182019282820191908785111562004fb357600080fd5b8387015b858110156200508d5781818a03121562004fd15760008081fd5b62004fdb620042b8565b62004fe68262004f22565b815262004ff586830162004f22565b868201526040620050088184016200470e565b9082015260606200501b8382016200470e565b9082015260806200502e838201620045dd565b9082015260a06200504183820162004f2f565b9082015260c06200505483820162004f50565b9082015260e0620050678382016200470e565b908201526101006200507b838201620045dd565b90820152845292840192810162004fb7565b5090979650505050505050565b600082601f830112620050ac57600080fd5b81356020620050bf620043ca8362004842565b82815260059290921b84018101918181019086841115620050df57600080fd5b8286015b84811015620048f3578035620050f981620045ba565b8352918301918301620050e3565b600080600080600080600060e0888a0312156200512357600080fd5b873567ffffffffffffffff808211156200513c57600080fd5b6200514a8b838c0162004ec0565b985060208a01359150808211156200516157600080fd5b6200516f8b838c0162004f6d565b975060408a01359150808211156200518657600080fd5b620051948b838c016200509a565b965060608a0135915080821115620051ab57600080fd5b620051b98b838c016200509a565b955060808a0135915080821115620051d057600080fd5b620051de8b838c0162004869565b945060a08a0135915080821115620051f557600080fd5b620052038b838c0162004869565b935060c08a01359150808211156200521a57600080fd5b50620052298a828b0162004869565b91505092959891949750929550565b600063ffffffff808316818103620052545762005254620049e0565b6001019392505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60208152600062002a1c6020830184866200525e565b600060208284031215620052d057600080fd5b5051919050565b6bffffffffffffffffffffffff82811682821603908082111562003cc55762003cc5620049e0565b60408152600062005314604083018662004cce565b8281036020840152620053298185876200525e565b9695505050505050565b805169ffffffffffffffffffff81168114620045ea57600080fd5b600080600080600060a086880312156200536757600080fd5b620053728662005333565b9450602086015193506040860151925060608601519150620053976080870162005333565b90509295509295909350565b60ff8181168382160190811115620032665762003266620049e0565b8082028115828204841417620032665762003266620049e0565b60008060408385031215620053ed57600080fd5b505080516020909101519092909150565b6bffffffffffffffffffffffff81811683821601908082111562003cc55762003cc5620049e0565b60008154620054358162004afc565b8085526020600183811680156200545557600181146200548e57620054be565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008516838901528284151560051b8901019550620054be565b866000528260002060005b85811015620054b65781548a820186015290830190840162005499565b890184019650505b505050505092915050565b60208152600062003263602083018462005426565b600082601f830112620054f057600080fd5b8151602062005503620043ca8362004842565b82815260059290921b840181019181810190868411156200552357600080fd5b8286015b84811015620048f3578051835291830191830162005527565b6000602082840312156200555357600080fd5b815167ffffffffffffffff808211156200556c57600080fd5b9083019061010082860312156200558257600080fd5b6200558c620042e5565b8251815260208301516020820152604083015160408201526060830151606082015260808301516080820152620055c660a0840162004995565b60a082015260c083015182811115620055de57600080fd5b620055ec87828601620054de565b60c08301525060e0830151828111156200560557600080fd5b620056138782860162004a34565b60e08301525095945050505050565b600081518084526020808501945080840160005b8381101562004c7b5781518752958201959082019060010162005636565b60408152825160408201526020830151606082015260408301516080820152606083015160a0820152608083015160c082015273ffffffffffffffffffffffffffffffffffffffff60a08401511660e0820152600060c0840151610100808185015250620056c761014084018262005622565b905060e08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08483030161012085015262005705828262004474565b915050828103602084015262003a4f818562005426565b601f8211156200576a57600081815260208120601f850160051c81016020861015620057455750805b601f850160051c820191505b81811015620057665782815560010162005751565b5050505b505050565b815167ffffffffffffffff8111156200578c576200578c62004289565b620057a4816200579d845462004afc565b846200571c565b602080601f831160018114620057fa5760008415620057c35750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855562005766565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015620058495788860151825594840194600190910190840162005828565b50858210156200588657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b8051602080830151919081101562004b4b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b60008262005910577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfe60c060405234801561001057600080fd5b506040516103ca3803806103ca83398101604081905261002f91610076565b600080546001600160a01b0319166001600160a01b039384161790559181166080521660a0526100b9565b80516001600160a01b038116811461007157600080fd5b919050565b60008060006060848603121561008b57600080fd5b6100948461005a565b92506100a26020850161005a565b91506100b06040850161005a565b90509250925092565b60805160a0516102e76100e36000396000603801526000818160c4015261011701526102e76000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806379188d161461007b578063f00e6a2a146100aa575b7f00000000000000000000000000000000000000000000000000000000000000003660008037600080366000845af43d6000803e808015610076573d6000f35b3d6000fd5b61008e6100893660046101c1565b6100ee565b6040805192151583526020830191909152015b60405180910390f35b60405173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526020016100a1565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461011557600080fd5b7f00000000000000000000000000000000000000000000000000000000000000005a91505a61138881101561014957600080fd5b61138881039050856040820482031161016157600080fd5b50803b61016d57600080fd5b6000808551602087016000858af192505a610188908361029a565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156101d457600080fd5b82359150602083013567ffffffffffffffff808211156101f357600080fd5b818501915085601f83011261020757600080fd5b81358181111561021957610219610192565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561025f5761025f610192565b8160405282815288602084870101111561027857600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b818103818111156102d4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea164736f6c6343000813000aa164736f6c6343000813000a",
}

var AutomationRegistryLogicAABI = AutomationRegistryLogicAMetaData.ABI

var AutomationRegistryLogicABin = AutomationRegistryLogicAMetaData.Bin

func DeployAutomationRegistryLogicA(auth *bind.TransactOpts, backend bind.ContractBackend, logicB common.Address) (common.Address, *types.Transaction, *AutomationRegistryLogicA, error) {
	parsed, err := AutomationRegistryLogicAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AutomationRegistryLogicABin), backend, logicB)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AutomationRegistryLogicA{address: address, abi: *parsed, AutomationRegistryLogicACaller: AutomationRegistryLogicACaller{contract: contract}, AutomationRegistryLogicATransactor: AutomationRegistryLogicATransactor{contract: contract}, AutomationRegistryLogicAFilterer: AutomationRegistryLogicAFilterer{contract: contract}}, nil
}

type AutomationRegistryLogicA struct {
	address common.Address
	abi     abi.ABI
	AutomationRegistryLogicACaller
	AutomationRegistryLogicATransactor
	AutomationRegistryLogicAFilterer
}

type AutomationRegistryLogicACaller struct {
	contract *bind.BoundContract
}

type AutomationRegistryLogicATransactor struct {
	contract *bind.BoundContract
}

type AutomationRegistryLogicAFilterer struct {
	contract *bind.BoundContract
}

type AutomationRegistryLogicASession struct {
	Contract     *AutomationRegistryLogicA
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AutomationRegistryLogicACallerSession struct {
	Contract *AutomationRegistryLogicACaller
	CallOpts bind.CallOpts
}

type AutomationRegistryLogicATransactorSession struct {
	Contract     *AutomationRegistryLogicATransactor
	TransactOpts bind.TransactOpts
}

type AutomationRegistryLogicARaw struct {
	Contract *AutomationRegistryLogicA
}

type AutomationRegistryLogicACallerRaw struct {
	Contract *AutomationRegistryLogicACaller
}

type AutomationRegistryLogicATransactorRaw struct {
	Contract *AutomationRegistryLogicATransactor
}

func NewAutomationRegistryLogicA(address common.Address, backend bind.ContractBackend) (*AutomationRegistryLogicA, error) {
	abi, err := abi.JSON(strings.NewReader(AutomationRegistryLogicAABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAutomationRegistryLogicA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicA{address: address, abi: abi, AutomationRegistryLogicACaller: AutomationRegistryLogicACaller{contract: contract}, AutomationRegistryLogicATransactor: AutomationRegistryLogicATransactor{contract: contract}, AutomationRegistryLogicAFilterer: AutomationRegistryLogicAFilterer{contract: contract}}, nil
}

func NewAutomationRegistryLogicACaller(address common.Address, caller bind.ContractCaller) (*AutomationRegistryLogicACaller, error) {
	contract, err := bindAutomationRegistryLogicA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicACaller{contract: contract}, nil
}

func NewAutomationRegistryLogicATransactor(address common.Address, transactor bind.ContractTransactor) (*AutomationRegistryLogicATransactor, error) {
	contract, err := bindAutomationRegistryLogicA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicATransactor{contract: contract}, nil
}

func NewAutomationRegistryLogicAFilterer(address common.Address, filterer bind.ContractFilterer) (*AutomationRegistryLogicAFilterer, error) {
	contract, err := bindAutomationRegistryLogicA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAFilterer{contract: contract}, nil
}

func bindAutomationRegistryLogicA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AutomationRegistryLogicAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationRegistryLogicA.Contract.AutomationRegistryLogicACaller.contract.Call(opts, result, method, params...)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.AutomationRegistryLogicATransactor.contract.Transfer(opts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.AutomationRegistryLogicATransactor.contract.Transact(opts, method, params...)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationRegistryLogicA.Contract.contract.Call(opts, result, method, params...)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.contract.Transfer(opts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.contract.Transact(opts, method, params...)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicACaller) FallbackTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomationRegistryLogicA.contract.Call(opts, &out, "fallbackTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) FallbackTo() (common.Address, error) {
	return _AutomationRegistryLogicA.Contract.FallbackTo(&_AutomationRegistryLogicA.CallOpts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicACallerSession) FallbackTo() (common.Address, error) {
	return _AutomationRegistryLogicA.Contract.FallbackTo(&_AutomationRegistryLogicA.CallOpts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicACaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomationRegistryLogicA.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) Owner() (common.Address, error) {
	return _AutomationRegistryLogicA.Contract.Owner(&_AutomationRegistryLogicA.CallOpts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicACallerSession) Owner() (common.Address, error) {
	return _AutomationRegistryLogicA.Contract.Owner(&_AutomationRegistryLogicA.CallOpts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "acceptOwnership")
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) AcceptOwnership() (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.AcceptOwnership(&_AutomationRegistryLogicA.TransactOpts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.AcceptOwnership(&_AutomationRegistryLogicA.TransactOpts)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) CancelUpkeep(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "cancelUpkeep", id)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) CancelUpkeep(id *big.Int) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CancelUpkeep(&_AutomationRegistryLogicA.TransactOpts, id)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) CancelUpkeep(id *big.Int) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CancelUpkeep(&_AutomationRegistryLogicA.TransactOpts, id)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) CheckCallback(opts *bind.TransactOpts, id *big.Int, values [][]byte, extraData []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "checkCallback", id, values, extraData)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) CheckCallback(id *big.Int, values [][]byte, extraData []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CheckCallback(&_AutomationRegistryLogicA.TransactOpts, id, values, extraData)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) CheckCallback(id *big.Int, values [][]byte, extraData []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CheckCallback(&_AutomationRegistryLogicA.TransactOpts, id, values, extraData)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) CheckUpkeep(opts *bind.TransactOpts, id *big.Int, triggerData []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "checkUpkeep", id, triggerData)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) CheckUpkeep(id *big.Int, triggerData []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CheckUpkeep(&_AutomationRegistryLogicA.TransactOpts, id, triggerData)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) CheckUpkeep(id *big.Int, triggerData []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CheckUpkeep(&_AutomationRegistryLogicA.TransactOpts, id, triggerData)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) CheckUpkeep0(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "checkUpkeep0", id)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) CheckUpkeep0(id *big.Int) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CheckUpkeep0(&_AutomationRegistryLogicA.TransactOpts, id)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) CheckUpkeep0(id *big.Int) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.CheckUpkeep0(&_AutomationRegistryLogicA.TransactOpts, id)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) ExecuteCallback(opts *bind.TransactOpts, id *big.Int, payload []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "executeCallback", id, payload)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) ExecuteCallback(id *big.Int, payload []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.ExecuteCallback(&_AutomationRegistryLogicA.TransactOpts, id, payload)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) ExecuteCallback(id *big.Int, payload []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.ExecuteCallback(&_AutomationRegistryLogicA.TransactOpts, id, payload)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) MigrateUpkeeps(opts *bind.TransactOpts, ids []*big.Int, destination common.Address) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "migrateUpkeeps", ids, destination)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) MigrateUpkeeps(ids []*big.Int, destination common.Address) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.MigrateUpkeeps(&_AutomationRegistryLogicA.TransactOpts, ids, destination)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) MigrateUpkeeps(ids []*big.Int, destination common.Address) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.MigrateUpkeeps(&_AutomationRegistryLogicA.TransactOpts, ids, destination)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) ReceiveUpkeeps(opts *bind.TransactOpts, encodedUpkeeps []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "receiveUpkeeps", encodedUpkeeps)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) ReceiveUpkeeps(encodedUpkeeps []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.ReceiveUpkeeps(&_AutomationRegistryLogicA.TransactOpts, encodedUpkeeps)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) ReceiveUpkeeps(encodedUpkeeps []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.ReceiveUpkeeps(&_AutomationRegistryLogicA.TransactOpts, encodedUpkeeps)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) RegisterUpkeep(opts *bind.TransactOpts, target common.Address, gasLimit uint32, admin common.Address, triggerType uint8, billingToken common.Address, checkData []byte, triggerConfig []byte, offchainConfig []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "registerUpkeep", target, gasLimit, admin, triggerType, billingToken, checkData, triggerConfig, offchainConfig)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) RegisterUpkeep(target common.Address, gasLimit uint32, admin common.Address, triggerType uint8, billingToken common.Address, checkData []byte, triggerConfig []byte, offchainConfig []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.RegisterUpkeep(&_AutomationRegistryLogicA.TransactOpts, target, gasLimit, admin, triggerType, billingToken, checkData, triggerConfig, offchainConfig)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) RegisterUpkeep(target common.Address, gasLimit uint32, admin common.Address, triggerType uint8, billingToken common.Address, checkData []byte, triggerConfig []byte, offchainConfig []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.RegisterUpkeep(&_AutomationRegistryLogicA.TransactOpts, target, gasLimit, admin, triggerType, billingToken, checkData, triggerConfig, offchainConfig)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.Transact(opts, "transferOwnership", to)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.TransferOwnership(&_AutomationRegistryLogicA.TransactOpts, to)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.TransferOwnership(&_AutomationRegistryLogicA.TransactOpts, to)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.contract.RawTransact(opts, calldata)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicASession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.Fallback(&_AutomationRegistryLogicA.TransactOpts, calldata)
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicATransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AutomationRegistryLogicA.Contract.Fallback(&_AutomationRegistryLogicA.TransactOpts, calldata)
}

type AutomationRegistryLogicAAdminPrivilegeConfigSetIterator struct {
	Event *AutomationRegistryLogicAAdminPrivilegeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAAdminPrivilegeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAAdminPrivilegeConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAAdminPrivilegeConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAAdminPrivilegeConfigSetIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAAdminPrivilegeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAAdminPrivilegeConfigSet struct {
	Admin           common.Address
	PrivilegeConfig []byte
	Raw             types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterAdminPrivilegeConfigSet(opts *bind.FilterOpts, admin []common.Address) (*AutomationRegistryLogicAAdminPrivilegeConfigSetIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "AdminPrivilegeConfigSet", adminRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAAdminPrivilegeConfigSetIterator{contract: _AutomationRegistryLogicA.contract, event: "AdminPrivilegeConfigSet", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchAdminPrivilegeConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAAdminPrivilegeConfigSet, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "AdminPrivilegeConfigSet", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAAdminPrivilegeConfigSet)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "AdminPrivilegeConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseAdminPrivilegeConfigSet(log types.Log) (*AutomationRegistryLogicAAdminPrivilegeConfigSet, error) {
	event := new(AutomationRegistryLogicAAdminPrivilegeConfigSet)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "AdminPrivilegeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicABillingConfigOverriddenIterator struct {
	Event *AutomationRegistryLogicABillingConfigOverridden

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicABillingConfigOverriddenIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicABillingConfigOverridden)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicABillingConfigOverridden)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicABillingConfigOverriddenIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicABillingConfigOverriddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicABillingConfigOverridden struct {
	Id  *big.Int
	Raw types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterBillingConfigOverridden(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicABillingConfigOverriddenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "BillingConfigOverridden", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicABillingConfigOverriddenIterator{contract: _AutomationRegistryLogicA.contract, event: "BillingConfigOverridden", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchBillingConfigOverridden(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicABillingConfigOverridden, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "BillingConfigOverridden", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicABillingConfigOverridden)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "BillingConfigOverridden", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseBillingConfigOverridden(log types.Log) (*AutomationRegistryLogicABillingConfigOverridden, error) {
	event := new(AutomationRegistryLogicABillingConfigOverridden)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "BillingConfigOverridden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicABillingConfigOverrideRemovedIterator struct {
	Event *AutomationRegistryLogicABillingConfigOverrideRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicABillingConfigOverrideRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicABillingConfigOverrideRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicABillingConfigOverrideRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicABillingConfigOverrideRemovedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicABillingConfigOverrideRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicABillingConfigOverrideRemoved struct {
	Id  *big.Int
	Raw types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterBillingConfigOverrideRemoved(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicABillingConfigOverrideRemovedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "BillingConfigOverrideRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicABillingConfigOverrideRemovedIterator{contract: _AutomationRegistryLogicA.contract, event: "BillingConfigOverrideRemoved", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchBillingConfigOverrideRemoved(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicABillingConfigOverrideRemoved, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "BillingConfigOverrideRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicABillingConfigOverrideRemoved)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "BillingConfigOverrideRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseBillingConfigOverrideRemoved(log types.Log) (*AutomationRegistryLogicABillingConfigOverrideRemoved, error) {
	event := new(AutomationRegistryLogicABillingConfigOverrideRemoved)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "BillingConfigOverrideRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicABillingConfigSetIterator struct {
	Event *AutomationRegistryLogicABillingConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicABillingConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicABillingConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicABillingConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicABillingConfigSetIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicABillingConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicABillingConfigSet struct {
	Token  common.Address
	Config AutomationRegistryBase23BillingConfig
	Raw    types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterBillingConfigSet(opts *bind.FilterOpts, token []common.Address) (*AutomationRegistryLogicABillingConfigSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "BillingConfigSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicABillingConfigSetIterator{contract: _AutomationRegistryLogicA.contract, event: "BillingConfigSet", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchBillingConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicABillingConfigSet, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "BillingConfigSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicABillingConfigSet)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "BillingConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseBillingConfigSet(log types.Log) (*AutomationRegistryLogicABillingConfigSet, error) {
	event := new(AutomationRegistryLogicABillingConfigSet)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "BillingConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicACancelledUpkeepReportIterator struct {
	Event *AutomationRegistryLogicACancelledUpkeepReport

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicACancelledUpkeepReportIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicACancelledUpkeepReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicACancelledUpkeepReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicACancelledUpkeepReportIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicACancelledUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicACancelledUpkeepReport struct {
	Id      *big.Int
	Trigger []byte
	Raw     types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterCancelledUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicACancelledUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "CancelledUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicACancelledUpkeepReportIterator{contract: _AutomationRegistryLogicA.contract, event: "CancelledUpkeepReport", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchCancelledUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicACancelledUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "CancelledUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicACancelledUpkeepReport)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "CancelledUpkeepReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseCancelledUpkeepReport(log types.Log) (*AutomationRegistryLogicACancelledUpkeepReport, error) {
	event := new(AutomationRegistryLogicACancelledUpkeepReport)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "CancelledUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAChainSpecificModuleUpdatedIterator struct {
	Event *AutomationRegistryLogicAChainSpecificModuleUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAChainSpecificModuleUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAChainSpecificModuleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAChainSpecificModuleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAChainSpecificModuleUpdatedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAChainSpecificModuleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAChainSpecificModuleUpdated struct {
	NewModule common.Address
	Raw       types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterChainSpecificModuleUpdated(opts *bind.FilterOpts) (*AutomationRegistryLogicAChainSpecificModuleUpdatedIterator, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "ChainSpecificModuleUpdated")
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAChainSpecificModuleUpdatedIterator{contract: _AutomationRegistryLogicA.contract, event: "ChainSpecificModuleUpdated", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchChainSpecificModuleUpdated(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAChainSpecificModuleUpdated) (event.Subscription, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "ChainSpecificModuleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAChainSpecificModuleUpdated)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "ChainSpecificModuleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseChainSpecificModuleUpdated(log types.Log) (*AutomationRegistryLogicAChainSpecificModuleUpdated, error) {
	event := new(AutomationRegistryLogicAChainSpecificModuleUpdated)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "ChainSpecificModuleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicADedupKeyAddedIterator struct {
	Event *AutomationRegistryLogicADedupKeyAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicADedupKeyAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicADedupKeyAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicADedupKeyAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicADedupKeyAddedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicADedupKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicADedupKeyAdded struct {
	DedupKey [32]byte
	Raw      types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterDedupKeyAdded(opts *bind.FilterOpts, dedupKey [][32]byte) (*AutomationRegistryLogicADedupKeyAddedIterator, error) {

	var dedupKeyRule []interface{}
	for _, dedupKeyItem := range dedupKey {
		dedupKeyRule = append(dedupKeyRule, dedupKeyItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "DedupKeyAdded", dedupKeyRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicADedupKeyAddedIterator{contract: _AutomationRegistryLogicA.contract, event: "DedupKeyAdded", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchDedupKeyAdded(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicADedupKeyAdded, dedupKey [][32]byte) (event.Subscription, error) {

	var dedupKeyRule []interface{}
	for _, dedupKeyItem := range dedupKey {
		dedupKeyRule = append(dedupKeyRule, dedupKeyItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "DedupKeyAdded", dedupKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicADedupKeyAdded)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "DedupKeyAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseDedupKeyAdded(log types.Log) (*AutomationRegistryLogicADedupKeyAdded, error) {
	event := new(AutomationRegistryLogicADedupKeyAdded)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "DedupKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAFeesWithdrawnIterator struct {
	Event *AutomationRegistryLogicAFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAFeesWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAFeesWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAFeesWithdrawn struct {
	Recipient    common.Address
	AssetAddress common.Address
	Amount       *big.Int
	Raw          types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts, recipient []common.Address, assetAddress []common.Address) (*AutomationRegistryLogicAFeesWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "FeesWithdrawn", recipientRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAFeesWithdrawnIterator{contract: _AutomationRegistryLogicA.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAFeesWithdrawn, recipient []common.Address, assetAddress []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var assetAddressRule []interface{}
	for _, assetAddressItem := range assetAddress {
		assetAddressRule = append(assetAddressRule, assetAddressItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "FeesWithdrawn", recipientRule, assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAFeesWithdrawn)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseFeesWithdrawn(log types.Log) (*AutomationRegistryLogicAFeesWithdrawn, error) {
	event := new(AutomationRegistryLogicAFeesWithdrawn)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAFundsAddedIterator struct {
	Event *AutomationRegistryLogicAFundsAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAFundsAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAFundsAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAFundsAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAFundsAddedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAFundsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAFundsAdded struct {
	Id     *big.Int
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterFundsAdded(opts *bind.FilterOpts, id []*big.Int, from []common.Address) (*AutomationRegistryLogicAFundsAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "FundsAdded", idRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAFundsAddedIterator{contract: _AutomationRegistryLogicA.contract, event: "FundsAdded", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAFundsAdded, id []*big.Int, from []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "FundsAdded", idRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAFundsAdded)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "FundsAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseFundsAdded(log types.Log) (*AutomationRegistryLogicAFundsAdded, error) {
	event := new(AutomationRegistryLogicAFundsAdded)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "FundsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAFundsWithdrawnIterator struct {
	Event *AutomationRegistryLogicAFundsWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAFundsWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAFundsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAFundsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAFundsWithdrawnIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAFundsWithdrawn struct {
	Id     *big.Int
	Amount *big.Int
	To     common.Address
	Raw    types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAFundsWithdrawnIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "FundsWithdrawn", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAFundsWithdrawnIterator{contract: _AutomationRegistryLogicA.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAFundsWithdrawn, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "FundsWithdrawn", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAFundsWithdrawn)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseFundsWithdrawn(log types.Log) (*AutomationRegistryLogicAFundsWithdrawn, error) {
	event := new(AutomationRegistryLogicAFundsWithdrawn)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAInsufficientFundsUpkeepReportIterator struct {
	Event *AutomationRegistryLogicAInsufficientFundsUpkeepReport

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAInsufficientFundsUpkeepReportIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAInsufficientFundsUpkeepReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAInsufficientFundsUpkeepReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAInsufficientFundsUpkeepReportIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAInsufficientFundsUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAInsufficientFundsUpkeepReport struct {
	Id      *big.Int
	Trigger []byte
	Raw     types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterInsufficientFundsUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAInsufficientFundsUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "InsufficientFundsUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAInsufficientFundsUpkeepReportIterator{contract: _AutomationRegistryLogicA.contract, event: "InsufficientFundsUpkeepReport", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchInsufficientFundsUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAInsufficientFundsUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "InsufficientFundsUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAInsufficientFundsUpkeepReport)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "InsufficientFundsUpkeepReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseInsufficientFundsUpkeepReport(log types.Log) (*AutomationRegistryLogicAInsufficientFundsUpkeepReport, error) {
	event := new(AutomationRegistryLogicAInsufficientFundsUpkeepReport)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "InsufficientFundsUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicANOPsSettledOffchainIterator struct {
	Event *AutomationRegistryLogicANOPsSettledOffchain

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicANOPsSettledOffchainIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicANOPsSettledOffchain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicANOPsSettledOffchain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicANOPsSettledOffchainIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicANOPsSettledOffchainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicANOPsSettledOffchain struct {
	Payees   []common.Address
	Payments []*big.Int
	Raw      types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterNOPsSettledOffchain(opts *bind.FilterOpts) (*AutomationRegistryLogicANOPsSettledOffchainIterator, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "NOPsSettledOffchain")
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicANOPsSettledOffchainIterator{contract: _AutomationRegistryLogicA.contract, event: "NOPsSettledOffchain", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchNOPsSettledOffchain(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicANOPsSettledOffchain) (event.Subscription, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "NOPsSettledOffchain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicANOPsSettledOffchain)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "NOPsSettledOffchain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseNOPsSettledOffchain(log types.Log) (*AutomationRegistryLogicANOPsSettledOffchain, error) {
	event := new(AutomationRegistryLogicANOPsSettledOffchain)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "NOPsSettledOffchain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAOwnershipTransferRequestedIterator struct {
	Event *AutomationRegistryLogicAOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationRegistryLogicAOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAOwnershipTransferRequestedIterator{contract: _AutomationRegistryLogicA.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAOwnershipTransferRequested)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseOwnershipTransferRequested(log types.Log) (*AutomationRegistryLogicAOwnershipTransferRequested, error) {
	event := new(AutomationRegistryLogicAOwnershipTransferRequested)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAOwnershipTransferredIterator struct {
	Event *AutomationRegistryLogicAOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationRegistryLogicAOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAOwnershipTransferredIterator{contract: _AutomationRegistryLogicA.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAOwnershipTransferred)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseOwnershipTransferred(log types.Log) (*AutomationRegistryLogicAOwnershipTransferred, error) {
	event := new(AutomationRegistryLogicAOwnershipTransferred)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAPausedIterator struct {
	Event *AutomationRegistryLogicAPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAPausedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterPaused(opts *bind.FilterOpts) (*AutomationRegistryLogicAPausedIterator, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAPausedIterator{contract: _AutomationRegistryLogicA.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPaused) (event.Subscription, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAPaused)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParsePaused(log types.Log) (*AutomationRegistryLogicAPaused, error) {
	event := new(AutomationRegistryLogicAPaused)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAPayeesUpdatedIterator struct {
	Event *AutomationRegistryLogicAPayeesUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAPayeesUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAPayeesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAPayeesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAPayeesUpdatedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAPayeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAPayeesUpdated struct {
	Transmitters []common.Address
	Payees       []common.Address
	Raw          types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterPayeesUpdated(opts *bind.FilterOpts) (*AutomationRegistryLogicAPayeesUpdatedIterator, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "PayeesUpdated")
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAPayeesUpdatedIterator{contract: _AutomationRegistryLogicA.contract, event: "PayeesUpdated", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchPayeesUpdated(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPayeesUpdated) (event.Subscription, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "PayeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAPayeesUpdated)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PayeesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParsePayeesUpdated(log types.Log) (*AutomationRegistryLogicAPayeesUpdated, error) {
	event := new(AutomationRegistryLogicAPayeesUpdated)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PayeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAPayeeshipTransferRequestedIterator struct {
	Event *AutomationRegistryLogicAPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAPayeeshipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAPayeeshipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAPayeeshipTransferRequested struct {
	Transmitter common.Address
	From        common.Address
	To          common.Address
	Raw         types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, from []common.Address, to []common.Address) (*AutomationRegistryLogicAPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAPayeeshipTransferRequestedIterator{contract: _AutomationRegistryLogicA.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPayeeshipTransferRequested, transmitter []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAPayeeshipTransferRequested)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParsePayeeshipTransferRequested(log types.Log) (*AutomationRegistryLogicAPayeeshipTransferRequested, error) {
	event := new(AutomationRegistryLogicAPayeeshipTransferRequested)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAPayeeshipTransferredIterator struct {
	Event *AutomationRegistryLogicAPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAPayeeshipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAPayeeshipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAPayeeshipTransferred struct {
	Transmitter common.Address
	From        common.Address
	To          common.Address
	Raw         types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, from []common.Address, to []common.Address) (*AutomationRegistryLogicAPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAPayeeshipTransferredIterator{contract: _AutomationRegistryLogicA.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPayeeshipTransferred, transmitter []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAPayeeshipTransferred)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParsePayeeshipTransferred(log types.Log) (*AutomationRegistryLogicAPayeeshipTransferred, error) {
	event := new(AutomationRegistryLogicAPayeeshipTransferred)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAPaymentWithdrawnIterator struct {
	Event *AutomationRegistryLogicAPaymentWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAPaymentWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAPaymentWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAPaymentWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAPaymentWithdrawnIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAPaymentWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAPaymentWithdrawn struct {
	Transmitter common.Address
	Amount      *big.Int
	To          common.Address
	Payee       common.Address
	Raw         types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterPaymentWithdrawn(opts *bind.FilterOpts, transmitter []common.Address, amount []*big.Int, to []common.Address) (*AutomationRegistryLogicAPaymentWithdrawnIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "PaymentWithdrawn", transmitterRule, amountRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAPaymentWithdrawnIterator{contract: _AutomationRegistryLogicA.contract, event: "PaymentWithdrawn", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchPaymentWithdrawn(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPaymentWithdrawn, transmitter []common.Address, amount []*big.Int, to []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "PaymentWithdrawn", transmitterRule, amountRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAPaymentWithdrawn)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PaymentWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParsePaymentWithdrawn(log types.Log) (*AutomationRegistryLogicAPaymentWithdrawn, error) {
	event := new(AutomationRegistryLogicAPaymentWithdrawn)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "PaymentWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAReorgedUpkeepReportIterator struct {
	Event *AutomationRegistryLogicAReorgedUpkeepReport

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAReorgedUpkeepReportIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAReorgedUpkeepReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAReorgedUpkeepReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAReorgedUpkeepReportIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAReorgedUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAReorgedUpkeepReport struct {
	Id      *big.Int
	Trigger []byte
	Raw     types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterReorgedUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAReorgedUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "ReorgedUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAReorgedUpkeepReportIterator{contract: _AutomationRegistryLogicA.contract, event: "ReorgedUpkeepReport", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchReorgedUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAReorgedUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "ReorgedUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAReorgedUpkeepReport)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "ReorgedUpkeepReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseReorgedUpkeepReport(log types.Log) (*AutomationRegistryLogicAReorgedUpkeepReport, error) {
	event := new(AutomationRegistryLogicAReorgedUpkeepReport)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "ReorgedUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAStaleUpkeepReportIterator struct {
	Event *AutomationRegistryLogicAStaleUpkeepReport

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAStaleUpkeepReportIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAStaleUpkeepReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAStaleUpkeepReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAStaleUpkeepReportIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAStaleUpkeepReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAStaleUpkeepReport struct {
	Id      *big.Int
	Trigger []byte
	Raw     types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterStaleUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAStaleUpkeepReportIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "StaleUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAStaleUpkeepReportIterator{contract: _AutomationRegistryLogicA.contract, event: "StaleUpkeepReport", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchStaleUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAStaleUpkeepReport, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "StaleUpkeepReport", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAStaleUpkeepReport)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "StaleUpkeepReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseStaleUpkeepReport(log types.Log) (*AutomationRegistryLogicAStaleUpkeepReport, error) {
	event := new(AutomationRegistryLogicAStaleUpkeepReport)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "StaleUpkeepReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUnpausedIterator struct {
	Event *AutomationRegistryLogicAUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUnpausedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AutomationRegistryLogicAUnpausedIterator, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUnpausedIterator{contract: _AutomationRegistryLogicA.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUnpaused) (event.Subscription, error) {

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUnpaused)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUnpaused(log types.Log) (*AutomationRegistryLogicAUnpaused, error) {
	event := new(AutomationRegistryLogicAUnpaused)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepAdminTransferRequestedIterator struct {
	Event *AutomationRegistryLogicAUpkeepAdminTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepAdminTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepAdminTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepAdminTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepAdminTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepAdminTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepAdminTransferRequested struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepAdminTransferRequested(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*AutomationRegistryLogicAUpkeepAdminTransferRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepAdminTransferRequested", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepAdminTransferRequestedIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepAdminTransferRequested", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepAdminTransferRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepAdminTransferRequested, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepAdminTransferRequested", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepAdminTransferRequested)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepAdminTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepAdminTransferRequested(log types.Log) (*AutomationRegistryLogicAUpkeepAdminTransferRequested, error) {
	event := new(AutomationRegistryLogicAUpkeepAdminTransferRequested)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepAdminTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepAdminTransferredIterator struct {
	Event *AutomationRegistryLogicAUpkeepAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepAdminTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepAdminTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepAdminTransferred struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepAdminTransferred(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*AutomationRegistryLogicAUpkeepAdminTransferredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepAdminTransferred", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepAdminTransferredIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepAdminTransferred", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepAdminTransferred(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepAdminTransferred, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepAdminTransferred", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepAdminTransferred)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepAdminTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepAdminTransferred(log types.Log) (*AutomationRegistryLogicAUpkeepAdminTransferred, error) {
	event := new(AutomationRegistryLogicAUpkeepAdminTransferred)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepCanceledIterator struct {
	Event *AutomationRegistryLogicAUpkeepCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepCanceledIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepCanceled struct {
	Id            *big.Int
	AtBlockHeight uint64
	Raw           types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepCanceled(opts *bind.FilterOpts, id []*big.Int, atBlockHeight []uint64) (*AutomationRegistryLogicAUpkeepCanceledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var atBlockHeightRule []interface{}
	for _, atBlockHeightItem := range atBlockHeight {
		atBlockHeightRule = append(atBlockHeightRule, atBlockHeightItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepCanceled", idRule, atBlockHeightRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepCanceledIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepCanceled", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepCanceled(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepCanceled, id []*big.Int, atBlockHeight []uint64) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var atBlockHeightRule []interface{}
	for _, atBlockHeightItem := range atBlockHeight {
		atBlockHeightRule = append(atBlockHeightRule, atBlockHeightItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepCanceled", idRule, atBlockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepCanceled)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepCanceled(log types.Log) (*AutomationRegistryLogicAUpkeepCanceled, error) {
	event := new(AutomationRegistryLogicAUpkeepCanceled)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepCheckDataSetIterator struct {
	Event *AutomationRegistryLogicAUpkeepCheckDataSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepCheckDataSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepCheckDataSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepCheckDataSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepCheckDataSetIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepCheckDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepCheckDataSet struct {
	Id           *big.Int
	NewCheckData []byte
	Raw          types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepCheckDataSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepCheckDataSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepCheckDataSet", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepCheckDataSetIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepCheckDataSet", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepCheckDataSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepCheckDataSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepCheckDataSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepCheckDataSet)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepCheckDataSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepCheckDataSet(log types.Log) (*AutomationRegistryLogicAUpkeepCheckDataSet, error) {
	event := new(AutomationRegistryLogicAUpkeepCheckDataSet)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepCheckDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepGasLimitSetIterator struct {
	Event *AutomationRegistryLogicAUpkeepGasLimitSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepGasLimitSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepGasLimitSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepGasLimitSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepGasLimitSetIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepGasLimitSet struct {
	Id       *big.Int
	GasLimit *big.Int
	Raw      types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepGasLimitSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepGasLimitSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepGasLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepGasLimitSetIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepGasLimitSet", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepGasLimitSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepGasLimitSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepGasLimitSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepGasLimitSet)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepGasLimitSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepGasLimitSet(log types.Log) (*AutomationRegistryLogicAUpkeepGasLimitSet, error) {
	event := new(AutomationRegistryLogicAUpkeepGasLimitSet)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepMigratedIterator struct {
	Event *AutomationRegistryLogicAUpkeepMigrated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepMigratedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepMigrated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepMigrated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepMigratedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepMigrated struct {
	Id               *big.Int
	RemainingBalance *big.Int
	Destination      common.Address
	Raw              types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepMigrated(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepMigratedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepMigrated", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepMigratedIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepMigrated", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepMigrated(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepMigrated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepMigrated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepMigrated)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepMigrated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepMigrated(log types.Log) (*AutomationRegistryLogicAUpkeepMigrated, error) {
	event := new(AutomationRegistryLogicAUpkeepMigrated)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepOffchainConfigSetIterator struct {
	Event *AutomationRegistryLogicAUpkeepOffchainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepOffchainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepOffchainConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepOffchainConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepOffchainConfigSetIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepOffchainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepOffchainConfigSet struct {
	Id             *big.Int
	OffchainConfig []byte
	Raw            types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepOffchainConfigSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepOffchainConfigSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepOffchainConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepOffchainConfigSetIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepOffchainConfigSet", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepOffchainConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepOffchainConfigSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepOffchainConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepOffchainConfigSet)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepOffchainConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepOffchainConfigSet(log types.Log) (*AutomationRegistryLogicAUpkeepOffchainConfigSet, error) {
	event := new(AutomationRegistryLogicAUpkeepOffchainConfigSet)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepOffchainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepPausedIterator struct {
	Event *AutomationRegistryLogicAUpkeepPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepPausedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepPaused struct {
	Id  *big.Int
	Raw types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepPaused(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepPausedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepPaused", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepPausedIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepPaused", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepPaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepPaused, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepPaused", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepPaused)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepPaused(log types.Log) (*AutomationRegistryLogicAUpkeepPaused, error) {
	event := new(AutomationRegistryLogicAUpkeepPaused)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepPerformedIterator struct {
	Event *AutomationRegistryLogicAUpkeepPerformed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepPerformedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepPerformedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepPerformed struct {
	Id           *big.Int
	Success      bool
	TotalPayment *big.Int
	GasUsed      *big.Int
	GasOverhead  *big.Int
	Trigger      []byte
	Raw          types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepPerformed(opts *bind.FilterOpts, id []*big.Int, success []bool) (*AutomationRegistryLogicAUpkeepPerformedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepPerformed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepPerformedIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepPerformed", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepPerformed(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepPerformed, id []*big.Int, success []bool) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepPerformed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepPerformed)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepPerformed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepPerformed(log types.Log) (*AutomationRegistryLogicAUpkeepPerformed, error) {
	event := new(AutomationRegistryLogicAUpkeepPerformed)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepPrivilegeConfigSetIterator struct {
	Event *AutomationRegistryLogicAUpkeepPrivilegeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepPrivilegeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepPrivilegeConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepPrivilegeConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepPrivilegeConfigSetIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepPrivilegeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepPrivilegeConfigSet struct {
	Id              *big.Int
	PrivilegeConfig []byte
	Raw             types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepPrivilegeConfigSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepPrivilegeConfigSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepPrivilegeConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepPrivilegeConfigSetIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepPrivilegeConfigSet", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepPrivilegeConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepPrivilegeConfigSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepPrivilegeConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepPrivilegeConfigSet)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepPrivilegeConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepPrivilegeConfigSet(log types.Log) (*AutomationRegistryLogicAUpkeepPrivilegeConfigSet, error) {
	event := new(AutomationRegistryLogicAUpkeepPrivilegeConfigSet)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepPrivilegeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepReceivedIterator struct {
	Event *AutomationRegistryLogicAUpkeepReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepReceivedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepReceived struct {
	Id              *big.Int
	StartingBalance *big.Int
	ImportedFrom    common.Address
	Raw             types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepReceived(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepReceivedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepReceived", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepReceivedIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepReceived", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepReceived(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepReceived, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepReceived", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepReceived)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepReceived(log types.Log) (*AutomationRegistryLogicAUpkeepReceived, error) {
	event := new(AutomationRegistryLogicAUpkeepReceived)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepRegisteredIterator struct {
	Event *AutomationRegistryLogicAUpkeepRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepRegisteredIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepRegistered struct {
	Id         *big.Int
	PerformGas uint32
	Admin      common.Address
	Raw        types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepRegistered(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepRegisteredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepRegisteredIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepRegistered", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepRegistered(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepRegistered, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepRegistered", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepRegistered)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepRegistered(log types.Log) (*AutomationRegistryLogicAUpkeepRegistered, error) {
	event := new(AutomationRegistryLogicAUpkeepRegistered)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepTriggerConfigSetIterator struct {
	Event *AutomationRegistryLogicAUpkeepTriggerConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepTriggerConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepTriggerConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepTriggerConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepTriggerConfigSetIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepTriggerConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepTriggerConfigSet struct {
	Id            *big.Int
	TriggerConfig []byte
	Raw           types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepTriggerConfigSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepTriggerConfigSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepTriggerConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepTriggerConfigSetIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepTriggerConfigSet", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepTriggerConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepTriggerConfigSet, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepTriggerConfigSet", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepTriggerConfigSet)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepTriggerConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepTriggerConfigSet(log types.Log) (*AutomationRegistryLogicAUpkeepTriggerConfigSet, error) {
	event := new(AutomationRegistryLogicAUpkeepTriggerConfigSet)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepTriggerConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AutomationRegistryLogicAUpkeepUnpausedIterator struct {
	Event *AutomationRegistryLogicAUpkeepUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AutomationRegistryLogicAUpkeepUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistryLogicAUpkeepUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistryLogicAUpkeepUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *AutomationRegistryLogicAUpkeepUnpausedIterator) Error() error {
	return it.fail
}

func (it *AutomationRegistryLogicAUpkeepUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AutomationRegistryLogicAUpkeepUnpaused struct {
	Id  *big.Int
	Raw types.Log
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) FilterUpkeepUnpaused(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepUnpausedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.FilterLogs(opts, "UpkeepUnpaused", idRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistryLogicAUpkeepUnpausedIterator{contract: _AutomationRegistryLogicA.contract, event: "UpkeepUnpaused", logs: logs, sub: sub}, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) WatchUpkeepUnpaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepUnpaused, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AutomationRegistryLogicA.contract.WatchLogs(opts, "UpkeepUnpaused", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AutomationRegistryLogicAUpkeepUnpaused)
				if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicAFilterer) ParseUpkeepUnpaused(log types.Log) (*AutomationRegistryLogicAUpkeepUnpaused, error) {
	event := new(AutomationRegistryLogicAUpkeepUnpaused)
	if err := _AutomationRegistryLogicA.contract.UnpackLog(event, "UpkeepUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicA) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _AutomationRegistryLogicA.abi.Events["AdminPrivilegeConfigSet"].ID:
		return _AutomationRegistryLogicA.ParseAdminPrivilegeConfigSet(log)
	case _AutomationRegistryLogicA.abi.Events["BillingConfigOverridden"].ID:
		return _AutomationRegistryLogicA.ParseBillingConfigOverridden(log)
	case _AutomationRegistryLogicA.abi.Events["BillingConfigOverrideRemoved"].ID:
		return _AutomationRegistryLogicA.ParseBillingConfigOverrideRemoved(log)
	case _AutomationRegistryLogicA.abi.Events["BillingConfigSet"].ID:
		return _AutomationRegistryLogicA.ParseBillingConfigSet(log)
	case _AutomationRegistryLogicA.abi.Events["CancelledUpkeepReport"].ID:
		return _AutomationRegistryLogicA.ParseCancelledUpkeepReport(log)
	case _AutomationRegistryLogicA.abi.Events["ChainSpecificModuleUpdated"].ID:
		return _AutomationRegistryLogicA.ParseChainSpecificModuleUpdated(log)
	case _AutomationRegistryLogicA.abi.Events["DedupKeyAdded"].ID:
		return _AutomationRegistryLogicA.ParseDedupKeyAdded(log)
	case _AutomationRegistryLogicA.abi.Events["FeesWithdrawn"].ID:
		return _AutomationRegistryLogicA.ParseFeesWithdrawn(log)
	case _AutomationRegistryLogicA.abi.Events["FundsAdded"].ID:
		return _AutomationRegistryLogicA.ParseFundsAdded(log)
	case _AutomationRegistryLogicA.abi.Events["FundsWithdrawn"].ID:
		return _AutomationRegistryLogicA.ParseFundsWithdrawn(log)
	case _AutomationRegistryLogicA.abi.Events["InsufficientFundsUpkeepReport"].ID:
		return _AutomationRegistryLogicA.ParseInsufficientFundsUpkeepReport(log)
	case _AutomationRegistryLogicA.abi.Events["NOPsSettledOffchain"].ID:
		return _AutomationRegistryLogicA.ParseNOPsSettledOffchain(log)
	case _AutomationRegistryLogicA.abi.Events["OwnershipTransferRequested"].ID:
		return _AutomationRegistryLogicA.ParseOwnershipTransferRequested(log)
	case _AutomationRegistryLogicA.abi.Events["OwnershipTransferred"].ID:
		return _AutomationRegistryLogicA.ParseOwnershipTransferred(log)
	case _AutomationRegistryLogicA.abi.Events["Paused"].ID:
		return _AutomationRegistryLogicA.ParsePaused(log)
	case _AutomationRegistryLogicA.abi.Events["PayeesUpdated"].ID:
		return _AutomationRegistryLogicA.ParsePayeesUpdated(log)
	case _AutomationRegistryLogicA.abi.Events["PayeeshipTransferRequested"].ID:
		return _AutomationRegistryLogicA.ParsePayeeshipTransferRequested(log)
	case _AutomationRegistryLogicA.abi.Events["PayeeshipTransferred"].ID:
		return _AutomationRegistryLogicA.ParsePayeeshipTransferred(log)
	case _AutomationRegistryLogicA.abi.Events["PaymentWithdrawn"].ID:
		return _AutomationRegistryLogicA.ParsePaymentWithdrawn(log)
	case _AutomationRegistryLogicA.abi.Events["ReorgedUpkeepReport"].ID:
		return _AutomationRegistryLogicA.ParseReorgedUpkeepReport(log)
	case _AutomationRegistryLogicA.abi.Events["StaleUpkeepReport"].ID:
		return _AutomationRegistryLogicA.ParseStaleUpkeepReport(log)
	case _AutomationRegistryLogicA.abi.Events["Unpaused"].ID:
		return _AutomationRegistryLogicA.ParseUnpaused(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepAdminTransferRequested"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepAdminTransferRequested(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepAdminTransferred"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepAdminTransferred(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepCanceled"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepCanceled(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepCheckDataSet"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepCheckDataSet(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepGasLimitSet"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepGasLimitSet(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepMigrated"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepMigrated(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepOffchainConfigSet"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepOffchainConfigSet(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepPaused"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepPaused(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepPerformed"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepPerformed(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepPrivilegeConfigSet"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepPrivilegeConfigSet(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepReceived"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepReceived(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepRegistered"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepRegistered(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepTriggerConfigSet"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepTriggerConfigSet(log)
	case _AutomationRegistryLogicA.abi.Events["UpkeepUnpaused"].ID:
		return _AutomationRegistryLogicA.ParseUpkeepUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (AutomationRegistryLogicAAdminPrivilegeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x7c44b4eb59ee7873514e7e43e7718c269d872965938b288aa143befca62f99d2")
}

func (AutomationRegistryLogicABillingConfigOverridden) Topic() common.Hash {
	return common.HexToHash("0x6b29e7a94c351cd97eae4333dff559076e4fad67f37febac9cdf75b6f7ab6c50")
}

func (AutomationRegistryLogicABillingConfigOverrideRemoved) Topic() common.Hash {
	return common.HexToHash("0x97d0ef3f46a56168af653f547bdb6f77ec2b1d7d9bc6ba0193c2b340ec68064a")
}

func (AutomationRegistryLogicABillingConfigSet) Topic() common.Hash {
	return common.HexToHash("0x720a5849025dc4fd0061aed1bb30efd713cde64ce7f8d807953ecca27c8f143c")
}

func (AutomationRegistryLogicACancelledUpkeepReport) Topic() common.Hash {
	return common.HexToHash("0xc3237c8807c467c1b39b8d0395eff077313e691bf0a7388106792564ebfd5636")
}

func (AutomationRegistryLogicAChainSpecificModuleUpdated) Topic() common.Hash {
	return common.HexToHash("0xdefc28b11a7980dbe0c49dbbd7055a1584bc8075097d1e8b3b57fb7283df2ad7")
}

func (AutomationRegistryLogicADedupKeyAdded) Topic() common.Hash {
	return common.HexToHash("0xa4a4e334c0e330143f9437484fe516c13bc560b86b5b0daf58e7084aaac228f2")
}

func (AutomationRegistryLogicAFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (AutomationRegistryLogicAFundsAdded) Topic() common.Hash {
	return common.HexToHash("0xafd24114486da8ebfc32f3626dada8863652e187461aa74d4bfa734891506203")
}

func (AutomationRegistryLogicAFundsWithdrawn) Topic() common.Hash {
	return common.HexToHash("0xf3b5906e5672f3e524854103bcafbbdba80dbdfeca2c35e116127b1060a68318")
}

func (AutomationRegistryLogicAInsufficientFundsUpkeepReport) Topic() common.Hash {
	return common.HexToHash("0x377c8b0c126ae5248d27aca1c76fac4608aff85673ee3caf09747e1044549e02")
}

func (AutomationRegistryLogicANOPsSettledOffchain) Topic() common.Hash {
	return common.HexToHash("0x5af23b715253628d12b660b27a4f3fc626562ea8a55040aa99ab3dc178989fad")
}

func (AutomationRegistryLogicAOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (AutomationRegistryLogicAOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (AutomationRegistryLogicAPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (AutomationRegistryLogicAPayeesUpdated) Topic() common.Hash {
	return common.HexToHash("0xa46de38886467c59be07a0675f14781206a5477d871628af46c2443822fcb725")
}

func (AutomationRegistryLogicAPayeeshipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e38367")
}

func (AutomationRegistryLogicAPayeeshipTransferred) Topic() common.Hash {
	return common.HexToHash("0x78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b3")
}

func (AutomationRegistryLogicAPaymentWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x9819093176a1851202c7bcfa46845809b4e47c261866550e94ed3775d2f40698")
}

func (AutomationRegistryLogicAReorgedUpkeepReport) Topic() common.Hash {
	return common.HexToHash("0x6aa7f60c176da7af894b384daea2249497448137f5943c1237ada8bc92bdc301")
}

func (AutomationRegistryLogicAStaleUpkeepReport) Topic() common.Hash {
	return common.HexToHash("0x405288ea7be309e16cfdf481367f90a413e1d4634fcdaf8966546db9b93012e8")
}

func (AutomationRegistryLogicAUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (AutomationRegistryLogicAUpkeepAdminTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xb1cbb2c4b8480034c27e06da5f096b8233a8fd4497028593a41ff6df79726b35")
}

func (AutomationRegistryLogicAUpkeepAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x5cff4db96bef051785e999f44bfcd21c18823e034fb92dd376e3db4ce0feeb2c")
}

func (AutomationRegistryLogicAUpkeepCanceled) Topic() common.Hash {
	return common.HexToHash("0x91cb3bb75cfbd718bbfccc56b7f53d92d7048ef4ca39a3b7b7c6d4af1f791181")
}

func (AutomationRegistryLogicAUpkeepCheckDataSet) Topic() common.Hash {
	return common.HexToHash("0xcba2d5723b2ee59e53a8e8a82a4a7caf4fdfe70e9f7c582950bf7e7a5c24e83d")
}

func (AutomationRegistryLogicAUpkeepGasLimitSet) Topic() common.Hash {
	return common.HexToHash("0xc24c07e655ce79fba8a589778987d3c015bc6af1632bb20cf9182e02a65d972c")
}

func (AutomationRegistryLogicAUpkeepMigrated) Topic() common.Hash {
	return common.HexToHash("0xb38647142fbb1ea4c000fc4569b37a4e9a9f6313317b84ee3e5326c1a6cd06ff")
}

func (AutomationRegistryLogicAUpkeepOffchainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3e8740446213c8a77d40e08f79136ce3f347d13ed270a6ebdf57159e0faf4850")
}

func (AutomationRegistryLogicAUpkeepPaused) Topic() common.Hash {
	return common.HexToHash("0x8ab10247ce168c27748e656ecf852b951fcaac790c18106b19aa0ae57a8b741f")
}

func (AutomationRegistryLogicAUpkeepPerformed) Topic() common.Hash {
	return common.HexToHash("0xad8cc9579b21dfe2c2f6ea35ba15b656e46b4f5b0cb424f52739b8ce5cac9c5b")
}

func (AutomationRegistryLogicAUpkeepPrivilegeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2fd8d70753a007014349d4591843cc031c2dd7a260d7dd82eca8253686ae7769")
}

func (AutomationRegistryLogicAUpkeepReceived) Topic() common.Hash {
	return common.HexToHash("0x74931a144e43a50694897f241d973aecb5024c0e910f9bb80a163ea3c1cf5a71")
}

func (AutomationRegistryLogicAUpkeepRegistered) Topic() common.Hash {
	return common.HexToHash("0xbae366358c023f887e791d7a62f2e4316f1026bd77f6fb49501a917b3bc5d012")
}

func (AutomationRegistryLogicAUpkeepTriggerConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2b72ac786c97e68dbab71023ed6f2bdbfc80ad9bb7808941929229d71b7d5664")
}

func (AutomationRegistryLogicAUpkeepUnpaused) Topic() common.Hash {
	return common.HexToHash("0x7bada562044eb163f6b4003c4553e4e62825344c0418eea087bed5ee05a47456")
}

func (_AutomationRegistryLogicA *AutomationRegistryLogicA) Address() common.Address {
	return _AutomationRegistryLogicA.address
}

type AutomationRegistryLogicAInterface interface {
	FallbackTo(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelUpkeep(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error)

	CheckCallback(opts *bind.TransactOpts, id *big.Int, values [][]byte, extraData []byte) (*types.Transaction, error)

	CheckUpkeep(opts *bind.TransactOpts, id *big.Int, triggerData []byte) (*types.Transaction, error)

	CheckUpkeep0(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error)

	ExecuteCallback(opts *bind.TransactOpts, id *big.Int, payload []byte) (*types.Transaction, error)

	MigrateUpkeeps(opts *bind.TransactOpts, ids []*big.Int, destination common.Address) (*types.Transaction, error)

	ReceiveUpkeeps(opts *bind.TransactOpts, encodedUpkeeps []byte) (*types.Transaction, error)

	RegisterUpkeep(opts *bind.TransactOpts, target common.Address, gasLimit uint32, admin common.Address, triggerType uint8, billingToken common.Address, checkData []byte, triggerConfig []byte, offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterAdminPrivilegeConfigSet(opts *bind.FilterOpts, admin []common.Address) (*AutomationRegistryLogicAAdminPrivilegeConfigSetIterator, error)

	WatchAdminPrivilegeConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAAdminPrivilegeConfigSet, admin []common.Address) (event.Subscription, error)

	ParseAdminPrivilegeConfigSet(log types.Log) (*AutomationRegistryLogicAAdminPrivilegeConfigSet, error)

	FilterBillingConfigOverridden(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicABillingConfigOverriddenIterator, error)

	WatchBillingConfigOverridden(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicABillingConfigOverridden, id []*big.Int) (event.Subscription, error)

	ParseBillingConfigOverridden(log types.Log) (*AutomationRegistryLogicABillingConfigOverridden, error)

	FilterBillingConfigOverrideRemoved(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicABillingConfigOverrideRemovedIterator, error)

	WatchBillingConfigOverrideRemoved(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicABillingConfigOverrideRemoved, id []*big.Int) (event.Subscription, error)

	ParseBillingConfigOverrideRemoved(log types.Log) (*AutomationRegistryLogicABillingConfigOverrideRemoved, error)

	FilterBillingConfigSet(opts *bind.FilterOpts, token []common.Address) (*AutomationRegistryLogicABillingConfigSetIterator, error)

	WatchBillingConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicABillingConfigSet, token []common.Address) (event.Subscription, error)

	ParseBillingConfigSet(log types.Log) (*AutomationRegistryLogicABillingConfigSet, error)

	FilterCancelledUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicACancelledUpkeepReportIterator, error)

	WatchCancelledUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicACancelledUpkeepReport, id []*big.Int) (event.Subscription, error)

	ParseCancelledUpkeepReport(log types.Log) (*AutomationRegistryLogicACancelledUpkeepReport, error)

	FilterChainSpecificModuleUpdated(opts *bind.FilterOpts) (*AutomationRegistryLogicAChainSpecificModuleUpdatedIterator, error)

	WatchChainSpecificModuleUpdated(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAChainSpecificModuleUpdated) (event.Subscription, error)

	ParseChainSpecificModuleUpdated(log types.Log) (*AutomationRegistryLogicAChainSpecificModuleUpdated, error)

	FilterDedupKeyAdded(opts *bind.FilterOpts, dedupKey [][32]byte) (*AutomationRegistryLogicADedupKeyAddedIterator, error)

	WatchDedupKeyAdded(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicADedupKeyAdded, dedupKey [][32]byte) (event.Subscription, error)

	ParseDedupKeyAdded(log types.Log) (*AutomationRegistryLogicADedupKeyAdded, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts, recipient []common.Address, assetAddress []common.Address) (*AutomationRegistryLogicAFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAFeesWithdrawn, recipient []common.Address, assetAddress []common.Address) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*AutomationRegistryLogicAFeesWithdrawn, error)

	FilterFundsAdded(opts *bind.FilterOpts, id []*big.Int, from []common.Address) (*AutomationRegistryLogicAFundsAddedIterator, error)

	WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAFundsAdded, id []*big.Int, from []common.Address) (event.Subscription, error)

	ParseFundsAdded(log types.Log) (*AutomationRegistryLogicAFundsAdded, error)

	FilterFundsWithdrawn(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAFundsWithdrawnIterator, error)

	WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAFundsWithdrawn, id []*big.Int) (event.Subscription, error)

	ParseFundsWithdrawn(log types.Log) (*AutomationRegistryLogicAFundsWithdrawn, error)

	FilterInsufficientFundsUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAInsufficientFundsUpkeepReportIterator, error)

	WatchInsufficientFundsUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAInsufficientFundsUpkeepReport, id []*big.Int) (event.Subscription, error)

	ParseInsufficientFundsUpkeepReport(log types.Log) (*AutomationRegistryLogicAInsufficientFundsUpkeepReport, error)

	FilterNOPsSettledOffchain(opts *bind.FilterOpts) (*AutomationRegistryLogicANOPsSettledOffchainIterator, error)

	WatchNOPsSettledOffchain(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicANOPsSettledOffchain) (event.Subscription, error)

	ParseNOPsSettledOffchain(log types.Log) (*AutomationRegistryLogicANOPsSettledOffchain, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationRegistryLogicAOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*AutomationRegistryLogicAOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationRegistryLogicAOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*AutomationRegistryLogicAOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*AutomationRegistryLogicAPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*AutomationRegistryLogicAPaused, error)

	FilterPayeesUpdated(opts *bind.FilterOpts) (*AutomationRegistryLogicAPayeesUpdatedIterator, error)

	WatchPayeesUpdated(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPayeesUpdated) (event.Subscription, error)

	ParsePayeesUpdated(log types.Log) (*AutomationRegistryLogicAPayeesUpdated, error)

	FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, from []common.Address, to []common.Address) (*AutomationRegistryLogicAPayeeshipTransferRequestedIterator, error)

	WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPayeeshipTransferRequested, transmitter []common.Address, from []common.Address, to []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferRequested(log types.Log) (*AutomationRegistryLogicAPayeeshipTransferRequested, error)

	FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, from []common.Address, to []common.Address) (*AutomationRegistryLogicAPayeeshipTransferredIterator, error)

	WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPayeeshipTransferred, transmitter []common.Address, from []common.Address, to []common.Address) (event.Subscription, error)

	ParsePayeeshipTransferred(log types.Log) (*AutomationRegistryLogicAPayeeshipTransferred, error)

	FilterPaymentWithdrawn(opts *bind.FilterOpts, transmitter []common.Address, amount []*big.Int, to []common.Address) (*AutomationRegistryLogicAPaymentWithdrawnIterator, error)

	WatchPaymentWithdrawn(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAPaymentWithdrawn, transmitter []common.Address, amount []*big.Int, to []common.Address) (event.Subscription, error)

	ParsePaymentWithdrawn(log types.Log) (*AutomationRegistryLogicAPaymentWithdrawn, error)

	FilterReorgedUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAReorgedUpkeepReportIterator, error)

	WatchReorgedUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAReorgedUpkeepReport, id []*big.Int) (event.Subscription, error)

	ParseReorgedUpkeepReport(log types.Log) (*AutomationRegistryLogicAReorgedUpkeepReport, error)

	FilterStaleUpkeepReport(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAStaleUpkeepReportIterator, error)

	WatchStaleUpkeepReport(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAStaleUpkeepReport, id []*big.Int) (event.Subscription, error)

	ParseStaleUpkeepReport(log types.Log) (*AutomationRegistryLogicAStaleUpkeepReport, error)

	FilterUnpaused(opts *bind.FilterOpts) (*AutomationRegistryLogicAUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*AutomationRegistryLogicAUnpaused, error)

	FilterUpkeepAdminTransferRequested(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*AutomationRegistryLogicAUpkeepAdminTransferRequestedIterator, error)

	WatchUpkeepAdminTransferRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepAdminTransferRequested, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseUpkeepAdminTransferRequested(log types.Log) (*AutomationRegistryLogicAUpkeepAdminTransferRequested, error)

	FilterUpkeepAdminTransferred(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*AutomationRegistryLogicAUpkeepAdminTransferredIterator, error)

	WatchUpkeepAdminTransferred(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepAdminTransferred, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseUpkeepAdminTransferred(log types.Log) (*AutomationRegistryLogicAUpkeepAdminTransferred, error)

	FilterUpkeepCanceled(opts *bind.FilterOpts, id []*big.Int, atBlockHeight []uint64) (*AutomationRegistryLogicAUpkeepCanceledIterator, error)

	WatchUpkeepCanceled(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepCanceled, id []*big.Int, atBlockHeight []uint64) (event.Subscription, error)

	ParseUpkeepCanceled(log types.Log) (*AutomationRegistryLogicAUpkeepCanceled, error)

	FilterUpkeepCheckDataSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepCheckDataSetIterator, error)

	WatchUpkeepCheckDataSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepCheckDataSet, id []*big.Int) (event.Subscription, error)

	ParseUpkeepCheckDataSet(log types.Log) (*AutomationRegistryLogicAUpkeepCheckDataSet, error)

	FilterUpkeepGasLimitSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepGasLimitSetIterator, error)

	WatchUpkeepGasLimitSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepGasLimitSet, id []*big.Int) (event.Subscription, error)

	ParseUpkeepGasLimitSet(log types.Log) (*AutomationRegistryLogicAUpkeepGasLimitSet, error)

	FilterUpkeepMigrated(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepMigratedIterator, error)

	WatchUpkeepMigrated(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepMigrated, id []*big.Int) (event.Subscription, error)

	ParseUpkeepMigrated(log types.Log) (*AutomationRegistryLogicAUpkeepMigrated, error)

	FilterUpkeepOffchainConfigSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepOffchainConfigSetIterator, error)

	WatchUpkeepOffchainConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepOffchainConfigSet, id []*big.Int) (event.Subscription, error)

	ParseUpkeepOffchainConfigSet(log types.Log) (*AutomationRegistryLogicAUpkeepOffchainConfigSet, error)

	FilterUpkeepPaused(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepPausedIterator, error)

	WatchUpkeepPaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepPaused, id []*big.Int) (event.Subscription, error)

	ParseUpkeepPaused(log types.Log) (*AutomationRegistryLogicAUpkeepPaused, error)

	FilterUpkeepPerformed(opts *bind.FilterOpts, id []*big.Int, success []bool) (*AutomationRegistryLogicAUpkeepPerformedIterator, error)

	WatchUpkeepPerformed(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepPerformed, id []*big.Int, success []bool) (event.Subscription, error)

	ParseUpkeepPerformed(log types.Log) (*AutomationRegistryLogicAUpkeepPerformed, error)

	FilterUpkeepPrivilegeConfigSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepPrivilegeConfigSetIterator, error)

	WatchUpkeepPrivilegeConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepPrivilegeConfigSet, id []*big.Int) (event.Subscription, error)

	ParseUpkeepPrivilegeConfigSet(log types.Log) (*AutomationRegistryLogicAUpkeepPrivilegeConfigSet, error)

	FilterUpkeepReceived(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepReceivedIterator, error)

	WatchUpkeepReceived(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepReceived, id []*big.Int) (event.Subscription, error)

	ParseUpkeepReceived(log types.Log) (*AutomationRegistryLogicAUpkeepReceived, error)

	FilterUpkeepRegistered(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepRegisteredIterator, error)

	WatchUpkeepRegistered(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepRegistered, id []*big.Int) (event.Subscription, error)

	ParseUpkeepRegistered(log types.Log) (*AutomationRegistryLogicAUpkeepRegistered, error)

	FilterUpkeepTriggerConfigSet(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepTriggerConfigSetIterator, error)

	WatchUpkeepTriggerConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepTriggerConfigSet, id []*big.Int) (event.Subscription, error)

	ParseUpkeepTriggerConfigSet(log types.Log) (*AutomationRegistryLogicAUpkeepTriggerConfigSet, error)

	FilterUpkeepUnpaused(opts *bind.FilterOpts, id []*big.Int) (*AutomationRegistryLogicAUpkeepUnpausedIterator, error)

	WatchUpkeepUnpaused(opts *bind.WatchOpts, sink chan<- *AutomationRegistryLogicAUpkeepUnpaused, id []*big.Int) (event.Subscription, error)

	ParseUpkeepUnpaused(log types.Log) (*AutomationRegistryLogicAUpkeepUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
