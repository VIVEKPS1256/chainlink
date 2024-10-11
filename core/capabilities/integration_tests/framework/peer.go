package framework

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/mr-tron/base58"

	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
)

func getKeyBundlesAndPeerIDs(numNodes int) ([]ocr2key.KeyBundle, []peer, error) {
	var keyBundles []ocr2key.KeyBundle
	var donPeerIDs []peer
	for i := 0; i < numNodes; i++ {
		peerID := NewPeerID()

		keyBundle, err := ocr2key.New(chaintype.EVM)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create key bundle: %w", err)
		}

		keyBundles = append(keyBundles, keyBundle)

		pk := keyBundle.PublicKey()

		p := peer{
			PeerID: peerID,
			Signer: fmt.Sprintf("0x%x", pk),
		}

		donPeerIDs = append(donPeerIDs, p)
	}
	return keyBundles, donPeerIDs, nil
}

type peerWrapper struct {
	peer p2pPeer
}

func (t peerWrapper) Start(ctx context.Context) error {
	return nil
}

func (t peerWrapper) Close() error {
	return nil
}

func (t peerWrapper) Ready() error {
	return nil
}

func (t peerWrapper) HealthReport() map[string]error {
	return nil
}

func (t peerWrapper) Name() string {
	return "peerWrapper"
}

func (t peerWrapper) GetPeer() p2ptypes.Peer {
	return t.peer
}

type p2pPeer struct {
	id p2ptypes.PeerID
}

func (t p2pPeer) Start(ctx context.Context) error {
	return nil
}

func (t p2pPeer) Close() error {
	return nil
}

func (t p2pPeer) Ready() error {
	return nil
}

func (t p2pPeer) HealthReport() map[string]error {
	return nil
}

func (t p2pPeer) Name() string {
	return "p2pPeer"
}

func (t p2pPeer) ID() p2ptypes.PeerID {
	return t.id
}

func (t p2pPeer) UpdateConnections(peers map[p2ptypes.PeerID]p2ptypes.StreamConfig) error {
	return nil
}

func (t p2pPeer) Send(peerID p2ptypes.PeerID, msg []byte) error {
	return nil
}

func (t p2pPeer) Receive() <-chan p2ptypes.Message {
	return nil
}

func NewPeerID() string {
	var privKey [32]byte
	_, err := rand.Read(privKey[:])
	if err != nil {
		panic(err)
	}

	peerID := append(libp2pMagic(), privKey[:]...)

	return base58.Encode(peerID[:])
}

func libp2pMagic() []byte {
	return []byte{0x00, 0x24, 0x08, 0x01, 0x12, 0x20}
}

func ptr[T any](t T) *T { return &t }
