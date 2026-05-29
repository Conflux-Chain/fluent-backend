package service

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
)

type ERC20TokenStub struct {
	Caller   *contract.ERC20Caller
	Name     string
	Symbol   string
	Decimals uint8
}

type TokenPayConfig struct {
	Recipient           common.Address
	abiEncodedRecipient string

	Tokens           map[string]common.Address // name => token address
	normalizedTokens map[common.Address]ERC20TokenStub

	MinGasFeeRatio       uint64 `default:"120"` // 120% of base fee, e.g. 20 Gdrip * 120% = 24 Gdrip
	minGasFeeRatioBig    *big.Int
	MinGasTipRatio       uint64 `default:"20"` // 20% of base fee, e.g. 20 Gdrip * 20% = 4 Gdrip
	minGasTipRatioBig    *big.Int
	MaxGasCost           uint64 `default:"100000000000000000"` // 0.1 CFX
	maxGasCostBig        *big.Int
	MinSponsorBalance    uint64 `default:"1000000000000000000"` // 1 CFX
	minSponsorBalanceBig *big.Int

	CheckReceiptInterval        time.Duration `default:"1s"`
	CheckFundingReceiptInterval time.Duration `default:"30ms"`
}

func (config *TokenPayConfig) Normalize(client *web3go.Client) error {
	// validate
	if config.Recipient == (common.Address{}) {
		return errors.New("Recipient not specified")
	}

	if len(config.Tokens) == 0 {
		return errors.New("Tokens not specified")
	}

	// normalize recipient
	var buf [32]byte
	copy(buf[12:], config.Recipient.Bytes())
	config.abiEncodedRecipient = hexutil.Encode(buf[:])

	// normalize tokens
	caller, _ := client.ToClientForContract()
	config.normalizedTokens = make(map[common.Address]ERC20TokenStub)
	for _, tokenAddr := range config.Tokens {
		erc20Caller, err := contract.NewERC20Caller(tokenAddr, caller)
		if err != nil {
			return errors.WithMessage(err, "Failed to create ERC20 caller")
		}

		name, err := erc20Caller.Name(nil)
		if err != nil {
			return errors.WithMessage(err, "Failed to get token name")
		}

		symbol, err := erc20Caller.Symbol(nil)
		if err != nil {
			return errors.WithMessage(err, "Failed to get token symbol")
		}

		decimals, err := erc20Caller.Decimals(nil)
		if err != nil {
			return errors.WithMessage(err, "Failed to get token decimals")
		}

		config.normalizedTokens[tokenAddr] = ERC20TokenStub{
			Caller:   erc20Caller,
			Name:     name,
			Symbol:   symbol,
			Decimals: decimals,
		}
	}

	// normalize numeric configs
	config.minGasFeeRatioBig = new(big.Int).SetUint64(config.MinGasFeeRatio)
	config.minGasTipRatioBig = new(big.Int).SetUint64(config.MinGasTipRatio)
	config.maxGasCostBig = new(big.Int).SetUint64(config.MaxGasCost)
	config.minSponsorBalanceBig = new(big.Int).SetUint64(config.MinSponsorBalance)

	return nil
}
