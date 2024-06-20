package client

import (
	"context"
	"sync"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/v2/common/types"
)

type TransactionSender[TX any] interface {
	SendTransaction(ctx context.Context, tx TX) (SendTxReturnCode, error)
}

// TxErrorClassifier - defines interface of a function that transforms raw RPC error into the SendTxReturnCode enum
// (e.g. Successful, Fatal, Retryable, etc.)
type TxErrorClassifier[TX any] func(tx TX, err error) SendTxReturnCode

// SendTxRPCClient - defines interface of an RPC used by TransactionSender to broadcast transaction
type SendTxRPCClient[TX any] interface {
	SendTransaction(ctx context.Context, tx TX) error
}

func NewTransactionSender[TX any, CHAIN_ID types.ID, RPC SendTxRPCClient[TX]](
	lggr logger.Logger,
	chainID CHAIN_ID,
	multiNode *MultiNode[CHAIN_ID, RPC],
	txErrorClassifier TxErrorClassifier[TX],
) TransactionSender[TX] {
	return &transactionSender[TX, CHAIN_ID, RPC]{
		chainID:           chainID,
		lggr:              logger.Sugared(lggr).Named("TransactionSender").With("chainID", chainID.String()),
		multiNode:         multiNode,
		txErrorClassifier: txErrorClassifier,
	}
}

type transactionSender[TX any, CHAIN_ID types.ID, RPC SendTxRPCClient[TX]] struct {
	chainID           CHAIN_ID
	lggr              logger.Logger
	multiNode         *MultiNode[CHAIN_ID, RPC]
	txErrorClassifier TxErrorClassifier[TX]
}

func (txSender *transactionSender[TX, CHAIN_ID, RPC]) SendTransaction(ctx context.Context, tx TX) (SendTxReturnCode, error) {
	txResults := make(chan SendTxReturnCode, len(txSender.multiNode.primaryNodes))
	txResultsToReport := make(chan SendTxReturnCode, len(txSender.multiNode.primaryNodes))
	primaryBroadcastWg := sync.WaitGroup{}

	err := txSender.multiNode.DoAll(ctx, func(ctx context.Context, rpc RPC, isSendOnly bool) bool {
		if isSendOnly {
			txSender.multiNode.wg.Add(1)
			go func() {
				defer txSender.multiNode.wg.Done()
				// Send-only nodes' results are ignored as they tend to return false-positive responses.
				// Broadcast to them is necessary to speed up the propagation of TX in the network.
				_ = rpc.SendTransaction(ctx, tx)
			}()
			return true
		}

		// Primary Nodes
		primaryBroadcastWg.Add(1)
		go func() {
			defer primaryBroadcastWg.Done()
			txErr := rpc.SendTransaction(ctx, tx)
			result := txSender.txErrorClassifier(tx, txErr)

			txResultsToReport <- result
			txResults <- result
		}()
		return true
	})

	// Wait for all sends to finish
	primaryBroadcastWg.Wait()

	if err != nil {
		return 0, err
	}

	// TODO: Collect Tx Results

	return 0, nil
}
