package store

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/store"
	"github.com/ethereum/go-ethereum/common"
)

type UserOpStore struct {
	inner *store.Store
}

func NewUserOpStore(store *store.Store) *UserOpStore {
	return &UserOpStore{
		inner: store,
	}
}

func (store *UserOpStore) GetPendingCount(sender common.Address) (int64, error) {
	db := store.inner.DB.Model(&UserOp{}).
		Where("sender = ?", sender.Hex()).
		Where("status = ?", UserOpStatusSigned)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, api.ErrDatabaseCause(err, "Failed to query pending userOp count")
	}

	return count, nil
}

func (store *UserOpStore) Create(userOp *UserOp) error {
	if err := store.inner.DB.Create(userOp).Error; err != nil {
		return api.ErrDatabaseCause(err, "Failed to create user operation")
	}

	return nil
}

func (store *UserOpStore) DeleteExpired(timeout time.Duration) (int64, error) {
	db := store.inner.DB.
		Where("status = ?", UserOpStatusSigned).
		Where("valid_until < ?", time.Now().Add(-timeout)).
		Delete(&UserOp{})

	if err := db.Error; err != nil {
		return 0, api.ErrDatabaseCause(err, "Failed to delete expired user operations")
	}

	return db.RowsAffected, nil
}
