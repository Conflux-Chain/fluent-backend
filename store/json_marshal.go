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
		VerificationGasLimit: new(big.Int).SetBytes(userOp.AccountGasLimits[:16]),
		CallGasLimit:         new(big.Int).SetBytes(userOp.AccountGasLimits[16:]),
		PreVerificationGas:   userOp.PreVerificationGas,
		MaxPriorityFeePerGas: new(big.Int).SetBytes(userOp.GasFees[:16]),
		MaxFeePerGas:         new(big.Int).SetBytes(userOp.GasFees[16:]),
		Signature:            userOp.Signature,
	}

	if len(userOp.PaymasterAndData) >= 52 {
		result.Paymaster = common.BytesToAddress(userOp.PaymasterAndData[:20])
		result.PaymasterVerificationGasLimit = new(big.Int).SetBytes(userOp.PaymasterAndData[20:36])
		result.PaymasterPostOpGasLimit = new(big.Int).SetBytes(userOp.PaymasterAndData[36:52])
		result.PaymasterData = userOp.PaymasterAndData[52:]
	}

	return result
}
