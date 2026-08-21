package store

import "time"

const (
	UserOpStatusSigned  = "signed"
	UserOpStatusSuccess = "success"
	UserOpStatusFailed  = "failed"
)

var AllTables = []any{&UserOp{}}

type Model struct {
	ID        uint64    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type UserOp struct {
	Model
	UserOpHash            string    `gorm:"size:66;not null;unique"`
	Sender                string    `gorm:"size:42;not null;index:idx_sender_status"`
	Nonce                 string    `gorm:"size:66;not null"`
	ValidUntil            time.Time `gorm:"not null;index:idx_status_valid_until,priority:2"`
	Status                string    `gorm:"size:32;not null;index:idx_sender_status;index:idx_status_valid_until,priority:1"`
	ActualGasCost         string    `gorm:"size:32;not null"`
	ActualUserOpFeePerGas string    `gorm:"size:32;not null"`
}
