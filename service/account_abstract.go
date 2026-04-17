package service

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	emptyAddress        = common.Address{}
	delegatedCodePrefix = "0xef0100" // EIP-7702 standard
	setCodeTxTo         = common.HexToAddress("0x7702770277027702770277027702770277027702")
)

type AccountAbstract struct {
	client *web3go.Client

	sender            common.Address
	chainId           uint64
	delegatedContract common.Address

	mu sync.Mutex
}

func NewAccountAbstract(client *web3go.Client, delegatedContract common.Address) (*AccountAbstract, error) {
	// get the default signer
	sm, err := client.GetSignerManager()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get signer manager from RPC client")
	}

	signers := sm.List()
	if len(signers) == 0 {
		return nil, errors.New("No signer found")
	}

	// check sender balance
	sender := signers[0].Address()
	balance, err := client.Eth.Balance(sender, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve balance of AA tx sender")
	}

	if balance.Sign() == 0 {
		return nil, errors.Errorf("AA tx sender balance is 0, address = %v", sender)
	}

	// retrieve chain ID
	chainId, err := client.Eth.ChainId()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve chain ID")
	}

	if chainId == nil {
		return nil, errors.New("Chain ID unavailable on fullnode")
	}

	return &AccountAbstract{
		client:            client,
		sender:            signers[0].Address(),
		chainId:           *chainId,
		delegatedContract: delegatedContract,
	}, nil
}

func (aa *AccountAbstract) SendSetCodeTransaction(auth gethTypes.SetCodeAuthorization) (common.Hash, error) {
	if err := aa.validateAuth(auth); err != nil {
		return common.Hash{}, err
	}

	tx := types.TransactionArgs{
		From:              &aa.sender,
		To:                &setCodeTxTo,
		AuthorizationList: []gethTypes.SetCodeAuthorization{auth},
		ChainID:           (*hexutil.Big)(new(big.Int).SetUint64(aa.chainId)),
	}

	// use mutex to ensure correct tx nonce, and need to enhance for high QPS case
	aa.mu.Lock()
	defer aa.mu.Unlock()

	txHash, err := aa.client.Eth.SendTransactionByArgs(tx)
	if err != nil {
		return common.Hash{}, NewRPCError(err, "Failed to send transaction")
	}

	return txHash, nil
}

func (aa *AccountAbstract) validateAuth(auth gethTypes.SetCodeAuthorization) error {
	// validate chain ID
	if chainId := auth.ChainID.Uint64(); chainId > 0 && chainId != aa.chainId {
		return api.ErrValidationStrf("Invalid chain ID, expected = %v, got = %v", aa.chainId, chainId)
	}

	// recover authority
	authority, err := auth.Authority()
	if err != nil {
		return api.ErrValidationStrf("Failed to recover authority from signature: %v", err.Error())
	}

	// validate nonce against the latest nonce
	nonce, err := aa.client.Eth.TransactionCount(authority, nil)
	if err != nil {
		return NewRPCError(err, "Failed to retrieve authority nonce")
	}

	if auth.Nonce != nonce.Uint64() {
		return api.ErrValidationStrf("Invalid nonce, expected = %v, got = %v", nonce.Uint64(), auth.Nonce)
	}

	// validate delegated contract
	onChainCode, err := aa.getDelegatedContract(authority)
	if err != nil {
		return ErrRPCError.WithData(err.Error())
	}

	if auth.Address == emptyAddress {
		// no delegated contract yet
		if onChainCode == emptyAddress {
			return api.ErrValidationStr("Authority is not delegated to any contract yet")
		}

		return nil
	}

	// restrict delegated contract if required
	if aa.delegatedContract != emptyAddress && auth.Address != aa.delegatedContract {
		return api.ErrValidationStrf("Invalid delegated contract, expected = %v, got = %v", aa.delegatedContract, auth.Address)
	}

	// already delegated
	if auth.Address == onChainCode {
		return api.ErrValidationStr("Authority is already delegated to the specified contract")
	}

	return nil
}

func (aa *AccountAbstract) getDelegatedContract(authority common.Address) (common.Address, error) {
	code, err := aa.client.Eth.CodeAt(authority, nil)
	if err != nil {
		return emptyAddress, errors.WithMessage(err, "Failed to retrieve authority code")
	}

	codeLen := len(code)
	if codeLen == 0 {
		return emptyAddress, nil
	}

	if codeLen != 23 {
		return emptyAddress, fmt.Errorf(
			"Invalid code length, expected = 23, got = %v, authority = %v, code = %v",
			codeLen, authority, hexutil.Encode(code),
		)
	}

	if prefix := hexutil.Encode(code[0:3]); prefix != delegatedCodePrefix {
		return emptyAddress, fmt.Errorf(
			"Invalid code prefix, expected = %v, got = %v, authority = %v, code = %v",
			delegatedCodePrefix, prefix, authority, hexutil.Encode(code),
		)
	}

	return common.BytesToAddress(code[3:]), nil
}

func (aa *AccountAbstract) GetSetCodeResult(txHash common.Hash) (types.LocalizedSetAuthTrace, bool, error) {
	tx, err := aa.client.Eth.TransactionByHash(txHash)
	if err != nil {
		return types.LocalizedSetAuthTrace{}, false, NewRPCError(err, "Failed to retrieve transaction by hash %v", txHash)
	}

	if tx == nil {
		return types.LocalizedSetAuthTrace{}, false, ErrAccountAbstractTxNotFound.WithData(txHash)
	}

	if tx.From != aa.sender {
		return types.LocalizedSetAuthTrace{}, false, api.ErrValidationStrf("Invalid tx sender, expected = %v, got = %v", aa.sender, tx.From)
	}

	receipt, err := aa.client.Eth.TransactionReceipt(txHash)
	if err != nil {
		return types.LocalizedSetAuthTrace{}, false, NewRPCError(err, "Failed to retrieve transaction receipt")
	}

	if receipt == nil {
		return types.LocalizedSetAuthTrace{}, false, nil
	}

	// check tx outcome
	if receipt.Status == nil {
		logrus.WithField("tx", txHash).Error("Receipt status is nil")
	} else if *receipt.Status != gethTypes.ReceiptStatusSuccessful {
		logger := logrus.WithField("tx", txHash)

		if receipt.TxExecErrorMsg != nil {
			logger.WithField("error", *receipt.TxExecErrorMsg)
		}

		logger.Error("Set code transaction failed")
	}

	// check auth result
	bn := types.BlockNumberOrHashWithNumber(types.BlockNumber(receipt.BlockNumber))
	auths, err := aa.client.Trace.BlockSetAuthTraces(bn)
	if err != nil {
		return types.LocalizedSetAuthTrace{}, false, NewRPCError(err, "Failed to retrieve block auths")
	}

	for _, v := range auths {
		if v.TransactionHash != txHash {
			continue
		}

		if v.BlockHash != receipt.BlockHash {
			return types.LocalizedSetAuthTrace{}, false, ErrRPCError.WithData("Chain reorg detected due to block hash mismatch")
		}

		return v, true, nil
	}

	return types.LocalizedSetAuthTrace{}, false, ErrRPCError.WithData("Chain reorg detected due to auth not found")
}
