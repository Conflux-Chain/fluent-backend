package service

// TokenPay Implementation Notes
//
// IMPORTANT: For security design considerations and known risks, see TOKEN_PAY_DESIGN.md in this directory.
// This file implements a token pay mechanism that allows users to pay gas fees with ERC20 tokens.
// The implementation accepts certain security risks by design - refer to TOKEN_PAY_DESIGN.md for detailed explanation.

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var txpoolErrMsgKeywords = map[string]bool{
	"insufficient funds": true,
	"nonce":              true,
	"underpriced":        true,
}

const (
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
	logger := logrus.WithFields(logrus.Fields{
		"module": "TokenPay",
		"user":   result.Sender,
		"nonce":  result.Nonce,
	})
	logger.Info("Begin to funding ETH")

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

	go tp.monitor(logger, TokenPayMonitorContext{
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

func (tp *TokenPay) monitor(logger *logrus.Entry, context TokenPayMonitorContext) {
	defer tp.inflight.Delete(context.checkResult.Sender)

	// funding ETH tx
	logger.WithField("txHash", context.fundingTxHash).Info("Succeeded to send Funding ETH tx")

	if err := tp.waitForFundingTx(logger, context.fundingTxHash, 5*defaultReceiptTimeout); err != nil {
		logger.WithError(err).Error("Failed to wait for receipt of Funding ETH tx")
		return
	}

	logger.Info("Succeeded to execute Funding ETH tx")

	// transfer token tx
	transferTokenTxHash, err := tp.SendRawTransactionWithRetry(context.rawTransferTokenTx, defaultSendTxRetryTimes, defaultSendTxRetryInterval)
	if err != nil {
		logger.WithError(err).WithField("txHash", crypto.Keccak256Hash(context.rawTransferTokenTx)).Error("Failed to send transfer token tx")

		// Blacklist the sender and client IP if the transfer token tx failed due to known txpool errors, which is most likely caused by malicious users.
		errMsg := err.Error()
		for v := range txpoolErrMsgKeywords {
			if strings.Contains(errMsg, v) {
				tp.addBlacklist(logger, context.checkResult.Sender, context.ip)
				break
			}
		}

		return
	}

	logger.WithField("txHash", transferTokenTxHash).Info("Succeeded to send Transfer token tx")

	// business tx
	businessTxHash, err := tp.SendRawTransactionWithRetry(context.rawBusinessTx, defaultSendTxRetryTimes, defaultSendTxRetryInterval)
	if err != nil {
		logger.WithError(err).WithField("txHash", crypto.Keccak256Hash(context.rawBusinessTx)).Error("Failed to send business tx")
		return
	}

	logger.WithField("txHash", businessTxHash).Info("Succeeded to send Business tx")

	// check for transfer token tx receipt
	if success, errMsg, expired := tp.WaitForReceipt(transferTokenTxHash, tp.config.CheckReceiptInterval, defaultReceiptTimeout); !success {
		logger.WithFields(logrus.Fields{
			"txHash":  transferTokenTxHash,
			"errMsg":  errMsg,
			"expired": expired,
		}).Error("Failed to wait for receipt of Transfer token tx")

		// Blacklist the sender and client IP if the transfer token tx failed.
		if !expired {
			tp.addBlacklist(logger, context.checkResult.Sender, context.ip)
		}

		return
	}

	logger.Info("Succeeded to execute Transfer token tx")

	// do not care about the execution result of business tx
	tp.WaitForReceipt(businessTxHash, tp.config.CheckReceiptInterval, defaultReceiptTimeout)

	logger.Info("Succeeded to execute Business tx")
}

func (tp *TokenPay) addBlacklist(logger *logrus.Entry, user common.Address, ip string) {
	// Note, since the IP address is limited, no worry about the memory usage of the blacklist map. On the other hand,
	// the user address is also bounded by the number of client IP.
	//
	// If necessary, we can add admin API to update the blacklist map, or add a TTL to the blacklist map in future.
	tp.blacklisted.Store(user, true)
	tp.blacklisted.Store(ip, true)

	logger.WithField("ip", ip).Warn("Blacklisted the user and client IP")
}

func (tp *TokenPay) waitForFundingTx(logger *logrus.Entry, txHash common.Hash, timeout time.Duration) error {
	// retrieve the funding ETH
	tx, err := tp.client.Eth.TransactionByHash(txHash)
	if err != nil {
		return errors.WithMessage(err, "Failed to get transaction by hash")
	}

	if tx == nil {
		return fmt.Errorf("Transaction not found, txHash = %v", txHash)
	}

	// wait for nonce to be used
	txs := []common.Hash{txHash}
	start := time.Now()
	startIter := time.Now()
	price := tx.GasPrice

	for {
		// timeout to wait for nonce used
		if time.Since(start) > timeout {
			return fmt.Errorf("Timeout to wait for funding ETH tx nonce, nonce = %v", tx.Nonce)
		}

		time.Sleep(tp.config.CheckFundingReceiptInterval)

		nonce, err := tp.client.Eth.TransactionCount(tx.From, nil)
		if err != nil {
			logger.WithError(err).Info("Failed to get state nonce, retry again")
			continue
		}

		// nonce used
		if nonce.Uint64() > tx.Nonce {
			break
		}

		if time.Since(startIter) < defaultReceiptTimeout {
			continue
		}

		startIter = time.Now()

		// re-send the funding ETH tx with higher gas price (2x)
		price.Mul(price, big.NewInt(2))
		txArgs := types.TransactionArgs{
			To:       tx.To,
			Gas:      (*hexutil.Uint64)(&tx.Gas),
			GasPrice: (*hexutil.Big)(price),
			Value:    (*hexutil.Big)(tx.Value),
			Nonce:    (*hexutil.Uint64)(&tx.Nonce),
		}

		newTxHash, err := tp.Send(txArgs)
		if err != nil {
			logger.WithError(err).WithField("price", price).Info("Failed to re-send funding ETH tx with higher gas price")
		} else {
			txs = append(txs, newTxHash)
			logger.WithField("txHash", newTxHash).WithField("price", price).Info("Succeeded to re-send funding ETH tx with higher gas price")
		}
	}

	// check receipt of all sent txs
	var receipt *types.Receipt
	for _, v := range txs {
		if receipt, err = tp.client.Eth.TransactionReceipt(v); err != nil {
			logger.WithError(err).WithField("txHash", v).Info("Failed to get receipt of funding ETH tx")
		} else if receipt != nil {
			break
		}
	}

	if receipt == nil || receipt.Status == nil {
		return fmt.Errorf("Funding ETH tx receipt not found or receipt status is nil, txs = %v", txs)
	}

	if *receipt.Status != gethTypes.ReceiptStatusSuccessful {
		errMsg := "N/A"
		if receipt.TxExecErrorMsg != nil {
			errMsg = *receipt.TxExecErrorMsg
		}
		return fmt.Errorf("Funding ETH tx failed, txHash = %v, errMsg = %v", receipt.TransactionHash, errMsg)
	}

	return nil
}
