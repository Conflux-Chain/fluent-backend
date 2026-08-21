package service

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go/signers"
	"github.com/stretchr/testify/assert"
)

var testVerifyingPaymasterConfig = VerifyingPaymasterConfig{
	Address: common.HexToAddress("0x6666"),
	contractWhitelist: map[common.Address]bool{
		common.HexToAddress("0x1111"): true,
		common.HexToAddress("0x2222"): true,
	},
	maxGasCost:       big.NewInt(100000000000000000),
	SignatureTimeout: time.Minute * 5,
}

func assertNewTestVerifyingPaymaster(t *testing.T) *VerifyingPaymaster {
	smartAccountABI, err := abi.JSON(strings.NewReader(contract.SimpleSmartAccount7702MetaData.ABI))
	assert.NoError(t, err)
	executeMethod, ok := smartAccountABI.Methods["execute"]
	assert.True(t, ok)
	executeBatchMethod, ok := smartAccountABI.Methods["executeBatch"]
	assert.True(t, ok)

	return &VerifyingPaymaster{
		config:             testVerifyingPaymasterConfig,
		executeMethod:      &executeMethod,
		executeBatchMethod: &executeBatchMethod,
		signer:             signers.MustNewRandomPrivateKeySigner(),
	}
}

func assertVerifyingPaymasterValidateBizErr(t *testing.T, expectedBizErr *api.BusinessError, userOp contract.PackedUserOperation, delegation ...common.Address) {
	paymaster := assertNewTestVerifyingPaymaster(t)

	var delegatedContract common.Address
	if len(delegation) > 0 {
		delegatedContract = delegation[0]
	}
	err := paymaster.validate(&userOp, delegatedContract)
	assert.Error(t, err)

	bizErr, ok := err.(*api.BusinessError)
	assert.True(t, ok)
	assert.Equal(t, expectedBizErr.Code, bizErr.Code)
	assert.Equal(t, expectedBizErr.Message, bizErr.Message)

	bizErrData, ok := bizErr.Data.(string)
	assert.True(t, ok)
	assert.True(t, len(bizErrData) > 0)

	if expectedBizErr.Data == nil {
		return
	}

	expectedBizErrData, ok := expectedBizErr.Data.(string)
	assert.True(t, ok)
	assert.True(t, len(expectedBizErrData) > 0)
	assert.True(t, strings.Contains(bizErrData, expectedBizErrData), fmt.Sprintf("Expected: %v, actual = %v", expectedBizErrData, bizErrData))
}

func TestVerifyingPaymasterValidate(t *testing.T) {
	// paymasterAndData - invalid length
	assertVerifyingPaymasterValidateBizErr(t, api.ErrValidationStr("Invalid paymaster data length"), contract.PackedUserOperation{
		PaymasterAndData: nil,
	})

	assertVerifyingPaymasterValidateBizErr(t, api.ErrValidationStr("Invalid paymaster data length"), contract.PackedUserOperation{
		PaymasterAndData: make([]byte, verifyingPaymasterDataLength+1),
	})

	// paymasterAndData - invalid address
	var data [verifyingPaymasterDataLength]byte
	copy(data[:20], common.HexToAddress("0x6667").Bytes())
	assertVerifyingPaymasterValidateBizErr(t, api.ErrValidationStr("Invalid paymaster address"), contract.PackedUserOperation{
		PaymasterAndData: data[:],
	})

	// initCode - not empty
	paymasterAndData := make([]byte, verifyingPaymasterDataLength)
	copy(paymasterAndData[:20], testVerifyingPaymasterConfig.Address.Bytes())

	assertVerifyingPaymasterValidateBizErr(t, api.ErrValidationStr("Invalid initCode"), contract.PackedUserOperation{
		PaymasterAndData: paymasterAndData,
		InitCode:         []byte{1},
	})

	// initCode - invalid 7702 marker
	assertVerifyingPaymasterValidateBizErr(t, api.ErrValidationStr("Invalid initCode"), contract.PackedUserOperation{
		PaymasterAndData: paymasterAndData,
		InitCode:         []byte{1},
	}, common.HexToAddress("0x9999"))

	// maxCost - exceed maxGasCost
	userOp := contract.PackedUserOperation{
		PaymasterAndData:   paymasterAndData,
		PreVerificationGas: testVerifyingPaymasterConfig.maxGasCost,
	}

	big.NewInt(1).FillBytes(userOp.AccountGasLimits[:16])   // verification gas limit
	big.NewInt(2).FillBytes(userOp.AccountGasLimits[16:32]) // call gas limit
	big.NewInt(4).FillBytes(userOp.PaymasterAndData[20:36]) // paymaster verification gas limit
	big.NewInt(5).FillBytes(userOp.PaymasterAndData[36:52]) // paymaster postOp gas limit
	big.NewInt(1).FillBytes(userOp.GasFees[16:])            // maxFeePerGas

	var tmp VerifyingPaymaster
	expectedMaxCost := new(big.Int).Add(testVerifyingPaymasterConfig.maxGasCost, big.NewInt(12))
	assert.Equal(t, expectedMaxCost, tmp.maxCost(&userOp))

	assertVerifyingPaymasterValidateBizErr(t, ErrVerifyingPaymasterMaxGasCostExceeded, userOp)
}

func TestVerifyingPaymasterValidateCallData(t *testing.T) {
	paymaster := assertNewTestVerifyingPaymaster(t)

	// nil calldata
	assert.Error(t, paymaster.validateCallData(nil))

	// invalid selector
	assert.Error(t, paymaster.validateCallData(hexutil.MustDecode("0xa9059cbb"))) // transfer(address,uint256)

	abi, err := abi.JSON(strings.NewReader(contract.SimpleSmartAccount7702MetaData.ABI))
	assert.NoError(t, err)

	// execute - target not whitelisted
	callData, err := abi.Pack("execute", common.HexToAddress("0x1112"), big.NewInt(10), []byte{1, 2, 3})
	assert.NoError(t, err)
	expectedBizErr := ErrVerifyingPaymasterContractNotWhitelisted.WithData(common.HexToAddress("0x1112"))
	assert.Equal(t, expectedBizErr, paymaster.validateCallData(callData))

	// execute - valid
	callData, err = abi.Pack("execute", common.HexToAddress("0x1111"), big.NewInt(10), []byte{1, 2, 3})
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(callData))

	// executeBatch - empty
	callData, err = abi.Pack("executeBatch", []contract.Execution{})
	assert.NoError(t, err)
	assert.Error(t, paymaster.validateCallData(callData))

	// executeBatch - target not whitelisted
	callData, err = abi.Pack("executeBatch", []contract.Execution{
		{Target: common.HexToAddress("0x1111"), Value: big.NewInt(10), CallData: []byte{1, 2, 3}},
		{Target: common.HexToAddress("0x2223"), Value: big.NewInt(20), CallData: []byte{4, 5, 6}},
	})
	assert.NoError(t, err)
	expectedBizErr = ErrVerifyingPaymasterContractNotWhitelisted.WithData(common.HexToAddress("0x2223"))
	assert.Equal(t, expectedBizErr, paymaster.validateCallData(callData))

	// executeBatch - valid
	callData, err = abi.Pack("executeBatch", []contract.Execution{
		{Target: common.HexToAddress("0x1111"), Value: big.NewInt(10), CallData: []byte{1, 2, 3}},
		{Target: common.HexToAddress("0x2222"), Value: big.NewInt(20), CallData: []byte{4, 5, 6}},
	})
	assert.NoError(t, err)
	assert.NoError(t, paymaster.validateCallData(callData))
}
