// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exposed_configurator

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

type ConfiguratorConfigurationState struct {
	ConfigCount             uint64
	LatestConfigBlockNumber uint32
	IsFlipped               bool
	ConfigDigest            [2][32]byte
}

var ExposedConfiguratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSigners\",\"type\":\"uint256\"}],\"name\":\"ExcessSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FaultToleranceMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSigners\",\"type\":\"uint256\"}],\"name\":\"InsufficientSigners\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signers\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFlipped\",\"type\":\"bool\"}],\"name\":\"ProductionConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"retiredConfigDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFlipped\",\"type\":\"bool\"}],\"name\":\"PromoteStagingConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"signers\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFlipped\",\"type\":\"bool\"}],\"name\":\"StagingConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_configId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_configCount\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_signers\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_offchainTransmitters\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_encodedConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_encodedConfig\",\"type\":\"bytes\"}],\"name\":\"exposedConfigDigestFromConfigData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"donId\",\"type\":\"bytes32\"}],\"name\":\"exposedReadConfigurationStates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"latestConfigBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isFlipped\",\"type\":\"bool\"},{\"internalType\":\"bytes32[2]\",\"name\":\"configDigest\",\"type\":\"bytes32[2]\"}],\"internalType\":\"structConfigurator.ConfigurationState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"donId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isFlipped\",\"type\":\"bool\"}],\"name\":\"promoteStagingConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"donId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"signers\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setProductionConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"donId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"signers\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setStagingConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isVerifier\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6117ed806101576000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80638da5cb5b11610076578063dfb533d01161005b578063dfb533d0146101f2578063e6e7c5a414610205578063f2fde38b1461021857600080fd5b80638da5cb5b146101aa57806399a07340146101d257600080fd5b806360e72ec9116100a757806360e72ec91461016c578063790464e01461018d57806379ba5097146101a257600080fd5b806301ffc9a7146100c3578063181f5a771461012d575b600080fd5b6101186100d1366004610f00565b7fffffffff00000000000000000000000000000000000000000000000000000000167f40569294000000000000000000000000000000000000000000000000000000001490565b60405190151581526020015b60405180910390f35b604080518082018252601281527f436f6e666967757261746f7220302e342e300000000000000000000000000000602082015290516101249190610fad565b61017f61017a366004611276565b61022b565b604051908152602001610124565b6101a061019b366004611387565b610287565b005b6101a061039a565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610124565b6101e56101e036600461145f565b610497565b6040516101249190611478565b6101a0610200366004611387565b61053a565b6101a06102133660046114dd565b6108dc565b6101a0610226366004611512565b610a4a565b60006102778c8c8c8c8c8c8c8c8c8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508e92508d9150610a5e9050565b9c9b505050505050505050505050565b85518460ff16806000036102c7576040517f0743bae600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601f821115610311576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101839052601f60248201526044015b60405180910390fd5b61031c81600361155c565b8211610374578161032e82600361155c565b610339906001611579565b6040517f9dd9e6d800000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401610308565b61037c610b0c565b61038f8946308b8b8b8b8b8b6001610b8f565b505050505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461041b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610308565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61049f610eb4565b6000828152600260208181526040928390208351608081018552815467ffffffffffffffff8116825268010000000000000000810463ffffffff16938201939093526c0100000000000000000000000090920460ff161515828501528351808501948590529193909260608501929160018501919082845b815481526020019060010190808311610517575050505050815250509050919050565b85518460ff168060000361057a576040517f0743bae600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601f8211156105bf576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101839052601f6024820152604401610308565b6105ca81600361155c565b82116105dc578161032e82600361155c565b6105e4610b0c565b845160401461064f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c6964206f6e636861696e436f6e666967206c656e677468000000006044820152606401610308565b60208501516040860151600182146106e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f556e737570706f72746564206f6e636861696e436f6e6669672076657273696f60448201527f6e000000000000000000000000000000000000000000000000000000000000006064820152608401610308565b60008b81526002602081815260408084208151608081018352815467ffffffffffffffff8116825268010000000000000000810463ffffffff16948201949094526c0100000000000000000000000090930460ff161515838301528151808301928390529293909260608501929091600185019182845b81548152602001906001019080831161076057505050505081525050905060008160400151610790576000610793565b60015b60ff169050600260008e815260200190815260200160002060010181600281106107bf576107bf61158c565b01548314610829576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f57726f6e67207072656465636573736f72436f6e6669674469676573740000006044820152606401610308565b600260008e815260200190815260200160002060010181600281106108505761085061158c565b015483146108ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f496e76616c6964207072656465636573736f72436f6e666967446967657374006044820152606401610308565b6108cd8d46308f8f8f8f8f8f6000610b8f565b50505050505050505050505050565b6108e4610b0c565b600082815260026020526040902080546c01000000000000000000000000900460ff161515821515036109c25780547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff1682156c0100000000000000000000000002178155600181018261095957600061095c565b60015b60ff166002811061096f5761096f61158c565b015481546040516c0100000000000000000000000090910460ff161515815284907f1062aa08ac6046a0e69e3eafdf12d1eba63a67b71a874623e86eb06348a1d84f9060200160405180910390a3505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f50726f6d6f746553746167696e67436f6e6669673a206973466c69707065642060448201527f6d757374206d617463682063757272656e7420737461746500000000000000006064820152608401610308565b610a52610b0c565b610a5b81610dbf565b50565b6000808b8b8b8b8b8b8b8b8b8b604051602001610a849a9998979695949392919061166c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e09000000000000000000000000000000000000000000000000000000000000179150509a9950505050505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b8d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610308565b565b60008a8152600260205260408120805490919082908290610bb99067ffffffffffffffff16611719565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905590506000610bf48d8d8d858e8e8e8e8e8e610a5e565b90508315610cbc578c7f261b20c2ecd99d86d6e936279e4f78db34603a3de3a4a84d6f3d4e0dd55e24788460000160089054906101000a900463ffffffff1683858e8e8e8e8e8e8d600001600c9054906101000a900460ff16604051610c639a99989796959493929190611740565b60405180910390a260008d815260026020526040902083548291600101906c01000000000000000000000000900460ff16610c9f576000610ca2565b60015b60ff1660028110610cb557610cb561158c565b0155610d78565b8c7fef1b5f9d1b927b0fe871b12c7e7846457602d67b2bc36b0bc95feaf480e890568460000160089054906101000a900463ffffffff1683858e8e8e8e8e8e8d600001600c9054906101000a900460ff16604051610d239a99989796959493929190611740565b60405180910390a260008d815260026020526040902083548291600101906c01000000000000000000000000900460ff16610d5f576001610d62565b60005b60ff1660028110610d7557610d7561158c565b01555b505080547fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff16680100000000000000004363ffffffff160217905550505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610e3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610308565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604080516080810182526000808252602082018190529181019190915260608101610edd610ee2565b905290565b60405180604001604052806002906020820280368337509192915050565b600060208284031215610f1257600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610f4257600080fd5b9392505050565b6000815180845260005b81811015610f6f57602081850181015186830182015201610f53565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610f426020830184610f49565b803573ffffffffffffffffffffffffffffffffffffffff81168114610fe457600080fd5b919050565b803567ffffffffffffffff81168114610fe457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561107757611077611001565b604052919050565b600067ffffffffffffffff82111561109957611099611001565b5060051b60200190565b600082601f8301126110b457600080fd5b813567ffffffffffffffff8111156110ce576110ce611001565b6110ff60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611030565b81815284602083860101111561111457600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261114257600080fd5b813560206111576111528361107f565b611030565b82815260059290921b8401810191818101908684111561117657600080fd5b8286015b848110156111b657803567ffffffffffffffff81111561119a5760008081fd5b6111a88986838b01016110a3565b84525091830191830161117a565b509695505050505050565b600082601f8301126111d257600080fd5b813560206111e26111528361107f565b82815260059290921b8401810191818101908684111561120157600080fd5b8286015b848110156111b65780358352918301918301611205565b803560ff81168114610fe457600080fd5b60008083601f84011261123f57600080fd5b50813567ffffffffffffffff81111561125757600080fd5b60208301915083602082850101111561126f57600080fd5b9250929050565b60008060008060008060008060008060006101408c8e03121561129857600080fd5b8b359a5060208c013599506112af60408d01610fc0565b98506112bd60608d01610fe9565b975067ffffffffffffffff8060808e013511156112d957600080fd5b6112e98e60808f01358f01611131565b97508060a08e013511156112fc57600080fd5b61130c8e60a08f01358f016111c1565b965061131a60c08e0161121c565b95508060e08e0135111561132d57600080fd5b61133d8e60e08f01358f0161122d565b909550935061134f6101008e01610fe9565b9250806101208e0135111561136357600080fd5b506113758d6101208e01358e016110a3565b90509295989b509295989b9093969950565b600080600080600080600060e0888a0312156113a257600080fd5b87359650602088013567ffffffffffffffff808211156113c157600080fd5b6113cd8b838c01611131565b975060408a01359150808211156113e357600080fd5b6113ef8b838c016111c1565b96506113fd60608b0161121c565b955060808a013591508082111561141357600080fd5b61141f8b838c016110a3565b945061142d60a08b01610fe9565b935060c08a013591508082111561144357600080fd5b506114508a828b016110a3565b91505092959891949750929550565b60006020828403121561147157600080fd5b5035919050565b600060a08201905067ffffffffffffffff8351168252602063ffffffff81850151168184015260408401511515604084015260608401516060840160005b60028110156114d3578251825291830191908301906001016114b6565b5050505092915050565b600080604083850312156114f057600080fd5b823591506020830135801515811461150757600080fd5b809150509250929050565b60006020828403121561152457600080fd5b610f4282610fc0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820281158282048414176115735761157361152d565b92915050565b808201808211156115735761157361152d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825180855260208086019550808260051b84010181860160005b84811015611624577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018952611612838351610f49565b988401989250908301906001016115d8565b5090979650505050505050565b600081518084526020808501945080840160005b8381101561166157815187529582019590820190600101611645565b509495945050505050565b60006101408c83528b602084015273ffffffffffffffffffffffffffffffffffffffff8b16604084015267ffffffffffffffff808b1660608501528160808501526116b98285018b6115bb565b915083820360a08501526116cd828a611631565b915060ff881660c085015283820360e08501526116ea8288610f49565b90861661010085015283810361012085015290506117088185610f49565b9d9c50505050505050505050505050565b600067ffffffffffffffff8083168181036117365761173661152d565b6001019392505050565b600061014063ffffffff8d1683528b602084015267ffffffffffffffff808c1660408501528160608501526117778285018c6115bb565b9150838203608085015261178b828b611631565b915060ff891660a085015283820360c08501526117a88289610f49565b90871660e085015283810361010085015290506117c58186610f49565b9150508215156101208301529b9a505050505050505050505056fea164736f6c6343000813000a",
}

var ExposedConfiguratorABI = ExposedConfiguratorMetaData.ABI

var ExposedConfiguratorBin = ExposedConfiguratorMetaData.Bin

func DeployExposedConfigurator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExposedConfigurator, error) {
	parsed, err := ExposedConfiguratorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExposedConfiguratorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExposedConfigurator{address: address, abi: *parsed, ExposedConfiguratorCaller: ExposedConfiguratorCaller{contract: contract}, ExposedConfiguratorTransactor: ExposedConfiguratorTransactor{contract: contract}, ExposedConfiguratorFilterer: ExposedConfiguratorFilterer{contract: contract}}, nil
}

type ExposedConfigurator struct {
	address common.Address
	abi     abi.ABI
	ExposedConfiguratorCaller
	ExposedConfiguratorTransactor
	ExposedConfiguratorFilterer
}

type ExposedConfiguratorCaller struct {
	contract *bind.BoundContract
}

type ExposedConfiguratorTransactor struct {
	contract *bind.BoundContract
}

type ExposedConfiguratorFilterer struct {
	contract *bind.BoundContract
}

type ExposedConfiguratorSession struct {
	Contract     *ExposedConfigurator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ExposedConfiguratorCallerSession struct {
	Contract *ExposedConfiguratorCaller
	CallOpts bind.CallOpts
}

type ExposedConfiguratorTransactorSession struct {
	Contract     *ExposedConfiguratorTransactor
	TransactOpts bind.TransactOpts
}

type ExposedConfiguratorRaw struct {
	Contract *ExposedConfigurator
}

type ExposedConfiguratorCallerRaw struct {
	Contract *ExposedConfiguratorCaller
}

type ExposedConfiguratorTransactorRaw struct {
	Contract *ExposedConfiguratorTransactor
}

func NewExposedConfigurator(address common.Address, backend bind.ContractBackend) (*ExposedConfigurator, error) {
	abi, err := abi.JSON(strings.NewReader(ExposedConfiguratorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindExposedConfigurator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExposedConfigurator{address: address, abi: abi, ExposedConfiguratorCaller: ExposedConfiguratorCaller{contract: contract}, ExposedConfiguratorTransactor: ExposedConfiguratorTransactor{contract: contract}, ExposedConfiguratorFilterer: ExposedConfiguratorFilterer{contract: contract}}, nil
}

func NewExposedConfiguratorCaller(address common.Address, caller bind.ContractCaller) (*ExposedConfiguratorCaller, error) {
	contract, err := bindExposedConfigurator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorCaller{contract: contract}, nil
}

func NewExposedConfiguratorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExposedConfiguratorTransactor, error) {
	contract, err := bindExposedConfigurator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorTransactor{contract: contract}, nil
}

func NewExposedConfiguratorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExposedConfiguratorFilterer, error) {
	contract, err := bindExposedConfigurator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorFilterer{contract: contract}, nil
}

func bindExposedConfigurator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExposedConfiguratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ExposedConfigurator *ExposedConfiguratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExposedConfigurator.Contract.ExposedConfiguratorCaller.contract.Call(opts, result, method, params...)
}

func (_ExposedConfigurator *ExposedConfiguratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.ExposedConfiguratorTransactor.contract.Transfer(opts)
}

func (_ExposedConfigurator *ExposedConfiguratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.ExposedConfiguratorTransactor.contract.Transact(opts, method, params...)
}

func (_ExposedConfigurator *ExposedConfiguratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExposedConfigurator.Contract.contract.Call(opts, result, method, params...)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.contract.Transfer(opts)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.contract.Transact(opts, method, params...)
}

func (_ExposedConfigurator *ExposedConfiguratorCaller) ExposedConfigDigestFromConfigData(opts *bind.CallOpts, _configId [32]byte, _chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers [][]byte, _offchainTransmitters [][32]byte, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error) {
	var out []interface{}
	err := _ExposedConfigurator.contract.Call(opts, &out, "exposedConfigDigestFromConfigData", _configId, _chainId, _contractAddress, _configCount, _signers, _offchainTransmitters, _f, _onchainConfig, _encodedConfigVersion, _encodedConfig)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_ExposedConfigurator *ExposedConfiguratorSession) ExposedConfigDigestFromConfigData(_configId [32]byte, _chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers [][]byte, _offchainTransmitters [][32]byte, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error) {
	return _ExposedConfigurator.Contract.ExposedConfigDigestFromConfigData(&_ExposedConfigurator.CallOpts, _configId, _chainId, _contractAddress, _configCount, _signers, _offchainTransmitters, _f, _onchainConfig, _encodedConfigVersion, _encodedConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorCallerSession) ExposedConfigDigestFromConfigData(_configId [32]byte, _chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers [][]byte, _offchainTransmitters [][32]byte, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error) {
	return _ExposedConfigurator.Contract.ExposedConfigDigestFromConfigData(&_ExposedConfigurator.CallOpts, _configId, _chainId, _contractAddress, _configCount, _signers, _offchainTransmitters, _f, _onchainConfig, _encodedConfigVersion, _encodedConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorCaller) ExposedReadConfigurationStates(opts *bind.CallOpts, donId [32]byte) (ConfiguratorConfigurationState, error) {
	var out []interface{}
	err := _ExposedConfigurator.contract.Call(opts, &out, "exposedReadConfigurationStates", donId)

	if err != nil {
		return *new(ConfiguratorConfigurationState), err
	}

	out0 := *abi.ConvertType(out[0], new(ConfiguratorConfigurationState)).(*ConfiguratorConfigurationState)

	return out0, err

}

func (_ExposedConfigurator *ExposedConfiguratorSession) ExposedReadConfigurationStates(donId [32]byte) (ConfiguratorConfigurationState, error) {
	return _ExposedConfigurator.Contract.ExposedReadConfigurationStates(&_ExposedConfigurator.CallOpts, donId)
}

func (_ExposedConfigurator *ExposedConfiguratorCallerSession) ExposedReadConfigurationStates(donId [32]byte) (ConfiguratorConfigurationState, error) {
	return _ExposedConfigurator.Contract.ExposedReadConfigurationStates(&_ExposedConfigurator.CallOpts, donId)
}

func (_ExposedConfigurator *ExposedConfiguratorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExposedConfigurator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ExposedConfigurator *ExposedConfiguratorSession) Owner() (common.Address, error) {
	return _ExposedConfigurator.Contract.Owner(&_ExposedConfigurator.CallOpts)
}

func (_ExposedConfigurator *ExposedConfiguratorCallerSession) Owner() (common.Address, error) {
	return _ExposedConfigurator.Contract.Owner(&_ExposedConfigurator.CallOpts)
}

func (_ExposedConfigurator *ExposedConfiguratorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ExposedConfigurator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ExposedConfigurator *ExposedConfiguratorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExposedConfigurator.Contract.SupportsInterface(&_ExposedConfigurator.CallOpts, interfaceId)
}

func (_ExposedConfigurator *ExposedConfiguratorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExposedConfigurator.Contract.SupportsInterface(&_ExposedConfigurator.CallOpts, interfaceId)
}

func (_ExposedConfigurator *ExposedConfiguratorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExposedConfigurator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ExposedConfigurator *ExposedConfiguratorSession) TypeAndVersion() (string, error) {
	return _ExposedConfigurator.Contract.TypeAndVersion(&_ExposedConfigurator.CallOpts)
}

func (_ExposedConfigurator *ExposedConfiguratorCallerSession) TypeAndVersion() (string, error) {
	return _ExposedConfigurator.Contract.TypeAndVersion(&_ExposedConfigurator.CallOpts)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExposedConfigurator.contract.Transact(opts, "acceptOwnership")
}

func (_ExposedConfigurator *ExposedConfiguratorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.AcceptOwnership(&_ExposedConfigurator.TransactOpts)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.AcceptOwnership(&_ExposedConfigurator.TransactOpts)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactor) PromoteStagingConfig(opts *bind.TransactOpts, donId [32]byte, isFlipped bool) (*types.Transaction, error) {
	return _ExposedConfigurator.contract.Transact(opts, "promoteStagingConfig", donId, isFlipped)
}

func (_ExposedConfigurator *ExposedConfiguratorSession) PromoteStagingConfig(donId [32]byte, isFlipped bool) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.PromoteStagingConfig(&_ExposedConfigurator.TransactOpts, donId, isFlipped)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactorSession) PromoteStagingConfig(donId [32]byte, isFlipped bool) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.PromoteStagingConfig(&_ExposedConfigurator.TransactOpts, donId, isFlipped)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactor) SetProductionConfig(opts *bind.TransactOpts, donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedConfigurator.contract.Transact(opts, "setProductionConfig", donId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorSession) SetProductionConfig(donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.SetProductionConfig(&_ExposedConfigurator.TransactOpts, donId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactorSession) SetProductionConfig(donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.SetProductionConfig(&_ExposedConfigurator.TransactOpts, donId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactor) SetStagingConfig(opts *bind.TransactOpts, donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedConfigurator.contract.Transact(opts, "setStagingConfig", donId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorSession) SetStagingConfig(donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.SetStagingConfig(&_ExposedConfigurator.TransactOpts, donId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactorSession) SetStagingConfig(donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.SetStagingConfig(&_ExposedConfigurator.TransactOpts, donId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ExposedConfigurator.contract.Transact(opts, "transferOwnership", to)
}

func (_ExposedConfigurator *ExposedConfiguratorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.TransferOwnership(&_ExposedConfigurator.TransactOpts, to)
}

func (_ExposedConfigurator *ExposedConfiguratorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExposedConfigurator.Contract.TransferOwnership(&_ExposedConfigurator.TransactOpts, to)
}

type ExposedConfiguratorOwnershipTransferRequestedIterator struct {
	Event *ExposedConfiguratorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedConfiguratorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedConfiguratorOwnershipTransferRequested)
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
		it.Event = new(ExposedConfiguratorOwnershipTransferRequested)
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

func (it *ExposedConfiguratorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ExposedConfiguratorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedConfiguratorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedConfiguratorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorOwnershipTransferRequestedIterator{contract: _ExposedConfigurator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedConfiguratorOwnershipTransferRequested)
				if err := _ExposedConfigurator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ExposedConfigurator *ExposedConfiguratorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ExposedConfiguratorOwnershipTransferRequested, error) {
	event := new(ExposedConfiguratorOwnershipTransferRequested)
	if err := _ExposedConfigurator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedConfiguratorOwnershipTransferredIterator struct {
	Event *ExposedConfiguratorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedConfiguratorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedConfiguratorOwnershipTransferred)
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
		it.Event = new(ExposedConfiguratorOwnershipTransferred)
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

func (it *ExposedConfiguratorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ExposedConfiguratorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedConfiguratorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedConfiguratorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorOwnershipTransferredIterator{contract: _ExposedConfigurator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedConfiguratorOwnershipTransferred)
				if err := _ExposedConfigurator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ExposedConfigurator *ExposedConfiguratorFilterer) ParseOwnershipTransferred(log types.Log) (*ExposedConfiguratorOwnershipTransferred, error) {
	event := new(ExposedConfiguratorOwnershipTransferred)
	if err := _ExposedConfigurator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedConfiguratorProductionConfigSetIterator struct {
	Event *ExposedConfiguratorProductionConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedConfiguratorProductionConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedConfiguratorProductionConfigSet)
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
		it.Event = new(ExposedConfiguratorProductionConfigSet)
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

func (it *ExposedConfiguratorProductionConfigSetIterator) Error() error {
	return it.fail
}

func (it *ExposedConfiguratorProductionConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedConfiguratorProductionConfigSet struct {
	ConfigId                  [32]byte
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   [][]byte
	OffchainTransmitters      [][32]byte
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	IsFlipped                 bool
	Raw                       types.Log
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) FilterProductionConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ExposedConfiguratorProductionConfigSetIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.FilterLogs(opts, "ProductionConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorProductionConfigSetIterator{contract: _ExposedConfigurator.contract, event: "ProductionConfigSet", logs: logs, sub: sub}, nil
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) WatchProductionConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorProductionConfigSet, configId [][32]byte) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.WatchLogs(opts, "ProductionConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedConfiguratorProductionConfigSet)
				if err := _ExposedConfigurator.contract.UnpackLog(event, "ProductionConfigSet", log); err != nil {
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

func (_ExposedConfigurator *ExposedConfiguratorFilterer) ParseProductionConfigSet(log types.Log) (*ExposedConfiguratorProductionConfigSet, error) {
	event := new(ExposedConfiguratorProductionConfigSet)
	if err := _ExposedConfigurator.contract.UnpackLog(event, "ProductionConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedConfiguratorPromoteStagingConfigIterator struct {
	Event *ExposedConfiguratorPromoteStagingConfig

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedConfiguratorPromoteStagingConfigIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedConfiguratorPromoteStagingConfig)
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
		it.Event = new(ExposedConfiguratorPromoteStagingConfig)
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

func (it *ExposedConfiguratorPromoteStagingConfigIterator) Error() error {
	return it.fail
}

func (it *ExposedConfiguratorPromoteStagingConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedConfiguratorPromoteStagingConfig struct {
	ConfigId            [32]byte
	RetiredConfigDigest [32]byte
	IsFlipped           bool
	Raw                 types.Log
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) FilterPromoteStagingConfig(opts *bind.FilterOpts, configId [][32]byte, retiredConfigDigest [][32]byte) (*ExposedConfiguratorPromoteStagingConfigIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var retiredConfigDigestRule []interface{}
	for _, retiredConfigDigestItem := range retiredConfigDigest {
		retiredConfigDigestRule = append(retiredConfigDigestRule, retiredConfigDigestItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.FilterLogs(opts, "PromoteStagingConfig", configIdRule, retiredConfigDigestRule)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorPromoteStagingConfigIterator{contract: _ExposedConfigurator.contract, event: "PromoteStagingConfig", logs: logs, sub: sub}, nil
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) WatchPromoteStagingConfig(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorPromoteStagingConfig, configId [][32]byte, retiredConfigDigest [][32]byte) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var retiredConfigDigestRule []interface{}
	for _, retiredConfigDigestItem := range retiredConfigDigest {
		retiredConfigDigestRule = append(retiredConfigDigestRule, retiredConfigDigestItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.WatchLogs(opts, "PromoteStagingConfig", configIdRule, retiredConfigDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedConfiguratorPromoteStagingConfig)
				if err := _ExposedConfigurator.contract.UnpackLog(event, "PromoteStagingConfig", log); err != nil {
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

func (_ExposedConfigurator *ExposedConfiguratorFilterer) ParsePromoteStagingConfig(log types.Log) (*ExposedConfiguratorPromoteStagingConfig, error) {
	event := new(ExposedConfiguratorPromoteStagingConfig)
	if err := _ExposedConfigurator.contract.UnpackLog(event, "PromoteStagingConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExposedConfiguratorStagingConfigSetIterator struct {
	Event *ExposedConfiguratorStagingConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExposedConfiguratorStagingConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExposedConfiguratorStagingConfigSet)
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
		it.Event = new(ExposedConfiguratorStagingConfigSet)
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

func (it *ExposedConfiguratorStagingConfigSetIterator) Error() error {
	return it.fail
}

func (it *ExposedConfiguratorStagingConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExposedConfiguratorStagingConfigSet struct {
	ConfigId                  [32]byte
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   [][]byte
	OffchainTransmitters      [][32]byte
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	IsFlipped                 bool
	Raw                       types.Log
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) FilterStagingConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ExposedConfiguratorStagingConfigSetIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.FilterLogs(opts, "StagingConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return &ExposedConfiguratorStagingConfigSetIterator{contract: _ExposedConfigurator.contract, event: "StagingConfigSet", logs: logs, sub: sub}, nil
}

func (_ExposedConfigurator *ExposedConfiguratorFilterer) WatchStagingConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorStagingConfigSet, configId [][32]byte) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _ExposedConfigurator.contract.WatchLogs(opts, "StagingConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExposedConfiguratorStagingConfigSet)
				if err := _ExposedConfigurator.contract.UnpackLog(event, "StagingConfigSet", log); err != nil {
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

func (_ExposedConfigurator *ExposedConfiguratorFilterer) ParseStagingConfigSet(log types.Log) (*ExposedConfiguratorStagingConfigSet, error) {
	event := new(ExposedConfiguratorStagingConfigSet)
	if err := _ExposedConfigurator.contract.UnpackLog(event, "StagingConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ExposedConfigurator *ExposedConfigurator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ExposedConfigurator.abi.Events["OwnershipTransferRequested"].ID:
		return _ExposedConfigurator.ParseOwnershipTransferRequested(log)
	case _ExposedConfigurator.abi.Events["OwnershipTransferred"].ID:
		return _ExposedConfigurator.ParseOwnershipTransferred(log)
	case _ExposedConfigurator.abi.Events["ProductionConfigSet"].ID:
		return _ExposedConfigurator.ParseProductionConfigSet(log)
	case _ExposedConfigurator.abi.Events["PromoteStagingConfig"].ID:
		return _ExposedConfigurator.ParsePromoteStagingConfig(log)
	case _ExposedConfigurator.abi.Events["StagingConfigSet"].ID:
		return _ExposedConfigurator.ParseStagingConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ExposedConfiguratorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ExposedConfiguratorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ExposedConfiguratorProductionConfigSet) Topic() common.Hash {
	return common.HexToHash("0x261b20c2ecd99d86d6e936279e4f78db34603a3de3a4a84d6f3d4e0dd55e2478")
}

func (ExposedConfiguratorPromoteStagingConfig) Topic() common.Hash {
	return common.HexToHash("0x1062aa08ac6046a0e69e3eafdf12d1eba63a67b71a874623e86eb06348a1d84f")
}

func (ExposedConfiguratorStagingConfigSet) Topic() common.Hash {
	return common.HexToHash("0xef1b5f9d1b927b0fe871b12c7e7846457602d67b2bc36b0bc95feaf480e89056")
}

func (_ExposedConfigurator *ExposedConfigurator) Address() common.Address {
	return _ExposedConfigurator.address
}

type ExposedConfiguratorInterface interface {
	ExposedConfigDigestFromConfigData(opts *bind.CallOpts, _configId [32]byte, _chainId *big.Int, _contractAddress common.Address, _configCount uint64, _signers [][]byte, _offchainTransmitters [][32]byte, _f uint8, _onchainConfig []byte, _encodedConfigVersion uint64, _encodedConfig []byte) ([32]byte, error)

	ExposedReadConfigurationStates(opts *bind.CallOpts, donId [32]byte) (ConfiguratorConfigurationState, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	PromoteStagingConfig(opts *bind.TransactOpts, donId [32]byte, isFlipped bool) (*types.Transaction, error)

	SetProductionConfig(opts *bind.TransactOpts, donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetStagingConfig(opts *bind.TransactOpts, donId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedConfiguratorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExposedConfiguratorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExposedConfiguratorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExposedConfiguratorOwnershipTransferred, error)

	FilterProductionConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ExposedConfiguratorProductionConfigSetIterator, error)

	WatchProductionConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorProductionConfigSet, configId [][32]byte) (event.Subscription, error)

	ParseProductionConfigSet(log types.Log) (*ExposedConfiguratorProductionConfigSet, error)

	FilterPromoteStagingConfig(opts *bind.FilterOpts, configId [][32]byte, retiredConfigDigest [][32]byte) (*ExposedConfiguratorPromoteStagingConfigIterator, error)

	WatchPromoteStagingConfig(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorPromoteStagingConfig, configId [][32]byte, retiredConfigDigest [][32]byte) (event.Subscription, error)

	ParsePromoteStagingConfig(log types.Log) (*ExposedConfiguratorPromoteStagingConfig, error)

	FilterStagingConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ExposedConfiguratorStagingConfigSetIterator, error)

	WatchStagingConfigSet(opts *bind.WatchOpts, sink chan<- *ExposedConfiguratorStagingConfigSet, configId [][32]byte) (event.Subscription, error)

	ParseStagingConfigSet(log types.Log) (*ExposedConfiguratorStagingConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
