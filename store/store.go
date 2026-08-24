package store

import (
	"github.com/Conflux-Chain/go-conflux-util/store"
	"gorm.io/gorm"
)

type Store struct {
	*store.Store

	UserOp *UserOpStore
	Config *ConfigStore
}

func NewStore(db *gorm.DB) *Store {
	store := store.NewStore(db)

	return &Store{
		Store:  store,
		UserOp: NewUserOpStore(store),
		Config: NewConfigStore(store),
	}
}
