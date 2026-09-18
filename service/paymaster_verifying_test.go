package service

import (
	"math/big"
	"testing"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go/signers"
	"github.com/stretchr/testify/assert"
)

var testVerifyingPaymasterConfig = VerifyingPaymasterConfig{
	PaymasterConfig: PaymasterConfig{
		Address:          common.HexToAddress("0x6666"),
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
