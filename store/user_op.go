package store

import (
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm"
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

func (store *UserOpStore) Update(event *contract.VerifyingPaymasterSponsored, tx ...*gorm.DB) (bool, error) {
	db := store.inner.DB
	if len(tx) > 0 {
		db = tx[0]
	}

	var status string
	if event.Success {
		status = UserOpStatusSucceeded
	} else {
		status = UserOpStatusFailed
	}

	// Note, the correct ActualGasCost value comes from the UserOperationEvent of EntryPoint contract instead of
	// the Sponsored event of Paymaster contract. However, the gap between the two values is very small, so we can
	// use the Sponsored event value as an approximation.
	//
	// In the future, we may consider to retrieve the UserOperationEvent of EntryPoint if business requires more accuracy.
	result := db.Model(&UserOp{}).
		Where("hash = ?", hexutil.Encode(event.UserOpHash[:])).
		Updates(UserOp{
			Status:                status,
			ActualGasCost:         event.ActualGasCost.String(),
			ActualUserOpFeePerGas: event.ActualUserOpFeePerGas.String(),
		})

	return result.RowsAffected > 0, result.Error
}
