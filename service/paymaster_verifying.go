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

const (
	// VerifyingPaymaster Encoding: address(20) || validationGasLimit(16) || postOpGasLimit(16) || delegation(20) || validAfter(6) || validUntil(6) || signature(65)
	verifyingPaymasterDataLength = 149

	initCode7702Marker = "0x7702000000000000000000000000000000000000"
)

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
	}

	// check whitelist
	if !paymaster.config.smartAccountMap[delegation] {
		return nil, ErrVerifyingPaymasterInvalidSmartAccount.WithData(delegation)
	}

	// assemble the paymaster data
	var buf [verifyingPaymasterDataLength]byte

	validUntil := time.Now().Add(paymaster.config.SignatureTimeout).Unix()

	copy(buf[:20], paymaster.config.Address.Bytes()) // address
	copy(buf[52:72], delegation.Bytes())             // delegation
	big.NewInt(validUntil).FillBytes(buf[78:84])     // validUntil
	copy(buf[84:], dummySignature)                   // dummy signature

	return buf[:], nil
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

	// re-assemble paymasterData for signing, including validAfter and validUntil
	validUntil := time.Now().Add(paymaster.config.SignatureTimeout).Unix()
	big.NewInt(0).FillBytes(userOp.PaymasterAndData[72:78])          // validAfter
	big.NewInt(validUntil).FillBytes(userOp.PaymasterAndData[78:84]) // validUntil

	// compute the paymaster signature
	hash, err := paymaster.caller.GetPaymasterHash(nil, userOp)
	if err != nil {
		return nil, NewRPCError(err, "Failed to get paymaster hash")
	}

	signature, err := paymaster.signer.SignHash(hash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to sign paymaster hash")
	}

	// re-assemble signature into paymasterAndData
	copy(userOp.PaymasterAndData[84:], signature)

	return userOp.PaymasterAndData, nil
}

func (paymaster *VerifyingPaymaster) validate(userOp *contract.PackedUserOperation) error {
	// validate the sender address is not empty
	if userOp.Sender == (common.Address{}) {
		return api.ErrValidationStr("Invalid sender address")
	}

	// validate the paymasterAndData length at first, to avoid panic when accessing the slice
	if len(userOp.PaymasterAndData) != verifyingPaymasterDataLength {
		return api.ErrValidationStrf("Invalid paymaster data length: %d, expected %d", len(userOp.PaymasterAndData), verifyingPaymasterDataLength)
	}

	// check paymaster address
	if paymasterAddress := common.BytesToAddress(userOp.PaymasterAndData[:20]); paymasterAddress != paymaster.config.Address {
		return api.ErrValidationStrf("Invalid paymaster address: %s, expected %s", paymasterAddress, paymaster.config.Address)
	}

	// check delegation address
	delegation := common.BytesToAddress(userOp.PaymasterAndData[52:72])
	if !paymaster.config.smartAccountMap[delegation] {
		return api.ErrValidationStrf("Invalid delegation address: %s", delegation)
	}

	// check max cost
	maxCost := paymaster.maxCost(userOp)
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

// maxCost calculates the maximum cost of a user operation based on its gas limits and fees.
func (paymaster *VerifyingPaymaster) maxCost(userOp *contract.PackedUserOperation) *big.Int {
	maxCost := big.NewInt(0)

	maxCost.Add(maxCost, userOp.PreVerificationGas)                             // pre-verification gas
	maxCost.Add(maxCost, new(big.Int).SetBytes(userOp.AccountGasLimits[0:16]))  // account verification gas limit
	maxCost.Add(maxCost, new(big.Int).SetBytes(userOp.AccountGasLimits[16:32])) // account call gas limit
	maxCost.Add(maxCost, new(big.Int).SetBytes(userOp.PaymasterAndData[20:36])) // paymaster verification gas limit
	maxCost.Add(maxCost, new(big.Int).SetBytes(userOp.PaymasterAndData[36:52])) // paymaster postOp gas limit

	maxCost.Mul(maxCost, new(big.Int).SetBytes(userOp.GasFees[16:32])) // multiply by maxFeePerGas

	return maxCost
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

	// requires empty initCode if delegatoin unchanged
	if currentDelegation == delegation && len(initCode) > 0 {
		return api.ErrValidationStr("Invalid initCode, empty value required")
	}

	// requires 7702 marker if delegation changed (0 -> A or A -> B)
	if currentDelegation != delegation && hexutil.Encode(initCode) != initCode7702Marker {
		return api.ErrValidationStrf("Invalid initCode, expected %v", initCode7702Marker)
	}

	return nil
}
