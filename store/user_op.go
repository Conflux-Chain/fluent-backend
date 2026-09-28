package store

import (
	"encoding/json"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/blockchain/contract/account"
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

func (store *UserOpStore) GetCountByBlockTimestamp(sender common.Address, since time.Time) (int64, error) {
	return store.getCount("sender = ? AND block_time >= ?", sender.Hex(), since)
}

// Create records the finalized event even if the raw UserOperation is unavailable.
func (store *UserOpStore) Create(userOp *contract.PackedUserOperation, event *account.EntryPointUserOperationEvent, blockTime time.Time, tx ...*gorm.DB) error {
	db := store.inner.DB
	if len(tx) > 0 {
		db = tx[0]
	}

	var rawUserOp []byte
	if userOp != nil {
		var err error
		if rawUserOp, err = json.Marshal(convertPackedUserOp(userOp)); err != nil {
			return errors.WithMessage(err, "Failed to JSON marshal user op")
		}
	}

	entity := UserOp{
		Hash:    common.Hash(event.UserOpHash).Hex(),
		Sender:  event.Sender.Hex(),
		Nonce:   hexutil.EncodeBig(event.Nonce),
		Success: event.Success,

		ActualGasCost: decimal.NewFromBigInt(event.ActualGasCost, 0),
		ActualGasUsed: event.ActualGasUsed.Uint64(),
		BlockTime:     blockTime,

		RawUserOp: string(rawUserOp),
	}

	if err := db.Create(&entity).Error; err != nil {
		return api.ErrDatabaseCause(err, "Failed to create user operation")
	}

	return nil
}
