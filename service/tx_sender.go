package service

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/go-rpc-provider/utils"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const errMsgTxAlreadyExist = "tx already exist"

var (
	defaultBalanceCheckInterval = time.Minute
	defaultBalanceThreshold     = big.NewInt(1e18) // 1 CFX
)

type TxSender struct {
	client *web3go.Client

	sender common.Address

	chainId    uint64
	chainIdBig *hexutil.Big

	mu sync.Mutex
}

func NewTxSender(client *web3go.Client) (*TxSender, error) {
	// get the default signer
	sm, err := client.GetSignerManager()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get signer manager from RPC client")
	}

	signers := sm.List()
	if len(signers) == 0 {
		return nil, errors.New("No signer found")
	}

	sender := signers[0].Address()

	// check sender balance
	balance, err := client.Eth.Balance(sender, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve sender balance")
	}

	if balance.Sign() == 0 {
		return nil, errors.Errorf("Sender balance is 0, address = %v", sender)
	}

	// retrieve chain ID
	chainId, err := client.Eth.ChainId()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve chain ID")
	}

	if chainId == nil {
		return nil, errors.New("Chain ID unavailable on fullnode")
	}

	txSender := TxSender{
		client:     client,
		sender:     sender,
		chainId:    *chainId,
		chainIdBig: (*hexutil.Big)(new(big.Int).SetUint64(*chainId)),
	}

	// starts to monitor the balance of the sender address in the entire process lifetime,
	// and graceful shutdown is unnecessary.
	go txSender.monitorBalance()

	return &txSender, nil
}

func (s *TxSender) monitorBalance() {
	ticker := time.NewTicker(defaultBalanceCheckInterval)
	defer ticker.Stop()

	healthCounter := health.NewTimedCounter()
	balanceCounter := health.NewTimedCounter(health.TimedCounterConfig{
		Threshold: time.Second,
		Remind:    time.Hour,
	})

	for range ticker.C {
		balance, err := s.client.Eth.Balance(s.sender, nil)
		healthCounter.LogOnError(err, "Monitoring tx sender balance")
		if err != nil {
			continue
		}

		if balance.Cmp(defaultBalanceThreshold) < 0 {
			err = fmt.Errorf("Tx sender balance not enough, address = %v, balance = %v, threshold = %v", s.sender, balance, defaultBalanceThreshold)
		}

		balanceCounter.LogOnError(err, "Monitoring tx sender balance")
	}
}

func (s *TxSender) Send(tx types.TransactionArgs) (common.Hash, error) {
	if tx.From == nil {
		tx.From = &s.sender
	}

	if tx.ChainID == nil {
		tx.ChainID = s.chainIdBig
	}

	// use mutex to ensure correct tx nonce, and need to enhance for high QPS case
	s.mu.Lock()
	defer s.mu.Unlock()

	txHash, err := s.client.Eth.SendTransactionByArgs(tx)
	if err != nil {
		return common.Hash{}, NewRPCError(err, "Failed to send transaction")
	}

	return txHash, nil
}

func (s *TxSender) SendRawTransactionWithRetry(rawTx []byte, maxRetry int, interval time.Duration) (common.Hash, error) {
	var retry int

	for {
		txHash, err := s.client.Eth.SendRawTransaction(rawTx)
		if err == nil {
			return txHash, nil
		}

		// do not retry for RPC error
		if utils.IsRPCJSONError(err) {
			// Sometimes, tx sent failed due to io error, and retry again.
			// However, the tx already inserted into txpool successfully, which means request succeeded but response failed.
			// In this case, the retry will failed with "tx already exist", so we can treat it as success and return the tx hash.
			if strings.Contains(err.Error(), errMsgTxAlreadyExist) {
				logrus.WithError(err).WithField("retry", retry).Warn("Tx already sent, treat as success")
				return crypto.Keccak256Hash(rawTx), nil
			}

			return common.Hash{}, err
		}

		// retry again for other errors (e.g. io error)
		retry++
		if retry > maxRetry {
			return common.Hash{}, err
		}

		logrus.WithError(err).WithField("retry", retry).Debug("Failed to send tx, retrying...")

		time.Sleep(interval)
	}
}

func (s *TxSender) WaitForReceipt(txHash common.Hash, interval time.Duration, timeout time.Duration) (success bool, errMsg string, expired bool) {
	startTime := time.Now()

	for {
		if time.Since(startTime) > timeout {
			return false, "", true
		}

		time.Sleep(interval)

		receipt, err := s.client.Eth.TransactionReceipt(txHash)
		if err != nil {
			logrus.WithError(err).WithField("txHash", txHash).Info("Failed to get tx receipt")
			continue
		}

		if receipt == nil || receipt.Status == nil {
			continue
		}

		if *receipt.Status == gethTypes.ReceiptStatusSuccessful {
			return true, "", false
		}

		if receipt.TxExecErrorMsg == nil {
			return false, "", false
		}

		return false, *receipt.TxExecErrorMsg, false
	}
}
