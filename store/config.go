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

// Get returns the value of the config by name. If the config does not exist, it returns false.
func (store *ConfigStore) Get(name string) (string, bool, error) {
	var config Config

	exists, err := store.inner.Get(&config, "name = ?", name)
	if err != nil {
		return "", false, api.ErrDatabaseCause(err, "Failed to get config by name")
	}

	if !exists {
		return "", false, nil
	}

	return config.Value, true, nil
}

// Upsert updates the value of the config by name. If the config does not exist, it creates a new config with the given name and value.
func (store *ConfigStore) Upsert(name, value string, tx ...*gorm.DB) error {
	db := store.inner.DB
	if len(tx) > 0 {
		db = tx[0]
	}

	err := db.Model(&Config{}).
		Where("name = ?", name).
		Assign(Config{Value: value}).
		FirstOrCreate(&Config{Name: name}).
		Error

	if err != nil {
		return api.ErrDatabaseCause(err, "Failed to upsert config by name")
	}

	return nil
}
