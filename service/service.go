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
		PrivateKey     string
		RequestTimeout time.Duration `default:"3s"`
		LogEnabled     bool
	}

	AccountAbstract struct {
		DelegatedContract common.Address
	}

	GasTank GasTankPaymasterConfig

	TokenPay TokenPayConfig
}

type Services struct {
	AccountAbstract *AccountAbstract // may be nil if the delegated contract address is not specified
	PriceOracle     *PriceOracle
	GasTank         *GasTankPaymaster // may be nil if the gas tank paymaster is not specified
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

	priceOracle := NewPriceOracle(config.TokenPay.normalizedTokens)

	services := Services{
		PriceOracle: priceOracle,
		TokenPay:    NewTokenPay(config.TokenPay, txSender, priceOracle),
	}

	// AccountAbstract service is optional, only create it if the delegated contract address is specified
	if config.AccountAbstract.DelegatedContract != (common.Address{}) {
		services.AccountAbstract = NewAccountAbstract(txSender, config.AccountAbstract.DelegatedContract)
	}

	// GasTankPaymaster service is optional, only create it if the gas tank paymaster address is specified
	if config.GasTank.Address != (common.Address{}) {
		if services.GasTank, err = NewGasTankPaymaster(config.GasTank, priceOracle, client); err != nil {
			return Services{}, errors.WithMessage(err, "Failed to create gas tank paymaster service")
		}
	}

	return services, nil
}
