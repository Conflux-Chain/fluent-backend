package store

import (
	"math/big"
	"testing"
	"time"

	"github.com/Conflux-Chain/fluent-backend/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func assertCreateUserOp(t *testing.T, store *Store, hash string, sender string, nonce int64, blockTimeOffset ...time.Duration) {
	var paymasterAndData [60]byte
	copy(paymasterAndData[:], []byte{0x01, 0x02, 0x03})

	userOp := contract.PackedUserOperation{
		Sender:             common.BytesToAddress([]byte(sender)),
		Nonce:              big.NewInt(nonce),
		PreVerificationGas: big.NewInt(666),
		PaymasterAndData:   paymasterAndData[:],
	}

	event := contract.EntryPointUserOperationEvent{
		UserOpHash:    common.BytesToHash([]byte(hash)),
		Sender:        common.BytesToAddress([]byte(sender)),
		Nonce:         big.NewInt(nonce),
		Success:       true,
		ActualGasCost: big.NewInt(30),
		ActualGasUsed: big.NewInt(6),
	}

	blockTime := time.Now()
	if len(blockTimeOffset) > 0 {
		blockTime = blockTime.Add(blockTimeOffset[0])
	}

	err := store.UserOp.Create(&userOp, &event, blockTime)
	assert.NoError(t, err)
}

func TestUserOpCreate(t *testing.T) {
	store := newTestStore()

	assertCreateUserOp(t, store, "hash-1", "user-1", 1, -10*time.Minute)
	assertCreateUserOp(t, store, "hash-2", "user-1", 2, -5*time.Minute)
	assertCreateUserOp(t, store, "hash-3", "user-2", 1)

	count, err := store.UserOp.GetCountByBlockTimestamp(common.BytesToAddress([]byte("user-1")), time.Now().Add(-time.Hour))
	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)

	count, err = store.UserOp.GetCountByBlockTimestamp(common.BytesToAddress([]byte("user-1")), time.Now().Add(-7*time.Minute))
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}
