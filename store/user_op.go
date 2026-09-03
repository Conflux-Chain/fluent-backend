package store

import (
	"encoding/json"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
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

func (store *UserOpStore) getCount(whereClause string, args ...any) (int64, error) {
	var count int64

	db := store.inner.DB.Model(&UserOp{}).Where(whereClause, args...)

	if err := db.Count(&count).Error; err != nil {
		return 0, api.ErrDatabaseCause(err, "Failed to query user op count by filter")
	}

	return count, nil
}

func (store *UserOpStore) GetCountByValidUntil(sender common.Address, since time.Time) (int64, error) {
	return store.getCount("sender = ? AND valid_until >= ?", sender.Hex(), since)
}

func (store *UserOpStore) GetPendingCount(sender common.Address) (int64, error) {
	return store.getCount("sender = ? AND status = ?", sender.Hex(), UserOpStatusSigned)
}

func (store *UserOpStore) GetPendingCountByIP(ip string) (int64, error) {
	return store.getCount("ip_address = ? AND status = ?", ip, UserOpStatusSigned)
}

func (store *UserOpStore) Create(userOp *contract.PackedUserOperation, hash string, validUntil time.Time, ip string) error {
	rawUserOp, err := json.Marshal(convertPackedUserOp(userOp))
	if err != nil {
		return errors.WithMessage(err, "Failed to JSON marshal user op")
	}

	entity := UserOp{
		Hash:       hash,
		IPAddress:  ip,
		Sender:     userOp.Sender.Hex(),
		Nonce:      hexutil.EncodeBig(userOp.Nonce),
		Status:     UserOpStatusSigned,
		ValidUntil: validUntil,

		ActualGasCost:         decimal.Zero,
		ActualUserOpFeePerGas: decimal.Zero,

		RawUserOp: string(rawUserOp),
	}

	if err := store.inner.DB.Create(&entity).Error; err != nil {
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

func (store *UserOpStore) Update(event *contract.VerifyingPaymasterSponsored, blockTimestamp uint64, tx ...*gorm.DB) (bool, error) {
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
			ActualGasCost:         decimal.NewFromBigInt(event.ActualGasCost, 0),
			ActualUserOpFeePerGas: decimal.NewFromBigInt(event.ActualUserOpFeePerGas, 0),
			BlockTimestamp:        blockTimestamp,
		})

	return result.RowsAffected > 0, result.Error
}
