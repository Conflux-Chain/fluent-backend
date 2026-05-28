package service

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const binancePriceUrlCFXUSDT = "https://api.binance.com/api/v3/ticker/price?symbol=CFXUSDT"

type PriceOracle struct {
	client *resty.Client

	allowedTokens map[common.Address]decimal.Decimal // token address => 10^decimals
}

func NewPriceOracle(tokens map[common.Address]ERC20TokenStub) *PriceOracle {
	allowedTokens := make(map[common.Address]decimal.Decimal)
	for token, stub := range tokens {
		allowedTokens[token] = decimal.New(1, int32(stub.Decimals))
	}

	return &PriceOracle{
		client:        resty.New(),
		allowedTokens: allowedTokens,
	}
}

// GetETHPrice returns the price of ETH/token.
func (oracle *PriceOracle) GetETHPrice(quoteToken common.Address) (*big.Int, error) {
	exp, ok := oracle.allowedTokens[quoteToken]
	if !ok {
		return nil, api.ErrValidationStr("Unsupported token")
	}

	var result struct {
		Symbol string
		Price  string
	}

	resp, err := oracle.client.R().SetResult(&result).Get(binancePriceUrlCFXUSDT)
	if err != nil {
		return nil, NewRPCError(err, "Failed to fetch price from Binance")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Invalid binance response status: %v", resp.Status())
	}

	price, err := decimal.NewFromString(result.Price)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse binance price")
	}

	return price.Mul(exp).BigInt(), nil
}
