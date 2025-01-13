package model

type CategoryModel struct {
	Model
	ProductId uint32 `gorm:"size:32;not null;column:product_id" json:"product_id"`
	Category  string `gorm:"size:256;not null;column:category" json:"category"`
}
