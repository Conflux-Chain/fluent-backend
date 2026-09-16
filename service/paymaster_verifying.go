package service

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
)

const initCode7702Marker = "0x7702000000000000000000000000000000000000"

type VerifyingPaymasterConfig struct {
	PaymasterConfig `mapstructure:",squash"`

	SmartAccountWhitelist []common.Address
	smartAccountMap       map[common.Address]bool
	ContractWhitelist     []common.Address
	contractMap           map[common.Address]bool
	MaxGasCost            uint64 `default:"100000000000000000"` // 0.1 CFX by default, and could up to 1 CFX
	maxGasCostBig         *big.Int

	Limiter LimitConfig
}

func (config *VerifyingPaymasterConfig) validateAndNormalize() error {
	// validate
	if config.Address == (common.Address{}) {
		return errors.New("Address is required")
	}

	if len(config.SmartAccountWhitelist) == 0 {
		return errors.New("Smart account whitelist is required")
	}

	if len(config.ContractWhitelist) == 0 {
		return errors.New("Contract whitelist is required")
	}

	// normalize
	config.smartAccountMap = make(map[common.Address]bool)
	for _, addr := range config.SmartAccountWhitelist {
		config.smartAccountMap[addr] = true
	}

	config.contractMap = make(map[common.Address]bool)
	for _, addr := range config.ContractWhitelist {
		config.contractMap[addr] = true
	}

	config.maxGasCostBig = new(big.Int).SetUint64(config.MaxGasCost)

	return nil
}

type VerifyingPaymaster struct {
	inner              *Paymaster[*contract.VerifyingPaymasterCaller]
	config             VerifyingPaymasterConfig
	client             *web3go.Client
	executeMethod      abi.Method
	executeBatchMethod abi.Method
	limiter            Limiter
}

func NewVerifyingPaymaster(config VerifyingPaymasterConfig, client *web3go.Client, store *store.Store) (*VerifyingPaymaster, error) {
	if err := config.validateAndNormalize(); err != nil {
		return nil, errors.WithMessage(err, "Invalid VerifyingPaymaster config")
	}

	paymaster, err := NewPaymaster(config.PaymasterConfig, client, contract.NewVerifyingPaymasterCaller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create VerifyingPaymaster")
	}

	// smart account execute ABI
	smartAccountABI, err := abi.JSON(strings.NewReader(contract.SimpleSmartAccount7702MetaData.ABI))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse SimpleSmartAccount7702 ABI")
	}

	executeMethod, ok := smartAccountABI.Methods["execute"]
	if !ok {
		return nil, errors.New("Failed to get execute method from SimpleSmartAccount7702 ABI")
	}

	executeBatchMethod, ok := smartAccountABI.Methods["executeBatch"]
	if !ok {
		return nil, errors.New("Failed to get executeBatch method from SimpleSmartAccount7702 ABI")
	}

	// check if the signer is whitelisted by the paymaster
	signerAllowed, err := paymaster.caller.IsSignerAllowed(nil, paymaster.signer.Address())
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to check if signer is allowed by VerifyingPaymaster")
	}

	if !signerAllowed {
		return nil, fmt.Errorf("Signer is not allowed by VerifyingPaymaster: %v", paymaster.signer.Address())
	}

	return &VerifyingPaymaster{
		inner:              paymaster,
		config:             config,
		client:             client,
		executeMethod:      executeMethod,
		executeBatchMethod: executeBatchMethod,
		limiter:            NewLimiter(config.Limiter, store),
	}, nil
}

// Stub returns a stub paymasterAndData for gas estimation.
func (paymaster *VerifyingPaymaster) Stub(sender, delegation common.Address) ([]byte, error) {
	// retrieve the delegated contract if not provided
	if delegation == (common.Address{}) {
		var err error

		if delegation, err = GetDelegatedContract(paymaster.client, sender); err != nil {
			return nil, err
		}

		if delegation == (common.Address{}) {
			return nil, api.ErrValidationStr("Delegation not provided for the EOA sender")
		}
	}

	// check whitelist
	if !paymaster.config.smartAccountMap[delegation] {
		return nil, ErrVerifyingPaymasterInvalidSmartAccount.WithData(delegation)
	}

	return paymaster.inner.generateStub(delegation), nil
}

// Sign validates the user operation and signs the user operation with the paymaster's private key.
// It returns the signed paymasterAndData, which includes the paymaster address, gas limits, delegation, validAfter, validUntil and signature.
func (paymaster *VerifyingPaymaster) Sign(userOp contract.PackedUserOperation) ([]byte, error) {
	// rate limit
	if err := paymaster.limiter.Limit(&userOp); err != nil {
		return nil, err
	}

	// validate the user operation
	//
	// NOTE To prevent abuse in future, we could blacklist the sender address and
	// its IP address for a while, if failed to validate too many times.
	if err := paymaster.validate(&userOp); err != nil {
		return nil, err
	}

	return paymaster.inner.sign(userOp)
}

func (paymaster *VerifyingPaymaster) validate(userOp *contract.PackedUserOperation) error {
	// validate the sender address is not empty
	if userOp.Sender == (common.Address{}) {
		return api.ErrValidationStr("Invalid sender address")
	}

	// validate the paymasterAndData field and extract the custom data.
	customData, err := paymaster.inner.validatePaymasterAndData(userOp)
	if err != nil {
		return err
	}

	// check delegation address
	if len(customData) != common.AddressLength {
		return api.ErrValidationStr("Invalid paymasterAndData length")
	}

	delegation := common.BytesToAddress(customData)

	if !paymaster.config.smartAccountMap[delegation] {
		return ErrVerifyingPaymasterInvalidSmartAccount.WithData(delegation)
	}

	// check max cost
	maxCost := userOp.MaxGasCost()
	if paymaster.config.maxGasCostBig.Cmp(maxCost) < 0 {
		return ErrVerifyingPaymasterMaxGasCostExceeded.WithData(fmt.Sprintf("max = %v, actual = %v", paymaster.config.maxGasCostBig, maxCost))
	}

	// check calldata
	if err := paymaster.validateCallData(userOp.CallData); err != nil {
		return err
	}

	// check init code based on the delegated contract
	if err := paymaster.validateInitCode(userOp.Sender, delegation, userOp.InitCode); err != nil {
		return err
	}

	return nil
}

// validateCallData checks if the call data is valid for the user operation.
func (paymaster *VerifyingPaymaster) validateCallData(callData []byte) error {
	// must call execute or executeBatch functions of the smart account
	if len(callData) < 4 {
		return api.ErrValidationStr("Invalid callData, too short to parse function selector")
	}

	selector, args := callData[:4], callData[4:]

	if bytes.Equal(paymaster.executeMethod.ID, selector) {
		// single execute
		var execution contract.Execution
		if err := UnpackArguments(paymaster.executeMethod.Inputs, args, &execution); err != nil {
			return api.ErrValidation(errors.WithMessage(err, "Failed to unpack callData for execute method"))
		}

		if !paymaster.config.contractMap[execution.Target] {
			return ErrVerifyingPaymasterContractNotWhitelisted.WithData(execution.Target)
		}
	} else if bytes.Equal(paymaster.executeBatchMethod.ID, selector) {
		// batch execute
		var executions []contract.Execution
		if err := UnpackArguments(paymaster.executeBatchMethod.Inputs, args, &executions); err != nil {
			return api.ErrValidation(errors.WithMessage(err, "Failed to unpack callData for executeBatch method"))
		}

		if len(executions) == 0 {
			return api.ErrValidationStr("Invalid callData, batch is empty")
		}

		for _, execution := range executions {
			if !paymaster.config.contractMap[execution.Target] {
				return ErrVerifyingPaymasterContractNotWhitelisted.WithData(execution.Target)
			}
		}
	} else {
		return api.ErrValidationStr("Invalid callData, unsupported function selector")
	}

	return nil
}

// validateInitCode checks if the init code is valid for the smart account delegation.
func (paymaster *VerifyingPaymaster) validateInitCode(sender, delegation common.Address, initCode []byte) error {
	// retrieve the current delegation for the sender
	currentDelegation, err := GetDelegatedContract(paymaster.client, sender)
	if err != nil {
		return err
	}

	// requires empty initCode if delegation unchanged
	if currentDelegation == delegation && len(initCode) > 0 {
		return api.ErrValidationStr("Invalid initCode, empty value required")
	}

	// requires 7702 marker if delegation changed (0 -> A or A -> B)
	if currentDelegation != delegation && hexutil.Encode(initCode) != initCode7702Marker {
		return api.ErrValidationStrf("Invalid initCode, expected %v", initCode7702Marker)
	}

	return nil
}
