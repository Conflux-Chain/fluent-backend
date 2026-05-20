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
}

type Services struct {
	AccountAbstract *AccountAbstract
	GasTank         *GasTankPaymaster
}

func New(config Config) (Services, error) {
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

	aa, err := NewAccountAbstract(client, config.AccountAbstract.DelegatedContract)
	if err != nil {
		return Services{}, errors.WithMessage(err, "Failed to create account abstract service")
	}

	gasTank, err := NewGasTankPaymaster(config.GasTank, client)
	if err != nil {
		return Services{}, errors.WithMessage(err, "Failed to create gas tank paymaster service")
	}

	return Services{
		AccountAbstract: aa,
		GasTank:         gasTank,
	}, nil
}
