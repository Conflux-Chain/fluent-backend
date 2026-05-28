package service

import (
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/signers"
	"github.com/pkg/errors"
)

type Config struct {
	RPC struct {
		URL            string
		RequestTimeout time.Duration `default:"3s"`
		PrivateKey     string
		LogEnabled     bool
	}

	AccountAbstract struct {
		DelegatedContract common.Address
	}

	GasTank GasTankPaymasterConfig

	TokenPay TokenPayConfig
}

type Services struct {
	AccountAbstract *AccountAbstract
	PriceOracle     *PriceOracle
	GasTank         *GasTankPaymaster
	TokenPay        *TokenPay
}

func New(config Config) (Services, error) {
	// RPC client
	if len(config.RPC.URL) == 0 {
		return Services{}, errors.New("RPC URL not specified")
	}

	var opt web3go.ClientOption
	opt.RequestTimeout = config.RPC.RequestTimeout
	if len(config.RPC.PrivateKey) > 0 {
		opt.SignerManager = signers.MustNewSignerManagerByPrivateKeyStrings([]string{config.RPC.PrivateKey})
	}

	if config.RPC.LogEnabled {
		opt.Logger = os.Stdout
	}

	client, err := web3go.NewClientWithOption(config.RPC.URL, opt)
	if err != nil {
		return Services{}, errors.WithMessage(err, "Failed to create RPC client")
	}

	// normalize config
	if err = config.TokenPay.Normalize(client); err != nil {
		return Services{}, errors.WithMessage(err, "Failed to normalize token-pay config")
	}

	// services
	txSender, err := NewTxSender(client)
	if err != nil {
		return Services{}, errors.WithMessage(err, "Failed to create transaction sender")
	}

	aa := NewAccountAbstract(txSender, config.AccountAbstract.DelegatedContract)

	gasTank, err := NewGasTankPaymaster(config.GasTank, client)
	if err != nil {
		return Services{}, errors.WithMessage(err, "Failed to create gas tank paymaster service")
	}

	priceOracle := NewPriceOracle(config.TokenPay.normalizedTokens)
	tokenPay := NewTokenPay(config.TokenPay, txSender, priceOracle)

	return Services{
		AccountAbstract: aa,
		PriceOracle:     priceOracle,
		GasTank:         gasTank,
		TokenPay:        tokenPay,
	}, nil
}
