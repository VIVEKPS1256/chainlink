package rollups

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml/daoracle"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/rollups/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestOPL1Oracle_CalculateCustomCalldataGasPrice(t *testing.T) {
	oracleAddress := common.HexToAddress("0x0000000000000000000000000000000044433322").String()

	t.Parallel()

	t.Run("correctly fetches gas price if chain has custom calldata", func(t *testing.T) {
		ethClient := mocks.NewL1OracleClient(t)
		expectedPriceHex := "0x0000000000000000000000000000000000000000000000000000000000000032"

		daOracle := CreateTestDAOracle(t, daoracle.OPStack, oracleAddress, "0x0000000000000000000000000000000000001234")
		oracle := NewCustomCalldataDAOracle(logger.Test(t), ethClient, daOracle)

		ethClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			require.NotNil(t, callMsg.To)
			require.Equal(t, oracleAddress, callMsg.To.String())
			require.Nil(t, blockNumber)
		}).Return(hexutil.MustDecode(expectedPriceHex), nil).Once()

		price, err := oracle.getCustomCalldataGasPrice(tests.Context(t))
		require.NoError(t, err)
		require.Equal(t, big.NewInt(50), price)
	})

	t.Run("throws error if custom calldata fails to decode", func(t *testing.T) {
		ethClient := mocks.NewL1OracleClient(t)

		daOracle := CreateTestDAOracle(t, daoracle.OPStack, oracleAddress, "0xblahblahblah")
		oracle := NewCustomCalldataDAOracle(logger.Test(t), ethClient, daOracle)

		_, err := oracle.getCustomCalldataGasPrice(tests.Context(t))
		require.Error(t, err)
	})
}
