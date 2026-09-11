package api

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

type SetCodeAuth struct {
	// ChainID is the chain ID the authorization is bound to. 0 means the auth is valid on any chain.
	ChainId uint64 `json:"chainId"`
	// Contract is the 20-byte hex address (with 0x prefix) of the smart-contract to delegate the EOA to.
	// Set to "0x0000000000000000000000000000000000000000" to revoke an existing delegation.
	Contract string `json:"contract" binding:"required,hex,len=42"`
	// Nonce is the current on-chain nonce of the signing authority (EOA). Must match exactly.
	Nonce uint64 `json:"nonce"`
	// V is the recovery identifier of the EIP-7702 authorization signature (0 or 1).
	V uint8 `json:"v"`
	// R is the R component of the EIP-7702 authorization signature, as a 0x-prefixed hex string.
	R string `json:"r" binding:"required,hex,len=66"`
	// S is the S component of the EIP-7702 authorization signature, as a 0x-prefixed hex string.
	S string `json:"s" binding:"required,hex,len=66"`
}

// mustToGeth converts the SetCodeAuth to the Geth SetCodeAuthorization type.
//
// Note, the auth should be validated before calling this function, as it will panic if the R or S values are not valid hex strings.
func (auth *SetCodeAuth) mustToGeth() types.SetCodeAuthorization {
	rBytes32 := hexutil.MustDecode(auth.R)
	rU256, _ := uint256.FromBig(new(big.Int).SetBytes(rBytes32)) // never overflow

	sBytes32 := hexutil.MustDecode(auth.S)
	sU256, _ := uint256.FromBig(new(big.Int).SetBytes(sBytes32)) // never overflow

	return types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(auth.ChainId),
		Address: common.HexToAddress(auth.Contract),
		Nonce:   auth.Nonce,
		V:       auth.V,
		R:       *rU256,
		S:       *sU256,
	}
}

type SetCodeResult struct {
	// Executed indicates whether the transaction has been executed.
	Executed bool `json:"executed"`
	// Success indicates whether the set-code authorization succeeded. Only meaningful when Executed is true.
	Success bool `json:"success"`
	// Error contains the failure reason reported by the EVM when Success is false.
	Error string `json:"error"`
}

type GasTankPrepareCreditRequest struct {
	// ERC20 token address to deposit for gas fee payment.
	Token string `json:"token" binding:"required,hex,len=42"`
	// Amount of tokens to deposit for gas fee payment.
	Amount string `json:"amount" binding:"required,min=1,max=77"`
}

type GasTankPrepareRefundRequest struct {
	// Smart account address in hex format with 0x prefix.
	Sender string `json:"sender" binding:"required,hex,len=42"`
	// ERC20 token address to pay gas fee.
	Token string `json:"token" binding:"required,hex,len=42"`
}

type VerifyingPaymasterStubRequest struct {
	// Smart account address in hex format with 0x prefix.
	Sender string `json:"sender" form:"sender" binding:"required,hex,len=42"`
	// Delegated contract address in hex format with 0x prefix. Set to non-zero value if user op carries a 7702 auth message.
	Delegation string `json:"delegation" form:"delegation" binding:"required,hex,len=42"`
}

type PaymasterAndDataStub struct {
	Address string `json:"address"`
	Data    string `json:"data"`
}

func ToPaymasterAndDataStub(paymasterAndData []byte) (stub PaymasterAndDataStub) {
	if len(paymasterAndData) >= 20 {
		stub.Address = common.BytesToAddress(paymasterAndData[:20]).Hex()
	}

	if len(paymasterAndData) >= 52 {
		stub.Data = hexutil.Encode(paymasterAndData[52:])
	} else {
		stub.Data = "0x"
	}

	return
}

type UserOperation struct {
	Sender               string `json:"sender" binding:"required,hex,len=42"`
	Nonce                string `json:"nonce" binding:"required,hexadecimal"`
	Factory              string `json:"factory" binding:"omitempty,hex,len=42"`
	FactoryData          string `json:"factoryData" binding:"omitempty,hex,min=2"`
	CallData             string `json:"callData" binding:"required,hex,min=2"`
	VerificationGasLimit string `json:"verificationGasLimit" binding:"required,hexadecimal,max=32"`
	CallGasLimit         string `json:"callGasLimit" binding:"required,hexadecimal,max=32"`
	PreVerificationGas   string `json:"preVerificationGas" binding:"required,hexadecimal,max=32"`
	MaxFeePerGas         string `json:"maxFeePerGas" binding:"required,hexadecimal,max=32"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas" binding:"required,hexadecimal,max=32"`
	Signature            string `json:"signature" binding:"required,hex,len=132"`

	// Paymaster
	Paymaster                     string `json:"paymaster" binding:"required,hex,len=42"`
	PaymasterVerificationGasLimit string `json:"paymasterVerificationGasLimit" binding:"required,hexadecimal,max=32"`
	PaymasterPostOpGasLimit       string `json:"paymasterPostOpGasLimit" binding:"required,hexadecimal,max=32"`
	PaymasterData                 string `json:"paymasterData" binding:"required,hex,min=156"` // at least validAfter (6) || validUntil (6) || signature (65)
}

func hexToBig(hexNum string) *big.Int {
	if strings.HasPrefix(hexNum, "0x") || strings.HasPrefix(hexNum, "0X") {
		hexNum = hexNum[2:]
	}

	if len(hexNum)%2 == 1 {
		hexNum = "0" + hexNum
	}

	dataBytes, _ := hex.DecodeString(hexNum)

	if len(dataBytes) == 0 {
		return common.Big0
	}

	return new(big.Int).SetBytes(dataBytes)
}

func shouldDecodeHex(hex string) []byte {
	if len(hex) == 0 {
		return nil
	}

	if !strings.HasPrefix(hex, "0x") && !strings.HasPrefix(hex, "0X") {
		return nil
	}

	decoded, _ := hexutil.Decode(hex)

	return decoded
}

func (userOp *UserOperation) ToPackedUserOperation() contract.PackedUserOperation {
	var initCode []byte
	if len(userOp.Factory) > 2 {
		initCode = append(initCode, shouldDecodeHex(userOp.Factory)...)

		if len(userOp.FactoryData) > 2 {
			initCode = append(initCode, shouldDecodeHex(userOp.FactoryData)...)
		}
	}

	packed := contract.PackedUserOperation{
		Sender:             common.HexToAddress(userOp.Sender),
		Nonce:              hexToBig(userOp.Nonce),
		InitCode:           initCode,
		CallData:           shouldDecodeHex(userOp.CallData),
		PreVerificationGas: hexToBig(userOp.PreVerificationGas),
		Signature:          shouldDecodeHex(userOp.Signature),
	}

	packed.SetAccountGasLimits(hexToBig(userOp.VerificationGasLimit), hexToBig(userOp.CallGasLimit))
	packed.SetGasFees(hexToBig(userOp.MaxPriorityFeePerGas), hexToBig(userOp.MaxFeePerGas))
	packed.SetPaymasterAndData(
		common.HexToAddress(userOp.Paymaster),
		hexToBig(userOp.PaymasterVerificationGasLimit),
		hexToBig(userOp.PaymasterPostOpGasLimit),
		shouldDecodeHex(userOp.PaymasterData),
	)

	return packed
}

type TokenPayConfig struct {
	// Tokens is the list of ERC20 token contracts supported for token-pay. Note, the tokens[0] is the default USDT token used for quoting and payment.
	Tokens []string `json:"tokens"`
	// Recipient is the configured recipient address that receives token payments.
	Recipient string `json:"recipient"`
	// MinGasFeeRatio is the minimum allowed gas fee ratio (gas fee / base fee) in percentage. For example, 120 means the gas fee must be at least 120% of the base fee.
	MinGasFeeRatio uint64 `json:"minGasFeeRatio"`
	// MinGasTipRatio is the minimum allowed gas tip ratio (priority fee / base fee) in percentage. For example, 20 means the priority fee must be at least 20% of the base fee.
	MinGasTipRatio uint64 `json:"minGasTipRatio"`
	// MaxGasCost is the maximum allowed gas cost in wei.
	MaxGasCost uint64 `json:"maxGasCost"`
	// SuggestedGasPriceBumpRatio is the percentage by which to bump the suggested gas price to increase the chance of timely inclusion in blocks.
	SuggestedGasPriceBumpRatio uint64 `json:"suggestedGasPriceBumpRatio"`
	// SuggestedTokenPriceBumpRatio is the percentage by which to bump the suggested token price to increase the chance of timely inclusion in blocks when the price is volatile.
	SuggestedTokenPriceBumpRatio uint64 `json:"suggestedTokenPriceBumpRatio"`
}

func NewTokenPayConfig(priceConfig service.PriceConfig, tokenPayConfig service.TokenPayConfig) TokenPayConfig {
	var tokens []string

	for _, v := range priceConfig.USDT {
		tokens = append(tokens, v.Hex())
	}

	for _, v := range priceConfig.CNH {
		tokens = append(tokens, v.Hex())
	}

	return TokenPayConfig{
		Tokens:                       tokens,
		Recipient:                    tokenPayConfig.Recipient.Hex(),
		MinGasFeeRatio:               tokenPayConfig.MinGasFeeRatio,
		MinGasTipRatio:               tokenPayConfig.MinGasTipRatio,
		MaxGasCost:                   tokenPayConfig.MaxGasCost,
		SuggestedGasPriceBumpRatio:   tokenPayConfig.SuggestedGasPriceBumpRatio,
		SuggestedTokenPriceBumpRatio: tokenPayConfig.SuggestedTokenPriceBumpRatio,
	}
}

type TokenPayPriceRequest struct {
	Token string `json:"token" form:"token" binding:"required,hex,len=42"`
}

type TokenPayRequest struct {
	RawTransferTokenTx string `json:"rawTransferTokenTx" binding:"required,hex"`
	RawBusinessTx      string `json:"rawBusinessTx" binding:"required,hex"`
}
