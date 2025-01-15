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

type UserModel struct {
	Model
	Password string `gorm:"size:64;not null;column:password"`
	Status   int8   `gorm:"size:4;not null;default:0;column:status" json:"status"`
	Email    string `gorm:"size:64;column:email" json:"email"`
	Verified bool   `gorm:"default:false;column:verified"`
	Currency string `gorm:"size:16;column:currency" json:"currency"`
}
