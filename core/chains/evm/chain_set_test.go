package evm_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	evmclimocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	configtest "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/configtest/v2"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestChainRelayExtenders(t *testing.T) {
	t.Parallel()

	newId := testutils.NewRandomEVMChainID()
	cfg := configtest.NewGeneralConfig(t, func(c *chainlink.Config, s *chainlink.Secrets) {
		one := uint32(1)
		c.EVM[0].MinIncomingConfirmations = &one
		t := true
		c.EVM = append(c.EVM, &toml.EVMConfig{ChainID: utils.NewBig(newId), Enabled: &t, Chain: toml.Defaults(nil)})
	})
	db := pgtest.NewSqlxDB(t)
	kst := cltest.NewKeyStore(t, db, cfg.Database())
	require.NoError(t, kst.Unlock(cltest.Password))

	opts := evmtest.NewChainRelayExtOpts(t, evmtest.TestChainOpts{DB: db, KeyStore: kst.Eth(), GeneralConfig: cfg})
	opts.GenEthClient = func(*big.Int) evmclient.Client {
		return cltest.NewEthMocksWithStartupAssertions(t)
	}
	chainSet, err := evm.NewChainRelayerExtenders(testutils.Context(t), opts)
	require.NoError(t, err)

	require.Len(t, chainSet, 2)
	for _, c := range chainSet {
		require.NoError(t, c.Start(testutils.Context(t)))
		require.NoError(t, c.Ready())
	}

	require.NotEqual(t, chainSet[0].Chain().ID().String(), chainSet[1].Chain().ID().String())

	for _, c := range chainSet {
		require.NoError(t, c.Close())
	}

	chainSet[0].Chain().Client().(*evmclimocks.Client).AssertCalled(t, "Close")
	chainSet[1].Chain().Client().(*evmclimocks.Client).AssertCalled(t, "Close")

	assert.Error(t, chainSet[0].Chain().Ready())
	assert.Error(t, chainSet[1].Chain().Ready())
}
