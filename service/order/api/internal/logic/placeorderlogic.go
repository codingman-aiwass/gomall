package logic

import (
	"context"
	"gomall/service/order/api/internal/svc"
	"gomall/service/order/api/internal/types"
	"gomall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceOrderLogic {
	return &PlaceOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceOrderLogic) PlaceOrder(req *types.PlaceOrderRequest) (resp *types.PlaceOrderResponse, err error) {
	OrderItems := make([]*order.OrderItem, 0)
	for _, item := range req.OrderItems {
		OrderItems = append(OrderItems, &order.OrderItem{
			Cost: float32(item.Cost),
			Item: &order.CartItem{
				ProductId: item.Item.ProductId,
				Quantity:  item.Item.Quantity,
			},
		})
	}
	res, err := l.svcCtx.OrderRpc.PlaceOrder(l.ctx, &order.PlaceOrderReq{
		UserId:       uint32(l.ctx.Value("userId").(int64)),
		UserCurrency: req.UserCurrency,
		Address: &order.Address{
			StreetAddress: req.UserAddress.StreetAddress,
			City:          req.UserAddress.City,
			State:         req.UserAddress.State,
			Country:       req.UserAddress.Country,
			ZipCode:       req.UserAddress.ZipCode,
		},
		Email:      req.Email,
		OrderItems: OrderItems,
	})
	if err != nil {
		return nil, err
	}
	return &types.PlaceOrderResponse{
		Order: types.OrderResult{OrderId: res.Order.OrderId},
	}, nil
}
