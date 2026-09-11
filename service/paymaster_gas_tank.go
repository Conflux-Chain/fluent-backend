package service

import (
	"bytes"
	"fmt"
	"math/big"
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

func (paymaster *GasTankPaymaster) StubCredit(token common.Address, amount *big.Int) ([]byte, error) {
	// check if token is allowed
	tokenAllowed, err := paymaster.gasTankCaller.IsTokenAllowed(nil, token)
	if err != nil {
		return nil, NewRPCError(err, "Failed to check if token is allowed")
	}

	if !tokenAllowed {
		return nil, api.ErrValidationStr("Token is not allowed")
	}

	return contract.GeneratePaymasterAndDataStub(
		paymaster.config.Address,
		paymaster.config.SignatureTimeout,
		GasTankData{
			Mode:         gasTankPaymasterModeCredit,
			Token:        token,
			MaxTokenCost: amount,
		},
	), nil
}

func (paymaster *GasTankPaymaster) StubRefund(sender, token common.Address) ([]byte, error) {
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

	return contract.GeneratePaymasterAndDataStub(
		paymaster.config.Address,
		paymaster.config.SignatureTimeout,
		GasTankData{
			Mode:         gasTankPaymasterModeRefund,
			Token:        token,
			MaxTokenCost: balance,
		},
	), nil
}

func (paymaster *GasTankPaymaster) Sign(userOp contract.PackedUserOperation) ([]byte, error) {
	// EOA should be upgraded to smart account via 7702 transaction
	if len(userOp.InitCode) > 0 {
		return nil, api.ErrValidationStr("InitCode is not supported")
	}

	// Parse and validate paymaster data.
	gasTankData, err := ParseGasTankData(userOp.PaymasterCustomData())
	if err != nil {
		return nil, err
	}

	if err = paymaster.validatePaymasterData(&userOp, &gasTankData); err != nil {
		return nil, err
	}

	// calculate the max token cost
	maxTokenCost, err := paymaster.calculateMaxTokenCost(&userOp, &gasTankData)
	if err != nil {
		return nil, err
	}

	// validate sender balance in refund mode
	if gasTankData.Mode == gasTankPaymasterModeRefund {
		if err = paymaster.validateSenderBalance(userOp.Sender, gasTankData.Token, maxTokenCost); err != nil {
			return nil, err
		}
	}

	// validate calldata in credit mode
	if gasTankData.Mode == gasTankPaymasterModeCredit {
		if err = paymaster.validateCallData(&userOp, &gasTankData, maxTokenCost); err != nil {
			return nil, err
		}
	}

	// re-construct the paymaster data in user op
	gasTankData.MaxTokenCost = maxTokenCost
	userOp.UpdateCustomData(gasTankData.Bytes())
	userOp.UpdatePaymasterPreSign(paymaster.config.SignatureTimeout)

	// compute the paymaster signature
	hash, err := paymaster.gasTankCaller.GetPaymasterHash(nil, userOp)
	if err != nil {
		return nil, NewRPCError(err, "Failed to retrieve paymaster hash from blockchain")
	}

	signature, err := paymaster.signer.SignHash(hash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to sign paymaster hash")
	}

	userOp.UpdatePaymasterSignature(signature)

	return userOp.PaymasterAndData, nil
}

func (paymaster *GasTankPaymaster) calculateMaxTokenCost(userOp *contract.PackedUserOperation, gasTankData *GasTankData) (*big.Int, error) {
	price, err := paymaster.priceOracle.GetETHPrice(gasTankData.Token)
	if err != nil {
		return nil, err
	}

	maxGasCost := userOp.MaxGasCost()
	maxTokenCost := new(big.Int).Mul(maxGasCost, price)
	maxTokenCost.Div(maxTokenCost, big10Exp18)

	return maxTokenCost, nil
}

func (paymaster *GasTankPaymaster) validatePaymasterData(userOp *contract.PackedUserOperation, gasTankData *GasTankData) error {
	// address
	if userOp.Paymaster() != paymaster.config.Address {
		return api.ErrValidationStr("Invalid paymaster address")
	}

	// mode
	if gasTankData.Mode != gasTankPaymasterModeRefund && gasTankData.Mode != gasTankPaymasterModeCredit {
		return api.ErrValidationStr("Invalid paymaster mode")
	}

	// token
	allowed, err := paymaster.gasTankCaller.IsTokenAllowed(nil, gasTankData.Token)
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

func (paymaster *GasTankPaymaster) validateCallData(userOp *contract.PackedUserOperation, gasTankData *GasTankData, maxTokenCost *big.Int) error {
	depositAmount := gasTankData.MaxTokenCost

	if depositAmount.Cmp(maxTokenCost) < 0 {
		return ErrGasTankInsufficientBalance.WithData(fmt.Sprintf("maxTokenCost = %v, depositAmount = %v", maxTokenCost, depositAmount))
	}

	packedCallData, err := paymaster.packApproveAndDeposit(gasTankData.Token, depositAmount)
	if err != nil {
		return errors.WithMessage(err, "Failed to pack approve + deposit calldata")
	}

	if !bytes.Equal(userOp.CallData, packedCallData) {
		return api.ErrValidationStr("Invalid user operation calldata")
	}

	erc20Caller, err := paymaster.loadOrCreateERC20Caller(gasTankData.Token)
	if err != nil {
		return errors.WithMessage(err, "Failed to load or create ERC20 caller")
	}

	balance, err := erc20Caller.BalanceOf(nil, userOp.Sender)
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

type GasTankData struct {
	Mode         byte
	Token        common.Address
	MaxTokenCost *big.Int
}

func (data GasTankData) Bytes() []byte {
	var buf [53]byte

	buf[0] = data.Mode
	copy(buf[1:21], data.Token.Bytes())
	if data.MaxTokenCost != nil {
		data.MaxTokenCost.FillBytes(buf[21:53])
	}

	return buf[:]
}

func ParseGasTankData(data []byte) (GasTankData, error) {
	if len(data) != 53 {
		return GasTankData{}, api.ErrValidationStrf("Invalid GasTankData length, expected 53, got %d", len(data))
	}

	return GasTankData{
		Mode:         data[0],
		Token:        common.BytesToAddress(data[1:21]),
		MaxTokenCost: new(big.Int).SetBytes(data[21:53]),
	}, nil
}
