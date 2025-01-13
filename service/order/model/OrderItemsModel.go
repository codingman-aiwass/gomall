package model

type OrderItemModel struct {
	Model
	OrderId   uint32 `gorm:"size:32;not null;column:order_id" json:"order_id"`
	ProductId uint32 `gorm:"size:32;not null;column:product_id" json:"product_id"`
	Quantity  uint32 `gorm:"size:32;not null;column:quantity" json:"quantity"`
	Cost      uint32 `gorm:"size:32;not null;column:cost" json:"cost"`
}
