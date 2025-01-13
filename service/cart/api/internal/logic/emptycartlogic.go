package logic

import (
	"context"
	"gomall/service/cart/rpc/types/cart"

	"gomall/service/cart/api/internal/svc"
	"gomall/service/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmptyCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyCartLogic {
	return &EmptyCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmptyCartLogic) EmptyCart(req *types.EmptyCartRequest) (resp *types.EmptyCartResponse, err error) {
	_, err = l.svcCtx.CartRpc.EmptyCart(l.ctx, &cart.EmptyCartReq{UserId: uint32(l.ctx.Value("userId").(int64))})
	if err != nil {
		return &types.EmptyCartResponse{Success: false}, err
	}

	return &types.EmptyCartResponse{Success: true}, nil
}
