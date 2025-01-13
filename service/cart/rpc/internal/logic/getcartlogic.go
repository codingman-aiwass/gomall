package logic

import (
	"context"
	"gomall/service/cart/model"
	"google.golang.org/grpc/status"

	"gomall/service/cart/rpc/internal/svc"
	"gomall/service/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartLogic) GetCart(in *cart.GetCartReq) (*cart.GetCartResp, error) {
	err := l.svcCtx.DB.First(&model.CartModel{}, "user_id = ?", in.UserId).Error
	if err != nil {
		return nil, status.Error(100, "cart is empty")
	}
	var cartItems []*model.CartModel
	err = l.svcCtx.DB.Find(&cartItems, "user_id = ?", in.UserId).Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	var resItems []*cart.CartItem
	for _, item := range cartItems {
		resItems = append(resItems, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: in.UserId,
			Items:  resItems,
		},
	}, nil
}
