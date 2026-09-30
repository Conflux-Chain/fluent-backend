package service

import (
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/Conflux-Chain/fluent-backend/contract"
	uniswapv3 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/stretchr/testify/require"
)

func TestNewExecutionPolicy(t *testing.T) {
	routerV2 := common.HexToAddress("0x2001")
	routerV3 := common.HexToAddress("0x2002")
	weth := common.HexToAddress("0x2003")
	for _, test := range []struct {
		name      string
		v2        common.Address
		v3        common.Address
		whitelist map[common.Address]bool
		rpcError  bool
		wantCalls int32
		wantErr   string
	}{
		{name: "disabled"},
		{name: "V2 only", v2: routerV2, wantCalls: 1},
		{name: "V3 only", v3: routerV3, wantCalls: 1},
		{name: "both", v2: routerV2, v3: routerV3, wantCalls: 2},
		{name: "false whitelist entries", v2: routerV2, v3: routerV3, whitelist: map[common.Address]bool{routerV2: false, routerV3: false}, wantCalls: 2},
		{name: "V2 whitelist conflict", v2: routerV2, whitelist: map[common.Address]bool{routerV2: true}, wantErr: "Contract whitelist already contains the Uniswap V2 router"},
		{name: "V3 whitelist conflict", v3: routerV3, whitelist: map[common.Address]bool{routerV3: true}, wantErr: "Contract whitelist already contains the Uniswap V3 router"},
		{name: "duplicate router", v2: routerV2, v3: routerV2, wantCalls: 1, wantErr: "Composite execution policy already contains the Uniswap V3 router"},
		{name: "V2 RPC failure", v2: routerV2, rpcError: true, wantCalls: 1, wantErr: "Failed to create Uniswap V2 execution policy: Failed to get Uniswap V2 WETH"},
		{name: "V3 RPC failure", v3: routerV3, rpcError: true, wantCalls: 1, wantErr: "Failed to create Uniswap V3 execution policy: Failed to get Uniswap V3 WETH"},
	} {
		t.Run(test.name, func(t *testing.T) {
			var calls atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				var body struct {
					ID     json.RawMessage   `json:"id"`
					Method string            `json:"method"`
					Params []json.RawMessage `json:"params"`
				}
				if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
					t.Errorf("decode RPC request: %v", err)
					writer.WriteHeader(http.StatusBadRequest)
					return
				}
				calls.Add(1)
				if body.Method != "eth_call" || len(body.Params) == 0 {
					t.Errorf("unexpected RPC request: %s", body.Method)
					writer.WriteHeader(http.StatusBadRequest)
					return
				}
				var call struct {
					To    common.Address `json:"to"`
					Data  hexutil.Bytes  `json:"data"`
					Input hexutil.Bytes  `json:"input"`
				}
				if err := json.Unmarshal(body.Params[0], &call); err != nil {
					t.Errorf("decode contract call: %v", err)
					writer.WriteHeader(http.StatusBadRequest)
					return
				}
				data := call.Data
				if len(data) == 0 {
					data = call.Input
				}
				if call.To == test.v2 && test.v2 != (common.Address{}) {
					requireSelector := contract.UniswapV2RouterABI.Methods["WETH"].ID
					if hexutil.Encode(data) != hexutil.Encode(requireSelector) {
						t.Errorf("unexpected V2 selector: %x", data)
					}
				} else if call.To == test.v3 && test.v3 != (common.Address{}) {
					if hexutil.Encode(data) != hexutil.Encode(contract.UniswapV3RouterABI.Methods["WETH9"].ID) {
						t.Errorf("unexpected V3 selector: %x", data)
					}
				} else {
					t.Errorf("unexpected router: %v", call.To)
				}
				response := map[string]any{"jsonrpc": "2.0", "id": body.ID}
				if test.rpcError {
					response["error"] = map[string]any{"code": -32000, "message": "test RPC failure"}
				} else {
					response["result"] = hexutil.Encode(common.LeftPadBytes(weth.Bytes(), 32))
				}
				writer.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(writer).Encode(response); err != nil {
					t.Errorf("encode RPC response: %v", err)
				}
			}))
			t.Cleanup(server.Close)
			client, err := web3go.NewClient(server.URL)
			require.NoError(t, err)
			var config DeFiConfig
			config.Uniswap.V2.Router = test.v2
			config.Uniswap.V3.Router = test.v3
			policy, err := NewExecutionPolicy(config, client, test.whitelist)
			require.Equal(t, test.wantCalls, calls.Load())
			if test.wantErr != "" {
				require.ErrorContains(t, err, test.wantErr)
				require.Nil(t, policy)
				if test.rpcError {
					require.ErrorContains(t, err, "test RPC failure")
				}
				return
			}
			require.NoError(t, err)
			require.IsType(t, CompositeExecutionPolicy{}, policy)
			composite := policy.(CompositeExecutionPolicy)
			require.Len(t, composite, int(test.wantCalls))
			if test.v2 != (common.Address{}) {
				require.Equal(t, &UniswapV2ExecutionPolicy{router: test.v2, weth: weth}, composite[test.v2])
			}
			if test.v3 != (common.Address{}) {
				require.Equal(t, &UniswapV3ExecutionPolicy{router: test.v3, weth: weth}, composite[test.v3])
			}
		})
	}
}

type executionPolicyFunc func(contract.Execution, map[common.Address]bool) error

func (policy executionPolicyFunc) Validate(execution contract.Execution, whitelist map[common.Address]bool) error {
	return policy(execution, whitelist)
}

func TestCompositeExecutionPolicyValidate(t *testing.T) {
	target := common.HexToAddress("0x1001")
	execution := contract.Execution{Target: target, Value: big.NewInt(7), CallData: []byte{1, 2, 3, 4}}
	for _, test := range []struct {
		name       string
		whitelist  map[common.Address]bool
		registered bool
		childErr   error
		wantCalled bool
		wantErr    error
	}{
		{name: "whitelist bypasses child", whitelist: map[common.Address]bool{target: true}, registered: true, childErr: errors.New("rejected")},
		{name: "whitelist without child", whitelist: map[common.Address]bool{target: true}},
		{name: "unknown target", wantErr: ErrVerifyingPaymasterContractNotWhitelisted.WithData(target)},
		{name: "false whitelist entry", whitelist: map[common.Address]bool{target: false}, wantErr: ErrVerifyingPaymasterContractNotWhitelisted.WithData(target)},
		{name: "child accepts", registered: true, wantCalled: true},
		{name: "false entry delegates", whitelist: map[common.Address]bool{target: false}, registered: true, wantCalled: true},
		{name: "child rejects", registered: true, childErr: errors.New("rejected"), wantCalled: true, wantErr: ErrVerifyingPaymasterContractNotWhitelisted.WithData("rejected")},
	} {
		t.Run(test.name, func(t *testing.T) {
			called := false
			composite := CompositeExecutionPolicy{}
			if test.registered {
				composite[target] = executionPolicyFunc(func(got contract.Execution, whitelist map[common.Address]bool) error {
					called = true
					require.Equal(t, execution, got)
					require.Equal(t, test.whitelist, whitelist)
					return test.childErr
				})
			}
			var policy ExecutionPolicy = composite
			require.Equal(t, test.wantErr, policy.Validate(execution, test.whitelist))
			require.Equal(t, test.wantCalled, called)
		})
	}
}

func TestUniswapV2ExecutionPolicyValidate(t *testing.T) {
	router := common.HexToAddress("0x1001")
	weth := common.HexToAddress("0x1002")
	input := common.HexToAddress("0x1003")
	output := common.HexToAddress("0x1004")
	unknown := common.HexToAddress("0x1005")
	var policy ExecutionPolicy = &UniswapV2ExecutionPolicy{router: router, weth: weth}
	whitelist := map[common.Address]bool{input: true, output: true, unknown: false}
	for _, method := range []struct {
		name      string
		ethInput  bool
		ethOutput bool
	}{
		{name: "swapExactTokensForTokens"},
		{name: "swapTokensForExactTokens"},
		{name: "swapExactTokensForTokensSupportingFeeOnTransferTokens"},
		{name: "swapExactTokensForETH", ethOutput: true},
		{name: "swapTokensForExactETH", ethOutput: true},
		{name: "swapExactTokensForETHSupportingFeeOnTransferTokens", ethOutput: true},
		{name: "swapExactETHForTokens", ethInput: true},
		{name: "swapETHForExactTokens", ethInput: true},
		{name: "swapExactETHForTokensSupportingFeeOnTransferTokens", ethInput: true},
	} {
		t.Run(method.name, func(t *testing.T) {
			first, last := input, output
			var value *big.Int
			if method.ethInput {
				first, value = weth, big.NewInt(1)
			}
			if method.ethOutput {
				last = weth
			}
			pack := func(path []common.Address) []byte {
				args := []any{big.NewInt(1)}
				if !method.ethInput {
					args = append(args, big.NewInt(1))
				}
				args = append(args, path, input, big.NewInt(100))
				data, err := contract.UniswapV2RouterABI.Pack(method.name, args...)
				require.NoError(t, err)
				return data
			}
			for _, test := range []struct {
				name    string
				path    []common.Address
				value   *big.Int
				wantErr string
			}{
				{name: "valid", path: []common.Address{first, last}, value: value},
				{name: "unlisted intermediate token", path: []common.Address{first, unknown, last}, value: value},
				{name: "empty path", value: value, wantErr: "Invalid path"},
				{name: "single token path", path: []common.Address{first}, value: value, wantErr: "Invalid path"},
				{name: "invalid input", path: []common.Address{unknown, last}, value: value, wantErr: "Input token"},
				{name: "invalid output", path: []common.Address{first, unknown}, value: value, wantErr: "Output token"},
			} {
				t.Run(test.name, func(t *testing.T) {
					err := policy.Validate(contract.Execution{Target: router, Value: test.value, CallData: pack(test.path)}, whitelist)
					if test.wantErr == "" {
						require.NoError(t, err)
					} else {
						require.ErrorContains(t, err, test.wantErr)
					}
				})
			}
			data := pack([]common.Address{first, last})
			for _, amount := range []*big.Int{nil, big.NewInt(0), big.NewInt(1), big.NewInt(-1)} {
				err := policy.Validate(contract.Execution{Target: router, Value: amount, CallData: data}, whitelist)
				if method.ethInput {
					if amount == nil || amount.Sign() <= 0 {
						require.ErrorContains(t, err, "Positive msg.value required")
					} else {
						require.NoError(t, err)
					}
				} else if amount != nil && amount.Sign() != 0 {
					require.ErrorContains(t, err, "Zero msg.value required")
				} else {
					require.NoError(t, err)
				}
			}
			require.ErrorContains(t, policy.Validate(contract.Execution{Target: router, Value: value, CallData: data[:4]}, whitelist), "Failed to unpack calldata")
			if !method.ethInput && !method.ethOutput {
				for _, path := range [][]common.Address{{weth, output}, {input, weth}} {
					execution := contract.Execution{Target: router, CallData: pack(path)}
					require.ErrorContains(t, policy.Validate(execution, whitelist), "token not in whitelist")
					require.NoError(t, policy.Validate(execution, map[common.Address]bool{input: true, output: true, weth: true}))
				}
			}
		})
	}
}

func packV3Execution(t *testing.T, method string, input, output common.Address, path []byte) []byte {
	t.Helper()
	var params any
	switch method {
	case "exactInputSingle":
		params = uniswapv3.ISwapRouterExactInputSingleParams{TokenIn: input, TokenOut: output, Fee: big.NewInt(3000), Recipient: input, Deadline: big.NewInt(100), AmountIn: big.NewInt(1), AmountOutMinimum: big.NewInt(1), SqrtPriceLimitX96: big.NewInt(0)}
	case "exactOutputSingle":
		params = uniswapv3.ISwapRouterExactOutputSingleParams{TokenIn: input, TokenOut: output, Fee: big.NewInt(3000), Recipient: input, Deadline: big.NewInt(100), AmountOut: big.NewInt(1), AmountInMaximum: big.NewInt(1), SqrtPriceLimitX96: big.NewInt(0)}
	case "exactInput":
		params = uniswapv3.ISwapRouterExactInputParams{Path: path, Recipient: input, Deadline: big.NewInt(100), AmountIn: big.NewInt(1), AmountOutMinimum: big.NewInt(1)}
	case "exactOutput":
		params = uniswapv3.ISwapRouterExactOutputParams{Path: path, Recipient: input, Deadline: big.NewInt(100), AmountOut: big.NewInt(1), AmountInMaximum: big.NewInt(1)}
	default:
		t.Fatalf("unsupported test method %q", method)
	}
	data, err := contract.UniswapV3RouterABI.Pack(method, params)
	require.NoError(t, err)
	return data
}

func TestUniswapV3ExecutionPolicyValidate(t *testing.T) {
	router := common.HexToAddress("0x1001")
	weth := common.HexToAddress("0x1002")
	input := common.HexToAddress("0x1003")
	output := common.HexToAddress("0x1004")
	unknown := common.HexToAddress("0x1005")
	var policy ExecutionPolicy = &UniswapV3ExecutionPolicy{router: router, weth: weth}
	whitelist := map[common.Address]bool{input: true, output: true, unknown: false}
	for _, method := range []string{"exactInputSingle", "exactOutputSingle", "exactInput", "exactOutput"} {
		t.Run(method, func(t *testing.T) {
			for _, test := range []struct {
				name     string
				input    common.Address
				output   common.Address
				multiHop bool
				wantErr  string
			}{
				{name: "valid", input: input, output: output},
				{name: "WETH input", input: weth, output: output},
				{name: "WETH output", input: input, output: weth},
				{name: "invalid input", input: unknown, output: output, wantErr: "Input token"},
				{name: "invalid output", input: input, output: unknown, wantErr: "Output token"},
				{name: "unlisted intermediate token", input: input, output: output, multiHop: true},
			} {
				t.Run(test.name, func(t *testing.T) {
					first, last := test.input, test.output
					if method == "exactOutput" {
						first, last = last, first
					}
					path := append(first.Bytes(), 0, 11, 184)
					if test.multiHop {
						path = append(path, unknown.Bytes()...)
						path = append(path, 0, 1, 244)
					}
					path = append(path, last.Bytes()...)
					data := packV3Execution(t, method, test.input, test.output, path)
					err := policy.Validate(contract.Execution{Target: router, CallData: data}, whitelist)
					if test.wantErr == "" {
						require.NoError(t, err)
					} else {
						require.ErrorContains(t, err, test.wantErr)
					}
				})
			}
			selector := contract.UniswapV3RouterABI.Methods[method].ID
			require.ErrorContains(t, policy.Validate(contract.Execution{Target: router, CallData: selector}, whitelist), "Failed to unpack calldata")
			if method == "exactInput" || method == "exactOutput" {
				for _, length := range []int{0, 20, 42, 44, 65, 67} {
					data := packV3Execution(t, method, input, output, make([]byte, length))
					require.ErrorContains(t, policy.Validate(contract.Execution{Target: router, CallData: data}, whitelist), "Invalid path length")
				}
			}
		})
	}
}

func TestExecutionPolicyInvalidCallData(t *testing.T) {
	router := common.HexToAddress("0x1001")
	for _, test := range []struct {
		name        string
		policy      ExecutionPolicy
		unsupported []byte
	}{
		{name: "V2", policy: &UniswapV2ExecutionPolicy{router: router}, unsupported: contract.UniswapV2RouterABI.Methods["WETH"].ID},
		{name: "V3", policy: &UniswapV3ExecutionPolicy{router: router}, unsupported: contract.UniswapV3RouterABI.Methods["WETH9"].ID},
	} {
		t.Run(test.name, func(t *testing.T) {
			require.Len(t, test.unsupported, 4)
			require.ErrorContains(t, test.policy.Validate(contract.Execution{}, nil), "Target contract not in whitelist")
			for _, data := range [][]byte{nil, {1}, {1, 2}, {1, 2, 3}} {
				require.ErrorContains(t, test.policy.Validate(contract.Execution{Target: router, CallData: data}, nil), "Call data too short")
			}
			require.ErrorContains(t, test.policy.Validate(contract.Execution{Target: router, CallData: []byte{0xff, 0xff, 0xff, 0xff}}, nil), "Failed to find method by ID")
			require.ErrorContains(t, test.policy.Validate(contract.Execution{Target: router, CallData: test.unsupported}, nil), "Method not in whitelist")
		})
	}
}
