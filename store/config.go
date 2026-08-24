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

// Get returns the value of the config by key. If the config does not exist, it returns false.
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

// Upsert updates the value of the config by key. If the config does not exist, it creates a new config with the given key and value.
func (store *ConfigStore) Upsert(key string, value string, tx ...*gorm.DB) error {
	db := store.inner.DB
	if len(tx) > 0 {
		db = tx[0]
	}

	// update by key
	result := db.Model(&Config{}).Where("key = ?", key).Update("value", value)
	if err := result.Error; err != nil {
		return api.ErrDatabaseCausef(err, "Failed to update config by key %v", key)
	}

	if result.RowsAffected > 0 {
		return nil
	}

	// create if absent
	config := Config{
		Key:   key,
		Value: value,
	}

	if er := db.Create(&config).Error; er != nil {
		return api.ErrDatabaseCausef(er, "Failed to create config of key %v", key)
	}

	return nil
}
