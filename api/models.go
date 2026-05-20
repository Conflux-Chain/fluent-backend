package api

import (
	"math/big"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type SetCodeAuth struct {
	// ChainID is the chain ID the authorization is bound to. 0 means the auth is valid on any chain.
	ChainId uint64 `json:"chainId"`
	// Contract is the 20-byte hex address (with 0x prefix) of the smart-contract to delegate the EOA to.
	// Set to "0x0000000000000000000000000000000000000000" to revoke an existing delegation.
	Contract string `json:"contract" binding:"required,len=42"`
	// Nonce is the current on-chain nonce of the signing authority (EOA). Must match exactly.
	Nonce uint64 `json:"nonce"`
	// V is the recovery identifier of the EIP-7702 authorization signature (0 or 1).
	V uint8 `json:"v"`
	// R is the R component of the EIP-7702 authorization signature, as a 0x-prefixed hex string.
	R string `json:"r" binding:"required"`
	// S is the S component of the EIP-7702 authorization signature, as a 0x-prefixed hex string.
	S string `json:"s" binding:"required"`
}

type SetCodeResult struct {
	// Executed indicates whether the transaction has been executed.
	Executed bool `json:"executed"`
	// Success indicates whether the set-code authorization succeeded. Only meaningful when Executed is true.
	Success bool `json:"success"`
	// Error contains the failure reason reported by the EVM when Success is false.
	Error string `json:"error"`
}

type GasTankPrepareDepositRequest struct {
	// ERC20 token address to deposit for gas fee payment.
	Token string `json:"token" binding:"required,len=42"`
	// Amount of tokens to deposit for gas fee payment.
	Amount string `json:"amount" binding:"required"`
}

type GasTankPrepareRequest struct {
	// Smart account address in hex format with 0x prefix.
	Sender string `json:"sender" binding:"required,len=42"`
	// ERC20 token address to pay gas fee.
	Token string `json:"token" binding:"required,len=42"`
}

type UserOperation struct {
	Sender               string `json:"sender" binding:"hex,len=42"`
	Nonce                string `json:"nonce" binding:"hex,min=4"`
	CallData             string `json:"callData" binding:"hex"`
	VerificationGasLimit string `json:"verificationGasLimit" binding:"hex,min=4,max=34"`
	CallGasLimit         string `json:"callGasLimit" binding:"hex,min=4,max=34"`
	PreVerificationGas   string `json:"preVerificationGas" binding:"hex,min=4,max=34"`
	MaxFeePerGas         string `json:"maxFeePerGas" binding:"hex,min=4,max=34"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas" binding:"hex,min=4,max=34"`
	Signature            string `json:"signature" binding:"hex,len=132"`

	// Paymaster
	Paymaster                     string `json:"paymaster" binding:"hex,len=42"`
	PaymasterVerificationGasLimit string `json:"paymasterVerificationGasLimit" binding:"hex,min=4,max=34"`
	PaymasterPostOpGasLimit       string `json:"paymasterPostOpGasLimit" binding:"hex,min=4,max=34"`
	PaymasterData                 string `json:"paymasterData" binding:"hex,min=156"` // at least validAfter (6) || validUntil (6) || signature (65)
}

func hexToBig(hex string) *big.Int {
	dataBytes, _ := hexutil.Decode(hex)

	if len(dataBytes) == 0 {
		return common.Big0
	}

	return new(big.Int).SetBytes(dataBytes)
}

func (userOp *UserOperation) ToPackedUserOperation() contract.PackedUserOperation {
	var accountGasLimits [32]byte
	hexToBig(userOp.VerificationGasLimit).FillBytes(accountGasLimits[0:16])
	hexToBig(userOp.CallGasLimit).FillBytes(accountGasLimits[16:32])

	var gasFees [32]byte
	hexToBig(userOp.MaxPriorityFeePerGas).FillBytes(gasFees[0:16])
	hexToBig(userOp.MaxFeePerGas).FillBytes(gasFees[16:32])

	var paymasterBuf [52]byte
	copy(paymasterBuf[0:20], common.HexToAddress(userOp.Paymaster).Bytes())
	hexToBig(userOp.PaymasterVerificationGasLimit).FillBytes(paymasterBuf[20:36])
	hexToBig(userOp.PaymasterPostOpGasLimit).FillBytes(paymasterBuf[36:52])

	paymasterData, _ := hexutil.Decode(userOp.PaymasterData)

	return contract.PackedUserOperation{
		Sender:             common.HexToAddress(userOp.Sender),
		Nonce:              hexToBig(userOp.Nonce),
		CallData:           hexutil.MustDecode(userOp.CallData),
		AccountGasLimits:   accountGasLimits,
		PreVerificationGas: hexToBig(userOp.PreVerificationGas),
		GasFees:            gasFees,
		PaymasterAndData:   append(paymasterBuf[:], paymasterData...),
		Signature:          hexutil.MustDecode(userOp.Signature),
	}
}
