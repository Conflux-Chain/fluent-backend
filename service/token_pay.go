package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

func (tp *TokenPay) Sponsor(rawTransferTokenTx, rawBusinessTx []byte) error {
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
		checkResult:        result,
		fundingTxHash:      fundingTxHash,
		rawTransferTokenTx: rawTransferTokenTx,
		rawBusinessTx:      rawBusinessTx,
	})

	return nil
}

type TokenPayMonitorContext struct {
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

	// TODO 优化异常处理流程：
	// 1. 发送交易 io error 时可以重试，且注意处理 response io error 后导致 tx already exists 的情况；
	// 2. 重试多次后考虑换备机发送交易；
	// 3. 如果发送交易失败原因是 rpc error，则需要进一步分析主因，但目前已做充分检查，先报警人工介入；
	// 4. 交易丢失：考虑重发，长时间丢失则报警人工处理。这种场景，主要是遇到 txpool 一直满的情况；
	// 5. 交易长时间不打包：虽然已经通过更高的 gas fee/tip 来提升交易优先级，但遇到该情况时，还需要支持重发逻辑，包括重发 funding 交易，用户重发 transfer token 和 business 交易；
	// 6. 交易执行失败：需要分情况处理，如果用户 nonce 或者 token balance 问题，需要拉黑；

	// funding ETH tx
	logger.WithField("txHash", context.fundingTxHash).Info("Funding ETH tx sent")

	if success, errMsg := tp.waitForReceipt(context.fundingTxHash, tp.config.CheckFundingReceiptInterval); !success {
		logger.WithField("txHash", context.fundingTxHash).WithField("errMsg", errMsg).Error("Funding ETH tx failed")
		return
	}

	// transfer token tx
	transferTokenTxHash, err := tp.client.Eth.SendRawTransaction(context.rawTransferTokenTx)
	if err != nil {
		logger.WithError(err).Error("Failed to send transfer token tx")
		return
	}

	logger.WithField("txHash", transferTokenTxHash).Info("Transfer token tx sent")

	// business tx
	businessTxHash, err := tp.client.Eth.SendRawTransaction(context.rawBusinessTx)
	if err != nil {
		logger.WithError(err).Error("Failed to send business tx")
		return
	}

	logger.WithField("txHash", businessTxHash).Info("Business tx sent")

	// check for transfer token tx receipt
	if success, errMsg := tp.waitForReceipt(transferTokenTxHash, tp.config.CheckReceiptInterval); !success {
		logger.WithField("txHash", transferTokenTxHash).WithField("errMsg", errMsg).Error("Transfer token tx failed")
		tp.blacklisted.Store(context.checkResult.Sender, true)
		return
	}

	// do not care about the execution result of business tx
	tp.waitForReceipt(businessTxHash, tp.config.CheckReceiptInterval)

	logger.Info("Token pay completed")
}

func (tp *TokenPay) waitForReceipt(txHash common.Hash, interval time.Duration) (bool, string) {
	for {
		time.Sleep(interval)

		receipt, err := tp.client.Eth.TransactionReceipt(txHash)
		if err != nil {
			continue
		}

		if receipt == nil || receipt.Status == nil {
			continue
		}

		if *receipt.Status == gethTypes.ReceiptStatusSuccessful {
			return true, ""
		}

		if receipt.TxExecErrorMsg == nil {
			return false, ""
		}

		return false, *receipt.TxExecErrorMsg
	}
}
