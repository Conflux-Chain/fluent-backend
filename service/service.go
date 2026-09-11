package service

import (
	"os"
	"time"

	"github.com/Conflux-Chain/fluent-backend/store"
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

	Price PriceConfig

	AccountAbstract struct {
		DelegatedContract common.Address
	}

	VerifyingPaymaster VerifyingPaymasterConfig
	GasTank            GasTankPaymasterConfig

	TokenPay TokenPayConfig
}

type Services struct {
	config Config

	AccountAbstract    *AccountAbstract // may be nil if the delegated contract address is not specified
	PriceOracle        *PriceOracle
	VerifyingPaymaster *VerifyingPaymaster
	GasTank            *GasTankPaymaster // may be nil if the gas tank paymaster is not specified
	TokenPay           *TokenPay

	client *web3go.Client
}

func New(config Config, store *store.Store) (Services, error) {
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

	// services
	txSender, err := NewTxSender(client)
	if err != nil {
		return Services{}, errors.WithMessage(err, "Failed to create transaction sender")
	}

	services := Services{
		config: config,
		client: client,
	}

	// create AccountAbstract service if the delegated contract address is specified
	if config.AccountAbstract.DelegatedContract != (common.Address{}) {
		services.AccountAbstract = NewAccountAbstract(txSender, config.AccountAbstract.DelegatedContract)
	}

	// create VerifyingPaymaster service if paymaster address and whitelist are specified
	if config.VerifyingPaymaster.Address != (common.Address{}) &&
		len(config.VerifyingPaymaster.SmartAccountWhitelist) > 0 &&
		len(config.VerifyingPaymaster.ContractWhitelist) > 0 {
		if services.VerifyingPaymaster, err = NewVerifyingPaymaster(config.VerifyingPaymaster, client, store); err != nil {
			return Services{}, errors.WithMessage(err, "Failed to create verifying paymaster service")
		}
	}

	// create price oracle service if at least one USDT configured
	if len(config.Price.USDT) > 0 {
		if services.PriceOracle, err = NewPriceOracle(config.Price, client); err != nil {
			return Services{}, errors.WithMessage(err, "Failed to create price oracle service")
		}

		// create gas tank paymaster service if the paymaster address is specified
		if config.GasTank.Address != (common.Address{}) {
			if services.GasTank, err = NewGasTankPaymaster(config.GasTank, services.PriceOracle, client); err != nil {
				return Services{}, errors.WithMessage(err, "Failed to create gas tank paymaster service")
			}
		}

		// create token pay service if the recipient is specified
		if config.TokenPay.Recipient != (common.Address{}) {
			if services.TokenPay, err = NewTokenPay(config.TokenPay, txSender, services.PriceOracle); err != nil {
				return Services{}, errors.WithMessage(err, "Failed to create token pay service")
			}
		}
	}

	return services, nil
}

func (s Services) Config() struct {
	Price    PriceConfig
	TokenPay TokenPayConfig
} {
	return struct {
		Price    PriceConfig
		TokenPay TokenPayConfig
	}{
		Price:    s.config.Price,
		TokenPay: s.config.TokenPay,
	}
}

func (s Services) Client() *web3go.Client {
	return s.client
}
