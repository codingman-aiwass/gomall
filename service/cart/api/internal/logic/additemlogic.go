package logic

import (
	"context"
	"gomall/service/cart/rpc/types/cart"

	"gomall/service/cart/api/internal/svc"
	"gomall/service/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddItemLogic) AddItem(req *types.AddItemRequest) (resp *types.AddItemResponse, err error) {
	_, err = l.svcCtx.CartRpc.AddItem(l.ctx, &cart.AddItemReq{UserId: uint32(l.ctx.Value("userId").(int64)), Item: &cart.CartItem{
		ProductId: req.CartItem.ProductId,
		Quantity:  req.CartItem.Quantity,
	}})
	if err != nil {
		return &types.AddItemResponse{Success: false}, err
	}

	return &types.AddItemResponse{Success: true}, nil
}
