package keystone

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/forwarder"
)

type KeystoneForwarderDeployer struct {
	lggr     logger.Logger
	contract *forwarder.KeystoneForwarder
}

func (c *KeystoneForwarderDeployer) deploy(req deployRequest) (*deployResponse, error) {
	forwarderAddr, tx, forwarder, err := forwarder.DeployKeystoneForwarder(
		req.Chain.DeployerKey,
		req.Chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy KeystoneForwarder: %w", err)
	}
	resp := &deployResponse{
		Address: forwarderAddr,
		Tx:      tx.Hash(),
		Tv: deployment.TypeAndVersion{
			Type:    KeystoneForwarder,
			Version: deployment.Version1_0_0,
		},
	}
	_, err = req.Chain.Confirm(tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to confirm and save KeystoneForwarder: %w", err)
	}
	c.contract = forwarder
	return resp, nil
}
