package store

import (
	"math/big"
	"testing"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func (store *UserOpStore) assertGet(t *testing.T, hash string) (*UserOp, bool) {
	var userOp UserOp

	ok, err := store.inner.Get(&userOp, "hash = ?", hash)
	assert.NoError(t, err)

	if !ok {
		return nil, false
	}

	return &userOp, true
}

func assertCreateUserOp(t *testing.T, store *Store, hash string, hexSender string, nonce int64, validUntilOffset ...time.Duration) {
	var paymasterAndData [60]byte
	copy(paymasterAndData[:], []byte{0x01, 0x02, 0x03})

	userOp := contract.PackedUserOperation{
		Sender:             common.HexToAddress(hexSender),
		Nonce:              big.NewInt(nonce),
		PreVerificationGas: big.NewInt(666),
		PaymasterAndData:   paymasterAndData[:],
	}

	validUntil := time.Now()
	if len(validUntilOffset) > 0 {
		validUntil = validUntil.Add(validUntilOffset[0])
	}

	err := store.UserOp.Create(&userOp, hash, validUntil)
	assert.NoError(t, err)
}

func TestUserOpCreate(t *testing.T) {
	store := newTestStore()

	assertCreateUserOp(t, store, "hash-1", "0x01", 1)

	userOp, ok := store.UserOp.assertGet(t, "hash-1")
	assert.True(t, ok)
	assert.NotNil(t, userOp)
	assert.Equal(t, UserOpStatusSigned, userOp.Status)
	assert.Greater(t, len(userOp.RawUserOp), 0)
}

func TestUserOpPendingCount(t *testing.T) {
	store := newTestStore()

	assertCreateUserOp(t, store, "hash-1", "0x01", 1)
	assertCreateUserOp(t, store, "hash-2", "0x01", 2)
	assertCreateUserOp(t, store, "hash-3", "0x02", 1)

	// 2 pending user ops for sender 0x01
	count, err := store.UserOp.GetPendingCount(common.HexToAddress("0x01"))
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)

	// 1 pending user ops for sender 0x02
	count, err = store.UserOp.GetPendingCount(common.HexToAddress("0x02"))
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestUserOpDeleteExpired(t *testing.T) {
	store := newTestStore()

	assertCreateUserOp(t, store, "hash-1", "0x01", 1, -time.Hour) // expired
	assertCreateUserOp(t, store, "hash-2", "0x01", 2, time.Hour)  // not expired
	assertCreateUserOp(t, store, "hash-3", "0x02", 1, -time.Hour) // expired

	// 2 expired user ops should be deleted
	deleted, err := store.UserOp.DeleteExpired(time.Minute)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), deleted)

	// user op "hash-1" should be deleted
	userOp, ok := store.UserOp.assertGet(t, "hash-1")
	assert.False(t, ok)
	assert.Nil(t, userOp)

	// user op "hash-2" should not be deleted
	userOp, ok = store.UserOp.assertGet(t, "hash-2")
	assert.True(t, ok)
	assert.NotNil(t, userOp)

	// user op "hash-3" should be deleted
	userOp, ok = store.UserOp.assertGet(t, "hash-3")
	assert.False(t, ok)
	assert.Nil(t, userOp)
}

func TestUserOpUpdate(t *testing.T) {
	store := newTestStore()

	assertCreateUserOp(t, store, "hash-1", "0x01", 1)

	// not found to update
	updated, err := store.UserOp.Update(&contract.VerifyingPaymasterSponsored{
		UserOpHash:            [32]byte{0x01},
		ActualGasCost:         big.NewInt(111),
		ActualUserOpFeePerGas: big.NewInt(222),
	})
	assert.NoError(t, err)
	assert.False(t, updated)

	// found to update - success
	event1 := contract.VerifyingPaymasterSponsored{
		UserOpHash:            [32]byte{0x01},
		Success:               true,
		ActualGasCost:         big.NewInt(111),
		ActualUserOpFeePerGas: big.NewInt(222),
		Raw: types.Log{
			BlockNumber: 555,
		},
	}

	assertCreateUserOp(t, store, hexutil.Encode(event1.UserOpHash[:]), "0x01", 1)

	updated, err = store.UserOp.Update(&event1)
	assert.NoError(t, err)
	assert.True(t, updated)

	userOp, ok := store.UserOp.assertGet(t, hexutil.Encode(event1.UserOpHash[:]))
	assert.True(t, ok)
	assert.NotNil(t, userOp)
	assert.Equal(t, UserOpStatusSucceeded, userOp.Status)
	assert.Equal(t, event1.ActualGasCost, userOp.ActualGasCost.BigInt())
	assert.Equal(t, event1.ActualUserOpFeePerGas, userOp.ActualUserOpFeePerGas.BigInt())
	assert.Equal(t, uint64(555), userOp.BlockTimestamp)

	// found to update - failed
	event2 := contract.VerifyingPaymasterSponsored{
		UserOpHash:            event1.UserOpHash,
		Success:               false,
		ActualGasCost:         big.NewInt(333),
		ActualUserOpFeePerGas: big.NewInt(444),
		Raw: types.Log{
			BlockNumber: 666,
		},
	}

	updated, err = store.UserOp.Update(&event2)
	assert.NoError(t, err)
	assert.True(t, updated)

	userOp, ok = store.UserOp.assertGet(t, hexutil.Encode(event2.UserOpHash[:]))
	assert.True(t, ok)
	assert.NotNil(t, userOp)
	assert.Equal(t, UserOpStatusFailed, userOp.Status)
	assert.Equal(t, event2.ActualGasCost, userOp.ActualGasCost.BigInt())
	assert.Equal(t, event2.ActualUserOpFeePerGas, userOp.ActualUserOpFeePerGas.BigInt())
	assert.Equal(t, uint64(666), userOp.BlockTimestamp)
}
