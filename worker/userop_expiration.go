package worker

import (
	"time"

	"github.com/Conflux-Chain/fluent-backend/store"
	"github.com/sirupsen/logrus"
)

// ExpireUserOps periodically deletes expired user operations from the store.
func ExpireUserOps(interval time.Duration, timeout time.Duration, store *store.Store) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if deleted, err := store.UserOp.DeleteExpired(timeout); err != nil {
			logrus.WithError(err).Warn("Failed to delete expired user ops")
		} else if deleted > 0 {
			logrus.WithField("count", deleted).Info("Deleted expired user ops")
		}
	}
}
