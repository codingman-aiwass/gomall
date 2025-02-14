package model

type CategoryModel struct {
	Model
	// 添加索引
	ProductId uint32 `gorm:"size:32;not null;column:product_id;index:idx_product_id" json:"product_id"`
	Category  string `gorm:"size:256;not null;column:category;index:idx_category" json:"category"`
}
