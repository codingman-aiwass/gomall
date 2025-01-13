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

type CartModel struct {
	Model
	UserId    uint32 `gorm:"size:32;not null;column:user_id" json:"user_id"`
	ProductId uint32 `gorm:"size:32;not null;column:product_id" json:"product_id"`
	Quantity  int32  `gorm:"size:64;column:quantity" json:"quantity"`
}
