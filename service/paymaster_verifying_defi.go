package service

import (
	"fmt"

	"github.com/Conflux-Chain/fluent-backend/contract"
	uniswapv2 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v2"
	uniswapv3 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
)

////////////////////////////////////////////////////////////////////////////////
//
// Config struct, interface and factory function
//
////////////////////////////////////////////////////////////////////////////////

type DeFiConfig struct {
	Uniswap struct {
		V2 struct {
			Router common.Address
		}
		V3 struct {
			Router common.Address
		}
	}
}

// ExecutionPolicy defines the interface for execution policies that validate contract executions.
type ExecutionPolicy interface {
	// Validate checks whether the given contract execution is valid according to the execution policy and the contract whitelist.
	// It returns an error if the execution is not valid or if any error occurs during the validation. Otherwise, it returns nil.
	Validate(execution contract.Execution, whitelist map[common.Address]bool) error
}

func NewExecutionPolicy(config DeFiConfig, client *web3go.Client, whitelist map[common.Address]bool) (ExecutionPolicy, error) {
	composite := make(map[common.Address]ExecutionPolicy)

	// uniswap v2
	if config.Uniswap.V2.Router != (common.Address{}) {
		if whitelist[config.Uniswap.V2.Router] {
			return nil, fmt.Errorf("Contract whitelist already contains the Uniswap V2 router %v", config.Uniswap.V2.Router)
		}

		policy, err := NewUniswapV2ExecutionPolicy(config.Uniswap.V2.Router, client)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to create Uniswap V2 execution policy")
		}

		composite[config.Uniswap.V2.Router] = policy
	}

	// uniswap v3
	if config.Uniswap.V3.Router != (common.Address{}) {
		if whitelist[config.Uniswap.V3.Router] {
			return nil, fmt.Errorf("Contract whitelist already contains the Uniswap V3 router %v", config.Uniswap.V3.Router)
		}

		if config.Uniswap.V3.Router == config.Uniswap.V2.Router {
			return nil, fmt.Errorf("Composite execution policy already contains the Uniswap V3 router %v", config.Uniswap.V3.Router)
		}

		policy, err := NewUniswapV3ExecutionPolicy(config.Uniswap.V3.Router, client)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to create Uniswap V3 execution policy")
		}

		composite[config.Uniswap.V3.Router] = policy
	}

	return CompositeExecutionPolicy(composite), nil
}

////////////////////////////////////////////////////////////////////////////////
//
// Composite execution policy
//
////////////////////////////////////////////////////////////////////////////////

type CompositeExecutionPolicy map[common.Address]ExecutionPolicy

func (policy CompositeExecutionPolicy) Validate(execution contract.Execution, whitelist map[common.Address]bool) error {
	if whitelist[execution.Target] {
		return nil
	}

	dappPolicy, ok := policy[execution.Target]
	if !ok {
		return ErrVerifyingPaymasterContractNotWhitelisted.WithData(execution.Target)
	}

	if err := dappPolicy.Validate(execution, whitelist); err != nil {
		return ErrVerifyingPaymasterContractNotWhitelisted.WithData(err.Error())
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
//
// Uniswap v2 execution policy
//
////////////////////////////////////////////////////////////////////////////////

type UniswapV2ExecutionPolicy struct {
	router common.Address
	weth   common.Address
}

func NewUniswapV2ExecutionPolicy(router common.Address, client *web3go.Client) (*UniswapV2ExecutionPolicy, error) {
	caller, _ := client.ToClientForContract()

	routerCaller, err := uniswapv2.NewRouterCaller(router, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create Uniswap V2 router caller")
	}

	weth, err := routerCaller.WETH(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get Uniswap V2 WETH")
	}

	return &UniswapV2ExecutionPolicy{
		router: router,
		weth:   weth,
	}, nil
}

func (policy *UniswapV2ExecutionPolicy) Validate(execution contract.Execution, whitelist map[common.Address]bool) error {
	if execution.Target != policy.router {
		return fmt.Errorf("Target contract not in whitelist: %v", execution.Target)
	}

	if len(execution.CallData) < 4 {
		return fmt.Errorf("Call data too short %v", hexutil.Encode(execution.CallData))
	}

	// check method whitelist and msg.value
	method, err := contract.UniswapV2RouterABI.MethodById(execution.CallData[:4])
	if err != nil {
		return errors.WithMessagef(err, "Failed to find method by ID %v", hexutil.Encode(execution.CallData[:4]))
	}

	switch method.RawName {
	case "swapExactTokensForTokens", "swapTokensForExactTokens", "swapExactTokensForETH", "swapTokensForExactETH",
		"swapExactTokensForTokensSupportingFeeOnTransferTokens", "swapExactTokensForETHSupportingFeeOnTransferTokens":
		if execution.Value != nil && execution.Value.Sign() != 0 {
			return fmt.Errorf("Zero msg.value required for method %v", method.RawName)
		}
	case "swapExactETHForTokens", "swapETHForExactTokens", "swapExactETHForTokensSupportingFeeOnTransferTokens":
		if execution.Value == nil || execution.Value.Sign() <= 0 {
			return fmt.Errorf("Positive msg.value required for method %v", method.RawName)
		}
	default:
		return fmt.Errorf("Method not in whitelist: %v", method.RawName)
	}

	// unpack calldata to parse input & output tokens
	values, err := method.Inputs.Unpack(execution.CallData[4:])
	if err != nil {
		return errors.WithMessage(err, "Failed to unpack calldata")
	}

	if len(values) < 3 {
		return fmt.Errorf("Unexpected number of input values: %v", len(values))
	}

	path, ok := values[len(values)-3].([]common.Address)
	if !ok || len(path) < 2 {
		return fmt.Errorf("Invalid path in calldata: %v", values[len(values)-3])
	}

	input, output := path[0], path[len(path)-1]

	// check against the whitelist & WETH
	switch method.RawName {
	case "swapExactTokensForTokens", "swapTokensForExactTokens", "swapExactTokensForTokensSupportingFeeOnTransferTokens":
		// Note, WETH is not allowed if not configured in whitelist
		if !whitelist[input] {
			return fmt.Errorf("Input token not in whitelist: %v", input)
		}

		if !whitelist[output] {
			return fmt.Errorf("Output token not in whitelist: %v", output)
		}
	case "swapExactTokensForETH", "swapTokensForExactETH", "swapExactTokensForETHSupportingFeeOnTransferTokens":
		if !whitelist[input] {
			return fmt.Errorf("Input token not in whitelist: %v", input)
		}

		if output != policy.weth {
			return fmt.Errorf("Output token of method %v must be WETH: %v", method.RawName, policy.weth)
		}
	case "swapExactETHForTokens", "swapETHForExactTokens", "swapExactETHForTokensSupportingFeeOnTransferTokens":
		if input != policy.weth {
			return fmt.Errorf("Input token of method %v must be WETH: %v", method.RawName, policy.weth)
		}

		if !whitelist[output] {
			return fmt.Errorf("Output token not in whitelist: %v", output)
		}
	default:
		return fmt.Errorf("Method not in whitelist: %v", method.RawName)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
//
// Uniswap v3 execution policy
//
////////////////////////////////////////////////////////////////////////////////

type UniswapV3ExecutionPolicy struct {
	router common.Address
	weth   common.Address
}

func NewUniswapV3ExecutionPolicy(router common.Address, client *web3go.Client) (*UniswapV3ExecutionPolicy, error) {
	caller, _ := client.ToClientForContract()

	routerCaller, err := uniswapv3.NewSwapRouterCaller(router, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create Uniswap V3 router caller")
	}

	weth, err := routerCaller.WETH9(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get Uniswap V3 WETH")
	}

	return &UniswapV3ExecutionPolicy{
		router: router,
		weth:   weth,
	}, nil
}

func (policy *UniswapV3ExecutionPolicy) Validate(execution contract.Execution, whitelist map[common.Address]bool) error {
	if execution.Target != policy.router {
		return fmt.Errorf("Target contract not in whitelist: %v", execution.Target)
	}

	if len(execution.CallData) < 4 {
		return fmt.Errorf("Call data too short %v", hexutil.Encode(execution.CallData))
	}

	// check method whitelist and unpack input/output tokens
	method, err := contract.UniswapV3RouterABI.MethodById(execution.CallData[:4])
	if err != nil {
		return errors.WithMessagef(err, "Failed to find method by ID %v", hexutil.Encode(execution.CallData[:4]))
	}

	decodePath := func(path []byte) (first common.Address, last common.Address, err error) {
		pathLen := len(path)
		if pathLen < 43 || (pathLen-20)%23 != 0 {
			return common.Address{}, common.Address{}, fmt.Errorf("Invalid path length: %v", pathLen)
		}

		return common.BytesToAddress(path[:20]), common.BytesToAddress(path[pathLen-20:]), nil
	}

	var input, output common.Address

	switch method.RawName {
	case "exactInputSingle":
		var args struct {
			Params uniswapv3.ISwapRouterExactInputSingleParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return errors.WithMessagef(err, "Failed to unpack calldata for method %v", method.RawName)
		}

		input, output = args.Params.TokenIn, args.Params.TokenOut
	case "exactInput":
		var args struct {
			Params uniswapv3.ISwapRouterExactInputParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return errors.WithMessagef(err, "Failed to unpack calldata for method %v", method.RawName)
		}

		if input, output, err = decodePath(args.Params.Path); err != nil {
			return errors.WithMessagef(err, "Failed to decode path for method %v", method.RawName)
		}
	case "exactOutputSingle":
		var args struct {
			Params uniswapv3.ISwapRouterExactOutputSingleParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return errors.WithMessagef(err, "Failed to unpack calldata for method %v", method.RawName)
		}

		input, output = args.Params.TokenIn, args.Params.TokenOut
	case "exactOutput":
		var args struct {
			Params uniswapv3.ISwapRouterExactOutputParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return errors.WithMessagef(err, "Failed to unpack calldata for method %v", method.RawName)
		}

		if output, input, err = decodePath(args.Params.Path); err != nil {
			return errors.WithMessagef(err, "Failed to decode path for method %v", method.RawName)
		}
	default:
		return fmt.Errorf("Method not in whitelist: %v", method.RawName)
	}

	if !whitelist[input] && input != policy.weth {
		return fmt.Errorf("Input token of method %v not in whitelist or not WETH: %v", method.RawName, input)
	}

	if !whitelist[output] && output != policy.weth {
		return fmt.Errorf("Output token of method %v not in whitelist or not WETH: %v", method.RawName, output)
	}

	return nil
}
