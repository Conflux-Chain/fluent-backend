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
	Log            *types.Log
	SponsoredEvent *contract.VerifyingPaymasterSponsored
	UserOpEvent    *contract.EntryPointUserOperationEvent
	UserOp         *contract.PackedUserOperation
}

// UserOpEventParser is responsible for parsing Sponsored event and relevant data from the blockchain log.
type UserOpEventParser struct {
	client *web3go.Client

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
		paymasterFilterer:  &verifyingPaymaster.VerifyingPaymasterFilterer,
		entryPointAddr:     entryPointAddr,
		entryPointFilterer: entryPointFilterer,
		entryPointABI:      entryPointABI,
	}, nil
}

func (parser *UserOpEventParser) Parse(log *types.Log) (*Sponsorship, error) {
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
	userOp, err := parser.unpackUserOp(sponsoredEvent)
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
		if v.Topics[0] != eventHashUserOperation {
			continue
		}

		event, err := parser.entryPointFilterer.ParseUserOperationEvent(*v.ToEthLog())
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to parse UserOperation event log")
		}

		if event.UserOpHash == sponsoredEvent.UserOpHash {
			return event, nil
		}
	}

	return nil, fmt.Errorf("UserOperation event not found for userOpHash %v", sponsoredEvent.UserOpHash)
}

func (parser *UserOpEventParser) unpackUserOp(sponsoredEvent *contract.VerifyingPaymasterSponsored) (*contract.PackedUserOperation, error) {
	// get the bundle transaction to parse input data
	tx, err := parser.client.Eth.TransactionByHash(sponsoredEvent.Raw.TxHash)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get bundle transaction by hash")
	}

	if tx == nil {
		return nil, fmt.Errorf("Bundle transaction not found by hash %v", sponsoredEvent.Raw.TxHash)
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
		return nil, fmt.Errorf("Transaction input data too short for tx hash %v", sponsoredEvent.Raw.TxHash)
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
		return nil, fmt.Errorf("Unsupported method %v in transaction input for tx hash %v", method.Name, sponsoredEvent.Raw.TxHash)
	}

	// find the user operation by hash
	for _, v := range userOps {
		userOpHash, err := parser.calculateUserOpHash(v, delegates[v.Sender])
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to calculate user operation hash")
		}

		if userOpHash == sponsoredEvent.UserOpHash {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("UserOperation not found for userOpHash %v", sponsoredEvent.UserOpHash)
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
