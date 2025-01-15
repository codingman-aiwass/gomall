package model

type PaymentStatus struct {
	PENDING  int
	SUCCESS  int
	FAILED   int
	CANCELED int
}

const (
	PENDING = iota
	SUCCESS
	FAILED
	CANCELED
)

const (
	CREATE = iota
	CHARGE
	CANCEL
	TIMEOUT_CANCEL
)

var PaymentStatusMap = map[int]string{
	PENDING:  "pending",
	SUCCESS:  "success",
	FAILED:   "failed",
	CANCELED: "canceled",
}

var PaymentActionMap = map[int]string{
	CREATE:         "create",
	CHARGE:         "charge",
	CANCEL:         "cancel",
	TIMEOUT_CANCEL: "timeout",
}
