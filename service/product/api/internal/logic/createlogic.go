package logic

import (
	"context"
	"gomall/service/product/rpc/types/product"

	"gomall/service/product/api/internal/svc"
	"gomall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.ProductRpc.CreateProduct(l.ctx, &product.CreateProductReq{
		Name:        req.Name,
		Description: req.Desc,
		Stock:       req.Stock,
		Price:       float32(req.Price),
		Status:      req.Status,
		Picture:     req.Picture,
		Categories:  req.Categories,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
