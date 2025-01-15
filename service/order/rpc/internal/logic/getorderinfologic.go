package logic

import (
	"context"
	"gomall/service/order/model"
	"gomall/service/order/rpc/internal/svc"
	"gomall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoLogic {
	return &GetOrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderInfoLogic) GetOrderInfo(in *order.GetOrderInfoReq) (*order.GetOrderInfoResp, error) {
	var orderInfo model.OrderModel
	err := l.svcCtx.DB.First(&orderInfo, "id = ?", in.OrderId).Error
	if err != nil {
		return nil, err
	}
	var address model.AddressModel
	err = l.svcCtx.DB.First(&address, "id = ?", orderInfo.AddressId).Error
	if err != nil {
		return nil, err
	}
	var orderItems []model.OrderItemModel
	err = l.svcCtx.DB.Find(&orderItems, "id =?", orderInfo.Id).Error
	if err != nil {
		return nil, err
	}
	resOrderItems := make([]*order.OrderItem, len(orderItems))
	for i, orderItem := range orderItems {
		resOrderItems[i] = &order.OrderItem{
			Item: &order.CartItem{
				ProductId: orderItem.ProductId,
				Quantity:  int32(orderItem.Quantity),
			},
			Cost: float32(orderItem.Cost),
		}
	}
	return &order.GetOrderInfoResp{Order: &order.Order{
		OrderId:      orderInfo.Id,
		UserId:       orderInfo.UserId,
		UserCurrency: orderInfo.UserCurrency,
		Address: &order.Address{
			StreetAddress: address.StreetAddress,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
		},
		Email:     orderInfo.Email,
		CreatedAt: orderInfo.CreateAt.String(),
		Status:    int32(orderInfo.Status),
	}}, nil
}
