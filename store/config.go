package store

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/store"
	"gorm.io/gorm"
)

type ConfigStore struct {
	inner *store.Store
}

func NewConfigStore(store *store.Store) *ConfigStore {
	return &ConfigStore{
		inner: store,
	}
}

func (store *ConfigStore) Get(key string) (string, bool, error) {
	var config Config

	exists, err := store.inner.Get(&config, "key = ?", key)
	if err != nil {
		return "", false, api.ErrDatabaseCause(err, "Failed to get config by key")
	}

	if !exists {
		return "", false, nil
	}

	return config.Value, true, nil
}

func (store *ConfigStore) Update(key string, value string, tx ...*gorm.DB) (bool, error) {
	db := store.inner.DB
	if len(tx) > 0 {
		db = tx[0]
	}

	result := db.Model(&Config{}).
		Where("key = ?", key).
		Where("value != ?", value).
		Update("value", value)

	if err := result.Error; err != nil {
		return false, api.ErrDatabaseCause(err, "Failed to update config by key")
	}

	return result.RowsAffected > 0, nil
}
