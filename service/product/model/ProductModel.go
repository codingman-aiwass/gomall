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

type ProductModel struct {
	Model
	Name        string  `gorm:"size:64;not null;column:name" json:"name"`
	Description string  `gorm:"size:256;not null;column:description" json:"description"`
	Picture     string  `gorm:"size:256;column:picture" json:"picture"`
	Price       float64 `gorm:"default:0;not null;column:price" json:"price"`
	Stock       int64   `gorm:"default:0;not null;column:stock" json:"stock"`
	Status      int8    `gorm:"size:4;not null;default:0;column:status" json:"status"`
}
