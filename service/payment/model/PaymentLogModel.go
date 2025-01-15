package model

type PaymentLogModel struct {
	Model
	TransactionId uint64 `gorm:"size:64;not null;column:transaction_id"`

	// 0: 'create', 1: 'charge', 2: 'cancel', 3: 'timeout_cancel'
	Action  int8   `gorm:"size:4;not null;column:action" json:"action"`
	Message string `gorm:"size:256;column:message" json:"message"`
	// 0: pending, 1: success, 2: failed 3: canceled
	Status int8 `gorm:"size:4;not null;default:0;column:status" json:"status"`
}
