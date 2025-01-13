package logic

import (
	"context"
	"errors"
	"gomall/service/product/model"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"gomall/service/product/rpc/internal/svc"
	"gomall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.GetProductReq) (*product.GetProductResp, error) {
	FindProduct := model.ProductModel{}
	result := l.svcCtx.DB.First(&FindProduct, in.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(100, "product not found")
		}
		return nil, status.Error(500, result.Error.Error())
	}
	var categories []model.CategoryModel
	err := l.svcCtx.DB.Where("product_id = ?", FindProduct.Id).Find(&categories).Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newProduct := product.Product{
		Id:          FindProduct.Id,
		Name:        FindProduct.Name,
		Description: FindProduct.Description,
		Picture:     FindProduct.Picture,
		Price:       float32(FindProduct.Price),
	}
	for _, category := range categories {
		newProduct.Categories = append(newProduct.Categories, category.Category)
	}

	return &product.GetProductResp{Product: &newProduct}, nil
}
