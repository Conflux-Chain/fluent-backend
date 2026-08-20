package service

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/interfaces"
	"github.com/pkg/errors"
)

const (
	gasTankPaymasterModeRefund = byte(0)
	gasTankPaymasterModeCredit = byte(1)
)

var dummySignature = slices.Repeat([]byte{0x1b}, 65)

type GasTankPaymasterConfig struct {
	Address common.Address

	SignatureTimeout time.Duration `default:"5m"`
}

type GasTankPaymaster struct {
	config GasTankPaymasterConfig

	priceOracle *PriceOracle

	caller           bind.ContractCaller
	erc20CallerCache sync.Map
	gasTankCaller    *contract.GasTankPaymasterCaller
	signer           interfaces.Signer

	erc20ABI        abi.ABI
	gasTankABI      abi.ABI
	smartAccountABI abi.ABI
}

func NewGasTankPaymaster(config GasTankPaymasterConfig, priceOracle *PriceOracle, client *web3go.Client) (*GasTankPaymaster, error) {
	// validate config
	if config.Address == (common.Address{}) {
		return nil, errors.New("GasTankPaymaster address is required")
	}

	// init ABI
	erc20ABI, err := abi.JSON(strings.NewReader(contract.ERC20MetaData.ABI))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse ERC20 ABI")
	}

	gasTankABI, err := abi.JSON(strings.NewReader(contract.GasTankPaymasterMetaData.ABI))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse GasTankPaymaster ABI")
	}

	smartAccountABI, err := abi.JSON(strings.NewReader(contract.SimpleSmartAccount7702MetaData.ABI))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse SimpleSmartAccount7702 ABI")
	}

	// get the default signer
	sm, err := client.GetSignerManager()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get signer manager from RPC client")
	}

	signers := sm.List()
	if len(signers) == 0 {
		return nil, errors.New("No signer found")
	}

	// init contract caller, note, the 2nd return value signer fn is ignored
	caller, _ := client.ToClientForContract()
	gasTankCaller, err := contract.NewGasTankPaymasterCaller(config.Address, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create GasTankPaymasterCaller")
	}

	return &GasTankPaymaster{
		config:          config,
		priceOracle:     priceOracle,
		caller:          caller,
		gasTankCaller:   gasTankCaller,
		signer:          signers[0],
		erc20ABI:        erc20ABI,
		gasTankABI:      gasTankABI,
		smartAccountABI: smartAccountABI,
	}, nil
}

func (paymaster *GasTankPaymaster) PrepareCredit(token common.Address, amount *big.Int) ([]byte, error) {
	// check if token is allowed
	tokenAllowed, err := paymaster.gasTankCaller.IsTokenAllowed(nil, token)
	if err != nil {
		return nil, NewRPCError(err, "Failed to check if token is allowed")
	}

	if !tokenAllowed {
		return nil, api.ErrValidationStr("Token is not allowed")
	}

	data := GasTankPaymasterData{
		Address:      paymaster.config.Address,
		Mode:         gasTankPaymasterModeCredit,
		Token:        token,
		MaxTokenCost: amount,
		ValidUntil:   time.Now().Add(paymaster.config.SignatureTimeout).Unix(),
		Signature:    dummySignature,
	}

	return data.encode(), nil
}

func (paymaster *GasTankPaymaster) PrepareRefund(sender, token common.Address) ([]byte, error) {
	// check if token is allowed
	tokenAllowed, err := paymaster.gasTankCaller.IsTokenAllowed(nil, token)
	if err != nil {
		return nil, NewRPCError(err, "Failed to check if token is allowed")
	}

	if !tokenAllowed {
		return nil, api.ErrValidationStr("Token is not allowed")
	}

	// check sender balance
	balance, err := paymaster.gasTankCaller.BalanceOf(nil, sender, token)
	if err != nil {
		return nil, NewRPCError(err, "Failed to retrieve sender token balance")
	}

	if balance.Sign() == 0 {
		return nil, api.ErrValidationStr("Insufficient token balance")
	}

	data := GasTankPaymasterData{
		Address:      paymaster.config.Address,
		Mode:         gasTankPaymasterModeRefund,
		Token:        token,
		MaxTokenCost: balance,
		ValidUntil:   time.Now().Add(paymaster.config.SignatureTimeout).Unix(),
		Signature:    dummySignature,
	}

	return data.encode(), nil
}

func (paymaster *GasTankPaymaster) Sign(userOp contract.PackedUserOperation) ([]byte, error) {
	// EOA should be upgraded to smart account via 7702 transaction
	if len(userOp.InitCode) > 0 {
		return nil, api.ErrValidationStr("InitCode is not supported")
	}

	// Parse and validate paymaster data.
	paymasterData, err := parseGasTankPaymasterData(userOp.PaymasterAndData)
	if err != nil {
		return nil, err
	}

	if err = paymaster.validatePaymasterData(&paymasterData); err != nil {
		return nil, err
	}

	// calculate the max token cost
	maxTokenCost, err := paymaster.calculateMaxTokenCost(&userOp, &paymasterData)
	if err != nil {
		return nil, err
	}

	// validate sender balance in refund mode
	if paymasterData.Mode == gasTankPaymasterModeRefund {
		if err = paymaster.validateSenderBalance(userOp.Sender, paymasterData.Token, maxTokenCost); err != nil {
			return nil, err
		}
	}

	// validate calldata in credit mode
	if paymasterData.Mode == gasTankPaymasterModeCredit {
		if err = paymaster.validateCallData(userOp.Sender, userOp.CallData, &paymasterData, maxTokenCost); err != nil {
			return nil, err
		}
	}

	// re-construct the paymaster data in user op
	paymasterData.MaxTokenCost = maxTokenCost
	paymasterData.ValidAfter = 0
	paymasterData.ValidUntil = time.Now().Add(paymaster.config.SignatureTimeout).Unix()
	paymasterData.Signature = dummySignature
	userOp.PaymasterAndData = paymasterData.encode()

	// compute the paymaster signature
	hash, err := paymaster.gasTankCaller.GetPaymasterHash(nil, userOp)
	if err != nil {
		return nil, NewRPCError(err, "Failed to retrieve paymaster hash from blockchain")
	}

	if paymasterData.Signature, err = paymaster.signer.SignHash(hash); err != nil {
		return nil, errors.WithMessage(err, "Failed to sign paymaster hash")
	}

	return paymasterData.encode(), nil
}

func (paymaster *GasTankPaymaster) calculateMaxTokenCost(userOp *contract.PackedUserOperation, paymasterData *GasTankPaymasterData) (*big.Int, error) {
	gasLimit := new(big.Int).Add(paymasterData.ValidationGas, paymasterData.PostOpGas)
	gasLimit.Add(gasLimit, userOp.PreVerificationGas)
	gasLimit.Add(gasLimit, new(big.Int).SetBytes(userOp.AccountGasLimits[0:16]))
	gasLimit.Add(gasLimit, new(big.Int).SetBytes(userOp.AccountGasLimits[16:32]))

	maxFeePerGas := new(big.Int).SetBytes(userOp.GasFees[16:32])
	maxGasCost := new(big.Int).Mul(maxFeePerGas, gasLimit)

	price, err := paymaster.priceOracle.GetETHPrice(paymasterData.Token)
	if err != nil {
		return nil, err
	}

	maxTokenCost := new(big.Int).Mul(maxGasCost, price)
	maxTokenCost.Div(maxTokenCost, big10Exp18)

	return maxTokenCost, nil
}

func (paymaster *GasTankPaymaster) validatePaymasterData(paymasterData *GasTankPaymasterData) error {
	// address
	if paymasterData.Address != paymaster.config.Address {
		return api.ErrValidationStr("Invalid paymaster address")
	}

	// mode
	if paymasterData.Mode != gasTankPaymasterModeRefund && paymasterData.Mode != gasTankPaymasterModeCredit {
		return api.ErrValidationStr("Invalid paymaster mode")
	}

	// token
	allowed, err := paymaster.gasTankCaller.IsTokenAllowed(nil, paymasterData.Token)
	if err != nil {
		return NewRPCError(err, "Failed to check if token is allowed")
	}

	if !allowed {
		return api.ErrValidationStr("Token is not allowed")
	}

	return nil
}

func (paymaster *GasTankPaymaster) validateSenderBalance(sender, token common.Address, maxTokenCost *big.Int) error {
	balance, err := paymaster.gasTankCaller.BalanceOf(nil, sender, token)
	if err != nil {
		return NewRPCError(err, "Failed to retrieve sender token balance")
	}

	if balance.Cmp(maxTokenCost) < 0 {
		return ErrGasTankInsufficientBalance.WithData(fmt.Sprintf("maxTokenCost = %v, balance = %v", maxTokenCost, balance))
	}

	return nil
}

func (paymaster *GasTankPaymaster) validateCallData(sender common.Address, userOpCallData []byte, paymasterData *GasTankPaymasterData, maxTokenCost *big.Int) error {
	depositAmount := paymasterData.MaxTokenCost

	if depositAmount.Cmp(maxTokenCost) < 0 {
		return ErrGasTankInsufficientBalance.WithData(fmt.Sprintf("maxTokenCost = %v, depositAmount = %v", maxTokenCost, depositAmount))
	}

	packedCallData, err := paymaster.packApproveAndDeposit(paymasterData.Token, depositAmount)
	if err != nil {
		return errors.WithMessage(err, "Failed to pack approve + deposit calldata")
	}

	if !bytes.Equal(userOpCallData, packedCallData) {
		return api.ErrValidationStr("Invalid user operation calldata")
	}

	erc20Caller, err := paymaster.loadOrCreateERC20Caller(paymasterData.Token)
	if err != nil {
		return errors.WithMessage(err, "Failed to load or create ERC20 caller")
	}

	balance, err := erc20Caller.BalanceOf(nil, sender)
	if err != nil {
		return NewRPCError(err, "Failed to retrieve sender token balance")
	}

	if balance.Cmp(depositAmount) < 0 {
		return ErrGasTankInsufficientBalance.WithData(fmt.Sprintf("depositAmount = %v, balance = %v", depositAmount, balance))
	}

	return nil
}

func (paymaster *GasTankPaymaster) packApproveAndDeposit(token common.Address, amount *big.Int) ([]byte, error) {
	approveCallData, err := paymaster.erc20ABI.Pack("approve", paymaster.config.Address, amount)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to pack approve call data")
	}

	depositTokenCallData, err := paymaster.gasTankABI.Pack("depositToken", token, amount)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to pack depositToken call data")
	}

	calls := []contract.Execution{
		{
			Target:   token,
			Value:    big.NewInt(0),
			CallData: approveCallData,
		},
		{
			Target:   paymaster.config.Address,
			Value:    big.NewInt(0),
			CallData: depositTokenCallData,
		},
	}

	executeBatchCallData, err := paymaster.smartAccountABI.Pack("executeBatch", calls)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to pack executeBatch call data")
	}

	return executeBatchCallData, nil
}

func (paymaster *GasTankPaymaster) loadOrCreateERC20Caller(token common.Address) (*contract.ERC20Caller, error) {
	if caller, ok := paymaster.erc20CallerCache.Load(token); ok {
		return caller.(*contract.ERC20Caller), nil
	}

	caller, err := contract.NewERC20Caller(token, paymaster.caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create ERC20 caller")
	}

	paymaster.erc20CallerCache.Store(token, caller)

	return caller, nil
}

type GasTankPaymasterData struct {
	Address       common.Address // 20
	ValidationGas *big.Int       // 16
	PostOpGas     *big.Int       // 16
	Mode          byte           // 1
	Token         common.Address // 20
	MaxTokenCost  *big.Int       // 32
	ValidAfter    int64          // 6
	ValidUntil    int64          // 6
	Signature     []byte         // 65
}

func parseGasTankPaymasterData(data []byte) (GasTankPaymasterData, error) {
	if len(data) != 182 {
		return GasTankPaymasterData{}, api.ErrValidationStrf("Invalid paymasterAndData length, expected 182, got %d", len(data))
	}

	return GasTankPaymasterData{
		Address:       common.BytesToAddress(data[0:20]),
		ValidationGas: new(big.Int).SetBytes(data[20:36]),
		PostOpGas:     new(big.Int).SetBytes(data[36:52]),
		Mode:          data[52],
		Token:         common.BytesToAddress(data[53:73]),
		MaxTokenCost:  new(big.Int).SetBytes(data[73:105]),
		ValidAfter:    new(big.Int).SetBytes(data[105:111]).Int64(),
		ValidUntil:    new(big.Int).SetBytes(data[111:117]).Int64(),
		Signature:     data[117:182],
	}, nil
}

func (data *GasTankPaymasterData) encode() []byte {
	var buf [182]byte

	copy(buf[0:20], data.Address.Bytes())
	if data.ValidationGas != nil {
		data.ValidationGas.FillBytes(buf[20:36])
	}
	if data.PostOpGas != nil {
		data.PostOpGas.FillBytes(buf[36:52])
	}
	buf[52] = data.Mode
	copy(buf[53:73], data.Token.Bytes())
	if data.MaxTokenCost != nil {
		data.MaxTokenCost.FillBytes(buf[73:105])
	}
	if data.ValidAfter > 0 {
		big.NewInt(data.ValidAfter).FillBytes(buf[105:111])
	}
	if data.ValidUntil > 0 {
		big.NewInt(data.ValidUntil).FillBytes(buf[111:117])
	}
	copy(buf[117:182], data.Signature)

	return buf[:]
}
