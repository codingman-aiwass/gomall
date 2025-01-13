package logic

import (
	"context"
	"gomall/service/product/rpc/types/product"

	"gomall/service/product/api/internal/svc"
	"gomall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.ProductRpc.DeleteProduct(l.ctx, &product.DeleteProductReq{
		Id: req.Id,
	})

	if err != nil {
		return &types.RemoveResponse{
			Success: false,
		}, err
	}
	return &types.RemoveResponse{
		Success: true,
	}, nil
}
