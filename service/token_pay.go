package service

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/go-rpc-provider/utils"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	errMsgTxAlreadyExist    = "tx already exist"
	errMsgInsufficientFunds = "insufficient funds for gas * price + value"
	errMsgNonceTooLow       = "nonce too low"

	defaultSendTxRetryTimes    = 10
	defaultSendTxRetryInterval = 3 * time.Second
	defaultReceiptTimeout      = time.Minute
)

type TokenPay struct {
	*TxSender

	config TokenPayConfig

	priceOracle *PriceOracle

	txSigner gethTypes.Signer

	inflight sync.Map

	blacklisted sync.Map
}

func NewTokenPay(config TokenPayConfig, sender *TxSender, priceOracle *PriceOracle) *TokenPay {
	return &TokenPay{
		TxSender:    sender,
		config:      config,
		priceOracle: priceOracle,
		txSigner:    gethTypes.LatestSignerForChainID(sender.chainIdBig.ToInt()),
	}
}

func (tp *TokenPay) Config() TokenPayConfig {
	return tp.config
}

func (tp *TokenPay) Sponsor(rawTransferTokenTx, rawBusinessTx []byte, ip string) error {
	// validate ip
	if ip = strings.TrimSpace(ip); len(ip) == 0 {
		return api.ErrValidationStr("Client IP address is empty")
	}

	if _, ok := tp.blacklisted.Load(ip); ok {
		return api.ErrValidationStrf("IP address is blacklisted, ip = %v", ip)
	}

	// unmarshal given txs
	var transferTokenTx, businessTx gethTypes.Transaction

	if err := transferTokenTx.UnmarshalBinary(rawTransferTokenTx); err != nil {
		return api.ErrValidation(errors.WithMessage(err, "Failed to decode transfer token tx"))
	}

	if err := businessTx.UnmarshalBinary(rawBusinessTx); err != nil {
		return api.ErrValidation(errors.WithMessage(err, "Failed to decode business tx"))
	}

	// check sponsor balance
	sponsorBalance, err := tp.client.Eth.Balance(tp.TxSender.sender, nil)
	if err != nil {
		return NewRPCError(err, "Failed to retrieve sponsor balance")
	}

	if tp.config.minSponsorBalanceBig.Cmp(sponsorBalance) > 0 {
		return ErrTokenPaySponsorBalanceNotEnough.WithData(fmt.Sprintf("min = %v, actual = %v", tp.config.minSponsorBalanceBig, sponsorBalance))
	}

	// check the validity of given 2 txs
	result, err := tp.check(&transferTokenTx, &businessTx)
	if err != nil {
		return err
	}

	// allow only 1 sponsor tx per sender
	if _, loaded := tp.inflight.LoadOrStore(result.Sender, struct{}{}); loaded {
		return api.ErrValidationStrf("Another transaction in process, sender = %v, nonce = %v", result.Sender, result.Nonce)
	}

	// send funding ETH tx
	logrus.WithFields(logrus.Fields{
		"user":  result.Sender,
		"nonce": result.Nonce,
	}).Info("Begin to funding ETH")

	fundingTxGasLimit := hexutil.Uint64(result.FundingGasLimit.Uint64())
	fundingTxArgs := types.TransactionArgs{
		To:       &result.Sender,
		Gas:      &fundingTxGasLimit,
		GasPrice: (*hexutil.Big)(result.FundingGasPrice),
		Value:    (*hexutil.Big)(result.FundingValue),
	}

	fundingTxHash, err := tp.Send(fundingTxArgs)
	if err != nil {
		// clear the inflight record if failed to send funding tx
		tp.inflight.Delete(result.Sender)
		return err
	}

	go tp.monitor(TokenPayMonitorContext{
		ip:                 ip,
		checkResult:        result,
		fundingTxHash:      fundingTxHash,
		rawTransferTokenTx: rawTransferTokenTx,
		rawBusinessTx:      rawBusinessTx,
	})

	return nil
}

type TokenPayMonitorContext struct {
	ip                 string
	checkResult        checkResult
	fundingTxHash      common.Hash
	rawTransferTokenTx []byte
	rawBusinessTx      []byte
}

func (tp *TokenPay) monitor(context TokenPayMonitorContext) {
	defer tp.inflight.Delete(context.checkResult.Sender)

	logger := logrus.WithFields(logrus.Fields{
		"user":  context.checkResult.Sender,
		"nonce": context.checkResult.Nonce,
	})

	// funding ETH tx
	logger.WithField("txHash", context.fundingTxHash).Info("Succeeded to send Funding ETH tx")

	if success, errMsg, expired := tp.waitForReceipt(context.fundingTxHash, tp.config.CheckFundingReceiptInterval, defaultReceiptTimeout); !success {
		logger.WithFields(logrus.Fields{
			"txHash":  context.fundingTxHash,
			"errMsg":  errMsg,
			"expired": expired,
		}).Error("Failed to wait for receipt of Funding ETH tx")
		return
	}

	// transfer token tx
	transferTokenTxHash, err := tp.sendRawTransactionWithRetry(context.rawTransferTokenTx, defaultSendTxRetryTimes, defaultSendTxRetryInterval)
	if err != nil {
		logger.WithError(err).WithField("txHash", crypto.Keccak256Hash(context.rawTransferTokenTx)).Error("Failed to send transfer token tx")

		// Blacklist the sender and client IP due to balance or nonce problems, which may be caused by malicious users.
		if errMsg := err.Error(); strings.Contains(errMsg, errMsgInsufficientFunds) || strings.Contains(errMsg, errMsgNonceTooLow) {
			tp.addBlacklist(context.checkResult.Sender, context.ip)
		}

		return
	}

	logger.WithField("txHash", transferTokenTxHash).Info("Succeeded to send Transfer token tx")

	// business tx
	businessTxHash, err := tp.sendRawTransactionWithRetry(context.rawBusinessTx, defaultSendTxRetryTimes, defaultSendTxRetryInterval)
	if err != nil {
		logger.WithError(err).WithField("txHash", crypto.Keccak256Hash(context.rawBusinessTx)).Error("Failed to send business tx")
		return
	}

	logger.WithField("txHash", businessTxHash).Info("Succeeded to send Business tx")

	// check for transfer token tx receipt
	if success, errMsg, expired := tp.waitForReceipt(transferTokenTxHash, tp.config.CheckReceiptInterval, defaultReceiptTimeout); !success {
		logger.WithFields(logrus.Fields{
			"txHash":  transferTokenTxHash,
			"errMsg":  errMsg,
			"expired": expired,
		}).Error("Failed to wait for receipt of Transfer token tx")

		// Blacklist the sender and client IP if the transfer token tx failed.
		if !expired {
			tp.addBlacklist(context.checkResult.Sender, context.ip)
		}

		return
	}

	// do not care about the execution result of business tx
	tp.waitForReceipt(businessTxHash, tp.config.CheckReceiptInterval, defaultReceiptTimeout)

	logger.Info("Token pay completed")
}

func (tp *TokenPay) addBlacklist(user common.Address, ip string) {
	// Note, since the IP address is limited, no worry about the memory usage of the blacklist map. On the other hand,
	// the user address is also bounded by the number of client IP.
	//
	// If necessary, we can add admin API to update the blacklist map, or add a TTL to the blacklist map in future.
	tp.blacklisted.Store(user, true)
	tp.blacklisted.Store(ip, true)
}

func (tp *TokenPay) waitForReceipt(txHash common.Hash, interval time.Duration, timeout time.Duration) (success bool, errMsg string, expired bool) {
	startTime := time.Now()

	for {
		if time.Since(startTime) > timeout {
			return false, "", true
		}

		time.Sleep(interval)

		receipt, err := tp.client.Eth.TransactionReceipt(txHash)
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

func (tp *TokenPay) sendRawTransactionWithRetry(rawTx []byte, maxRetry int, interval time.Duration) (common.Hash, error) {
	var retry int

	for {
		txHash, err := tp.client.Eth.SendRawTransaction(rawTx)
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
