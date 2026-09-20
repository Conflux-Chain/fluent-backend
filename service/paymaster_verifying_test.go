package service

import (
	"math/big"
	"testing"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	uniswapv3 "github.com/Conflux-Chain/go-conflux-util/blockchain/contract/defi/uniswap/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go/signers"
	"github.com/stretchr/testify/assert"
)

var testVerifyingPaymasterConfig = VerifyingPaymasterConfig{
	PaymasterConfig: PaymasterConfig{
		Address:          common.HexToAddress("0x7777"),
		SignatureTimeout: time.Minute * 5,
	},
	smartAccountMap: map[common.Address]bool{
		common.HexToAddress("0x9999"): true,
	},
	ContractWhitelist: []common.Address{common.HexToAddress("0x1111"), common.HexToAddress("0x2222")},
	contractMap: map[common.Address]bool{
		common.HexToAddress("0x1111"): true,
		common.HexToAddress("0x2222"): true,
	},
	maxGasCostBig: big.NewInt(100000000000000000),
}

var uniswapV2Policy = &UniswapV2ExecutionPolicy{
	router: common.HexToAddress("0x3333"),
	weth:   common.HexToAddress("0x4444"),
}

var uniswapV3Policy = &UniswapV3ExecutionPolicy{
	router: common.HexToAddress("0x5555"),
	weth:   common.HexToAddress("0x6666"),
}

func newTestVerifyingPaymaster() *VerifyingPaymaster {
	return &VerifyingPaymaster{
		inner: &Paymaster[*contract.VerifyingPaymasterCaller]{
			config: PaymasterConfig{
				Address:          testVerifyingPaymasterConfig.Address,
				SignatureTimeout: testVerifyingPaymasterConfig.SignatureTimeout,
			},
			caller: nil,
			signer: signers.MustNewRandomPrivateKeySigner(),
		},
		config: testVerifyingPaymasterConfig,
		executionPolicy: CompositeExecutionPolicy{
			TargetContractExecutionPolicy{},
			uniswapV2Policy,
			uniswapV3Policy,
		},
	}
}

func TestVerifyingPaymasterValidateCallData(t *testing.T) {
	paymaster := newTestVerifyingPaymaster()

	// nil calldata
	assert.Error(t, paymaster.validateCallData(nil))

	// invalid selector
	assert.Error(t, paymaster.validateCallData(hexutil.MustDecode("0xa9059cbb"))) // transfer(address,uint256)

	// execute - target not whitelisted
	callData, err := contract.SmartAccountABI.Pack("execute", common.HexToAddress("0x1112"), big.NewInt(10), []byte{1, 2, 3})
	assert.NoError(t, err)
	expectedBizErr := ErrVerifyingPaymasterContractNotWhitelisted.WithData(common.HexToAddress("0x1112"))
	assert.Equal(t, expectedBizErr, paymaster.validateCallData(callData))

	// execute - valid
	callData, err = contract.SmartAccountABI.Pack("execute", common.HexToAddress("0x1111"), big.NewInt(10), []byte{1, 2, 3})
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(callData))

	// executeBatch - empty
	callData, err = contract.SmartAccountABI.Pack("executeBatch", []contract.Execution{})
	assert.NoError(t, err)
	assert.Error(t, paymaster.validateCallData(callData))

	// executeBatch - target not whitelisted
	callData, err = contract.SmartAccountABI.Pack("executeBatch", []contract.Execution{
		{Target: common.HexToAddress("0x1111"), Value: big.NewInt(10), CallData: []byte{1, 2, 3}},
		{Target: common.HexToAddress("0x2223"), Value: big.NewInt(20), CallData: []byte{4, 5, 6}},
	})
	assert.NoError(t, err)
	expectedBizErr = ErrVerifyingPaymasterContractNotWhitelisted.WithData(common.HexToAddress("0x2223"))
	assert.Equal(t, expectedBizErr, paymaster.validateCallData(callData))

	// executeBatch - valid
	callData, err = contract.SmartAccountABI.Pack("executeBatch", []contract.Execution{
		{Target: common.HexToAddress("0x1111"), Value: big.NewInt(10), CallData: []byte{1, 2, 3}},
		{Target: common.HexToAddress("0x2222"), Value: big.NewInt(20), CallData: []byte{4, 5, 6}},
	})
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(callData))
}

func TestVerifyingPaymasterUniswapV2Policy(t *testing.T) {
	paymaster := newTestVerifyingPaymaster()

	// ETH to token - in whitelist
	path := []common.Address{uniswapV2Policy.weth, testVerifyingPaymasterConfig.ContractWhitelist[0]}
	swapCallData, err := contract.UniswapV2RouterABI.Pack("swapExactETHForTokens", big.NewInt(1), path, common.HexToAddress("0x0001"), big.NewInt(1))
	assert.NoError(t, err)
	executeCallData, err := contract.SmartAccountABI.Pack("execute", uniswapV2Policy.router, big.NewInt(1), swapCallData) // msg.value > 0
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(executeCallData))

	// token to ETH - in whitelist
	path = []common.Address{testVerifyingPaymasterConfig.ContractWhitelist[0], uniswapV2Policy.weth}
	swapCallData, err = contract.UniswapV2RouterABI.Pack("swapExactTokensForETH", big.NewInt(1), big.NewInt(1), path, common.HexToAddress("0x0001"), big.NewInt(1))
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV2Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(executeCallData))

	// token to token - in whitelist
	path = []common.Address{testVerifyingPaymasterConfig.ContractWhitelist[0], testVerifyingPaymasterConfig.ContractWhitelist[1]}
	swapCallData, err = contract.UniswapV2RouterABI.Pack("swapExactTokensForTokens", big.NewInt(1), big.NewInt(1), path, common.HexToAddress("0x0001"), big.NewInt(1))
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV2Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(executeCallData))

	// token to token - not in whitelist
	paymaster.config.contractMap = map[common.Address]bool{testVerifyingPaymasterConfig.ContractWhitelist[0]: true}
	assert.Error(t, paymaster.validateCallData(executeCallData))

	// unsupported function
	unsupportedCallData, err := contract.UniswapV2RouterABI.Pack("getAmountsOut", big.NewInt(1), []common.Address{testVerifyingPaymasterConfig.ContractWhitelist[0], testVerifyingPaymasterConfig.ContractWhitelist[1]})
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV2Policy.router, big.NewInt(0), unsupportedCallData)
	assert.NoError(t, err)
	assert.Error(t, paymaster.validateCallData(executeCallData))
}

func TestVerifyingPaymasterUniswapV3Policy(t *testing.T) {
	paymaster := newTestVerifyingPaymaster()

	// unsupported function
	swapCallData, err := contract.UniswapV3RouterABI.Pack("multicall", [][]byte{})
	assert.NoError(t, err)
	executeCallData, err := contract.SmartAccountABI.Pack("execute", uniswapV3Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.Error(t, paymaster.validateCallData(executeCallData))

	newUniswapV3Path := func(tokenIn common.Address, fee *big.Int, tokenOut common.Address) []byte {
		var path [43]byte
		copy(path[0:20], tokenIn.Bytes())
		fee.FillBytes(path[20:23])
		copy(path[23:43], tokenOut.Bytes())
		return path[:]
	}

	// exactInputSingle: token to token
	swapCallData, err = contract.UniswapV3RouterABI.Pack("exactInputSingle", uniswapv3.ISwapRouterExactInputSingleParams{
		TokenIn:           testVerifyingPaymasterConfig.ContractWhitelist[0],
		TokenOut:          testVerifyingPaymasterConfig.ContractWhitelist[1],
		Fee:               big.NewInt(1),
		Recipient:         common.HexToAddress("0x0001"),
		Deadline:          big.NewInt(1),
		AmountIn:          big.NewInt(1),
		AmountOutMinimum:  big.NewInt(1),
		SqrtPriceLimitX96: big.NewInt(0),
	})
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV3Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(executeCallData))

	// exactInput: token to ETH
	swapCallData, err = contract.UniswapV3RouterABI.Pack("exactInput", uniswapv3.ISwapRouterExactInputParams{
		Path:             newUniswapV3Path(testVerifyingPaymasterConfig.ContractWhitelist[0], big.NewInt(1), uniswapV3Policy.weth),
		Recipient:        common.HexToAddress("0x0001"),
		Deadline:         big.NewInt(1),
		AmountIn:         big.NewInt(1),
		AmountOutMinimum: big.NewInt(1),
	})
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV3Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(executeCallData))

	// exactOutputSingle: ETH to token
	swapCallData, err = contract.UniswapV3RouterABI.Pack("exactOutputSingle", uniswapv3.ISwapRouterExactOutputSingleParams{
		TokenIn:           uniswapV3Policy.weth,
		TokenOut:          testVerifyingPaymasterConfig.ContractWhitelist[0],
		Fee:               big.NewInt(1),
		Recipient:         common.HexToAddress("0x0001"),
		Deadline:          big.NewInt(1),
		AmountOut:         big.NewInt(1),
		AmountInMaximum:   big.NewInt(1),
		SqrtPriceLimitX96: big.NewInt(0),
	})
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV3Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(executeCallData))

	// exactOutput: token to token
	params := uniswapv3.ISwapRouterExactOutputParams{
		Path:            newUniswapV3Path(testVerifyingPaymasterConfig.ContractWhitelist[0], big.NewInt(1), testVerifyingPaymasterConfig.ContractWhitelist[1]),
		Recipient:       common.HexToAddress("0x0001"),
		Deadline:        big.NewInt(1),
		AmountOut:       big.NewInt(1),
		AmountInMaximum: big.NewInt(1),
	}
	swapCallData, err = contract.UniswapV3RouterABI.Pack("exactOutput", params)
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV3Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(executeCallData))

	// exactOutput: token to non-whitelisted token
	params.Path = newUniswapV3Path(testVerifyingPaymasterConfig.ContractWhitelist[0], big.NewInt(1), common.HexToAddress("0x0002"))
	swapCallData, err = contract.UniswapV3RouterABI.Pack("exactOutput", params)
	assert.NoError(t, err)
	executeCallData, err = contract.SmartAccountABI.Pack("execute", uniswapV3Policy.router, big.NewInt(0), swapCallData)
	assert.NoError(t, err)
	assert.Error(t, paymaster.validateCallData(executeCallData))
}
