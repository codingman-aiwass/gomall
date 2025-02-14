package logic

import (
	"context"
	"gomall/service/product/rpc/types/product"

	"gomall/service/product/api/internal/svc"
	"gomall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductLogic {
	return &ListProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProductLogic) ListProduct(req *types.ListProductRequest) (resp *types.ListProductResponse, err error) {
	res, err := l.svcCtx.ProductRpc.ListProducts(l.ctx, &product.ListProductsReq{
		Page:         req.Page,
		PageSize:     req.PageSize,
		CategoryName: req.CategoryName,
		LastId:       uint32(req.LastId),
	})
	if err != nil {
		return nil, err
	}
	var ProductList []*types.Product
	for _, perProduct := range res.Products {
		Product := &types.Product{
			Id:         perProduct.Id,
			Name:       perProduct.Name,
			Desc:       perProduct.Description,
			Price:      int64(perProduct.Price),
			Picture:    perProduct.Picture,
			Categories: perProduct.Categories,
		}
		ProductList = append(ProductList, Product)
	}

	return &types.ListProductResponse{
		ProductList: ProductList,
	}, nil
}
