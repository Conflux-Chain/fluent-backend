package worker

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// configKeyEventScanNextBlock is the key used to store the next block number to scan user op events in the database.
//
// NOTE if the paymaster contract address changes, we should remove this key from the database to avoid scanning events from an incorrect block number.
const configKeyEventScanNextBlock = "worker.userOpEventScan.nextBlock"

// eventHashSponsored is the hash of the Sponsored event signature, used to filter logs for this specific event.
//
// Solidity: event Sponsored(bytes32 indexed userOpHash, bool success, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
var eventHashSponsored = common.HexToHash("0x9e4d0db315fe76b23fdbfa2d345dd700140cae22039688064727b295782b0204")

type UserOpEventScanConfig struct {
	Contract common.Address

	// NextBlock is the next block number to start scanning events from, 0 means the finalized block.
	NextBlock uint64
	nextBlock uint64

	// Interval is the interval between each scan after catch-up phase.
	Interval time.Duration `default:"10s"`
}

// UserOpEventScanner is a worker that scans user op events from the blockchain and updates the database accordingly.
type UserOpEventScanner struct {
	config   UserOpEventScanConfig
	client   *web3go.Client
	store    *store.Store
	filterer *contract.VerifyingPaymasterFilterer
}

// NewUserOpEventScanner creates a new UserOpEventScanner with the given configuration, web3 client, and store.
// It initializes the scanner and loads the next block number to scan from the database or configuration.
// It will return an error if the config is invalid or initialization fails.
func NewUserOpEventScanner(config UserOpEventScanConfig, client *web3go.Client, store *store.Store) (*UserOpEventScanner, error) {
	if config.Contract == (common.Address{}) {
		return nil, errors.New("Contract address is required")
	}

	if config.Interval <= 0 {
		return nil, errors.New("Interval must be greater than 0")
	}

	caller, _ := client.ToClientForContract()
	verifyingPaymasterFilterer, err := contract.NewVerifyingPaymasterFilterer(config.Contract, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create VerifyingPaymasterFilterer")
	}

	scanner := UserOpEventScanner{
		config:   config,
		client:   client,
		store:    store,
		filterer: verifyingPaymasterFilterer,
	}

	if scanner.config.nextBlock, err = scanner.loadNextBlock(); err != nil {
		return nil, errors.WithMessage(err, "Failed to load the next block to scan event")
	}

	return &scanner, nil
}

// loadNextBlock loads the next block number to scan user op events from the database or configuration.
func (scanner *UserOpEventScanner) loadNextBlock() (uint64, error) {
	// load break point from database
	value, ok, err := scanner.store.Config.Get(configKeyEventScanNextBlock)
	if err != nil {
		return 0, errors.WithMessage(err, "Failed to load next block from database")
	}

	if ok {
		nextBlock, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return 0, errors.WithMessage(err, "Failed to parse next block number as uint64 from database")
		}

		logrus.WithField("next", nextBlock).Debug("Succeeded to load next block to scan user op event from database")

		return nextBlock, nil
	}

	// use the config value if set
	if scanner.config.NextBlock > 0 {
		logrus.WithField("next", scanner.config.NextBlock).Debug("Use configured next block value to scan user op event")
		return scanner.config.NextBlock, nil
	}

	// otherwise, fallback to the finalized block number
	block, err := scanner.client.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if err != nil {
		return 0, errors.WithMessage(err, "Failed to get finalized block number from blockchain")
	}

	nextBlock := block.Number.Uint64()

	logrus.WithField("next", nextBlock).Debug("Succeeded to retrieve finalized block number to scan user op event")

	return nextBlock, nil
}

// Work starts the scanning process for user op events. It first catches up to the latest finalized block,
// and then continues to scan for new events periodically.
func (scanner *UserOpEventScanner) Work() {
	logrus.WithField("next", scanner.config.nextBlock).Info("Begin to scan user op event")

	// Firstly, catch up to the latest finalized block.
	//
	// Note: there are only few event logs at early phase, so we can retrieve them in one request from Confura.
	for {
		ok, err := scanner.scan()
		if err != nil {
			logrus.WithError(err).WithField("next", scanner.config.nextBlock).Warn("Failed to scan user op event logs in catch-up phase")
			time.Sleep(scanner.config.Interval)
		} else if !ok {
			break
		}
	}

	// After catch-up phase, we can scan the event logs periodically.
	ticker := time.NewTicker(scanner.config.Interval)
	defer ticker.Stop()

	for range ticker.C {
		if _, err := scanner.scan(); err != nil {
			logrus.WithError(err).WithField("next", scanner.config.nextBlock).Warn("Failed to scan user op event logs periodically")
		}
	}
}

// scan retrieves user op event logs from the blockchain and updates the database accordingly.
func (scanner *UserOpEventScanner) scan() (bool, error) {
	// get the finalized block number
	finalizedBlock, err := scanner.client.Eth.BlockByNumber(types.FinalizedBlockNumber, false)
	if err != nil {
		return false, errors.WithMessage(err, "Failed to retrieve the finalized block")
	}

	// already catch up to the finalized block
	finalizedBlockNumber := finalizedBlock.Number.Uint64()
	if scanner.config.nextBlock > finalizedBlockNumber {
		return false, nil
	}

	logrus.WithFields(logrus.Fields{
		"from": scanner.config.nextBlock,
		"to":   finalizedBlockNumber,
	}).Debug("Scanning user op event logs from blockchain")

	// retrieve event logs between nextBlock and finalizedBlockNumber.
	logs, err := getLogs(scanner.client, scanner.config.Contract, scanner.config.nextBlock, finalizedBlockNumber, eventHashSponsored)
	if err != nil {
		return false, errors.WithMessagef(err, "Failed to retrieve event logs, next = %v, finalized = %v", scanner.config.nextBlock, finalizedBlockNumber)
	}

	// handle the retrieved event logs
	if err = scanner.handle(logs, finalizedBlockNumber+1); err != nil {
		return false, errors.WithMessage(err, "Failed to handle event logs")
	}

	logrus.WithFields(logrus.Fields{
		"from": scanner.config.nextBlock,
		"to":   finalizedBlockNumber,
	}).Debug("Succeeded to scan and handle user op events")

	return true, nil
}

// handle processes the retrieved event logs, updates the user ops and configuration in the database, and updates the next block number in memory.
func (scanner *UserOpEventScanner) handle(logs []types.Log, nextBlock uint64) error {
	// parse event logs
	var events []*contract.VerifyingPaymasterSponsored

	for _, v := range logs {
		event, err := scanner.filterer.ParseSponsored(*v.ToEthLog())
		if err != nil {
			return errors.WithMessage(err, "Failed to parse Sponsored event log")
		}

		events = append(events, event)
	}

	// update user ops and config in a transaction
	fc := func(tx *gorm.DB) error {
		// update user ops
		for _, v := range events {
			if updated, dbErr := scanner.store.UserOp.Update(v, tx); dbErr != nil {
				return errors.WithMessage(dbErr, "Failed to update user op in database")
			} else if !updated {
				logrus.WithField("userOpHash", hexutil.Encode(v.UserOpHash[:])).Error("Sponsored event userOpHash not found in database")
			}
		}

		// update config
		if dbErr := scanner.store.Config.Upsert(configKeyEventScanNextBlock, fmt.Sprint(nextBlock), tx); dbErr != nil {
			return errors.WithMessagef(dbErr, "Failed to update config in database by key %v", configKeyEventScanNextBlock)
		}

		return nil
	}

	// perform the transaction
	if err := scanner.store.DB.Transaction(fc); err != nil {
		return errors.WithMessage(err, "Failed to handle event logs in database transaction")
	}

	// update the next block number in memory
	scanner.config.nextBlock = nextBlock

	return nil
}

/////////////////////////////////////////////////////////////////////////////////
//
// Blockchain utils
//
/////////////////////////////////////////////////////////////////////////////////

const errNarrowDownPattern = "narrow down"

// getLogs retrieves logs from the blockchain for a given contract address and optional topic within a specified block range.
//
// Now, its implementation depends on the Confura that do not limit the block number range, and support to index by address + topic0.
func getLogs(client *web3go.Client, contract common.Address, blockFrom, blockTo uint64, topic0 ...common.Hash) ([]types.Log, error) {
	// address filter
	filter := types.FilterQuery{
		Addresses: []common.Address{contract},
	}

	if len(topic0) > 0 {
		filter.Topics = [][]common.Hash{{topic0[0]}}
	}

	from, to := blockFrom, blockTo
	var result []types.Log

	blockNumberConverter := func(number uint64) *types.BlockNumber {
		bn := types.NewBlockNumber(int64(number))
		return &bn
	}

	for from <= to {
		// set block number range in log filter
		filter.FromBlock = blockNumberConverter(from)
		filter.ToBlock = blockNumberConverter(to)

		logs, err := client.Eth.Logs(filter)
		if err == nil {
			// success and move forward
			result = append(result, logs...)

			from = to + 1
			to = blockTo
		} else if strings.Contains(err.Error(), errNarrowDownPattern) {
			// narrow down the block number range
			if from == to {
				return nil, fmt.Errorf("Failed to narrow down block number range, from == to == %v", from)
			}

			to = from + (to-from)/2
		} else {
			// other error
			return nil, errors.WithMessagef(err, "Failed to retrieve event logs, from = %v, to = %v", from, to)
		}
	}

	return result, nil
}
