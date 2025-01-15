package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id       uint32         `gorm:"size:32;primaryKey;autoIncrement;column:id" json:"id"`
	CreateAt time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:create_at" json:"create_at"`
	DeleteAt gorm.DeletedAt `gorm:"index;column:delete_at" json:"deleted_at,omitempty" swaggerignore:"true"`
}

type PaymentModel struct {
	Model
	UserId        uint32    `gorm:"size:32;not null;column:user_id"`
	OrderId       uint32    `gorm:"size:32;not null;column:order_id"`
	TransactionId uint64    `gorm:"size:64;not null;column:transaction_id"`
	Amount        float64   `gorm:"size:64;not null;column:amount"`
	Currency      string    `gorm:"size:16;column:currency" json:"currency"`
	UpdateAt      time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;column:update_at" json:"update_at"`
	// 0: pending, 1: success, 2: failed 3: canceled
	Status int8 `gorm:"size:4;not null;default:0;column:status" json:"status"`
}
