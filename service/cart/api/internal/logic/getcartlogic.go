package logic

import (
	"context"
	"gomall/service/cart/rpc/types/cart"

	"gomall/service/cart/api/internal/svc"
	"gomall/service/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartLogic) GetCart(req *types.GetCartRequest) (resp *types.GetCartResponse, err error) {
	res, err := l.svcCtx.CartRpc.GetCart(l.ctx, &cart.GetCartReq{UserId: uint32(l.ctx.Value("userId").(int64))})
	if err != nil {
		return nil, err
	}
	var items []types.CartItem
	for _, item := range res.Cart.Items {
		items = append(items, types.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	return &types.GetCartResponse{
		UserId:    res.Cart.UserId,
		CartItems: items,
	}, nil
}
