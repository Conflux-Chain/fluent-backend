package worker

import (
	"fmt"
	"strings"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

const (
	entryPointMethodHandleOps           = "handleOps"
	entryPointMethodHandleAggregatedOps = "handleAggregatedOps"
)

// eventHashUserOperation is the hash of the UserOperation event signature, used to filter logs for this specific event.
//
// See EntryPoint contract for the UserOperation event definition.
var eventHashUserOperation = common.HexToHash("0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f")

type Sponsorship struct {
	Log            types.Log
	SponsoredEvent *contract.VerifyingPaymasterSponsored
	UserOpEvent    *contract.EntryPointUserOperationEvent
	UserOp         *contract.PackedUserOperation
}

// UserOpEventParser is responsible for parsing Sponsored event and relevant data from the blockchain log.
//
// If bundle transaction contains many user operations too frequently, it's better to add a LRU cache for
// the retrieved bundle transaction and receipt.
type UserOpEventParser struct {
	client *web3go.Client

	paymaster common.Address

	paymasterFilterer *contract.VerifyingPaymasterFilterer

	entryPointAddr     common.Address
	entryPointFilterer *contract.EntryPointFilterer
	entryPointABI      abi.ABI
}

func NewUserOpEventParser(client *web3go.Client, paymaster common.Address) (*UserOpEventParser, error) {
	if paymaster == (common.Address{}) {
		return nil, errors.New("Paymaster address is required")
	}

	filterer, _ := client.ToClientForContract()

	verifyingPaymaster, err := contract.NewVerifyingPaymaster(paymaster, filterer)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create VerifyingPaymaster")
	}

	entryPointAddr, err := verifyingPaymaster.EntryPoint(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get entry point address from VerifyingPaymaster")
	}

	entryPointFilterer, err := contract.NewEntryPointFilterer(entryPointAddr, filterer)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create EntryPoint filterer")
	}

	entryPointABI, err := abi.JSON(strings.NewReader(contract.EntryPointMetaData.ABI))
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse EntryPoint ABI")
	}

	return &UserOpEventParser{
		client:             client,
		paymaster:          paymaster,
		paymasterFilterer:  &verifyingPaymaster.VerifyingPaymasterFilterer,
		entryPointAddr:     entryPointAddr,
		entryPointFilterer: entryPointFilterer,
		entryPointABI:      entryPointABI,
	}, nil
}

// Parse parses a Sponsored event log and retrieves the associated user operation and its details.
func (parser *UserOpEventParser) Parse(log types.Log) (*Sponsorship, error) {
	// validate the paymaster address and Sponsored event signature
	if log.Address != parser.paymaster {
		return nil, fmt.Errorf("Log is not from the expected paymaster contract: %v", parser.paymaster)
	}

	if len(log.Topics) == 0 || log.Topics[0] != eventHashSponsored {
		return nil, fmt.Errorf("Log is not a Sponsored event, expected event hash: %v", eventHashSponsored)
	}

	// parse VerifyingPaymaster.Sponsored event
	sponsoredEvent, err := parser.paymasterFilterer.ParseSponsored(*log.ToEthLog())
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse Sponsored event log")
	}

	// find the corresponding EntryPoint.UserOperationEvent
	userOpEvent, err := parser.findUserOperationEvent(sponsoredEvent)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to find corresponding UserOperation event")
	}

	// unpack user op from tx input data
	userOp, err := parser.unpackUserOp(userOpEvent)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to unpack user operation from transaction input data")
	}

	return &Sponsorship{
		Log:            log,
		SponsoredEvent: sponsoredEvent,
		UserOpEvent:    userOpEvent,
		UserOp:         userOp,
	}, nil
}

func (parser *UserOpEventParser) findUserOperationEvent(sponsoredEvent *contract.VerifyingPaymasterSponsored) (*contract.EntryPointUserOperationEvent, error) {
	receipt, err := parser.client.Eth.TransactionReceipt(sponsoredEvent.Raw.TxHash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get transaction receipt")
	}

	if receipt == nil {
		return nil, fmt.Errorf("Transaction receipt not found %v", sponsoredEvent.Raw.TxHash)
	}

	for _, v := range receipt.Logs {
		// should be from the expected EntryPoint contract
		if v.Address != parser.entryPointAddr {
			continue
		}

		// should be an EntryPoint.UserOperationEvent
		if len(v.Topics) == 0 || v.Topics[0] != eventHashUserOperation {
			continue
		}

		// UserOperationEvent.topic[1] should match the UserOpHash of the Sponsored event
		if len(v.Topics) < 2 || v.Topics[1] != sponsoredEvent.UserOpHash {
			continue
		}

		event, err := parser.entryPointFilterer.ParseUserOperationEvent(*v.ToEthLog())
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to parse UserOperation event log")
		}

		return event, nil
	}

	return nil, fmt.Errorf("UserOperation event not found for userOpHash %v", sponsoredEvent.UserOpHash)
}

func (parser *UserOpEventParser) unpackUserOp(userOpEvent *contract.EntryPointUserOperationEvent) (*contract.PackedUserOperation, error) {
	// get the bundle transaction to parse input data
	tx, err := parser.client.Eth.TransactionByHash(userOpEvent.Raw.TxHash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get bundle transaction by hash")
	}

	if tx == nil {
		return nil, fmt.Errorf("Bundle transaction not found by hash %v", userOpEvent.Raw.TxHash)
	}

	// handle 7702 auth messages
	delegates := make(map[common.Address]common.Address)
	for _, auth := range tx.AuthorizationList {
		sender, err := auth.Authority()
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to get authority")
		}

		delegates[sender] = auth.Address
	}

	// unpack the user operation from the bundle transaction input data
	if len(tx.Input) < 4 {
		return nil, fmt.Errorf("Transaction input data too short for tx hash %v", userOpEvent.Raw.TxHash)
	}

	method, err := parser.entryPointABI.MethodById(tx.Input[:4])
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get method by ID from transaction input")
	}

	var userOps []contract.PackedUserOperation

	switch method.Name {
	case entryPointMethodHandleOps:
		var args struct {
			Ops         []contract.PackedUserOperation
			Beneficiary common.Address
		}

		if err = service.UnpackArguments(method.Inputs, tx.Input[4:], &args); err != nil {
			return nil, errors.WithMessage(err, "Failed to unpack handleOps arguments from transaction input")
		}

		userOps = args.Ops
	case entryPointMethodHandleAggregatedOps:
		var args struct {
			OpsPerAggregator []contract.IEntryPointUserOpsPerAggregator
			Beneficiary      common.Address
		}

		if err = service.UnpackArguments(method.Inputs, tx.Input[4:], &args); err != nil {
			return nil, errors.WithMessage(err, "Failed to unpack handleAggregatedOps arguments from transaction input")
		}

		for _, v := range args.OpsPerAggregator {
			userOps = append(userOps, v.UserOps...)
		}
	default:
		return nil, fmt.Errorf("Unsupported method %v in transaction input for tx hash %v", method.Name, userOpEvent.Raw.TxHash)
	}

	// find the user operation by hash
	for i := range userOps {
		userOp := userOps[i]

		// skip user operations that do not match the sender and nonce of the user operation event
		if userOp.Sender != userOpEvent.Sender || userOp.Nonce.Cmp(userOpEvent.Nonce) != 0 {
			continue
		}

		// skip user operations not sponsored by the expected paymaster
		if userOp.Paymaster() != userOpEvent.Paymaster {
			continue
		}

		userOpHash, err := parser.calculateUserOpHash(userOp, delegates[userOp.Sender])
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to calculate user operation hash")
		}

		if userOpHash == userOpEvent.UserOpHash {
			return &userOp, nil
		}
	}

	return nil, fmt.Errorf("UserOperation not found for userOpHash %v", userOpEvent.UserOpHash)
}

func (parser *UserOpEventParser) calculateUserOpHash(userOp contract.PackedUserOperation, delegation common.Address) (common.Hash, error) {
	calldata, err := parser.entryPointABI.Pack("getUserOpHash", userOp)
	if err != nil {
		return common.Hash{}, errors.WithMessage(err, "Failed to pack calldata of getUserOpHash")
	}

	request := types.CallRequest{
		To:   &parser.entryPointAddr,
		Data: calldata,
	}

	latestBlockNumber := types.BlockNumberOrHashWithNumber(types.LatestBlockNumber)

	var overrides *types.StateOverride
	if delegation != (common.Address{}) {
		code := hexutil.Bytes(append(service.DelegatedCodePrefix, delegation.Bytes()...))

		overrides = &types.StateOverride{
			userOp.Sender: types.OverrideAccount{
				Code: &code,
			},
		}
	}

	result, err := parser.client.Eth.Call(request, &latestBlockNumber, overrides, nil)
	if err != nil {
		return common.Hash{}, errors.WithMessage(err, "Failed to call getUserOpHash")
	}

	if len(result) != 32 {
		return common.Hash{}, errors.New("Invalid result length from getUserOpHash")
	}

	return common.BytesToHash(result), nil
}
