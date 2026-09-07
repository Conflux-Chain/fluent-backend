package service

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const (
	binancePriceUrlCFXUSDT = "https://api.binance.com/api/v3/ticker/price?symbol=CFXUSDT"
	okxPriceUrlUSDTCNY     = "https://www.okx.com/v3/c2c/otc-ticker/quotedPrice?baseCurrency=USDT&quoteCurrency=CNY&side=sell"
)

type PriceConfig struct {
	USDT []common.Address // e.g. USDT0, USDT, USDC
	CNH  []common.Address // e.g. AxCNH, AxCNH0
}

type ERC20TokenStub struct {
	caller     *contract.ERC20Caller
	decimalExp decimal.Decimal // 10^decimals
}

type PriceOracle struct {
	client *resty.Client

	usdtTokens map[common.Address]ERC20TokenStub
	cnhTokens  map[common.Address]ERC20TokenStub
}

func NewPriceOracle(config PriceConfig, client *web3go.Client) (*PriceOracle, error) {
	if len(config.USDT) == 0 && len(config.CNH) == 0 {
		return nil, errors.New("PriceConfig must have at least one USDT or CNH token")
	}

	oracle := PriceOracle{
		client:     resty.New(),
		usdtTokens: make(map[common.Address]ERC20TokenStub),
		cnhTokens:  make(map[common.Address]ERC20TokenStub),
	}

	// initialize USDT token exponents
	caller, _ := client.ToClientForContract()
	for _, v := range config.USDT {
		if _, ok := oracle.usdtTokens[v]; ok {
			return nil, errors.Errorf("Duplicated USDT token configured %v", v)
		}

		erc20Caller, err := contract.NewERC20Caller(v, caller)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to create ERC20 caller for token %v", v)
		}

		decimals, err := erc20Caller.Decimals(nil)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to get decimals for token %v", v)
		}

		oracle.usdtTokens[v] = ERC20TokenStub{
			caller:     erc20Caller,
			decimalExp: decimal.New(1, int32(decimals)),
		}
	}

	// initialize CNH token exponents
	for _, v := range config.CNH {
		if _, ok := oracle.cnhTokens[v]; ok {
			return nil, errors.Errorf("Duplicated CNH token configured %v", v)
		}

		if _, ok := oracle.usdtTokens[v]; ok {
			return nil, errors.Errorf("Token mis-configured as both USDT and CNH token %v", v)
		}

		erc20Caller, err := contract.NewERC20Caller(v, caller)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to create ERC20 caller for token %v", v)
		}

		decimals, err := erc20Caller.Decimals(nil)
		if err != nil {
			return nil, errors.WithMessagef(err, "Failed to get decimals for token %v", v)
		}

		oracle.cnhTokens[v] = ERC20TokenStub{
			caller:     erc20Caller,
			decimalExp: decimal.New(1, int32(decimals)),
		}
	}

	return &oracle, nil
}

// GetERC20TokenStub returns the ERC20TokenStub and a boolean indicating if the token is supported.
func (oracle *PriceOracle) GetERC20TokenStub(token common.Address) (ERC20TokenStub, bool) {
	if stub, ok := oracle.usdtTokens[token]; ok {
		return stub, true
	}

	if stub, ok := oracle.cnhTokens[token]; ok {
		return stub, true
	}

	return ERC20TokenStub{}, false
}

// GetETHPrice returns the price of ETH/token.
func (oracle *PriceOracle) GetETHPrice(quoteToken common.Address) (*big.Int, error) {
	// USDT
	if stub, ok := oracle.usdtTokens[quoteToken]; ok {
		usdtPerCfx, err := oracle.getBinancePrice(binancePriceUrlCFXUSDT)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to get binance CFX/USDT price")
		}

		return usdtPerCfx.Mul(stub.decimalExp).BigInt(), nil
	}

	// CNH
	if stub, ok := oracle.cnhTokens[quoteToken]; ok {
		usdtPerCfx, err := oracle.getBinancePrice(binancePriceUrlCFXUSDT)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to get binance CFX/USDT price")
		}

		cnyPerUsdt, err := oracle.getOkxPrice(okxPriceUrlUSDTCNY)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to get OKX USDT/CNY price")
		}

		return cnyPerUsdt.Mul(usdtPerCfx).Mul(stub.decimalExp).BigInt(), nil
	}

	// Unsupported
	return nil, api.ErrValidationStrf("Unsupported token %v", quoteToken)
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
		Data          []struct {
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

	if len(result.Data) != 1 {
		return decimal.Zero, ErrRPCError.WithData(fmt.Sprintf("Unexpected OKX data: %+v", result))
	}

	price, err := decimal.NewFromString(result.Data[0].Price)
	if err != nil {
		return decimal.Zero, errors.WithMessage(err, "Failed to parse OKX price")
	}

	return price, nil
}
