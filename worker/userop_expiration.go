package worker

import (
	"time"

	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/sirupsen/logrus"
)

type UserOpExpirationConfig struct {
	Timeout  time.Duration `default:"24h"` // user op expiration time, used to clean up expired user ops
	Interval time.Duration `default:"10m"` // user op expiration interval, used to clean up expired user ops
}

// ExpireUserOps periodically deletes expired user operations from the store.
//
// NOTE in the future, we could punish the malicious clients who send too many expired user ops of the same
// IP address. For example, blacklist the client IP for a long time.
func ExpireUserOps(config UserOpExpirationConfig, store *store.Store) {
	if config.Timeout <= 0 {
		config.Timeout = 24 * time.Hour
	}

	if config.Interval <= 0 {
		config.Interval = 10 * time.Minute
	}

	ticker := time.NewTicker(config.Interval)
	defer ticker.Stop()

	for range ticker.C {
		logrus.Debug("Begin to delete expired user ops")

		if deleted, err := store.UserOp.DeleteExpired(config.Timeout); err != nil {
			logrus.WithError(err).Warn("Failed to delete expired user ops")
		} else if deleted > 0 {
			logrus.WithField("count", deleted).Info("Succeeded to delete expired user ops")
		} else {
			logrus.Debug("No expired user ops found to delete")
		}
	}
}
