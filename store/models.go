package store

import (
	"time"

	"github.com/shopspring/decimal"
)

var AllTables = []any{&UserOp{}, &Config{}}

type Model struct {
	ID        uint64    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type UserOp struct {
	Model

	Hash    string `gorm:"size:66;not null;unique"`
	Sender  string `gorm:"size:42;not null;index:idx_sender_time"`
	Nonce   string `gorm:"size:66;not null"`
	Success bool   `gorm:"not null"`

	ActualGasCost decimal.Decimal `gorm:"type:decimal(32,0);not null"`
	ActualGasUsed uint64          `gorm:"not null"`
	BlockTime     time.Time       `gorm:"not null;index:idx_sender_time;index:idx_block_time"`

	RawUserOp string `gorm:"type:text;not null"`
}

type Config struct {
	Model

	Name  string `gorm:"size:64;not null;unique"`
	Value string `gorm:"size:1024;not null"`
}
