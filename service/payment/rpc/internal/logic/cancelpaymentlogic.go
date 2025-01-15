package logic

import (
	"context"
	"gomall/service/order/rpc/types/order"
	"gomall/service/payment/model"
	"gomall/service/payment/rpc/internal/svc"
	"gomall/service/payment/rpc/types/payment"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelPaymentLogic {
	return &CancelPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelPaymentLogic) CancelPayment(in *payment.CancelPaymentReq) (*payment.CancelPaymentResp, error) {
	// 首先验证支付状态，只有pending的订单才可以取消支付
	paymentInfo := model.PaymentModel{}
	err := l.svcCtx.DB.Where("order_id = ? and status = ?", in.OrderId, model.PENDING).First(&paymentInfo).Error
	if err != nil {
		return nil, status.Error(100, "订单不存在或查询订单时发生错误")
	}
	if paymentInfo.Status != model.PENDING {
		return nil, status.Error(100, "订单状态不是待支付状态，无法取消支付")
	}
	// 更新支付状态
	paymentInfo.Status = model.CANCELED
	err = l.svcCtx.DB.Save(&paymentInfo).Error
	if err != nil {
		return nil, status.Error(100, "更新支付状态时发生错误")
	}
	// 获取user id信息
	orderInfo, err := l.svcCtx.OrderRpc.GetOrderInfo(l.ctx, &order.GetOrderInfoReq{
		OrderId: in.OrderId,
	})
	// 更新订单状态
	_, err = l.svcCtx.OrderRpc.MarkOrderCanceled(l.ctx, &order.MarkOrderCanceledReq{
		OrderId: in.OrderId,
		UserId:  orderInfo.Order.UserId,
	})

	// 记录到日志中
	paymentLog := model.PaymentLogModel{
		TransactionId: in.TransactionId,
		Action:        model.CANCEL,
		Message:       "Payment Canceled",
		Status:        model.CANCELED,
	}
	err = l.svcCtx.DB.Create(&paymentLog).Error
	if err != nil {
		return nil, status.Error(100, "记录支付日志时发生错误")
	}

	return &payment.CancelPaymentResp{Success: true, Message: "cancel payment successfully."}, nil
}
