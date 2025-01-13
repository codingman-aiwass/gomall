package logic

import (
	"context"
	"gomall/service/product/model"
	"google.golang.org/grpc/status"

	"gomall/service/product/rpc/internal/svc"
	"gomall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProductsLogic) ListProducts(in *product.ListProductsReq) (*product.ListProductsResp, error) {
	// todo: add your logic here and delete this line
	offset := (int64(in.Page) - 1) * in.PageSize
	limit := in.PageSize

	var categories []model.CategoryModel
	// 先去category中根据种类查询出product_id
	if err := l.svcCtx.DB.Limit(int(limit)).Offset(int(offset)).Find(&categories, "category = ?", in.CategoryName).Error; err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 再根据product_id去product中查询出product
	ids := []uint32{}
	for _, category := range categories {
		ids = append(ids, category.ProductId)
	}
	var products []model.ProductModel
	if err := l.svcCtx.DB.Where("id in (?)", ids).Find(&products).Error; err != nil {
		return nil, status.Error(500, err.Error())
	}
	var respProducts []*product.Product
	for _, db_product := range products {
		respProduct := product.Product{
			Id:          db_product.Id,
			Name:        db_product.Name,
			Description: db_product.Description,
			Picture:     db_product.Picture,
			Price:       float32(db_product.Price),
		}

		respProducts = append(respProducts, &respProduct)
	}

	return &product.ListProductsResp{Products: respProducts}, nil
}
