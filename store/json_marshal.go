package store

import (
	"math/big"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type PackedUserOperation struct {
	Sender                        common.Address `json:"sender"`
	Nonce                         *big.Int       `json:"nonce"`
	InitCode                      hexutil.Bytes  `json:"initCode"`
	CallData                      hexutil.Bytes  `json:"callData"`
	VerificationGasLimit          *big.Int       `json:"verificationGasLimit"`
	CallGasLimit                  *big.Int       `json:"callGasLimit"`
	PreVerificationGas            *big.Int       `json:"preVerificationGas"`
	MaxPriorityFeePerGas          *big.Int       `json:"maxPriorityFeePerGas"`
	MaxFeePerGas                  *big.Int       `json:"maxFeePerGas"`
	Signature                     hexutil.Bytes  `json:"signature"`
	Paymaster                     common.Address `json:"paymaster"`
	PaymasterVerificationGasLimit *big.Int       `json:"paymasterVerificationGasLimit"`
	PaymasterPostOpGasLimit       *big.Int       `json:"paymasterPostOpGasLimit"`
	PaymasterData                 hexutil.Bytes  `json:"paymasterData"`
}

func convertPackedUserOp(userOp *contract.PackedUserOperation) PackedUserOperation {
	result := PackedUserOperation{
		Sender:               userOp.Sender,
		Nonce:                userOp.Nonce,
		InitCode:             userOp.InitCode,
		CallData:             userOp.CallData,
		VerificationGasLimit: userOp.VerificationGasLimit(),
		CallGasLimit:         userOp.CallGasLimit(),
		PreVerificationGas:   userOp.PreVerificationGas,
		MaxPriorityFeePerGas: userOp.MaxPriorityFeePerGas(),
		MaxFeePerGas:         userOp.MaxFeePerGas(),
		Signature:            userOp.Signature,
	}

	if len(userOp.PaymasterAndData) >= contract.MinPaymasterAndDataLen {
		result.Paymaster = userOp.Paymaster()
		result.PaymasterVerificationGasLimit = userOp.PaymasterVerificationGasLimit()
		result.PaymasterPostOpGasLimit = userOp.PaymasterPostOpGasLimit()
		result.PaymasterData = userOp.PaymasterData()
	}

	return result
}
