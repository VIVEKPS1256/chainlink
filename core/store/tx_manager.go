package store

import (
	"bytes"
	"fmt"
	"math/big"
	"regexp"
	"sync"

	"github.com/pkg/errors"

	"chainlink/core/assets"
	"chainlink/core/eth"
	"chainlink/core/logger"
	"chainlink/core/store/models"
	"chainlink/core/store/orm"
	"chainlink/core/utils"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tevino/abool"
	"go.uber.org/multierr"
	"gopkg.in/guregu/null.v3"
)

// DefaultGasLimit sets the default gas limit for outgoing transactions.
// if updating DefaultGasLimit, be sure it matches with the
// DefaultGasLimit specified in evm/test/Oracle_test.js
const DefaultGasLimit uint64 = 500000
const nonceReloadLimit int = 1

// ErrPendingConnection is the error returned if TxManager is not connected.
var ErrPendingConnection = errors.New("Cannot talk to chain, pending connection")

// TxManager represents an interface for interacting with the blockchain
type TxManager interface {
	HeadTrackable
	Connected() bool
	Register(accounts []accounts.Account)

	CreateTx(to common.Address, data []byte) (*models.Tx, error)
	CreateTxWithGas(surrogateID null.String, to common.Address, data []byte, gasPriceWei *big.Int, gasLimit uint64) (*models.Tx, error)
	CreateTxWithEth(from, to common.Address, value *assets.Eth) (*models.Tx, error)
	CheckAttempt(txAttempt *models.TxAttempt, blockHeight uint64) (*eth.TxReceipt, AttemptState, error)

	BumpGasUntilSafe(hash common.Hash) (*eth.TxReceipt, AttemptState, error)

	ContractLINKBalance(wr models.WithdrawalRequest) (assets.Link, error)
	WithdrawLINK(wr models.WithdrawalRequest) (common.Hash, error)
	GetLINKBalance(address common.Address) (*assets.Link, error)
	NextActiveAccount() *ManagedAccount

	GetEthBalance(address common.Address) (*assets.Eth, error)
	SubscribeToNewHeads(channel chan<- eth.BlockHeader) (eth.Subscription, error)
	GetBlockByNumber(hex string) (eth.BlockHeader, error)
	eth.LogSubscriber
	GetTxReceipt(common.Hash) (*eth.TxReceipt, error)
	GetChainID() (*big.Int, error)
}

//go:generate mockery -name TxManager -output ../internal/mocks/ -case=underscore

// EthTxManager contains fields for the Ethereum client, the KeyStore,
// the local Config for the application, and the database.
type EthTxManager struct {
	eth.Client
	keyStore            *KeyStore
	config              orm.ConfigReader
	orm                 *orm.ORM
	registeredAccounts  []accounts.Account
	availableAccounts   []*ManagedAccount
	availableAccountIdx int
	accountsMutex       *sync.Mutex
	connected           *abool.AtomicBool
	currentHead         models.Head
}

// NewEthTxManager constructs an EthTxManager using the passed variables and
// initializing internal variables.
func NewEthTxManager(client eth.Client, config orm.ConfigReader, keyStore *KeyStore, orm *orm.ORM) *EthTxManager {
	return &EthTxManager{
		Client:        client,
		config:        config,
		keyStore:      keyStore,
		orm:           orm,
		accountsMutex: &sync.Mutex{},
		connected:     abool.New(),
	}
}

// Register activates accounts for outgoing transactions and client side
// nonce management.
func (txm *EthTxManager) Register(accts []accounts.Account) {
	txm.accountsMutex.Lock()
	defer txm.accountsMutex.Unlock()

	cp := make([]accounts.Account, len(accts))
	copy(cp, accts)
	txm.registeredAccounts = cp
}

// Connected returns a bool indicating whether or not it is connected.
func (txm *EthTxManager) Connected() bool {
	return txm.connected.IsSet()
}

// Connect iterates over the available accounts to retrieve their nonce
// for client side management.
func (txm *EthTxManager) Connect(bn *models.Head) error {
	var merr error
	func() {
		txm.accountsMutex.Lock()
		defer txm.accountsMutex.Unlock()

		txm.availableAccounts = []*ManagedAccount{}
		for _, a := range txm.registeredAccounts {
			ma, err := txm.activateAccount(a)
			merr = multierr.Append(merr, err)
			if err == nil {
				txm.availableAccounts = append(txm.availableAccounts, ma)
			}
		}

		if bn != nil {
			txm.currentHead = *bn
		}
		txm.connected.Set()
	}()

	// Upon connecting/reconnecting, rebroadcast any transactions that are still unconfirmed
	attempts, err := txm.orm.UnconfirmedTxAttempts()
	if err != nil {
		merr = multierr.Append(merr, err)
		return merr
	}

	attempts = models.HighestPricedTxAttemptPerTx(attempts)

	for _, attempt := range attempts {
		ma := txm.getAccount(attempt.Tx.From)
		if ma == nil {
			logger.Warnf("Trying to rebroadcast tx %v, could not find account %v", attempt.Hash.Hex(), attempt.Tx.From.Hex())
			continue
		} else if ma.Nonce() > attempt.Tx.Nonce {
			// Do not rebroadcast txs with nonces that are lower than our current nonce
			continue
		}

		logger.Infof("Rebroadcasting tx %v", attempt.Hash.Hex())

		_, err = txm.SendRawTx(attempt.SignedRawTx)
		if err != nil && !isNonceTooLowError(err) {
			logger.Warnf("Failed to rebroadcast tx %v: %v", attempt.Hash.Hex(), err)
		}
	}

	return merr
}

// Disconnect marks this instance as disconnected.
func (txm *EthTxManager) Disconnect() {
	txm.connected.UnSet()
}

// OnNewHead does nothing; exists to comply with interface.
func (txm *EthTxManager) OnNewHead(head *models.Head) {
	txm.currentHead = *head
}

// CreateTx signs and sends a transaction to the Ethereum blockchain.
func (txm *EthTxManager) CreateTx(to common.Address, data []byte) (*models.Tx, error) {
	return txm.CreateTxWithGas(null.String{}, to, data, txm.config.EthGasPriceDefault(), DefaultGasLimit)
}

// CreateTxWithGas signs and sends a transaction to the Ethereum blockchain.
func (txm *EthTxManager) CreateTxWithGas(surrogateID null.String, to common.Address, data []byte, gasPriceWei *big.Int, gasLimit uint64) (*models.Tx, error) {
	ma, err := txm.nextAccount()
	if err != nil {
		return nil, err
	}

	gasPriceWei, gasLimit = normalizeGasParams(gasPriceWei, gasLimit, txm.config)
	return txm.createTx(surrogateID, ma, to, data, gasPriceWei, gasLimit, nil)
}

// CreateTxWithEth signs and sends a transaction with some ETH to transfer.
func (txm *EthTxManager) CreateTxWithEth(from, to common.Address, value *assets.Eth) (*models.Tx, error) {
	ma := txm.getAccount(from)
	if ma == nil {
		return nil, errors.New("account does not exist")
	}

	return txm.createTx(null.String{}, ma, to, []byte{}, txm.config.EthGasPriceDefault(), DefaultGasLimit, value)
}

func (txm *EthTxManager) nextAccount() (*ManagedAccount, error) {
	if !txm.Connected() {
		return nil, errors.Wrap(ErrPendingConnection, "EthTxManager#nextAccount")
	}

	ma := txm.NextActiveAccount()
	if ma == nil {
		return nil, errors.New("Must connect and activate an account before creating a transaction")
	}

	return ma, nil
}

func normalizeGasParams(gasPriceWei *big.Int, gasLimit uint64, config orm.ConfigReader) (*big.Int, uint64) {
	if !config.Dev() {
		return config.EthGasPriceDefault(), DefaultGasLimit
	}

	if gasPriceWei == nil {
		gasPriceWei = config.EthGasPriceDefault()
	}

	if gasLimit == 0 {
		gasLimit = DefaultGasLimit
	}

	return gasPriceWei, gasLimit
}

// createTx creates an ethereum transaction, and retries to submit the
// transaction if a nonce too low error is returned
func (txm *EthTxManager) createTx(
	surrogateID null.String,
	ma *ManagedAccount,
	to common.Address,
	data []byte,
	gasPriceWei *big.Int,
	gasLimit uint64,
	value *assets.Eth) (*models.Tx, error) {

	for nrc := 0; nrc <= nonceReloadLimit; nrc++ {
		tx, err := txm.sendInitialTx(surrogateID, ma, to, data, gasPriceWei, gasLimit, value)
		if err == nil {
			return tx, nil
		}

		if !isNonceTooLowError(err) {
			return nil, errors.Wrap(err, "TxManager#retryInitialTx sendInitialTx")
		}

		logger.Warnw(
			"Tx #0: nonce too low, retrying with network nonce",
			"nonce", tx.Nonce, "error", err.Error(),
		)

		err = ma.ReloadNonce(txm)
		if err != nil {
			return nil, errors.Wrap(err, "TxManager#retryInitialTx ReloadNonce")
		}
	}

	return nil, fmt.Errorf(
		"Transaction reattempt limit reached for 'nonce is too low' error. Limit: %v",
		nonceReloadLimit,
	)
}

// sendInitialTx creates the initial Tx record + attempt for an Ethereum Tx,
// there should only ever be one of those for a "job"
func (txm *EthTxManager) sendInitialTx(
	surrogateID null.String,
	ma *ManagedAccount,
	to common.Address,
	data []byte,
	gasPriceWei *big.Int,
	gasLimit uint64,
	value *assets.Eth) (*models.Tx, error) {

	var err error
	var tx *models.Tx

	err = ma.GetAndIncrementNonce(func(nonce uint64) error {
		blockHeight := uint64(txm.currentHead.Number)
		tx, err = txm.newTx(
			ma.Account,
			nonce,
			to,
			value.ToInt(),
			gasLimit,
			gasPriceWei,
			data,
			&ma.Address,
			blockHeight,
		)
		if err != nil {
			return errors.Wrap(err, "TxManager#sendInitialTx newTx")
		}

		tx.SurrogateID = surrogateID
		tx, err = txm.orm.CreateTx(tx)
		if err != nil {
			return errors.Wrap(err, "TxManager#sendInitialTx CreateTx")
		}

		_, err = txm.SendRawTx(tx.SignedRawTx)
		if err != nil {
			return errors.Wrap(err, "TxManager#sendInitialTx SendRawTx")
		}

		txAttempt, err := txm.orm.AddTxAttempt(tx, tx)
		if err != nil {
			return errors.Wrap(err, "TxManager#sendInitialTx AddTxAttempt")
		}

		logger.Debugw("Added Tx attempt #0", "txID", tx.ID, "txAttemptID", txAttempt.ID)

		return nil
	})

	return tx, err
}

var (
	nonceTooLowRegex = regexp.MustCompile("(nonce .*too low|same hash was already imported)")
)

func isNonceTooLowError(err error) bool {
	return err != nil && nonceTooLowRegex.MatchString(err.Error())
}

// newTx returns a newly signed Ethereum Transaction
func (txm *EthTxManager) newTx(
	account accounts.Account,
	nonce uint64,
	to common.Address,
	amount *big.Int,
	gasLimit uint64,
	gasPrice *big.Int,
	data []byte,
	from *common.Address,
	sentAt uint64) (*models.Tx, error) {

	transaction := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	transaction, err := txm.keyStore.SignTx(account, transaction, txm.config.ChainID())
	if err != nil {
		return nil, errors.Wrap(err, "TxManager newTx.SignTx")
	}

	rlp := new(bytes.Buffer)
	if err := transaction.EncodeRLP(rlp); err != nil {
		return nil, errors.Wrap(err, "TxManager newTx.EncodeRLP")
	}

	return &models.Tx{
		From:        *from,
		SentAt:      sentAt,
		To:          *transaction.To(),
		Nonce:       transaction.Nonce(),
		Data:        transaction.Data(),
		Value:       utils.NewBig(transaction.Value()),
		GasLimit:    transaction.Gas(),
		GasPrice:    utils.NewBig(transaction.GasPrice()),
		Hash:        transaction.Hash(),
		SignedRawTx: hexutil.Encode(rlp.Bytes()),
	}, nil
}

// GetLINKBalance returns the balance of LINK at the given address
func (txm *EthTxManager) GetLINKBalance(address common.Address) (*assets.Link, error) {
	contractAddress := common.HexToAddress(txm.config.LinkContractAddress())
	balance, err := txm.GetERC20Balance(address, contractAddress)
	if err != nil {
		return assets.NewLink(0), err
	}
	return (*assets.Link)(balance), nil
}

// BumpGasUntilSafe process a collection of related TxAttempts, trying to get
// at least one TxAttempt into a safe state, bumping gas if needed
func (txm *EthTxManager) BumpGasUntilSafe(hash common.Hash) (*eth.TxReceipt, AttemptState, error) {
	tx, _, err := txm.orm.FindTxByAttempt(hash)
	if err != nil {
		return nil, Unknown, errors.Wrap(err, "BumpGasUntilSafe FindTxByAttempt")
	}

	receipt, state, err := txm.checkChainForConfirmation(tx)
	if err != nil || state != Unconfirmed {
		return receipt, state, err
	}

	return txm.checkAccountForConfirmation(tx)
}

func (txm *EthTxManager) checkChainForConfirmation(tx *models.Tx) (*eth.TxReceipt, AttemptState, error) {
	blockHeight := uint64(txm.currentHead.Number)

	var merr error
	// Process attempts in reverse, since the attempt with the highest gas is
	// likely to be confirmed first
	for attemptIndex := len(tx.Attempts) - 1; attemptIndex >= 0; attemptIndex-- {
		receipt, state, err := txm.processAttempt(tx, attemptIndex, blockHeight)
		if state == Safe || state == Confirmed {
			return receipt, state, err // success, so all other attempt errors can be ignored.
		}
		merr = multierr.Append(merr, err)
	}

	return nil, Unconfirmed, merr
}

func (txm *EthTxManager) checkAccountForConfirmation(tx *models.Tx) (*eth.TxReceipt, AttemptState, error) {
	ma := txm.GetAvailableAccount(tx.From)

	if ma != nil && ma.lastSafeNonce > tx.Nonce {
		tx.Confirmed = true
		tx.Hash = utils.EmptyHash
		if err := txm.orm.SaveTx(tx); err != nil {
			return nil, Safe, fmt.Errorf("BumpGasUntilSafe error saving Tx confirmation to the database")
		}
		return nil, Safe, fmt.Errorf("BumpGasUntilSafe a version of the Ethereum Transaction from %v with nonce %v", tx.From, tx.Nonce)
	}

	return nil, Unconfirmed, nil
}

// GetAvailableAccount retrieves a managed account if it one matches the address given.
func (txm *EthTxManager) GetAvailableAccount(from common.Address) *ManagedAccount {
	for _, a := range txm.availableAccounts {
		if a.Address == from {
			return a
		}
	}
	return nil
}

// ContractLINKBalance returns the balance for the contract associated with this
// withdrawal request, or any errors
func (txm *EthTxManager) ContractLINKBalance(wr models.WithdrawalRequest) (assets.Link, error) {
	contractAddress := &wr.ContractAddress
	if (*contractAddress == common.Address{}) {
		if txm.config.OracleContractAddress() == nil {
			return assets.Link{}, errors.New(
				"OracleContractAddress not set; cannot check LINK balance")
		}
		contractAddress = txm.config.OracleContractAddress()
	}

	linkBalance, err := txm.GetLINKBalance(*contractAddress)
	if err != nil {
		return assets.Link{}, multierr.Combine(
			fmt.Errorf("Could not check LINK balance for %v",
				contractAddress),
			err)
	}
	return *linkBalance, nil
}

// GetETHAndLINKBalances attempts to retrieve the ethereum node's perception of
// the latest ETH and LINK balances for the active account on the txm, or an
// error on failure.
func (txm *EthTxManager) GetETHAndLINKBalances(address common.Address) (*assets.Eth, *assets.Link, error) {
	linkBalance, linkErr := txm.GetLINKBalance(address)
	ethBalance, ethErr := txm.GetEthBalance(address)
	merr := multierr.Append(linkErr, ethErr)
	return ethBalance, linkBalance, merr
}

// WithdrawLINK withdraws the given amount of LINK from the contract to the
// configured withdrawal address. If wr.ContractAddress is empty (zero address),
// funds are withdrawn from configured OracleContractAddress.
func (txm *EthTxManager) WithdrawLINK(wr models.WithdrawalRequest) (common.Hash, error) {
	oracle, err := eth.GetContract("Oracle")
	if err != nil {
		return common.Hash{}, err
	}

	data, err := oracle.EncodeMessageCall("withdraw", wr.DestinationAddress, (*big.Int)(wr.Amount))
	if err != nil {
		return common.Hash{}, err
	}

	contractAddress := &wr.ContractAddress
	if (*contractAddress == common.Address{}) {
		if txm.config.OracleContractAddress() == nil {
			return common.Hash{}, errors.New(
				"OracleContractAddress not set; cannot withdraw")
		}
		contractAddress = txm.config.OracleContractAddress()
	}

	tx, err := txm.CreateTx(*contractAddress, data)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash, nil
}

// CheckAttempt retrieves a receipt for a TxAttempt, and check if it meets the
// minimum number of confirmations
func (txm *EthTxManager) CheckAttempt(txAttempt *models.TxAttempt, blockHeight uint64) (*eth.TxReceipt, AttemptState, error) {
	receipt, err := txm.GetTxReceipt(txAttempt.Hash)
	if err != nil {
		return nil, Unknown, errors.Wrap(err, "CheckAttempt GetTxReceipt failed")
	}

	if receipt.Unconfirmed() {
		return receipt, Unconfirmed, nil
	}

	minimumConfirmations := new(big.Int).SetUint64(txm.config.MinOutgoingConfirmations())
	confirmedAt := new(big.Int).Add(minimumConfirmations, receipt.BlockNumber.ToInt())

	confirmedAt.Sub(confirmedAt, big.NewInt(1)) // confirmed at block counts as 1 conf

	if new(big.Int).SetUint64(blockHeight).Cmp(confirmedAt) == -1 {
		return receipt, Confirmed, nil
	}

	return receipt, Safe, nil
}

// AttemptState enumerates the possible states of a transaction attempt as it
// gets accepted and confirmed by the blockchain
type AttemptState int

const (
	// Unknown is returned when the state of a transaction could not be
	// determined because of an error
	Unknown AttemptState = iota
	// Unconfirmed means that a transaction has had no confirmations at all
	Unconfirmed
	// Confirmed means that a transaftion has had at least one transaction, but
	// not enough to satisfy the minimum number of confirmations configuration
	// option
	Confirmed
	// Safe has the required number of confirmations or more
	Safe
)

// String conforms to the Stringer interface for AttemptState
func (a AttemptState) String() string {
	switch a {
	case Unconfirmed:
		return "unconfirmed"
	case Confirmed:
		return "confirmed"
	case Safe:
		return "safe"
	default:
		return "unknown"
	}
}

// processAttempt checks the state of a transaction attempt on the blockchain
// and decides if it is safe, needs bumping or more confirmations are needed to
// decide
func (txm *EthTxManager) processAttempt(
	tx *models.Tx,
	attemptIndex int,
	blockHeight uint64,
) (*eth.TxReceipt, AttemptState, error) {
	txAttempt := tx.Attempts[attemptIndex]

	receipt, state, err := txm.CheckAttempt(txAttempt, blockHeight)

	switch state {
	case Safe:
		txm.updateLastSafeNonce(tx)
		return receipt, state, txm.handleSafe(tx, attemptIndex)

	case Confirmed:
		logger.Debugw(
			fmt.Sprintf("Tx #%d is %s", attemptIndex, state),
			"txHash", txAttempt.Hash.String(),
			"txID", txAttempt.TxID,
			"receiptBlockNumber", receipt.BlockNumber.ToInt(),
			"currentBlockNumber", blockHeight,
			"receiptHash", receipt.Hash.Hex(),
		)

		return receipt, state, nil

	case Unconfirmed:
		attemptLimit := txm.config.TxAttemptLimit()
		if attemptIndex >= int(attemptLimit) {
			logger.Warnw(
				fmt.Sprintf("Tx #%d is %s, has met TxAttemptLimit", attemptIndex, state),
				"txAttemptLimit", attemptLimit,
				"txHash", txAttempt.Hash.String(),
				"txID", txAttempt.TxID,
			)
			return receipt, state, nil
		}

		if isLatestAttempt(tx, attemptIndex) && txm.hasTxAttemptMetGasBumpThreshold(tx, attemptIndex, blockHeight) {
			logger.Debugw(
				fmt.Sprintf("Tx #%d is %s, bumping gas", attemptIndex, state),
				"txHash", txAttempt.Hash.String(),
				"txID", txAttempt.TxID,
				"currentBlockNumber", blockHeight,
			)
			err = txm.bumpGas(tx, attemptIndex, blockHeight)
		} else {
			logger.Debugw(
				fmt.Sprintf("Tx #%d is %s", attemptIndex, state),
				"txHash", txAttempt.Hash.String(),
				"txID", txAttempt.TxID,
			)
		}

		return receipt, state, err

	default:
		logger.Debugw(
			fmt.Sprintf("Tx #%d is %s, error fetching receipt", attemptIndex, state),
			"txHash", txAttempt.Hash.String(),
			"txID", txAttempt.TxID,
			"error", err,
		)
		return nil, Unknown, errors.Wrap(err, "processAttempt CheckAttempt failed")
	}
}

func (txm *EthTxManager) updateLastSafeNonce(tx *models.Tx) {
	for _, a := range txm.availableAccounts {
		if tx.From == a.Address {
			a.updateLastSafeNonce(tx.Nonce)
		}
	}
}

// hasTxAttemptMetGasBumpThreshold returns true if the current block height
// exceeds the configured gas bump threshold, indicating that it is time for a
// new transaction attempt to be created with an increased gas price
func (txm *EthTxManager) hasTxAttemptMetGasBumpThreshold(
	tx *models.Tx,
	attemptIndex int,
	blockHeight uint64) bool {

	gasBumpThreshold := txm.config.EthGasBumpThreshold()
	txAttempt := tx.Attempts[attemptIndex]

	return blockHeight >= txAttempt.SentAt+gasBumpThreshold
}

// isLatestAttempt returns true only if the attempt is the last
// attempt associated with the transaction, alluding to the fact that
// it has the highest gas price after subsequent bumps.
func isLatestAttempt(tx *models.Tx, attemptIndex int) bool {
	return attemptIndex+1 == len(tx.Attempts)
}

// handleSafe marks a transaction as safe, no more work needs to be done
func (txm *EthTxManager) handleSafe(
	tx *models.Tx,
	attemptIndex int) error {
	txAttempt := tx.Attempts[attemptIndex]

	if err := txm.orm.MarkTxSafe(tx, txAttempt); err != nil {
		return errors.Wrap(err, "handleSafe MarkTxSafe failed")
	}

	minimumConfirmations := txm.config.MinOutgoingConfirmations()
	ethBalance, linkBalance, balanceErr := txm.GetETHAndLINKBalances(tx.From)

	logger.Infow(
		fmt.Sprintf("Tx #%d is safe", attemptIndex),
		"minimumConfirmations", minimumConfirmations,
		"txHash", txAttempt.Hash.String(),
		"txID", txAttempt.TxID,
		"ethBalance", ethBalance,
		"linkBalance", linkBalance,
		"err", balanceErr,
	)

	return nil
}

// bumpGas creates a new transaction attempt with an increased gas cost
func (txm *EthTxManager) bumpGas(tx *models.Tx, attemptIndex int, blockHeight uint64) error {
	txAttempt := tx.Attempts[attemptIndex]

	originalGasPrice := txAttempt.GasPrice.ToInt()
	bumpedGasPrice := new(big.Int).Add(originalGasPrice, txm.config.EthGasBumpWei())

	bumpedTxAttempt, err := txm.createAttempt(tx, bumpedGasPrice, blockHeight)
	if err != nil {
		return errors.Wrapf(err, "bumpGas from Tx #%s", txAttempt.Hash.Hex())
	}

	logger.Infow(
		fmt.Sprintf("Tx #%d created with bumped gas %v", attemptIndex+1, bumpedGasPrice),
		"originalTxHash", txAttempt.Hash,
		"newTxHash", bumpedTxAttempt.Hash)
	return nil
}

// createAttempt adds a new transaction attempt to a transaction record
func (txm *EthTxManager) createAttempt(
	tx *models.Tx,
	gasPriceWei *big.Int,
	blockHeight uint64,
) (*models.TxAttempt, error) {
	ma := txm.getAccount(tx.From)
	if ma == nil {
		return nil, fmt.Errorf("Unable to locate %v as an available account in EthTxManager. Has TxManager been started or has the address been removed?", tx.From.Hex())
	}

	newTxAttempt, err := txm.newTx(
		ma.Account,
		tx.Nonce,
		tx.To,
		tx.Value.ToInt(),
		tx.GasLimit,
		gasPriceWei,
		tx.Data,
		&ma.Address,
		blockHeight,
	)
	if err != nil {
		return nil, errors.Wrap(err, "createAttempt#newTx failed")
	}

	if _, err = txm.SendRawTx(newTxAttempt.SignedRawTx); err != nil {
		return nil, errors.Wrap(err, "createAttempt#SendRawTx failed")
	}

	txAttempt, err := txm.orm.AddTxAttempt(tx, newTxAttempt)
	if err != nil {
		return nil, errors.Wrap(err, "createAttempt#AddTxAttempt failed")
	}

	logger.Debugw(fmt.Sprintf("Added Tx attempt #%d", len(tx.Attempts)+1), "txID", tx.ID, "txAttemptID", txAttempt.ID)

	return txAttempt, nil
}

// NextActiveAccount uses round robin to select a managed account
// from the list of available accounts as defined in Register(...)
func (txm *EthTxManager) NextActiveAccount() *ManagedAccount {
	txm.accountsMutex.Lock()
	defer txm.accountsMutex.Unlock()

	if len(txm.availableAccounts) == 0 {
		return nil
	}

	account := txm.availableAccounts[txm.availableAccountIdx]
	txm.availableAccountIdx = (txm.availableAccountIdx + 1) % len(txm.availableAccounts)
	return account
}

func (txm *EthTxManager) getAccount(from common.Address) *ManagedAccount {
	txm.accountsMutex.Lock()
	defer txm.accountsMutex.Unlock()

	for _, a := range txm.availableAccounts {
		if a.Address == from {
			return a
		}
	}

	return nil
}

// ActivateAccount retrieves an account's nonce from the blockchain for client
// side management in ManagedAccount.
func (txm *EthTxManager) activateAccount(account accounts.Account) (*ManagedAccount, error) {
	nonce, err := txm.GetNonce(account.Address)
	if err != nil {
		return nil, err
	}

	return NewManagedAccount(account, nonce), nil
}

// ManagedAccount holds the account information alongside a client managed nonce
// to coordinate outgoing transactions.
type ManagedAccount struct {
	accounts.Account
	nonce         uint64
	lastSafeNonce uint64
	mutex         *sync.Mutex
}

// NewManagedAccount creates a managed account that handles nonce increments
// locally.
func NewManagedAccount(a accounts.Account, nonce uint64) *ManagedAccount {
	return &ManagedAccount{Account: a, nonce: nonce, mutex: &sync.Mutex{}}
}

// Nonce returns the client side managed nonce.
func (a *ManagedAccount) Nonce() uint64 {
	return a.nonce
}

// ReloadNonce fetch and update the current nonce via eth_getTransactionCount
func (a *ManagedAccount) ReloadNonce(txm *EthTxManager) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	nonce, err := txm.GetNonce(a.Address)
	if err != nil {
		return fmt.Errorf("TxManager ReloadNonce: %v", err)
	}
	logger.Debugw("Got new network nonce", "nonce", nonce)
	a.nonce = nonce
	return nil
}

// GetAndIncrementNonce will Yield the current nonce to a callback function and increment it once the
// callback has finished executing
func (a *ManagedAccount) GetAndIncrementNonce(callback func(uint64) error) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	err := callback(a.nonce)
	if err == nil {
		a.nonce = a.nonce + 1
	}

	return err
}

func (a *ManagedAccount) updateLastSafeNonce(latest uint64) {
	if latest > a.lastSafeNonce {
		a.lastSafeNonce = latest
	}
}
