package service

import (
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/interfaces"
	"github.com/pkg/errors"
)

// minPaymasterAndDataLen defines the minimum length of the paymasterAndData field, including the address, custom data, validity period, and signature.
// Encoding format: paymasterAndGasLimits(52) || customData || validAfter(6) || validUntil(6) || signature(65).
const minPaymasterAndDataLen = contract.MinPaymasterAndDataLen + 77

// dummySignature is a placeholder signature used for gas estimation and initial user operation setup.
var dummySignature = slices.Repeat([]byte{0x1b}, 65)

// PaymasterContractCaller defines the interface that a paymaster contract caller must implement.
type PaymasterContractCaller interface {
	Paused(opts *bind.CallOpts) (bool, error)
	Balance(opts *bind.CallOpts) (*big.Int, error)
	GetPaymasterHash(opts *bind.CallOpts, userOp contract.PackedUserOperation) ([32]byte, error)
}

// PaymasterConfig holds the configuration for a paymaster, including its address and signature timeout.
type PaymasterConfig struct {
	Address          common.Address
	SignatureTimeout time.Duration `default:"5m"`
}

// Paymaster represents a paymaster contract instance with its configuration, caller, and signer.
// It provides methods to generate stub paymasterAndData for gas estimation and to sign user operations.
type Paymaster[T PaymasterContractCaller] struct {
	config PaymasterConfig
	caller T
	signer interfaces.Signer
}

// NewPaymaster creates a new Paymaster instance with the specified configuration, RPC client, and caller factory.
// It validates the configuration, retrieves the default signer, and initializes the paymaster contract caller.
func NewPaymaster[T PaymasterContractCaller](config PaymasterConfig, client *web3go.Client, callerFactory func(common.Address, bind.ContractCaller) (T, error)) (*Paymaster[T], error) {
	// validate config
	if config.Address == (common.Address{}) {
		return nil, errors.New("Invalid paymaster address")
	}

	if config.SignatureTimeout <= 0 {
		return nil, errors.New("Invalid signature timeout")
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

	// create paymaster contract caller
	caller, _ := client.ToClientForContract()
	contractCaller, err := callerFactory(config.Address, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create paymaster contract caller")
	}

	return &Paymaster[T]{
		config: config,
		caller: contractCaller,
		signer: signers[0],
	}, nil
}

// generateStub generates a stub paymasterAndData for gas estimation with the specified custom data.
func (paymaster *Paymaster[T]) generateStub(customData interface{ Bytes() []byte }) []byte {
	dataBytes := customData.Bytes()
	dataLen := len(dataBytes)
	size := minPaymasterAndDataLen + dataLen

	buf := make([]byte, size)

	validUntil := time.Now().Add(paymaster.config.SignatureTimeout).Unix()

	copy(buf[:20], paymaster.config.Address.Bytes())                        // address
	copy(buf[52:52+dataLen], dataBytes)                                     // custom data
	contract.SafeBigFillBytes(big.NewInt(validUntil), buf[size-71:size-65]) // validUntil
	copy(buf[size-65:], dummySignature)                                     // dummy signature

	return buf
}

// validatePaymasterAndData validates the paymasterAndData field of the given user operation.
// It checks the length and the paymaster address, and returns the custom data if valid.
func (paymaster *Paymaster[T]) validatePaymasterAndData(userOp *contract.PackedUserOperation) ([]byte, error) {
	// validate length
	size := len(userOp.PaymasterAndData)
	if size < minPaymasterAndDataLen {
		return nil, api.ErrValidationStr("Invalid paymasterAndData length")
	}

	// check paymaster address
	if userOp.Paymaster() != paymaster.config.Address {
		return nil, api.ErrValidationStrf("Invalid paymaster address: %s, expected %s", userOp.Paymaster(), paymaster.config.Address)
	}

	return userOp.PaymasterAndData[52 : size-77], nil
}

// sign updates the paymasterAndData field of the given user operation with the specified custom data if provided,
// and signs it using the paymaster's signer. It returns the updated paymasterAndData.
// Before signing, it will check if the paymaster contract is paused and if the deposit balance is sufficient.
func (paymaster *Paymaster[T]) sign(userOp contract.PackedUserOperation, customData ...interface{ Bytes() []byte }) ([]byte, error) {
	// check if paymaster contract paused
	paused, err := paymaster.caller.Paused(nil)
	if err != nil {
		return nil, NewRPCError(err, "Failed to check if paymaster contract is paused")
	}

	if paused {
		return nil, ErrPaymasterPaused
	}

	// check paymaster deposit balance
	balance, err := paymaster.caller.Balance(nil)
	if err != nil {
		return nil, NewRPCError(err, "Failed to get paymaster deposit balance")
	}

	if maxCost := userOp.MaxGasCost(); maxCost.Cmp(balance) > 0 {
		return nil, ErrPaymasterInsufficientBalance.WithData(fmt.Sprintf("balance = %v, required = %v", balance, maxCost))
	}

	// update the custom data if provided
	if len(customData) > 0 {
		dataBytes := customData[0].Bytes()
		dataLen := len(dataBytes)
		size := minPaymasterAndDataLen + dataLen

		if len(userOp.PaymasterAndData) != size {
			return nil, api.ErrValidationStrf("Invalid paymasterAndData length, expected %d, got %d", size, len(userOp.PaymasterAndData))
		}

		copy(userOp.PaymasterAndData[52:52+dataLen], dataBytes)
	}

	// update signature time window before signing
	size := len(userOp.PaymasterAndData)
	validUntil := time.Now().Add(paymaster.config.SignatureTimeout).Unix()

	contract.SafeBigFillBytes(big.NewInt(0), userOp.PaymasterAndData[size-77:size-71])          // validAfter
	contract.SafeBigFillBytes(big.NewInt(validUntil), userOp.PaymasterAndData[size-71:size-65]) // validUntil
	copy(userOp.PaymasterAndData[size-65:], dummySignature)                                     // dummy signature

	// compute the paymaster signature
	hash, err := paymaster.caller.GetPaymasterHash(nil, userOp)
	if err != nil {
		return nil, NewRPCError(err, "Failed to retrieve paymaster hash from blockchain")
	}

	signature, err := paymaster.signer.SignHash(hash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to sign paymaster hash")
	}

	// update the paymaster signature in the user operation
	copy(userOp.PaymasterAndData[size-65:], signature)

	return userOp.PaymasterAndData, nil
}
