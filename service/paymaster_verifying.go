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
	// VerifyingPaymaster Encoding: address(20) || validationGasLimit(16) || postOpGasLimit(16) || validAfter(6) || validUntil(6) || signature(65)
	verifyingPaymasterDataLength = 129

	initCode7702Marker = "0x7702000000000000000000000000000000000000"
)

type VerifyingPaymasterConfig struct {
	Address           common.Address
	ContractWhitelist []common.Address
	contractWhitelist map[common.Address]bool
	MaxGasCost        int64 `default:"100000000000000000"` // 0.1 CFX by default, and could up to 1 CFX for int64 type
	maxGasCost        *big.Int
	SignatureTimeout  time.Duration `default:"5m"`

	MaxPendingOps int64 `default:"10"` // max pending user ops per sender
}

type VerifyingPaymaster struct {
	config             VerifyingPaymasterConfig
	client             *web3go.Client
	caller             *contract.VerifyingPaymasterCaller
	entryPointCaller   *contract.EntryPointCaller
	executeMethod      *abi.Method
	executeBatchMethod *abi.Method
	signer             interfaces.Signer
	userOpStore        *store.UserOpStore
}

func NewVerifyingPaymaster(config VerifyingPaymasterConfig, client *web3go.Client, userOpStore *store.UserOpStore) (*VerifyingPaymaster, error) {
	// check config and normalize it
	if config.Address == (common.Address{}) {
		return nil, errors.New("VerifyingPaymaster address is required")
	}

	if len(config.ContractWhitelist) == 0 {
		return nil, errors.New("Contract whitelist is required")
	}

	config.maxGasCost = big.NewInt(config.MaxGasCost)
	config.contractWhitelist = make(map[common.Address]bool)
	for _, addr := range config.ContractWhitelist {
		config.contractWhitelist[addr] = true
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

	entryPointAddress, err := verifyingPaymasterCaller.EntryPoint(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get EntryPoint address from VerifyingPaymaster")
	}

	entryPointCaller, err := contract.NewEntryPointCaller(entryPointAddress, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create EntryPoint contract caller")
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
		entryPointCaller:   entryPointCaller,
		executeMethod:      &executeMethod,
		executeBatchMethod: &executeBatchMethod,
		signer:             signers[0],
		userOpStore:        userOpStore,
	}, nil
}

// Stub returns a stub paymasterAndData for gas estimation.
func (paymaster *VerifyingPaymaster) Stub() []byte {
	var buf [verifyingPaymasterDataLength]byte

	validUntil := time.Now().Add(paymaster.config.SignatureTimeout).Unix()

	copy(buf[:20], paymaster.config.Address.Bytes()) // address
	big.NewInt(validUntil).FillBytes(buf[58:64])     // validUntil
	copy(buf[64:], dummySignature)                   // dummy signature

	return buf[:]
}

// Sign validates the user operation and signs the user operation with the paymaster's private key.
// It returns the signed paymasterAndData, which includes the paymaster address, gas limits, validAfter, validUntil, and signature.
func (paymaster *VerifyingPaymaster) Sign(userOp contract.PackedUserOperation, delegatedContract common.Address) ([]byte, error) {
	// limit the number of pending user ops
	if userOp.Sender == (common.Address{}) {
		return nil, api.ErrValidationStr("Invalid userOp, sender is empty")
	}

	pendings, err := paymaster.userOpStore.GetPendingCount(userOp.Sender)
	if err != nil {
		return nil, err
	}

	if pendings >= paymaster.config.MaxPendingOps {
		return nil, ErrVerifyingPaymasterTooManyPendingOps.WithData(fmt.Sprintf("max = %v", paymaster.config.MaxPendingOps))
	}

	// validate the user operation
	if err := paymaster.validate(&userOp, delegatedContract); err != nil {
		return nil, err
	}

	// re-assemble paymasterData for signing, including validAfter and validUntil
	validUntil := time.Now().Add(paymaster.config.SignatureTimeout)
	big.NewInt(0).FillBytes(userOp.PaymasterAndData[52:58])                 // validAfter
	big.NewInt(validUntil.Unix()).FillBytes(userOp.PaymasterAndData[58:64]) // validUntil

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
	copy(userOp.PaymasterAndData[64:], signature)

	// persistent the user op to database
	userOpHash, err := paymaster.entryPointCaller.GetUserOpHash(nil, userOp)
	if err != nil {
		return nil, NewRPCError(err, "Failed to get user operation hash")
	}

	entity := store.UserOp{
		UserOpHash: hexutil.Encode(userOpHash[:]),
		Sender:     userOp.Sender.Hex(),
		Nonce:      hexutil.Encode(userOp.Nonce.Bytes()),
		ValidUntil: validUntil,
		Status:     store.UserOpStatusSigned,
	}

	if err = paymaster.userOpStore.Create(&entity); err != nil {
		return nil, err
	}

	return userOp.PaymasterAndData, nil
}

func (paymaster *VerifyingPaymaster) validate(userOp *contract.PackedUserOperation, delegatedContract common.Address) error {
	// validate the paymasterAndData length at first, to avoid panic when accessing the slice
	if len(userOp.PaymasterAndData) != verifyingPaymasterDataLength {
		return api.ErrValidationStrf("Invalid paymaster data length: %d, expected %d", len(userOp.PaymasterAndData), verifyingPaymasterDataLength)
	}

	// check paymaster address
	if paymasterAddress := common.BytesToAddress(userOp.PaymasterAndData[:20]); paymasterAddress != paymaster.config.Address {
		return api.ErrValidationStrf("Invalid paymaster address: %s, expected %s", paymasterAddress, paymaster.config.Address)
	}

	// check init code
	if delegatedContract == (common.Address{}) && len(userOp.InitCode) > 0 {
		return api.ErrValidationStr("Invalid initCode, empty value required")
	}

	if delegatedContract != (common.Address{}) && hexutil.Encode(userOp.InitCode) != initCode7702Marker {
		return api.ErrValidationStrf("Invalid initCode, expected %s", initCode7702Marker)
	}

	// check max cost
	if maxCost := paymaster.maxCost(userOp); paymaster.config.maxGasCost.Cmp(maxCost) < 0 {
		return ErrVerifyingPaymasterMaxGasCostExceeded.WithData(fmt.Sprintf("max = %v, actual = %v", paymaster.config.maxGasCost, maxCost))
	}

	// check calldata
	if err := paymaster.validateCallData(userOp.CallData); err != nil {
		return err
	}

	// check if delegation is whitelisted
	if err := paymaster.validateSmartAccount(userOp.Sender, delegatedContract); err != nil {
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
		unpacked, err := paymaster.executeMethod.Inputs.Unpack(args)
		if err != nil {
			return api.ErrValidation(errors.WithMessage(err, "Failed to unpack callData for execute method"))
		}

		var execution contract.Execution
		if err = paymaster.executeMethod.Inputs.Copy(&execution, unpacked); err != nil {
			return api.ErrValidation(errors.WithMessage(err, "Failed to copy unpacked data to Execution struct"))
		}

		if !paymaster.config.contractWhitelist[execution.Target] {
			return ErrVerifyingPaymasterContractNotWhitelisted.WithData(execution.Target)
		}
	} else if bytes.Equal(paymaster.executeBatchMethod.ID, selector) {
		// batch execute
		unpacked, err := paymaster.executeBatchMethod.Inputs.Unpack(args)
		if err != nil {
			return api.ErrValidation(errors.WithMessage(err, "Failed to unpack callData for executeBatch method"))
		}

		var executions []contract.Execution
		if err = paymaster.executeBatchMethod.Inputs.Copy(&executions, unpacked); err != nil {
			return api.ErrValidation(errors.WithMessage(err, "Failed to copy unpacked data to []Execution struct"))
		}

		if len(executions) == 0 {
			return api.ErrValidationStr("Invalid callData, batch is empty")
		}

		for _, execution := range executions {
			if !paymaster.config.contractWhitelist[execution.Target] {
				return ErrVerifyingPaymasterContractNotWhitelisted.WithData(execution.Target)
			}
		}
	} else {
		return api.ErrValidationStr("Invalid callData, unsupported function selector")
	}

	return nil
}

// validateSmartAccount checks if the smart account is whitelisted by the paymaster.
func (paymaster *VerifyingPaymaster) validateSmartAccount(sender common.Address, delegatedContract common.Address) error {
	if delegatedContract == (common.Address{}) {
		delegation, err := GetDelegatedContract(paymaster.client, sender)
		if err != nil {
			return err
		}

		if delegation == (common.Address{}) {
			return api.ErrValidationStr("Delegated contract not found")
		}

		delegatedContract = delegation
	}

	whitelisted, err := paymaster.caller.SmartAccountWhitelist(nil, delegatedContract)
	if err != nil {
		return NewRPCError(err, "Failed to check if smart account is whitelisted")
	}

	if !whitelisted {
		return ErrVerifyingPaymasterInvalidSmartAccount.WithData(delegatedContract)
	}

	return nil
}
