package store

import (
	"testing"

	"github.com/Conflux-Chain/go-conflux-util/store"
	"github.com/stretchr/testify/assert"
)

func newTestStore() *Store {
	config := store.NewMemoryConfig()
	return NewStore(config.MustOpenOrCreate(AllTables...))
}

func TestConfigGet(t *testing.T) {
	store := newTestStore()

	// not found
	value, ok, err := store.Config.Get("aaa")
	assert.NoError(t, err)
	assert.False(t, ok)
	assert.Equal(t, "", value)

	// found
	err = store.Config.Upsert("aaa", "AAA")
	assert.NoError(t, err)

	value, ok, err = store.Config.Get("aaa")
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, "AAA", value)
}

func TestConfigUpsert(t *testing.T) {
	store := newTestStore()

	// insert
	err := store.Config.Upsert("aaa", "AAA")
	assert.NoError(t, err)

	value, _, _ := store.Config.Get("aaa")
	assert.Equal(t, "AAA", value)

	// update
	err = store.Config.Upsert("aaa", "BBB")
	assert.NoError(t, err)

	value, _, _ = store.Config.Get("aaa")
	assert.Equal(t, "BBB", value)
}
