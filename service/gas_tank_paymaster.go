package service

import (
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/interfaces"
	"github.com/pkg/errors"
)

var (
	gasTankPaymasterModeRefund = byte(0)
	gasTankPaymasterModeCredit = byte(1)

	dummySignature = slices.Repeat([]byte{0x1b}, 65)
)

type GasTankPaymasterConfig struct {
	Address common.Address

	SignatureTimeout time.Duration `default:"5m"`
}

type GasTankPaymaster struct {
	config GasTankPaymasterConfig
	caller *contract.GasTankPaymasterCaller
	signer interfaces.Signer
}

func NewGasTankPaymaster(config GasTankPaymasterConfig, client *web3go.Client) (*GasTankPaymaster, error) {
	if config.Address == (common.Address{}) {
		return nil, errors.New("GasTankPaymaster address is required")
	}

	// get the default signer
	sm, err := client.GetSignerManager()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get signer manager from RPC client")
	}

	signers := sm.List()
	if len(signers) == 0 {
		return nil, errors.New("No signer found")
	}

	caller, _ := client.ToClientForContract()

	gasTankCaller, err := contract.NewGasTankPaymasterCaller(config.Address, caller)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to create GasTankPaymasterCaller")
	}

	return &GasTankPaymaster{
		config: config,
		caller: gasTankCaller,
		signer: signers[0],
	}, nil
}

func (paymaster *GasTankPaymaster) PrepareDeposit(token common.Address, amount *big.Int) []byte {
	data := GasTankPaymasterData{
		Address:      paymaster.config.Address,
		Mode:         gasTankPaymasterModeCredit,
		Token:        token,
		MaxTokenCost: amount,
		Signature:    dummySignature,
	}

	return data.encode()
}

func (paymaster *GasTankPaymaster) Prepare(sender, token common.Address) ([]byte, error) {
	balance, err := paymaster.caller.BalanceOf(nil, sender, token)
	if err != nil {
		return nil, ErrRPCError.WithData(errors.WithMessage(err, "Failed to retrieve sender token balance"))
	}

	if balance.Sign() == 0 {
		return nil, ErrGasTankInsufficientBalance
	}

	data := GasTankPaymasterData{
		Address:      paymaster.config.Address,
		Mode:         gasTankPaymasterModeRefund,
		Token:        token,
		MaxTokenCost: balance,
		Signature:    dummySignature,
	}

	return data.encode(), nil
}

func (paymaster *GasTankPaymaster) Sign(userOp contract.PackedUserOperation) ([]byte, error) {
	// EOA should be upgraded to smart account via 7702 transaction
	if len(userOp.InitCode) > 0 {
		return nil, api.ErrValidationStr("InitCode is not supported")
	}

	// Parse and validate paymaster data. Note, the paymaster contract will validate the mode and token.
	paymasterData, err := parseGasTankPaymasterData(userOp.PaymasterAndData)
	if err != nil {
		return nil, err
	}

	if paymasterData.Address != paymaster.config.Address {
		return nil, api.ErrValidationStr("Invalid paymaster address")
	}

	// calculate the max token cost
	maxTokenCost := paymaster.calculateMaxTokenCost(&userOp, &paymasterData)

	// validate sender balance in refund mode
	if paymasterData.Mode == gasTankPaymasterModeRefund {
		if err := paymaster.validateSenderBalance(userOp.Sender, paymasterData.Token, maxTokenCost); err != nil {
			return nil, err
		}
	}

	// validate deposit amount in credit mode
	if paymasterData.Mode == gasTankPaymasterModeCredit {
		if err := paymaster.validateDepositAmount(userOp.CallData, maxTokenCost); err != nil {
			return nil, err
		}
	}

	// re-construct the paymaster data in user op
	paymasterData.MaxTokenCost = maxTokenCost
	paymasterData.ValidAfter = 0
	paymasterData.ValidUntil = uint64(time.Now().Add(paymaster.config.SignatureTimeout).Unix())
	paymasterData.Signature = dummySignature
	userOp.PaymasterAndData = paymasterData.encode()

	// compute the paymaster signature
	hash, err := paymaster.caller.GetPaymasterHash(nil, userOp)
	if err != nil {
		return nil, ErrRPCError.WithData(errors.WithMessage(err, "Failed to retrieve paymaster hash from blockchain"))
	}

	if paymasterData.Signature, err = paymaster.signer.SignHash(hash); err != nil {
		return nil, errors.WithMessage(err, "Failed to sign paymaster hash")
	}

	return paymasterData.encode(), nil
}

func (paymaster *GasTankPaymaster) calculateMaxTokenCost(userOp *contract.PackedUserOperation, paymasterData *GasTankPaymasterData) *big.Int {
	gasLimit := new(big.Int).Add(paymasterData.ValidationGas, paymasterData.PostOpGas)
	gasLimit.Add(gasLimit, userOp.PreVerificationGas)
	gasLimit.Add(gasLimit, new(big.Int).SetBytes(userOp.AccountGasLimits[0:16]))
	gasLimit.Add(gasLimit, new(big.Int).SetBytes(userOp.AccountGasLimits[16:32]))

	maxFeePerGas := new(big.Int).SetBytes(userOp.GasFees[16:32])
	maxGasCost := new(big.Int).Mul(maxFeePerGas, gasLimit)

	// TODO 集成 price oracle，根据价格计算 maxTokenCost, 这里先简单假设 token 价格是 0.06 token/CFX
	return new(big.Int).Div(new(big.Int).Mul(maxGasCost, big.NewInt(6)), big.NewInt(100))
}

func (paymaster *GasTankPaymaster) validateSenderBalance(sender, token common.Address, maxTokenCost *big.Int) error {
	balance, err := paymaster.caller.BalanceOf(nil, sender, token)
	if err != nil {
		return ErrRPCError.WithData(errors.WithMessage(err, "Failed to retrieve sender token balance"))
	}

	if balance.Cmp(maxTokenCost) < 0 {
		return ErrGasTankInsufficientBalance.WithData(fmt.Sprintf("maxTokenCost = %v, balance = %v", maxTokenCost, balance))
	}

	return nil
}

func (paymaster *GasTankPaymaster) validateDepositAmount(calldata []byte, maxTokenCost *big.Int) error {
	if len(calldata) != 580 {
		return api.ErrValidationStrf("Invalid callData length for approve + deposit, expected 580, got %d", len(calldata))
	}

	// executeBatch(Call[]): approve(address,uint256) + depositToken(address,uint256)
	const offsetApproveAmount = 296
	const offsetDepositAmount = 520
	approveAmount := new(big.Int).SetBytes(calldata[offsetApproveAmount : offsetApproveAmount+32])
	depositAmount := new(big.Int).SetBytes(calldata[offsetDepositAmount : offsetDepositAmount+32])

	if approveAmount.Cmp(depositAmount) != 0 {
		return api.ErrValidationStr("Approve amount must be equal to deposit amount")
	}

	if depositAmount.Cmp(maxTokenCost) < 0 {
		return ErrGasTankInsufficientBalance.WithData(fmt.Sprintf("maxTokenCost = %v, depositAmount = %v", maxTokenCost, depositAmount))
	}

	return nil
}

type GasTankPaymasterData struct {
	Address       common.Address // 20
	ValidationGas *big.Int       // 16
	PostOpGas     *big.Int       // 16
	Mode          byte           // 1
	Token         common.Address // 20
	MaxTokenCost  *big.Int       // 32
	ValidAfter    uint64         // 6
	ValidUntil    uint64         // 6
	Signature     []byte         // 65
}

func parseGasTankPaymasterData(data []byte) (GasTankPaymasterData, error) {
	if len(data) != 182 {
		return GasTankPaymasterData{}, api.ErrValidationStrf("Invalid paymasterAndData length, expected 182, got %d", len(data))
	}

	return GasTankPaymasterData{
		Address:       common.BytesToAddress(data[0:20]),
		ValidationGas: new(big.Int).SetBytes(data[20:36]),
		PostOpGas:     new(big.Int).SetBytes(data[36:52]),
		Mode:          data[52],
		Token:         common.BytesToAddress(data[53:73]),
		MaxTokenCost:  new(big.Int).SetBytes(data[73:105]),
		ValidAfter:    new(big.Int).SetBytes(data[105:111]).Uint64(),
		ValidUntil:    new(big.Int).SetBytes(data[111:117]).Uint64(),
		Signature:     data[117:182],
	}, nil
}

func (data *GasTankPaymasterData) encode() []byte {
	var buf [182]byte

	copy(buf[0:20], data.Address.Bytes())
	if data.ValidationGas != nil {
		data.ValidationGas.FillBytes(buf[20:36])
	}
	if data.PostOpGas != nil {
		data.PostOpGas.FillBytes(buf[36:52])
	}
	buf[52] = data.Mode
	copy(buf[53:73], data.Token.Bytes())
	if data.MaxTokenCost != nil {
		data.MaxTokenCost.FillBytes(buf[73:105])
	}
	new(big.Int).SetUint64(data.ValidAfter).FillBytes(buf[105:111])
	new(big.Int).SetUint64(data.ValidUntil).FillBytes(buf[111:117])
	copy(buf[117:182], data.Signature)

	return buf[:]
}
