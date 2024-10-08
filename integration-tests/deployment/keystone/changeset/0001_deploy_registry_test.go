package changeset_test

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/keystone/changeset"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/memory"
)

func TestApply0001(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	ab := deployment.NewMemoryAddressBook()
	cfg := memory.MemoryEnvironmentConfig{
		Nodes:  1,
		Chains: 2,
	}
	env := memory.NewMemoryEnvironment(t, lggr, zapcore.DebugLevel, cfg)

	registrySel := env.AllChainSelectors()[0]
	resp, err := changeset.Apply0001(lggr, env, ab, registrySel)
	require.NoError(t, err)
	require.NotNil(t, resp)
	// capabilities registry should be deployed on chain 0
	addrs, err := resp.AddressBook.AddressesForChain(registrySel)
	require.NoError(t, err)
	require.Len(t, addrs, 1)

	// no capabilities registry on chain 1
	_, err = resp.AddressBook.AddressesForChain(env.AllChainSelectors()[1])
	require.Error(t, err)

}
