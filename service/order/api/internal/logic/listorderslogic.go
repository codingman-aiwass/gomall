package logic

import (
	"context"
	"gomall/service/order/rpc/types/order"

	"gomall/service/order/api/internal/svc"
	"gomall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrdersLogic {
	return &ListOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOrdersLogic) ListOrders(req *types.ListOrderRequest) (resp *types.ListOrderResponse, err error) {
	res, err := l.svcCtx.OrderRpc.ListOrder(l.ctx, &order.ListOrderReq{
		UserId: uint32(l.ctx.Value("userId").(int64)),
	})
	if err != nil {
		return nil, err
	}

	var orders []types.Order
	orderItems := make([]types.OrderItem, 0)
	for _, item := range res.Orders {
		for _, orderItem := range item.OrderItems {
			orderItems = append(orderItems, types.OrderItem{
				Item: types.CartItem{
					ProductId: orderItem.Item.ProductId,
					Quantity:  orderItem.Item.Quantity,
				},
				Cost: float64(orderItem.Cost),
			})
		}
	}

	for _, item := range res.Orders {
		orders = append(orders, types.Order{
			OrderId:      item.OrderId,
			UserId:       item.UserId,
			UserCurrency: item.UserCurrency,
			UserAddress: types.Address{
				StreetAddress: item.Address.StreetAddress,
				City:          item.Address.City,
				State:         item.Address.State,
				Country:       item.Address.Country,
				ZipCode:       item.Address.ZipCode,
			},
			Email:      item.Email,
			OrderItems: orderItems,
			CreateAt:   string(item.CreatedAt),
		})
	}
	return &types.ListOrderResponse{
		Orders: orders,
	}, nil
}
