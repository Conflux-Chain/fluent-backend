package service

import (
	"fmt"
	"math/big"
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

type TokenPayConfig struct {
	Tokens    map[string]common.Address // name => token address
	Recipient common.Address

	MinGasPriceRatioPercentage uint64 `default:"80"`                  // 80% of current gas price
	MaxGasCost                 uint64 `default:"100000000000000000"`  // 0.1 CFX
	MinSponsorBalance          uint64 `default:"1000000000000000000"` // 1 CFX

	CheckReceiptInterval        time.Duration `default:"1s"`
	CheckFundingReceiptInterval time.Duration `default:"30ms"`
}

type TokenPay struct {
	*TxSender

	config TokenPayConfig

	priceOracle *PriceOracle

	txSigner                 gethTypes.Signer
	allowedTokens            map[common.Address]string // token address => token name
	abiEncodedTokenRecipient string

	inflight sync.Map
}

func NewTokenPay(config TokenPayConfig, sender *TxSender, priceOracle *PriceOracle) (*TokenPay, error) {
	// validate config
	if len(config.Tokens) == 0 {
		return nil, errors.New("Tokens not specified")
	}

	if config.Recipient == (common.Address{}) {
		return nil, errors.New("Recipient not specified")
	}

	// allowed tokens
	allowedTokens := make(map[common.Address]string)
	for name, tokenAddr := range config.Tokens {
		allowedTokens[tokenAddr] = name
	}

	// abi encoded token recipient
	var buf [32]byte
	copy(buf[12:], config.Recipient.Bytes())

	return &TokenPay{
		TxSender:                 sender,
		config:                   config,
		priceOracle:              priceOracle,
		txSigner:                 gethTypes.LatestSignerForChainID(sender.chainIdBig.ToInt()),
		allowedTokens:            allowedTokens,
		abiEncodedTokenRecipient: hexutil.Encode(buf[:]),
	}, nil
}

func (tp *TokenPay) Config() TokenPayConfig {
	return tp.config
}

func (tp *TokenPay) IsTokenAllowed(token common.Address) bool {
	_, ok := tp.allowedTokens[token]
	return ok
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

	if new(big.Int).SetUint64(tp.config.MinSponsorBalance).Cmp(sponsorBalance) > 0 {
		return ErrTokenPaySponsorBalanceNotEnough.WithData(fmt.Sprintf("min = %v, actual = %v", tp.config.MinSponsorBalance, sponsorBalance))
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

	// funding ETH tx
	logger.WithField("txHash", context.fundingTxHash).Info("Funding ETH tx sent")

	if !tp.waitForReceipt(context.fundingTxHash, tp.config.CheckFundingReceiptInterval) {
		// TODO 这种情况一般都是 sponsor balance 不够，需要报警充值
		logger.WithField("txHash", context.fundingTxHash).Error("Funding ETH tx failed")
		return
	}

	// transfer token tx
	transferTokenTxHash, err := tp.client.Eth.SendRawTransaction(context.rawTransferTokenTx)
	if err != nil {
		// TODO 错误容错处理，如果是 io error 则重试，如果是 rpc error，则需要根据情况特殊处理：
		// 1. 如果 balance 不够，可能是临时 chain reorg，也可能是钱被用户转走；
		// 2. 如果是 nonce 问题，可能是用户发起了其它交易；
		// 3. 其它错误则需要报警调查。
		logger.Error("Failed to send transfer token tx")
		return
	}

	logger.WithField("txHash", transferTokenTxHash).Info("Transfer token tx sent")

	if !tp.waitForReceipt(transferTokenTxHash, tp.config.CheckReceiptInterval) {
		// TODO 这种情况大概率是用户转走了 token，导致 token balance 不够，需要 “拉黑” 用户。
		// 但是，也可能是其他原因，导致 eth_call 成功，但是实际执行失败，这种情况很极端。
		logger.WithField("txHash", transferTokenTxHash).Error("Transfer token tx failed")
		return
	}

	// business tx
	businessTxHash, err := tp.client.Eth.SendRawTransaction(context.rawBusinessTx)
	if err != nil {
		// TODO 同 transfer token tx 的错误处理，这里不区分了，先简单处理成一样的。
		logger.Error("Failed to send business tx")
		return
	}

	logger.WithField("txHash", businessTxHash).Info("Business tx sent")

	// do not care about the execution result of business tx
	tp.waitForReceipt(businessTxHash, tp.config.CheckReceiptInterval)

	logger.Info("Token pay completed")
}

func (tp *TokenPay) waitForReceipt(txHash common.Hash, interval time.Duration) bool {
	for {
		time.Sleep(interval)

		// TODO 检查交易是否存在，防止 txpool 满了丢弃交易的情况，要考虑重发交易。

		receipt, err := tp.client.Eth.TransactionReceipt(txHash)
		if err != nil {
			// TODO 错误容错处理，一般都是 io error，需要重试，时间久了则需要报警
			continue
		}

		if receipt == nil {
			continue
		}

		return receipt.Status != nil && *receipt.Status == gethTypes.ReceiptStatusSuccessful
	}
}
