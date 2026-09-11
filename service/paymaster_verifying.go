package service

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/interfaces"
	"github.com/pkg/errors"
)

const initCode7702Marker = "0x7702000000000000000000000000000000000000"

type VerifyingPaymasterConfig struct {
	Address               common.Address
	SmartAccountWhitelist []common.Address
	smartAccountMap       map[common.Address]bool
	ContractWhitelist     []common.Address
	contractMap           map[common.Address]bool
	MaxGasCost            int64 `default:"100000000000000000"` // 0.1 CFX by default, and could up to 1 CFX for int64 type
	maxGasCostBig         *big.Int
	SignatureTimeout      time.Duration `default:"5m"`

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

	config.maxGasCostBig = big.NewInt(config.MaxGasCost)

	return nil
}

type VerifyingPaymaster struct {
	config             VerifyingPaymasterConfig
	client             *web3go.Client
	caller             *contract.VerifyingPaymasterCaller
	executeMethod      *abi.Method
	executeBatchMethod *abi.Method
	signer             interfaces.Signer
	limiter            Limiter
}

func NewVerifyingPaymaster(config VerifyingPaymasterConfig, client *web3go.Client, store *store.Store) (*VerifyingPaymaster, error) {
	if err := config.validateAndNormalize(); err != nil {
		return nil, errors.WithMessage(err, "Invalid VerifyingPaymaster config")
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

	// contract callers
	caller, _ := client.ToClientForContract()

	verifyingPaymasterCaller, err := contract.NewVerifyingPaymasterCaller(config.Address, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create VerifyingPaymaster contract caller")
	}

	// check if the signer is whitelisted by the paymaster
	signerAddr := signers[0].Address()

	signerAllowed, err := verifyingPaymasterCaller.IsSignerAllowed(nil, signerAddr)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to check if signer is allowed by VerifyingPaymaster")
	}

	if !signerAllowed {
		return nil, fmt.Errorf("Signer is not allowed by VerifyingPaymaster: %v", signerAddr)
	}

	return &VerifyingPaymaster{
		config:             config,
		client:             client,
		caller:             verifyingPaymasterCaller,
		executeMethod:      &executeMethod,
		executeBatchMethod: &executeBatchMethod,
		signer:             signers[0],
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

	return contract.GeneratePaymasterAndDataStub(paymaster.config.Address, paymaster.config.SignatureTimeout, delegation), nil
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

	userOp.UpdatePaymasterPreSign(paymaster.config.SignatureTimeout)

	// compute the paymaster signature
	hash, err := paymaster.caller.GetPaymasterHash(nil, userOp)
	if err != nil {
		return nil, NewRPCError(err, "Failed to get paymaster hash")
	}

	signature, err := paymaster.signer.SignHash(hash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to sign paymaster hash")
	}

	userOp.UpdatePaymasterSignature(signature)

	return userOp.PaymasterAndData, nil
}

func (paymaster *VerifyingPaymaster) validate(userOp *contract.PackedUserOperation) error {
	// validate the sender address is not empty
	if userOp.Sender == (common.Address{}) {
		return api.ErrValidationStr("Invalid sender address")
	}

	// validate the paymasterAndData length at first, to avoid panic when accessing the slice
	if len(userOp.PaymasterAndData) != contract.MinSignablePaymasterAndDataLen+common.AddressLength {
		return api.ErrValidationStr("Invalid paymaster data length")
	}

	// check paymaster address
	if userOp.Paymaster() != paymaster.config.Address {
		return api.ErrValidationStrf("Invalid paymaster address: %s, expected %s", userOp.Paymaster(), paymaster.config.Address)
	}

	// check delegation address
	delegation := common.BytesToAddress(userOp.PaymasterCustomData())
	if !paymaster.config.smartAccountMap[delegation] {
		return api.ErrValidationStrf("Invalid delegation address: %s", delegation)
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

	// check if paymaster contract paused
	paused, err := paymaster.caller.Paused(nil)
	if err != nil {
		return NewRPCError(err, "Failed to check if paymaster contract is paused")
	}

	if paused {
		return ErrVerifyingPaymasterPaused
	}

	// check paymaster deposit balance
	balance, err := paymaster.caller.Balance(nil)
	if err != nil {
		return NewRPCError(err, "Failed to get paymaster deposit balance")
	}

	if balance.Cmp(maxCost) < 0 {
		return ErrVerifyingPaymasterInsufficientBalance.WithData(fmt.Sprintf("balance = %v, required = %v", balance, maxCost))
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
