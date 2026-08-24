package worker

import (
	"time"

	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
)

type Config struct {
	UserOp struct {
		EventScan UserOpEventScanConfig

		Expiration struct {
			Timeout  time.Duration `default:"24h"` // user op expiration time, used to clean up expired user ops
			Interval time.Duration `default:"10m"` // user op expiration interval, used to clean up expired user ops
		}
	}
}

// Start starts all background workers in separate goroutines. It returns an error if any worker fails to initialize.
func Start(config Config, client *web3go.Client, store *store.Store) error {
	userOpEventScanner, err := NewUserOpEventScanner(config.UserOp.EventScan, client, store)
	if err != nil {
		return errors.WithMessage(err, "Failed to create user op event scanner")
	}

	go userOpEventScanner.Work()

	go ExpireUserOps(config.UserOp.Expiration.Interval, config.UserOp.Expiration.Timeout, store)

	return nil
}
