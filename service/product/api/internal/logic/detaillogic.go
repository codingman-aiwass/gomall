package logic

import (
	"context"
	"gomall/service/product/rpc/types/product"

	"gomall/service/product/api/internal/svc"
	"gomall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	res, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &product.GetProductReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}
	return &types.DetailResponse{
		Id:         res.Product.Id,
		Name:       res.Product.Name,
		Desc:       res.Product.Description,
		Price:      int64(res.Product.Price),
		Picture:    res.Product.Picture,
		Categories: res.Product.Categories,
	}, nil
}
