package service

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
)

type GasTankPaymaster struct {
	inner       *Paymaster[*contract.GasTankPaymasterCaller]
	priceOracle *PriceOracle
}

func NewGasTankPaymaster(config PaymasterConfig, client *web3go.Client, priceOracle *PriceOracle) (*GasTankPaymaster, error) {
	paymaster, err := NewPaymaster(config, client, contract.NewGasTankPaymasterCaller, "GasTankPaymaster")
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create GasTankPaymaster")
	}

	return &GasTankPaymaster{
		inner:       paymaster,
		priceOracle: priceOracle,
	}, nil
}

// Stub generates a stub paymaster and data for the given sender and token.
// It checks if the token is allowed and if the sender has sufficient available balance.
func (paymaster *GasTankPaymaster) Stub(sender, token common.Address) ([]byte, error) {
	// check if token is allowed
	tokenAllowed, err := paymaster.inner.caller.IsTokenAllowed(nil, token)
	if err != nil {
		return nil, NewRPCError(err, "Failed to check if token is allowed")
	}

	if !tokenAllowed {
		return nil, ErrGasTankTokenNotAllowed
	}

	// check sender available balance (excluding the withdraw amount)
	balance, err := paymaster.getAvailableTokenBalance(sender, token)
	if err != nil {
		return nil, err
	}

	if balance.Sign() == 0 {
		return nil, ErrGasTankInsufficientBalance
	}

	return paymaster.inner.generateStub(GasTankData{
		Token:        token,
		MaxTokenCost: balance,
	}), nil
}

// getAvailableTokenBalance retrieves the available token balance for a sender, excluding the withdraw amount.
func (paymaster *GasTankPaymaster) getAvailableTokenBalance(sender, token common.Address) (*big.Int, error) {
	account, err := paymaster.inner.caller.Accounts(nil, sender, token)
	if err != nil {
		return nil, NewRPCError(err, "Failed to retrieve sender token balance")
	}

	balance := new(big.Int).Sub(account.Balance, account.WithdrawAmount)
	if balance.Sign() < 0 {
		balance.SetInt64(0)
	}

	return balance, nil
}

// Sign signs the given user operation with the paymaster's signature. It validates the user operation,
// updates the max token cost, reconstructs the paymaster data, and computes the paymaster signature.
func (paymaster *GasTankPaymaster) Sign(userOp contract.PackedUserOperation) ([]byte, error) {
	// validate the user operation and parse the gas tank data
	gasTankData, err := paymaster.validate(&userOp)
	if err != nil {
		return nil, err
	}

	// calculate the max token cost
	maxTokenCost, err := paymaster.calculateMaxTokenCost(&userOp, gasTankData.Token)
	if err != nil {
		return nil, err
	}

	// validate the sender available balance
	balance, err := paymaster.getAvailableTokenBalance(userOp.Sender, gasTankData.Token)
	if err != nil {
		return nil, err
	}

	if balance.Cmp(maxTokenCost) < 0 {
		return nil, ErrGasTankInsufficientBalance.WithData(fmt.Sprintf("maxTokenCost = %v, available balance = %v", maxTokenCost, balance))
	}

	// update the max token cost and sign
	gasTankData.MaxTokenCost = maxTokenCost

	return paymaster.inner.sign(userOp, gasTankData)
}

// validate the userOp and return the parsed GasTankData if valid.
func (paymaster *GasTankPaymaster) validate(userOp *contract.PackedUserOperation) (*GasTankData, error) {
	// InitCode should be empty or match the 7702 marker
	if len(userOp.InitCode) > 0 && hexutil.Encode(userOp.InitCode) != initCode7702Marker {
		return nil, api.ErrValidationStr("Invalid InitCode")
	}

	// validate the paymasterAndData field and extract the custom data.
	customData, err := paymaster.inner.validatePaymasterAndData(userOp)
	if err != nil {
		return nil, err
	}

	// parse gas tank data
	gasTankData, err := ParseGasTankData(customData)
	if err != nil {
		return nil, err
	}

	// validate token
	allowed, err := paymaster.inner.caller.IsTokenAllowed(nil, gasTankData.Token)
	if err != nil {
		return nil, NewRPCError(err, "Failed to check if token is allowed")
	}

	if !allowed {
		return nil, ErrGasTankTokenNotAllowed.WithData(gasTankData.Token)
	}

	return &gasTankData, nil
}

// calculateMaxTokenCost calculates the maximum token cost for a given user operation and token.
func (paymaster *GasTankPaymaster) calculateMaxTokenCost(userOp *contract.PackedUserOperation, token common.Address) (*big.Int, error) {
	price, err := paymaster.priceOracle.GetETHPrice(token)
	if err != nil {
		return nil, err
	}

	maxGasCost := userOp.MaxGasCost()
	maxTokenCost := new(big.Int).Mul(maxGasCost, price)
	maxTokenCost.Div(maxTokenCost, big10Exp18)

	return maxTokenCost, nil
}

type GasTankData struct {
	Token        common.Address
	MaxTokenCost *big.Int
}

func (data GasTankData) Bytes() []byte {
	var buf [52]byte

	copy(buf[:20], data.Token.Bytes())
	contract.SafeBigFillBytes(data.MaxTokenCost, buf[20:52])

	return buf[:]
}

func ParseGasTankData(data []byte) (GasTankData, error) {
	if len(data) != 52 {
		return GasTankData{}, api.ErrValidationStrf("Invalid GasTankData length, expected 52, got %d", len(data))
	}

	return GasTankData{
		Token:        common.BytesToAddress(data[:20]),
		MaxTokenCost: new(big.Int).SetBytes(data[20:52]),
	}, nil
}
