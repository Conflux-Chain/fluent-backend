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

const (
	binancePriceUrlCFXUSDT = "https://api.binance.com/api/v3/ticker/price?symbol=CFXUSDT"
	okxPriceUrlUSDTCNY     = "https://www.okx.com/v3/c2c/otc-ticker/quotedPrice?baseCurrency=USDT&quoteCurrency=CNY&side=sell"
)

type PriceOracle struct {
	client *resty.Client

	allowedTokenSymbols map[common.Address]string          // token address => token symbol
	allowedTokenExps    map[common.Address]decimal.Decimal // token address => 10^decimals
}

func NewPriceOracle(tokens map[common.Address]ERC20TokenStub) *PriceOracle {
	allowedTokenSymbols := make(map[common.Address]string)
	allowedTokenExps := make(map[common.Address]decimal.Decimal)
	for token, stub := range tokens {
		allowedTokenSymbols[token] = stub.Symbol
		allowedTokenExps[token] = decimal.New(1, int32(stub.Decimals))
	}

	return &PriceOracle{
		client:              resty.New(),
		allowedTokenSymbols: allowedTokenSymbols,
		allowedTokenExps:    allowedTokenExps,
	}
}

// GetETHPrice returns the price of ETH/token.
func (oracle *PriceOracle) GetETHPrice(quoteToken common.Address) (*big.Int, error) {
	symbol, ok := oracle.allowedTokenSymbols[quoteToken]
	if !ok {
		return nil, api.ErrValidationStr("Unsupported token")
	}

	exp := oracle.allowedTokenExps[quoteToken]

	usdtPerCfx, err := oracle.getBinancePrice(binancePriceUrlCFXUSDT)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get binance CFX/USDT price")
	}

	// CFX/USDT, CFX/USDT0 or CFX/USDC
	if symbol == "USDT" || symbol == "USDT0" || symbol == "USDC" {
		return usdtPerCfx.Mul(exp).BigInt(), nil
	}

	// currently we do not support other tokens except AxCNH0 or AxCNH
	if symbol != "AxCNH" && symbol != "AxCNH0" {
		return nil, api.ErrValidationStrf("Unsupported token symbol: %v", symbol)
	}

	cnyPerUsdt, err := oracle.getOkxPrice(okxPriceUrlUSDTCNY)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get OKX USDT/CNY price")
	}

	return cnyPerUsdt.Mul(usdtPerCfx).Mul(exp).BigInt(), nil
}

func (oracle *PriceOracle) getBinancePrice(url string) (decimal.Decimal, error) {
	var result struct {
		Symbol string
		Price  string
	}

	resp, err := oracle.client.R().SetResult(&result).Get(url)
	if err != nil {
		return decimal.Zero, NewRPCError(err, "Failed to fetch price from Binance")
	}

	if resp.IsError() {
		return decimal.Zero, ErrRPCError.WithData(fmt.Sprintf("Invalid binance response status: %v", resp.Status()))
	}

	price, err := decimal.NewFromString(result.Price)
	if err != nil {
		return decimal.Zero, errors.WithMessage(err, "Failed to parse binance price")
	}

	return price, nil
}

func (oracle *PriceOracle) getOkxPrice(url string) (decimal.Decimal, error) {
	var result struct {
		Code          int    `json:"code"`
		Message       string `json:"msg"`
		DetailMessage string `json:"detailMsg"`
		ErrorCode     string `json:"error_code"`
		ErrorMessage  string `json:"error_message"`
		Data          struct {
			BestOption  bool   `json:"bestOption"`
			DepositName string `json:"depositName"`
			Payment     string `json:"payment"`
			Price       string `json:"price"`
		} `json:"data"`
	}

	resp, err := oracle.client.R().SetResult(&result).Get(url)
	if err != nil {
		return decimal.Zero, NewRPCError(err, "Failed to fetch price from OKX")
	}

	if resp.IsError() {
		return decimal.Zero, ErrRPCError.WithData(fmt.Sprintf("Invalid OKX response status: %v", resp.Status()))
	}

	if result.Code != 0 || result.ErrorCode != "0" {
		return decimal.Zero, ErrRPCError.WithData(fmt.Sprintf("OKX API error: %+v", result))
	}

	price, err := decimal.NewFromString(result.Data.Price)
	if err != nil {
		return decimal.Zero, errors.WithMessage(err, "Failed to parse OKX price")
	}

	return price, nil
}
