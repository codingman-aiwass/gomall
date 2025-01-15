package logic

import (
	"context"
	"encoding/json"
	"gomall/service/order/model"
	"gorm.io/gorm"

	"gomall/service/mq/rpc/types/mq"
	"gomall/service/order/rpc/internal/svc"
	"gomall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceOrderLogic {
	return &PlaceOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlaceOrderLogic) PlaceOrder(in *order.PlaceOrderReq) (*order.PlaceOrderResp, error) {
	var OrderItems []model.OrderItemModel
	address := model.AddressModel{
		UserId:        in.UserId,
		StreetAddress: in.Address.StreetAddress,
		City:          in.Address.City,
		State:         in.Address.State,
		Country:       in.Address.Country,
		ZipCode:       in.Address.ZipCode,
	}
	// 如果该地址在数据库中不存在，则创建该地址
	if err := l.svcCtx.DB.Where(&address).FirstOrCreate(&address).Error; err != nil {
		logx.Errorf("创建新收货地址失败：%v", err)
		return nil, err
	}

	for _, item := range in.OrderItems {
		OrderItems = append(OrderItems, model.OrderItemModel{
			ProductId: item.Item.ProductId,
			Quantity:  uint32(item.Item.Quantity),
			Cost:      uint32(item.Cost),
		})
	}
	orderModel := model.OrderModel{
		UserId:       in.UserId,
		UserCurrency: in.UserCurrency,
		Email:        in.Email,
		AddressId:    address.Id,
	}
	// 使用事务创建订单
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 设置订单状态为未支付
		orderModel.Status = 0
		if err := tx.Create(&orderModel).Error; err != nil {
			return err
		}
		for _, item := range OrderItems {
			item.OrderId = orderModel.Id
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}

		// 发送延时消息，15分钟后检查订单状态
		cancelPayload := struct {
			OrderId uint32 `json:"order_id"`
		}{
			OrderId: orderModel.Id,
		}
		payloadBytes, err := json.Marshal(cancelPayload)
		if err != nil {
			return err
		}

		_, err = l.svcCtx.MqRpc.SendDelayMessage(l.ctx, &mq.SendDelayMessageReq{
			Topic:        "order_timeout",
			Payload:      payloadBytes,
			DelaySeconds: 15 * 60, // 15分钟
			//DelaySeconds: 15 * 1, // 15s
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &order.PlaceOrderResp{Order: &order.OrderResult{OrderId: orderModel.Id}}, nil
}
