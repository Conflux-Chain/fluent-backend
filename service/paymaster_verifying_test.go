package service

import (
	"math/big"
	"testing"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go/signers"
	"github.com/stretchr/testify/require"
)

var testVerifyingPaymasterConfig = VerifyingPaymasterConfig{
	PaymasterConfig: PaymasterConfig{
		Address:          common.HexToAddress("0x7777"),
		SignatureTimeout: time.Minute * 5,
	},
	smartAccountMap: map[common.Address]bool{
		common.HexToAddress("0x9999"): true,
	},
	contractMap: map[common.Address]bool{
		common.HexToAddress("0x1111"): true,
		common.HexToAddress("0x2222"): true,
	},
	maxGasCostBig: big.NewInt(100000000000000000),
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
		},
	}
}

func newTestUserOp(paymaster *VerifyingPaymaster) contract.PackedUserOperation {
	userOp := contract.PackedUserOperation{
		Sender:             common.HexToAddress("0x0001"),
		Nonce:              big.NewInt(0),
		InitCode:           []byte{},
		CallData:           []byte{}, // see TestVerifyingPaymasterValidateCallData
		PreVerificationGas: big.NewInt(1),
		PaymasterAndData:   paymaster.inner.generateStub(common.HexToAddress("0x9999")),
	}

	userOp.SetAccountGasLimits(big.NewInt(2), big.NewInt(3))
	userOp.SetPaymasterGasLimits(big.NewInt(4), big.NewInt(5))
	userOp.SetGasFees(big.NewInt(0), big.NewInt(10))

	return userOp
}

func TestVerifyingPaymasterValidateStaticInvalid(t *testing.T) {
	paymaster := newTestVerifyingPaymaster()

	// invalid calldata as a sentinel error
	userOp := newTestUserOp(paymaster)
	require.Equal(t, api.ErrValidationStr("Invalid callData, too short to parse function selector"), paymaster.validate(&userOp))

	// invalid sender
	userOp = newTestUserOp(paymaster)
	userOp.Sender = common.Address{}
	require.Equal(t, api.ErrValidationStr("Invalid sender address"), paymaster.validate(&userOp))

	// invalid paymaster data length
	userOp = newTestUserOp(paymaster)
	userOp.PaymasterAndData = append(userOp.PaymasterAndData, []byte{1}...)
	require.Equal(t, api.ErrValidationStr("Invalid paymasterAndData length"), paymaster.validate(&userOp))

	// delegation not in whitelist
	userOp = newTestUserOp(paymaster)
	wrongDelegation := common.HexToAddress("0x9998")
	userOp.PaymasterAndData = paymaster.inner.generateStub(wrongDelegation)
	require.Equal(t, ErrVerifyingPaymasterInvalidSmartAccount.WithData(wrongDelegation), paymaster.validate(&userOp))

	// max gas cost exceeded
	userOp = newTestUserOp(paymaster)
	userOp.PreVerificationGas = testVerifyingPaymasterConfig.maxGasCostBig
	require.Equal(t, ErrVerifyingPaymasterMaxGasCostExceeded.Code, paymaster.validate(&userOp).(*api.BusinessError).Code)
}

func TestVerifyingPaymasterValidateCallData(t *testing.T) {
	paymaster := newTestVerifyingPaymaster()

	// nil calldata
	require.Error(t, paymaster.validateCallData(nil))

	// invalid selector
	require.Error(t, paymaster.validateCallData(hexutil.MustDecode("0xa9059cbb"))) // transfer(address,uint256)

	// helper function
	validate := func(method string, args ...any) error {
		callData, err := contract.SmartAccountABI.Pack(method, args...)
		if err != nil {
			panic(err)
		}

		return paymaster.validateCallData(callData)
	}

	// execute - target not whitelisted
	expectedBizErr := ErrVerifyingPaymasterContractNotWhitelisted.WithData(common.HexToAddress("0x1112"))
	require.Equal(t, expectedBizErr, validate("execute", common.HexToAddress("0x1112"), big.NewInt(0), []byte{1, 2, 3}))

	// execute - valid
	require.NoError(t, validate("execute", common.HexToAddress("0x1111"), big.NewInt(0), []byte{1, 2, 3}))

	// executeBatch - empty
	require.Error(t, validate("executeBatch", []contract.Execution{}))

	// executeBatch - target not whitelisted
	require.Equal(t, expectedBizErr, validate("executeBatch", []contract.Execution{
		{Target: common.HexToAddress("0x1112"), Value: big.NewInt(10), CallData: []byte{1, 2, 3}},
		{Target: common.HexToAddress("0x2222"), Value: big.NewInt(20), CallData: []byte{4, 5, 6}},
	}))

	// executeBatch - valid
	require.NoError(t, validate("executeBatch", []contract.Execution{
		{Target: common.HexToAddress("0x1111"), Value: big.NewInt(10), CallData: []byte{1, 2, 3}},
		{Target: common.HexToAddress("0x2222"), Value: big.NewInt(20), CallData: []byte{4, 5, 6}},
	}))
}
