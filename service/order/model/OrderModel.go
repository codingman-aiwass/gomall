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

type OrderModel struct {
	Model
	UserId       uint32 `gorm:"size:32;not null;column:user_id" json:"user_id"`
	UserCurrency string `gorm:"size:32;not null;column:user_currency" json:"user_currency"`
	AddressId    uint32 `gorm:"size:32;not null;column:address_id" json:"address_id"`
	Email        string `gorm:"size:256;column:email" json:"email"`
	// 0 表示未支付，1 表示已支付，2表示已取消，3表示已完成
	Status int8 `gorm:"size:4;not null;default:0;column:status" json:"status"`
}
