package service

import (
	"math/big"
	"testing"

	"github.com/Conflux-Chain/fluent-backend/contract"
	uniswapv3 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockExecutionPolicy struct {
	mock.Mock
}

func (policy *mockExecutionPolicy) IsAllowed(execution contract.Execution, whitelist map[common.Address]bool) bool {
	return policy.Called(execution, whitelist).Bool(0)
}

func TestCompositeExecutionPolicy(t *testing.T) {
	mockPolicy := new(mockExecutionPolicy)
	policy := CompositeExecutionPolicy{TargetContractExecutionPolicy{}, mockPolicy}
	whitelist := map[common.Address]bool{common.HexToAddress("0x1111"): true}

	// target contract in whitelist
	execution := contract.Execution{Target: common.HexToAddress("0x1111"), Value: big.NewInt(0), CallData: []byte{}}
	require.True(t, policy.IsAllowed(execution, whitelist))
	mockPolicy.AssertExpectations(t)

	// target contract not in whitelist, but dapp policy allows
	execution = contract.Execution{Target: common.HexToAddress("0x2222"), Value: big.NewInt(0), CallData: []byte{}}
	mockPolicy.On("IsAllowed", execution, whitelist).Return(true)
	require.True(t, policy.IsAllowed(execution, whitelist))
	mockPolicy.AssertExpectations(t)

	// target contract not in whitelist, and dapp policy disallows
	execution = contract.Execution{Target: common.HexToAddress("0x3333"), Value: big.NewInt(0), CallData: []byte{}}
	mockPolicy.On("IsAllowed", execution, whitelist).Return(false)
	require.False(t, policy.IsAllowed(execution, whitelist))
	mockPolicy.AssertExpectations(t)
}

func TestUniswapV2ExecutionPolicy(t *testing.T) {
	whitelist := map[common.Address]bool{
		common.HexToAddress("0x1111"): true,
		common.HexToAddress("0x2222"): true,
	}

	policy := UniswapV2ExecutionPolicy{
		router: common.HexToAddress("0x3333"),
		weth:   common.HexToAddress("0x4444"),
	}

	// target contract mismatch
	execution := contract.Execution{Target: common.HexToAddress("0x5555"), Value: big.NewInt(0), CallData: []byte{}}
	require.False(t, policy.IsAllowed(execution, whitelist))

	// invalid calldata length
	execution = contract.Execution{Target: common.HexToAddress("0x3333"), Value: big.NewInt(0), CallData: []byte{0x01}}
	require.False(t, policy.IsAllowed(execution, whitelist))

	// helper function
	isAllowed := func(value *big.Int, method string, args ...any) bool {
		callData, err := contract.UniswapV2RouterABI.Pack(method, args...)
		if err != nil {
			panic(err)
		}

		return policy.IsAllowed(contract.Execution{Target: common.HexToAddress("0x3333"), Value: value, CallData: callData}, whitelist)
	}

	// unsupported function
	require.False(t, isAllowed(big.NewInt(0), "getAmountsOut",
		big.NewInt(1), []common.Address{common.HexToAddress("0x1111"), common.HexToAddress("0x2222")}))

	// invalid msg.value - non-ETH
	require.False(t, isAllowed(big.NewInt(1), "swapExactTokensForTokens",
		big.NewInt(1), big.NewInt(1), []common.Address{common.HexToAddress("0x1111"), common.HexToAddress("0x2222")}, common.HexToAddress("0x01"), big.NewInt(1)))

	// invalid msg.value - ETH
	require.False(t, isAllowed(big.NewInt(0), "swapExactETHForTokens",
		big.NewInt(1), []common.Address{policy.weth, common.HexToAddress("0x2222")}, common.HexToAddress("0x01"), big.NewInt(1)))

	// ETH to token - in whitelist
	require.True(t, isAllowed(big.NewInt(1), "swapExactETHForTokens",
		big.NewInt(1), []common.Address{policy.weth, common.HexToAddress("0x1111")}, common.HexToAddress("0x01"), big.NewInt(1)))

	// token to ETH - in whitelist
	require.True(t, isAllowed(big.NewInt(0), "swapExactTokensForETH",
		big.NewInt(1), big.NewInt(1), []common.Address{common.HexToAddress("0x1111"), policy.weth}, common.HexToAddress("0x01"), big.NewInt(1)))

	// token to token - in whitelist
	require.True(t, isAllowed(big.NewInt(0), "swapExactTokensForTokens",
		big.NewInt(1), big.NewInt(1), []common.Address{common.HexToAddress("0x1111"), common.HexToAddress("0x2222")}, common.HexToAddress("0x01"), big.NewInt(1)))

	// token to token - not in whitelist
	require.False(t, isAllowed(big.NewInt(0), "swapExactTokensForTokens",
		big.NewInt(1), big.NewInt(1), []common.Address{common.HexToAddress("0x1111"), common.HexToAddress("0x222222")}, common.HexToAddress("0x01"), big.NewInt(1)))
}

func TestUniswapV2ExecutionPolicySupportingFeeOnTransferTokens(t *testing.T) {
	tokenIn := common.HexToAddress("0x1111")
	tokenOut := common.HexToAddress("0x2222")
	unknown := common.HexToAddress("0x5555")
	whitelist := map[common.Address]bool{tokenIn: true, tokenOut: true}
	policy := UniswapV2ExecutionPolicy{
		router: common.HexToAddress("0x3333"),
		weth:   common.HexToAddress("0x4444"),
	}

	for _, swap := range []struct {
		method string
		input  common.Address
		output common.Address
		eth    bool
	}{
		{"swapExactTokensForTokensSupportingFeeOnTransferTokens", tokenIn, tokenOut, false},
		{"swapExactTokensForETHSupportingFeeOnTransferTokens", tokenIn, policy.weth, false},
		{"swapExactETHForTokensSupportingFeeOnTransferTokens", policy.weth, tokenOut, true},
	} {
		t.Run(swap.method, func(t *testing.T) {
			value := big.NewInt(0)
			if swap.eth {
				value = big.NewInt(1)
			}
			for _, test := range []struct {
				name    string
				path    []common.Address
				value   *big.Int
				allowed bool
			}{
				{"valid", []common.Address{swap.input, swap.output}, value, true},
				{"intermediate token", []common.Address{swap.input, unknown, swap.output}, value, true},
				{"invalid input", []common.Address{unknown, swap.output}, value, false},
				{"invalid output", []common.Address{swap.input, unknown}, value, false},
				{"wrong WETH input", []common.Address{tokenOut, swap.output}, value, !swap.eth},
				{"wrong WETH output", []common.Address{swap.input, tokenIn}, value, swap.output != policy.weth},
				{"WETH endpoints", []common.Address{policy.weth, policy.weth}, value, false},
				{"empty path", nil, value, false},
				{"short path", []common.Address{swap.input}, value, false},
				{"nil value", []common.Address{swap.input, swap.output}, nil, !swap.eth},
				{"zero value", []common.Address{swap.input, swap.output}, big.NewInt(0), !swap.eth},
				{"positive value", []common.Address{swap.input, swap.output}, big.NewInt(1), swap.eth},
				{"negative value", []common.Address{swap.input, swap.output}, big.NewInt(-1), false},
			} {
				t.Run(test.name, func(t *testing.T) {
					args := []any{big.NewInt(1), test.path, common.HexToAddress("0x01"), big.NewInt(1)}
					if !swap.eth {
						args = append([]any{big.NewInt(1)}, args...)
					}
					callData, err := contract.UniswapV2RouterABI.Pack(swap.method, args...)
					require.NoError(t, err)
					execution := contract.Execution{Target: policy.router, Value: test.value, CallData: callData}
					require.Equal(t, test.allowed, policy.IsAllowed(execution, whitelist))
					execution.CallData = callData[:4]
					require.False(t, policy.IsAllowed(execution, whitelist))
				})
			}
		})
	}
}

func TestUniswapV3ExecutionPolicy(t *testing.T) {
	whitelist := map[common.Address]bool{
		common.HexToAddress("0x1111"): true,
		common.HexToAddress("0x2222"): true,
	}

	policy := UniswapV3ExecutionPolicy{
		router: common.HexToAddress("0x3333"),
		weth:   common.HexToAddress("0x4444"),
	}

	// target contract mismatch
	execution := contract.Execution{Target: common.HexToAddress("0x5555"), Value: big.NewInt(0), CallData: []byte{}}
	require.False(t, policy.IsAllowed(execution, whitelist))

	// invalid calldata length
	execution = contract.Execution{Target: common.HexToAddress("0x3333"), Value: big.NewInt(0), CallData: []byte{0x01}}
	require.False(t, policy.IsAllowed(execution, whitelist))

	// helper function
	isAllowed := func(value *big.Int, method string, args ...any) bool {
		callData, err := contract.UniswapV3RouterABI.Pack(method, args...)
		if err != nil {
			panic(err)
		}

		return policy.IsAllowed(contract.Execution{Target: common.HexToAddress("0x3333"), Value: value, CallData: callData}, whitelist)
	}

	newUniswapV3Path := func(tokenIn common.Address, fee *big.Int, tokenOut common.Address) []byte {
		var path [43]byte
		copy(path[0:20], tokenIn.Bytes())
		fee.FillBytes(path[20:23])
		copy(path[23:43], tokenOut.Bytes())
		return path[:]
	}

	// unsupported function
	require.False(t, isAllowed(big.NewInt(0), "multicall", [][]byte{}))

	// exactInputSingle: token to token
	require.True(t, isAllowed(big.NewInt(0), "exactInputSingle", uniswapv3.ISwapRouterExactInputSingleParams{
		TokenIn:           common.HexToAddress("0x1111"),
		TokenOut:          common.HexToAddress("0x2222"),
		Fee:               big.NewInt(1),
		Recipient:         common.HexToAddress("0x0001"),
		Deadline:          big.NewInt(1),
		AmountIn:          big.NewInt(1),
		AmountOutMinimum:  big.NewInt(1),
		SqrtPriceLimitX96: big.NewInt(0),
	}))

	// exactInput: token to ETH
	require.True(t, isAllowed(big.NewInt(0), "exactInput", uniswapv3.ISwapRouterExactInputParams{
		Path:             newUniswapV3Path(common.HexToAddress("0x1111"), big.NewInt(1), policy.weth),
		Recipient:        common.HexToAddress("0x0001"),
		Deadline:         big.NewInt(1),
		AmountIn:         big.NewInt(1),
		AmountOutMinimum: big.NewInt(1),
	}))

	// exactOutputSingle: ETH to token
	require.True(t, isAllowed(big.NewInt(0), "exactOutputSingle", uniswapv3.ISwapRouterExactOutputSingleParams{
		TokenIn:           policy.weth,
		TokenOut:          common.HexToAddress("0x1111"),
		Fee:               big.NewInt(1),
		Recipient:         common.HexToAddress("0x0001"),
		Deadline:          big.NewInt(1),
		AmountOut:         big.NewInt(1),
		AmountInMaximum:   big.NewInt(1),
		SqrtPriceLimitX96: big.NewInt(0),
	}))

	// exactOutput: token to token
	require.True(t, isAllowed(big.NewInt(0), "exactOutput", uniswapv3.ISwapRouterExactOutputParams{
		Path:            newUniswapV3Path(common.HexToAddress("0x1111"), big.NewInt(1), common.HexToAddress("0x2222")),
		Recipient:       common.HexToAddress("0x0001"),
		Deadline:        big.NewInt(1),
		AmountOut:       big.NewInt(1),
		AmountInMaximum: big.NewInt(1),
	}))

	// exactOutput: token to non-whitelisted token
	require.False(t, isAllowed(big.NewInt(0), "exactOutput", uniswapv3.ISwapRouterExactOutputParams{
		Path:            newUniswapV3Path(common.HexToAddress("0x222222"), big.NewInt(1), common.HexToAddress("0x1111")),
		Recipient:       common.HexToAddress("0x0001"),
		Deadline:        big.NewInt(1),
		AmountOut:       big.NewInt(1),
		AmountInMaximum: big.NewInt(1),
	}))
}
