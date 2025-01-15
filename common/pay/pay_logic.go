package pay

import (
	"errors"
	"math/rand"
	"strconv"
)

// PaymentRequest 定义支付请求参数结构体
type PaymentRequest struct {
	CreditCardNumber string
	CVV              int32
	ValidDateYear    int32
	ValidDateMonth   int32
	Amount           float64
}

// PaymentResponse 定义支付响应结构体
type PaymentResponse struct {
	Success bool
	Message string
}

// SimulatePayment 模拟发起支付请求的方法
func SimulatePayment(request PaymentRequest) (PaymentResponse, error) {
	// 检查参数合法性
	if request.CreditCardNumber == "" || len(request.CreditCardNumber) != 16 {
		return PaymentResponse{}, errors.New("invalid credit card number")
	}
	if len(strconv.Itoa(int(request.CVV))) != 3 {
		return PaymentResponse{}, errors.New("invalid CVV")
	}
	if request.Amount <= 0 {
		return PaymentResponse{}, errors.New("amount must be greater than zero")
	}

	// 模拟支付处理（随机返回成功或失败）
	isSuccess := rand.Intn(2) == 0 // 50% 成功概率
	//isSuccess := false // 50% 成功概率

	if isSuccess {
		return PaymentResponse{
			Success: true,
			Message: "Payment processed successfully",
		}, nil
	}

	return PaymentResponse{
		Success: false,
		Message: "Payment failed due to insufficient funds or network error",
	}, nil
}
