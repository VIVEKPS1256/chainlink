package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"

	commonclient "github.com/smartcontractkit/chainlink/v2/common/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
)

var _ TxmClient = (*evmTxmClient)(nil)

type evmTxmClient struct {
	client client.Client
}

func NewEvmTxmClient(c client.Client) *evmTxmClient {
	return &evmTxmClient{client: c}
}

func (c *evmTxmClient) PendingSequenceAt(ctx context.Context, addr common.Address) (evmtypes.Nonce, error) {
	return c.PendingNonceAt(ctx, addr)
}

func (c *evmTxmClient) ConfiguredChainID() *big.Int {
	return c.client.ConfiguredChainID()
}

func (c *evmTxmClient) BatchSendTransactions(
	ctx context.Context,
	attempts []TxAttempt,
	batchSize int,
	lggr logger.SugaredLogger,
) (
	codes []commonclient.SendTxReturnCode,
	txErrs []error,
	broadcastTime time.Time,
	successfulTxIDs []int64,
	err error,
) {
	// preallocate
	codes = make([]commonclient.SendTxReturnCode, len(attempts))
	txErrs = make([]error, len(attempts))

	reqs, broadcastTime, successfulTxIDs, batchErr := batchSendTransactions(ctx, attempts, batchSize, lggr, c.client)
	err = errors.Join(err, batchErr) // this error does not block processing

	// safety check - exits before processing
	if len(reqs) != len(attempts) {
		lenErr := fmt.Errorf("Returned request data length (%d) != number of tx attempts (%d)", len(reqs), len(attempts))
		err = errors.Join(err, lenErr)
		lggr.Criticalw("Mismatched length", "err", err)
		return
	}

	// for each batched tx convert response to standard error code
	var wg sync.WaitGroup
	wg.Add(len(reqs))
	processingErr := make([]error, len(attempts))
	for index := range reqs {
		go func(i int) {
			defer wg.Done()

			// convert to tx for logging purposes - exits early if error occurs
			tx, signedErr := GetGethSignedTx(attempts[i].SignedRawTx)
			if signedErr != nil {
				signedErrMsg := fmt.Sprintf("failed to process tx (index %d)", i)
				lggr.Errorw(signedErrMsg, "err", signedErr)
				processingErr[i] = fmt.Errorf("%s: %w", signedErrMsg, signedErr)
				return
			}
			sendErr := reqs[i].Error
			codes[i] = client.ClassifySendError(sendErr, lggr, tx, attempts[i].Tx.FromAddress, c.client.IsL2())
			txErrs[i] = sendErr
		}(index)
	}
	wg.Wait()
	err = errors.Join(err, errors.Join(processingErr...)) // merge errors together
	return
}

func (c *evmTxmClient) SendTransactionReturnCode(ctx context.Context, etx Tx, attempt TxAttempt, lggr logger.SugaredLogger) (commonclient.SendTxReturnCode, error) {
	signedTx, err := GetGethSignedTx(attempt.SignedRawTx)
	if err != nil {
		lggr.Criticalw("Fatal error signing transaction", "err", err, "etx", etx)
		return commonclient.Fatal, err
	}
	return c.client.SendTransactionReturnCode(ctx, signedTx, etx.FromAddress)
}

func (c *evmTxmClient) PendingNonceAt(ctx context.Context, fromAddress common.Address) (n evmtypes.Nonce, err error) {
	nextNonce, err := c.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return n, err
	}

	if nextNonce > math.MaxInt64 {
		return n, fmt.Errorf("nonce overflow, got: %v", nextNonce)
	}
	return evmtypes.Nonce(nextNonce), nil
}

func (c *evmTxmClient) SequenceAt(ctx context.Context, addr common.Address, blockNum *big.Int) (evmtypes.Nonce, error) {
	return c.client.SequenceAt(ctx, addr, blockNum)
}

func (c *evmTxmClient) BatchGetReceipts(ctx context.Context, attempts []TxAttempt) (txReceipt []*evmtypes.Receipt, txErr []error, funcErr error) {
	var reqs []rpc.BatchElem
	reqs, txReceipt, txErr = newGetBatchReceiptsReq(attempts)

	if err := c.client.BatchCallContext(ctx, reqs); err != nil {
		return nil, nil, fmt.Errorf("EthConfirmer#batchFetchReceipts error fetching receipts with BatchCallContext: %w", err)
	}

	for i, req := range reqs {
		txErr[i] = req.Error
	}
	return txReceipt, txErr, nil
}

func newGetBatchReceiptsReq(attempts []TxAttempt) (reqs []rpc.BatchElem, txReceipts []*evmtypes.Receipt, txErrs []error) {
	reqs = make([]rpc.BatchElem, len(attempts))
	txReceipts = make([]*evmtypes.Receipt, len(attempts))
	txErrs = make([]error, len(attempts))
	for i, attempt := range attempts {
		res := &evmtypes.Receipt{}
		req := rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{attempt.Hash},
			Result: res,
		}
		txReceipts[i] = res
		reqs[i] = req
	}

	return reqs, txReceipts, txErrs
}

func (c *evmTxmClient) FinalizedBlockHash(ctx context.Context) (common.Hash, *big.Int, error) {
	head, err := c.client.FinalizedBlock(ctx)
	if err != nil || head == nil {
		return common.Hash{}, big.NewInt(0), err
	}

	return head.Hash, big.NewInt(head.BlockNumber()), nil
}

// TODO: test me
func (c *evmTxmClient) BatchGetReceiptsWithFinalizedBlock(ctx context.Context, attempts []TxAttempt, useFinalityTag bool, finalityDepth uint32) (
	finalizedBlock *big.Int, txReceipt []*evmtypes.Receipt, txErr []error, funcErr error) {

	var reqs []rpc.BatchElem
	reqs, txReceipt, txErr = newGetBatchReceiptsReq(attempts)

	blockNumber := rpc.LatestBlockNumber
	if useFinalityTag {
		blockNumber = rpc.FinalizedBlockNumber
	}

	var head evmtypes.Head
	blockRequest := rpc.BatchElem{
		Method: "eth_getBlockByNumber",
		Args:   []interface{}{blockNumber, false},
		Result: &head,
	}
	reqs = append(reqs, blockRequest)

	if err := c.client.BatchCallContext(ctx, reqs); err != nil {
		return nil, nil, nil, fmt.Errorf("BatchGetReceiptsWithFinalizedBlock error fetching receipts with BatchCallContext: %w", err)
	}

	if blockRequest.Error != nil {
		return nil, nil, nil, fmt.Errorf("failed to fetch finalized block with BatchCallContext: %w", blockRequest.Error)
	}

	for i := range txErr {
		txErr[i] = reqs[i].Error
	}

	finalizedBlock = big.NewInt(head.BlockNumber() - int64(finalityDepth))
	if useFinalityTag {
		finalizedBlock = big.NewInt(head.BlockNumber())
	}
	return finalizedBlock, txReceipt, txErr, nil
}

// sendEmptyTransaction sends a transaction with 0 Eth and an empty payload to the burn address
// May be useful for clearing stuck nonces
func (c *evmTxmClient) SendEmptyTransaction(
	ctx context.Context,
	newTxAttempt func(seq evmtypes.Nonce, feeLimit uint32, fee gas.EvmFee, fromAddress common.Address) (attempt TxAttempt, err error),
	seq evmtypes.Nonce,
	gasLimit uint32,
	fee gas.EvmFee,
	fromAddress common.Address,
) (txhash string, err error) {
	defer utils.WrapIfError(&err, "sendEmptyTransaction failed")

	attempt, err := newTxAttempt(seq, gasLimit, fee, fromAddress)
	if err != nil {
		return txhash, err
	}

	signedTx, err := GetGethSignedTx(attempt.SignedRawTx)
	if err != nil {
		return txhash, err
	}

	_, err = c.client.SendTransactionReturnCode(ctx, signedTx, fromAddress)
	return signedTx.Hash().String(), err
}

func (c *evmTxmClient) CallContract(ctx context.Context, a TxAttempt, blockNumber *big.Int) (rpcErr fmt.Stringer, extractErr error) {
	_, errCall := c.client.CallContract(ctx, ethereum.CallMsg{
		From:       a.Tx.FromAddress,
		To:         &a.Tx.ToAddress,
		Gas:        uint64(a.Tx.FeeLimit),
		GasPrice:   a.TxFee.Legacy.ToInt(),
		GasFeeCap:  a.TxFee.DynamicFeeCap.ToInt(),
		GasTipCap:  a.TxFee.DynamicTipCap.ToInt(),
		Value:      nil,
		Data:       a.Tx.EncodedPayload,
		AccessList: nil,
	}, blockNumber)
	return client.ExtractRPCError(errCall)
}
