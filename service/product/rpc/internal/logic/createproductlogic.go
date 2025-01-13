package logic

import (
	"context"
	"gomall/service/product/model"
	"gorm.io/gorm"

	"gomall/service/product/rpc/internal/svc"
	"gomall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *product.CreateProductReq) (*product.CreateProductResp, error) {
	newProduct := model.ProductModel{
		Name:        in.Name,
		Description: in.Description,
		Stock:       in.Stock,
		Price:       float64(in.Price),
		Status:      int8(in.Status),
		Picture:     in.Picture,
	}

	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newProduct).Error; err != nil {
			return err
		}

		for _, category := range in.Categories {
			newProductCategory := model.CategoryModel{
				ProductId: newProduct.Id,
				Category:  category,
			}
			if err := tx.Create(&newProductCategory).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &product.CreateProductResp{Id: newProduct.Id}, nil
}
