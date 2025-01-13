package logic

import (
	"context"
	"gomall/service/product/rpc/types/product"

	"gomall/service/product/api/internal/svc"
	"gomall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	_, err = l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &product.UpdateProductReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Desc,
		Price:       float32(req.Price),
		Stock:       req.Stock,
		Status:      req.Status,
		Picture:     req.Picture,
		Categories:  req.Categories,
	})
	if err != nil {
		return &types.UpdateResponse{
			Success: false,
		}, err
	}

	return &types.UpdateResponse{
		Success: true,
	}, nil
}
