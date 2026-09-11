package contract

import (
	"math/big"
	"slices"
	"time"

	"github.com/Conflux-Chain/fluent-backend/util"
	"github.com/ethereum/go-ethereum/common"
)

// MinSignablePaymasterAndDataLen is the minimum length of the paymasterAndData field that can be signed,
// encoded as: paymasterAndGasLimits(52) || customData || validAfter(6) || validUntil(6) || signature(65).
const MinSignablePaymasterAndDataLen = MinPaymasterAndDataLen + 77

// DummySignature is a placeholder signature used in stub paymasterAndData for gas estimation.
var DummySignature = slices.Repeat([]byte{0x1b}, 65)

// GeneratePaymasterAndDataStub creates a stub paymasterAndData for gas estimation, including the paymaster address, custom data, validUntil, and a dummy signature.
func GeneratePaymasterAndDataStub(paymaster common.Address, timeout time.Duration, customData interface{ Bytes() []byte }) []byte {
	dataBytes := customData.Bytes()
	dataLen := len(dataBytes)
	size := MinSignablePaymasterAndDataLen + dataLen

	buf := make([]byte, size)

	validUntil := time.Now().Add(timeout).Unix()

	copy(buf[:20], paymaster.Bytes())                                   // address
	copy(buf[52:52+dataLen], dataBytes)                                 // custom data
	util.SafeBigFillBytes(big.NewInt(validUntil), buf[size-71:size-65]) // validUntil
	copy(buf[size-65:], DummySignature)                                 // dummy signature

	return buf
}

// PaymasterCustomData extracts the custom data from the paymasterAndData field of the user operation.
func (userOp *PackedUserOperation) PaymasterCustomData() []byte {
	if len(userOp.PaymasterAndData) < MinSignablePaymasterAndDataLen {
		return nil
	}

	return userOp.PaymasterAndData[52 : len(userOp.PaymasterAndData)-77]
}

func (userOp *PackedUserOperation) UpdateCustomData(customData []byte) bool {
	if len(userOp.PaymasterAndData) != MinSignablePaymasterAndDataLen+len(customData) {
		return false
	}

	copy(userOp.PaymasterAndData[52:52+len(customData)], customData)

	return true
}

// UpdatePaymasterPreSign updates the validAfter and validUntil fields in the paymasterAndData for pre-signing, and sets a dummy signature. Returns true if successful.
func (userOp *PackedUserOperation) UpdatePaymasterPreSign(timeout time.Duration) bool {
	if len(userOp.PaymasterAndData) < MinSignablePaymasterAndDataLen {
		return false
	}

	size := len(userOp.PaymasterAndData)
	validUntil := time.Now().Add(timeout).Unix()

	util.SafeBigFillBytes(big.NewInt(0), userOp.PaymasterAndData[size-77:size-71])          // validAfter
	util.SafeBigFillBytes(big.NewInt(validUntil), userOp.PaymasterAndData[size-71:size-65]) // validUntil
	copy(userOp.PaymasterAndData[size-65:], DummySignature)                                 // dummy signature

	return true
}

// UpdatePaymasterSignature updates the signature field in the paymasterAndData with the provided signature. Returns true if successful.
func (userOp *PackedUserOperation) UpdatePaymasterSignature(signature []byte) bool {
	if len(userOp.PaymasterAndData) < MinSignablePaymasterAndDataLen || len(signature) != 65 {
		return false
	}

	size := len(userOp.PaymasterAndData)
	copy(userOp.PaymasterAndData[size-65:], signature)

	return true
}
