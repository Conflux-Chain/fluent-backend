package service

import (
	"fmt"

	"github.com/Conflux-Chain/fluent-backend/contract"
	uniswapv2 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v2"
	uniswapv3 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v3"
	"github.com/ethereum/go-ethereum/common"
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

// ExecutionPolicy defines the interface for execution policies that determine whether a given contract execution is allowed.
type ExecutionPolicy interface {
	// IsAllowed determines whether the given contract execution is allowed based on the execution policy and the contract whitelist.
	// Note, any execution policy should return false if any error occurs during the evaluation, in which case the execution is
	// considered to be not allowed for sponsorship.
	IsAllowed(execution contract.Execution, whitelist map[common.Address]bool) bool
}

func NewExecutionPolicy(config DeFiConfig, client *web3go.Client, whitelist map[common.Address]bool) (ExecutionPolicy, error) {
	composite := []ExecutionPolicy{
		// check target contract whitelist at first
		TargetContractExecutionPolicy{},
	}

	// uniswap v2
	if config.Uniswap.V2.Router != (common.Address{}) {
		if whitelist[config.Uniswap.V2.Router] {
			return nil, fmt.Errorf("Contract whitelist already contains the Uniswap V2 router %v", config.Uniswap.V2.Router)
		}

		policy, err := NewUniswapV2ExecutionPolicy(config.Uniswap.V2.Router, client)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to create Uniswap V2 execution policy")
		}

		composite = append(composite, policy)
	}

	// uniswap v3
	if config.Uniswap.V3.Router != (common.Address{}) {
		if whitelist[config.Uniswap.V3.Router] {
			return nil, fmt.Errorf("Contract whitelist already contains the Uniswap V3 router %v", config.Uniswap.V3.Router)
		}

		policy, err := NewUniswapV3ExecutionPolicy(config.Uniswap.V3.Router, client)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to create Uniswap V3 execution policy")
		}

		composite = append(composite, policy)
	}

	return CompositeExecutionPolicy(composite), nil
}

////////////////////////////////////////////////////////////////////////////////
//
// Composite execution policy
//
////////////////////////////////////////////////////////////////////////////////

type CompositeExecutionPolicy []ExecutionPolicy

func (policy CompositeExecutionPolicy) IsAllowed(execution contract.Execution, whitelist map[common.Address]bool) bool {
	for _, v := range policy {
		if v.IsAllowed(execution, whitelist) {
			return true
		}
	}

	return false
}

////////////////////////////////////////////////////////////////////////////////
//
// Target contract whitelist execution policy
//
////////////////////////////////////////////////////////////////////////////////

type TargetContractExecutionPolicy struct{}

func (policy TargetContractExecutionPolicy) IsAllowed(execution contract.Execution, whitelist map[common.Address]bool) bool {
	return whitelist[execution.Target]
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

func (policy *UniswapV2ExecutionPolicy) IsAllowed(execution contract.Execution, whitelist map[common.Address]bool) bool {
	if execution.Target != policy.router {
		return false
	}

	if len(execution.CallData) < 4 {
		return false
	}

	// check method whitelist and msg.value
	method, err := contract.UniswapV2RouterABI.MethodById(execution.CallData[:4])
	if err != nil {
		return false
	}

	switch method.RawName {
	case "swapExactTokensForTokens", "swapTokensForExactTokens", "swapExactTokensForETH", "swapTokensForExactETH":
		if execution.Value != nil && execution.Value.Sign() != 0 {
			return false
		}
	case "swapExactETHForTokens", "swapETHForExactTokens":
		if execution.Value == nil || execution.Value.Sign() <= 0 {
			return false
		}
	default:
		return false
	}

	// unpack calldata to parse input & output tokens
	values, err := method.Inputs.Unpack(execution.CallData[4:])
	if err != nil {
		return false
	}

	if len(values) < 3 {
		return false
	}

	path, ok := values[len(values)-3].([]common.Address)
	if !ok || len(path) < 2 {
		return false
	}

	input, output := path[0], path[len(path)-1]

	// check against the whitelist & WETH
	switch method.RawName {
	case "swapExactTokensForTokens", "swapTokensForExactTokens":
		// Note, WETH is not allowed if not configured in whitelist
		return whitelist[input] && whitelist[output]
	case "swapExactTokensForETH", "swapTokensForExactETH":
		return whitelist[input] && output == policy.weth
	case "swapExactETHForTokens", "swapETHForExactTokens":
		return input == policy.weth && whitelist[output]
	default:
		return false
	}
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

func (policy *UniswapV3ExecutionPolicy) IsAllowed(execution contract.Execution, whitelist map[common.Address]bool) bool {
	if execution.Target != policy.router {
		return false
	}

	if len(execution.CallData) < 4 {
		return false
	}

	// check method whitelist and unpack input/output tokens
	method, err := contract.UniswapV3RouterABI.MethodById(execution.CallData[:4])
	if err != nil {
		return false
	}

	decodePath := func(path []byte) (first common.Address, last common.Address) {
		pathLen := len(path)
		if pathLen < 43 || (pathLen-20)%23 != 0 {
			return common.Address{}, common.Address{}
		}

		return common.BytesToAddress(path[:20]), common.BytesToAddress(path[pathLen-20:])
	}

	var input, output common.Address

	switch method.RawName {
	case "exactInputSingle":
		var args struct {
			Params uniswapv3.ISwapRouterExactInputSingleParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return false
		}

		input, output = args.Params.TokenIn, args.Params.TokenOut
	case "exactInput":
		var args struct {
			Params uniswapv3.ISwapRouterExactInputParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return false
		}

		input, output = decodePath(args.Params.Path)
	case "exactOutputSingle":
		var args struct {
			Params uniswapv3.ISwapRouterExactOutputSingleParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return false
		}

		input, output = args.Params.TokenIn, args.Params.TokenOut
	case "exactOutput":
		var args struct {
			Params uniswapv3.ISwapRouterExactOutputParams
		}

		if err = UnpackArguments(method.Inputs, execution.CallData[4:], &args); err != nil {
			return false
		}

		output, input = decodePath(args.Params.Path)
	default:
		return false
	}

	if input == (common.Address{}) || output == (common.Address{}) {
		return false
	}

	return (whitelist[input] || input == policy.weth) && (whitelist[output] || output == policy.weth)
}
