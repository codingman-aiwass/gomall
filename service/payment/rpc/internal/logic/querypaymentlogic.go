package logic

import (
	"context"
	"gomall/service/payment/model"
	"google.golang.org/grpc/status"

	"gomall/service/payment/rpc/internal/svc"
	"gomall/service/payment/rpc/types/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPaymentLogic {
	return &QueryPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryPaymentLogic) QueryPayment(in *payment.QueryPaymentReq) (*payment.QueryPaymentResp, error) {
	var paymentInfo model.PaymentModel
	err := l.svcCtx.DB.First(&paymentInfo, "transaction_id = ?", in.TransactionId).Error
	if err != nil {
		return nil, status.Error(100, "交易不存在或者查询交易时发生错误")
	}

	return &payment.QueryPaymentResp{
		OrderId: paymentInfo.OrderId,
		Amount:  float32(paymentInfo.Amount),
		Status:  model.PaymentStatusMap[int(paymentInfo.Status)],
	}, nil
}
