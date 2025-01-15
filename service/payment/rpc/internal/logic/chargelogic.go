package logic

import (
	"context"
	"encoding/json"
	"errors"
	"gomall/common/pay"
	"gomall/common/snowflake"
	"gomall/service/mq/rpc/types/mq"
	"gomall/service/order/rpc/types/order"
	"gomall/service/payment/model"
	"gomall/service/payment/rpc/internal/svc"
	"gomall/service/payment/rpc/types/payment"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChargeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChargeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChargeLogic {
	return &ChargeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChargeLogic) Charge(in *payment.ChargeReq) (*payment.ChargeResp, error) {
	transactionID := snowflake.GenerateTransactionID()
	orderInfo, err := l.svcCtx.OrderRpc.GetOrderInfo(l.ctx, &order.GetOrderInfoReq{
		OrderId: in.OrderId,
	})
	if err != nil {
		return nil, err
	}
	if orderInfo == nil {
		return nil, status.Error(100, "订单不存在")
	}

	if orderInfo.Order.UserId != in.UserId {
		return nil, status.Error(100, "用户id与订单id不匹配")
	}

	if orderInfo.Order.Status != 0 {
		return nil, status.Error(100, "订单不处于待支付状态，无法支付")
	}

	// 防止重复支付，查看payments表中是否有成功支付记录
	var paymentInfo model.PaymentModel
	err = l.svcCtx.DB.Where("order_id = ? and status = ?", in.OrderId, 1).First(&paymentInfo).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(100, "查询支付记录失败")
		}
	}
	if paymentInfo.Id > 0 {
		return nil, status.Error(100, "订单已支付，请勿重复支付")
	}

	paymentLog := model.PaymentLogModel{
		TransactionId: transactionID,
		Action:        model.CREATE,
		Message:       "Payment Created",
		Status:        model.PENDING,
	}
	paymentInfo = model.PaymentModel{
		TransactionId: transactionID,
		OrderId:       in.OrderId,
		UserId:        in.UserId,
		Amount:        float64(in.Amount),
		Status:        model.PENDING,
	}
	// 使用事务保存记录到数据库
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&paymentLog).Error; err != nil {
			return err
		}
		if err = tx.Create(&paymentInfo).Error; err != nil {
			return err
		}
		// 手动刷新 paymentInfo 的 CreateAt 和其他字段
		tx.First(&paymentInfo, paymentInfo.Id)
		tx.First(&paymentLog, paymentLog.Id)
		return nil
	})
	// 发送延时消息，15分钟后检查订单状态
	cancelPayload := struct {
		TransactionId uint64 `json:"transaction_id"`
		UserId        uint32 `json:"user_id"`
	}{
		TransactionId: transactionID,
		UserId:        in.UserId,
	}
	payloadBytes, err := json.Marshal(cancelPayload)
	if err != nil {
		return nil, status.Error(500, "创建支付消息失败")
	}

	_, err = l.svcCtx.MqRpc.SendDelayMessage(l.ctx, &mq.SendDelayMessageReq{
		Topic:   "transaction_timeout",
		Payload: payloadBytes,
		//DelaySeconds: 15 * 60, // 15分钟
		DelaySeconds: 15 * 1, // 15s
	})
	if err != nil {
		return nil, status.Error(500, "发送消息失败")
	}

	// 尝试发起支付请求
	resp, _ := pay.SimulatePayment(pay.PaymentRequest{
		Amount:           float64(in.Amount),
		CreditCardNumber: in.CreditCard.CreditCardNumber,
		CVV:              in.CreditCard.CreditCardCvv,
		ValidDateMonth:   in.CreditCard.CreditCardExpirationMonth,
		ValidDateYear:    in.CreditCard.CreditCardExpirationYear,
	})
	// TODO 支付失败的err没有成功传回给checkoutlogic,导致其认为支付成功，将订单标记为已支付
	logx.Infof("chargelogic.go resp: %v", resp)
	success := "true"
	if resp.Success {
		paymentLog.Status = model.SUCCESS
		paymentLog.Message = resp.Message
		paymentInfo.Status = model.SUCCESS
		logx.Infof("payment success: %v", resp.Message)
	} else {
		paymentLog.Status = model.FAILED
		paymentLog.Message = resp.Message
		paymentInfo.Status = model.FAILED
		logx.Errorf("payment failed: %v", resp.Message)
		success = "false"
	}

	// 使用事务保存修改记录到数据库
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		logx.Infof("payment log %v", paymentLog)
		if err = tx.Save(&paymentLog).Error; err != nil {
			return err
		}
		if err = tx.Save(&paymentInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return &payment.ChargeResp{TransactionId: transactionID, Success: success}, nil
}
