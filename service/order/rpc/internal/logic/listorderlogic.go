package logic

import (
	"context"
	"gomall/service/order/model"
	"strconv"

	"gomall/service/order/rpc/internal/svc"
	"gomall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOrderLogic) ListOrder(in *order.ListOrderReq) (*order.ListOrderResp, error) {
	var res_orders []*order.Order
	var orders []model.OrderModel
	var orderItems []model.OrderItemModel
	err := l.svcCtx.DB.Find(&orders, "user_id = ?", in.UserId).Error
	if err != nil {
		return nil, err
	}
	orderIds := make([]uint32, len(orders))
	var addressIds []uint32
	addressMap := map[uint32]model.AddressModel{}
	for i, order := range orders {
		orderIds[i] = order.Id
		addressMap[order.AddressId] = model.AddressModel{}
	}
	for id := range addressMap {
		addressIds = append(addressIds, id)
	}
	err = l.svcCtx.DB.Find(&orderItems, "id in (?)", orderIds).Error
	if err != nil {
		return nil, err
	}

	var addresses []model.AddressModel
	err = l.svcCtx.DB.Find(&addresses, "id in (?)", addressIds).Error
	if err != nil {
		return nil, err
	}
	for _, address := range addresses {
		addressMap[address.Id] = address
	}

	for _, per_order := range orders {
		for _, orderItem := range orderItems {
			if orderItem.OrderId == per_order.Id {
				res_orders = append(res_orders, &order.Order{
					OrderId:      strconv.Itoa(int(per_order.Id)),
					UserId:       per_order.UserId,
					UserCurrency: per_order.UserCurrency,
					Address: &order.Address{
						StreetAddress: addressMap[per_order.AddressId].StreetAddress,
						City:          addressMap[per_order.AddressId].City,
						State:         addressMap[per_order.AddressId].State,
						Country:       addressMap[per_order.AddressId].Country,
						ZipCode:       addressMap[per_order.AddressId].ZipCode,
					},
					Email:     per_order.Email,
					CreatedAt: per_order.CreateAt.String(),
					OrderItems: []*order.OrderItem{
						{
							Item: &order.CartItem{
								ProductId: orderItem.ProductId,
								Quantity:  int32(orderItem.Quantity),
							},
							Cost: float32(orderItem.Cost),
						},
					},
				})
			}
		}
	}

	return &order.ListOrderResp{Orders: res_orders}, nil
}
