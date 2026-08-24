package store

import (
	"time"
)

const (
	UserOpStatusSigned    = "signed"
	UserOpStatusSucceeded = "succeeded"
	UserOpStatusFailed    = "failed"
)

var AllTables = []any{&UserOp{}, &Config{}}

type Model struct {
	ID        uint64    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type UserOp struct {
	Model
	Hash                  string    `gorm:"size:66;not null;unique"`
	Sender                string    `gorm:"size:42;not null;index:idx_sender_status"`
	Nonce                 string    `gorm:"size:66;not null"`
	ValidUntil            time.Time `gorm:"not null;index:idx_status_valid_until,priority:2"`
	Status                string    `gorm:"size:32;not null;index:idx_sender_status;index:idx_status_valid_until,priority:1"`
	ActualGasCost         string    `gorm:"size:32;not null"`
	ActualUserOpFeePerGas string    `gorm:"size:32;not null"`
}

type Config struct {
	Model

	Key   string `gorm:"size:64;not null;unique"`
	Value string `gorm:"size:1024;not null"`
}
