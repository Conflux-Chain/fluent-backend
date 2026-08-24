package worker

import (
	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
)

type Config struct {
	UserOp struct {
		EventScan  UserOpEventScanConfig
		Expiration UserOpExpirationConfig
	}
}

// Start starts all background workers in separate goroutines. It returns an error if any worker fails to initialize.
func Start(config Config, client *web3go.Client, store *store.Store) error {
	if config.UserOp.EventScan.Contract != (common.Address{}) {
		userOpEventScanner, err := NewUserOpEventScanner(config.UserOp.EventScan, client, store)
		if err != nil {
			return errors.WithMessage(err, "Failed to create user op event scanner")
		}

		go userOpEventScanner.Work()
	}

	go ExpireUserOps(config.UserOp.Expiration, store)

	return nil
}
