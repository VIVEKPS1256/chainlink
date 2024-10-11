package framework

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	ragetypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/feeds_consumer"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/forwarder"

	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
)

func setupForwarderContract(t *testing.T, workflowDon DonConfiguration,
	backend *ethBlockchain) (common.Address, *forwarder.KeystoneForwarder) {
	addr, _, fwd, err := forwarder.DeployKeystoneForwarder(backend.transactionOpts, backend)
	require.NoError(t, err)
	backend.Commit()

	var signers []common.Address
	for _, p := range workflowDon.peerIDs {
		signers = append(signers, common.HexToAddress(p.Signer))
	}

	_, err = fwd.SetConfig(backend.transactionOpts, workflowDon.ID, workflowDon.ConfigVersion, workflowDon.F, signers)
	require.NoError(t, err)
	backend.Commit()

	return addr, fwd
}

func setupConsumerContract(t *testing.T, backend *ethBlockchain,
	forwarderAddress common.Address, workflowOwner string, workflowName string) (common.Address, *feeds_consumer.KeystoneFeedsConsumer) {
	addr, _, consumer, err := feeds_consumer.DeployKeystoneFeedsConsumer(backend.transactionOpts, backend)
	require.NoError(t, err)
	backend.Commit()

	var nameBytes [10]byte
	copy(nameBytes[:], workflowName)

	ownerAddr := common.HexToAddress(workflowOwner)

	_, err = consumer.SetConfig(backend.transactionOpts, []common.Address{forwarderAddress}, []common.Address{ownerAddr}, [][10]byte{nameBytes})
	require.NoError(t, err)

	backend.Commit()

	return addr, consumer
}

type peer struct {
	PeerID string
	Signer string
}

func peerIDToBytes(peerID string) ([32]byte, error) {
	var peerIDB ragetypes.PeerID
	err := peerIDB.UnmarshalText([]byte(peerID))
	if err != nil {
		return [32]byte{}, err
	}

	return peerIDB, nil
}

func peers(ps []peer) ([][32]byte, error) {
	out := [][32]byte{}
	for _, p := range ps {
		b, err := peerIDToBytes(p.PeerID)
		if err != nil {
			return nil, err
		}

		out = append(out, b)
	}

	return out, nil
}

func peerToNode(nopID uint32, p peer) (kcr.CapabilitiesRegistryNodeParams, error) {
	peerIDB, err := peerIDToBytes(p.PeerID)
	if err != nil {
		return kcr.CapabilitiesRegistryNodeParams{}, fmt.Errorf("failed to convert peerID: %w", err)
	}

	sig := strings.TrimPrefix(p.Signer, "0x")
	signerB, err := hex.DecodeString(sig)
	if err != nil {
		return kcr.CapabilitiesRegistryNodeParams{}, fmt.Errorf("failed to convert signer: %w", err)
	}

	var sigb [32]byte
	copy(sigb[:], signerB)

	return kcr.CapabilitiesRegistryNodeParams{
		NodeOperatorId: nopID,
		P2pId:          peerIDB,
		Signer:         sigb,
	}, nil
}
