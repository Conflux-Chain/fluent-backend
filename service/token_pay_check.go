package service

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	web3goTypes "github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

const (
	minTxGasLimit = 21000

	transferMethodId = "0xa9059cbb" // transfer(address,uint256)
)

var (
	big100     = big.NewInt(100)
	big10Exp18 = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
)

type checkResult struct {
	Sender          common.Address
	Nonce           uint64
	FundingGasLimit *big.Int
	FundingGasPrice *big.Int
	FundingValue    *big.Int
}

func (tp *TokenPay) check(transferTokenTx, businessTx *types.Transaction) (checkResult, error) {
	// nonce in sequence and same gas price
	if transferTokenTx.Nonce()+1 != businessTx.Nonce() {
		return checkResult{}, api.ErrValidationStr("Nonce not in sequence")
	}

	gasPrice := transferTokenTx.GasPrice()

	if businessTx.GasPrice().Cmp(gasPrice) != 0 {
		return checkResult{}, api.ErrValidationStr("Gas price mismatch")
	}

	// static check business tx
	businessTxSender, err := tp.staticCheckTx(businessTx)
	if err != nil {
		return checkResult{}, api.ErrValidation(errors.WithMessage(err, "Failed to static check business tx"))
	}

	// static check transfer token tx
	transferTokenTxSender, token, amount, err := tp.staticCheckTransferTokenTx(transferTokenTx)
	if err != nil {
		return checkResult{}, api.ErrValidation(errors.WithMessage(err, "Failed to static check transfer token tx"))
	}

	if transferTokenTxSender != businessTxSender {
		return checkResult{}, api.ErrValidationStrf("Tx sender mismatch: transferToken = %v, business = %v", transferTokenTxSender, businessTxSender)
	}

	// dynamic check transfer token tx
	if err = tp.dynamicCheckTransferTokenTx(transferTokenTx, transferTokenTxSender, token, amount); err != nil {
		return checkResult{}, err
	}

	// estimates the gas limit for funding ETH transaction
	senderGasLimit := new(big.Int).SetUint64(transferTokenTx.Gas() + businessTx.Gas())
	senderGasCost := new(big.Int).Mul(senderGasLimit, gasPrice)
	fundingGasLimit, err := tp.estimateFundingETHGas(transferTokenTxSender, senderGasCost)
	if err != nil {
		return checkResult{}, NewRPCError(err, "Failed to estimate gas limit for funding ETH transaction")
	}

	// limit the max gas cost for risk control
	totalGasLimit := new(big.Int).Add(senderGasLimit, fundingGasLimit)
	totalGasCost := new(big.Int).Mul(totalGasLimit, gasPrice)
	if maxGasCost := new(big.Int).SetUint64(tp.config.MaxGasCost); maxGasCost.Cmp(totalGasCost) < 0 {
		return checkResult{}, ErrTokenPayGasCostTooHigh.WithData(fmt.Sprintf("max = %v, actual = %v", maxGasCost, totalGasCost))
	}

	// ensure the transferred token amount could cover the gas fees of 3 txs: funding ETH, transfer token and business
	txEthPrice := new(big.Int).Mul(amount, big10Exp18)
	txEthPrice.Div(txEthPrice, totalGasCost)

	currentEthPrice, err := tp.priceOracle.GetETHPrice(token)
	if err != nil {
		return checkResult{}, err
	}

	if txEthPrice.Cmp(currentEthPrice) < 0 {
		return checkResult{}, ErrTokenPayTransferTokenAmountTooLow.WithData(fmt.Sprintf("ETH price by quote token, current = %v, tx = %v", currentEthPrice, txEthPrice))
	}

	// simulate the txs for successful execution
	if err = tp.simulateTx(transferTokenTx, transferTokenTxSender); err != nil {
		return checkResult{}, ErrTokenPaySimulateTxFailed.WithData(fmt.Sprintf("Failed to simulate transfer token transaction: %v", err.Error()))
	}

	if err = tp.simulateTx(businessTx, businessTxSender); err != nil {
		return checkResult{}, ErrTokenPaySimulateTxFailed.WithData(fmt.Sprintf("Failed to simulate business transaction: %v", err.Error()))
	}

	return checkResult{
		Sender:          transferTokenTxSender,
		Nonce:           transferTokenTx.Nonce(),
		FundingGasLimit: fundingGasLimit,
		FundingGasPrice: gasPrice,
		FundingValue:    senderGasCost,
	}, nil
}

// staticCheckTx performs basic static checks for the given transaction and returns the recovered sender address from signature.
func (tp *TokenPay) staticCheckTx(tx *types.Transaction) (sender common.Address, err error) {
	// txType: only legacy, 1559 and 7702 txs allowed
	if txType := tx.Type(); txType > types.SetCodeTxType || txType == types.AccessListTxType || txType == types.BlobTxType {
		return common.Address{}, errors.Errorf("Invalid tx type: %v", txType)
	}

	// chainId
	if chainId := tx.ChainId().Uint64(); chainId != tp.chainId {
		return common.Address{}, errors.Errorf("Invalid chain ID, expected = %v, got = %v", tp.chainId, chainId)
	}

	// gas limit: 21000 at least
	if gasLimit := tx.Gas(); gasLimit < minTxGasLimit {
		return common.Address{}, errors.Errorf("Invalid gas limit, minimum required is %v, got %v", minTxGasLimit, gasLimit)
	}

	// access list
	if numKyes := tx.AccessList().StorageKeys(); numKyes > 0 {
		return common.Address{}, errors.Errorf("Access list is not empty, keys = %v", numKyes)
	}

	// recover sender
	if sender, err = types.Sender(tp.txSigner, tx); err != nil {
		return common.Address{}, errors.Errorf("Failed to derive sender: %v", err)
	}

	return sender, nil
}

// staticCheckTransferTokenTx performs static checks specific for transfer token transaction. It returns the recovered sender address,
// transferred token and amount.
func (tp *TokenPay) staticCheckTransferTokenTx(tx *types.Transaction) (sender, token common.Address, amount *big.Int, err error) {
	// txType: legacy or 1559
	if txType := tx.Type(); txType != types.LegacyTxType && txType != types.DynamicFeeTxType {
		return common.Address{}, common.Address{}, nil, errors.Errorf("Invalid tx type: %v", txType)
	}

	// to address should be the transferred token address
	toAddr := tx.To()
	if toAddr == nil {
		return common.Address{}, common.Address{}, nil, errors.New("To address is nil")
	}
	token = *toAddr

	if _, ok := tp.allowedTokens[token]; !ok {
		return common.Address{}, common.Address{}, nil, errors.Errorf("Transferred token is not allowed: %v", token.Hex())
	}

	// value should be 0
	if txValue := tx.Value(); txValue.Sign() != 0 {
		return common.Address{}, common.Address{}, nil, errors.Errorf("Invalid tx value, expected 0, got %v", txValue)
	}

	// calldata
	if amount, err = tp.staticCheckTransferTokenCalldata(tx.Data()); err != nil {
		return common.Address{}, common.Address{}, nil, errors.WithMessage(err, "Failed to static check transfer token calldata")
	}

	// common check
	if sender, err = tp.staticCheckTx(tx); err != nil {
		return common.Address{}, common.Address{}, nil, errors.WithMessage(err, "Failed to static check tx in common")
	}

	return sender, token, amount, nil
}

// staticCheckTransferTokenCalldata performs static checks for the calldata of transfer token transaction. It returns the transferred amount.
func (tp *TokenPay) staticCheckTransferTokenCalldata(calldata []byte) (*big.Int, error) {
	// length
	if calldataLen := len(calldata); calldataLen != 68 {
		return nil, errors.Errorf("Invalid calldata length: expected = 68, got = %v", calldataLen)
	}

	// method ID
	if methodId := hexutil.Encode(calldata[:4]); methodId != transferMethodId {
		return nil, errors.Errorf("Invalid method ID: expected = %v, got = %v", transferMethodId, methodId)
	}

	// recipient
	if transferTo := hexutil.Encode(calldata[4:36]); transferTo != tp.abiEncodedTokenRecipient {
		return nil, errors.Errorf("Invalid transferTo address: expected = %v, got = %v", tp.abiEncodedTokenRecipient, transferTo)
	}

	// amount
	amount := new(big.Int).SetBytes(calldata[36:])

	return amount, nil
}

// dynamicCheckTransferTokenTx performs dynamic checks for the given transfer token transaction, including gas price, nonce and token balance.
func (tp *TokenPay) dynamicCheckTransferTokenTx(tx *types.Transaction, sender, token common.Address, amount *big.Int) error {
	// gas price
	currentGasPrice, err := tp.client.Eth.GasPrice()
	if err != nil {
		return NewRPCError(err, "Failed to retrieve current gas price")
	}

	minGasPrice := new(big.Int).Mul(currentGasPrice, new(big.Int).SetUint64(tp.config.MinGasPriceRatioPercentage))
	minGasPrice.Div(minGasPrice, big100)

	if txGasPrice := tx.GasPrice(); txGasPrice.Cmp(minGasPrice) < 0 {
		return ErrTokenPayPriceTooLow.WithData(fmt.Sprintf("Current gas price = %v, tx gas price = %v", currentGasPrice, txGasPrice))
	}

	// pending nonce
	pendingBlock := web3goTypes.BlockNumberOrHashWithNumber(web3goTypes.PendingBlockNumber)
	pendingNonce, err := tp.client.Eth.TransactionCount(sender, &pendingBlock)
	if err != nil {
		return NewRPCError(err, "Failed to retrieve pending nonce")
	}

	if txNonce := tx.Nonce(); txNonce != pendingNonce.Uint64() {
		return api.ErrValidationStrf("Invalid nonce, expected = %v, got = %v", pendingNonce, txNonce)
	}

	// state nonce
	stateNonce, err := tp.client.Eth.TransactionCount(sender, nil)
	if err != nil {
		return NewRPCError(err, "Failed to retrieve state nonce")
	}

	if stateNonce.Cmp(pendingNonce) != 0 {
		return api.ErrValidationStrf("Nonce and pending nonce mismatch, state = %v, pending = %v", stateNonce, pendingNonce)
	}

	// token balance
	caller, _ := tp.client.ToClientForContract()
	erc20Caller, err := contract.NewERC20Caller(token, caller)
	if err != nil {
		return errors.WithMessage(err, "Failed to create ERC20 caller")
	}

	balance, err := erc20Caller.BalanceOf(nil, sender)
	if err != nil {
		return NewRPCError(err, "Failed to retrieve token balance")
	}

	if balance.Cmp(amount) < 0 {
		return api.ErrValidationStrf("Insufficient token balance: required = %v, balance = %v", amount, balance)
	}

	return nil
}

// simulateTx simulates the execution of the given transaction via eth_call RPC.
// It will return error if the transaction execution failed.
func (tp *TokenPay) simulateTx(tx *types.Transaction, sender common.Address) error {
	gas := tx.Gas()
	nonce := tx.Nonce()
	accessList := tx.AccessList()
	txType := uint64(tx.Type())

	request := web3goTypes.CallRequest{
		From:              &sender,
		To:                tx.To(),
		Gas:               &gas,
		Value:             tx.Value(),
		Nonce:             &nonce,
		Data:              tx.Data(),
		AccessList:        &accessList,
		AuthorizationList: tx.SetCodeAuthorizations(),
		ChainID:           tx.ChainId(),
		Type:              &txType,
	}

	if txType == types.LegacyTxType {
		request.GasPrice = tx.GasPrice()
	} else {
		request.MaxFeePerGas = tx.GasFeeCap()
		request.MaxPriorityFeePerGas = tx.GasTipCap()
	}

	_, err := tp.client.Eth.Call(request, nil, nil, nil)

	return err
}

// estimateFundingETHGas estimates the required gas limit for the funding ETH transaction.
// This is because the sender may be delegated to a contract via 7702 transaction, and the
// 21000 gas limit is not enough in this case.
func (tp *TokenPay) estimateFundingETHGas(to common.Address, value *big.Int) (*big.Int, error) {
	return tp.client.Eth.EstimateGas(web3goTypes.CallRequest{
		From:  &tp.TxSender.sender,
		To:    &to,
		Value: value,
	}, nil, nil, nil)
}
