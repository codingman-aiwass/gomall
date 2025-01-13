package logic

import (
	"context"
	"errors"
	"gomall/service/product/model"
	"gomall/service/product/rpc/internal/svc"
	"gomall/service/product/rpc/types/product"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	oldProduct := model.ProductModel{}
	result := l.svcCtx.DB.First(&oldProduct, in.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(100, "product not found")
		}
		return nil, status.Error(500, result.Error.Error())
	}
	if in.Name != "" {
		oldProduct.Name = in.Name
	}

	if in.Description != "" {
		oldProduct.Description = in.Description
	}

	if in.Picture != "" {
		oldProduct.Picture = in.Picture
	}

	if in.Stock != 0 {
		oldProduct.Stock = in.Stock
	}

	if in.Price != 0 {
		oldProduct.Price = float64(in.Price)
	}

	if in.Status != 0 {
		oldProduct.Status = int8(in.Status)
	}

	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&oldProduct).Error; err != nil {
			return err
		}
		if len(in.Categories) > 0 {
			// 先删除所有的分类
			if err := tx.Where("product_id = ?", in.Id).Delete(&model.CategoryModel{}).Error; err != nil {
				return err
			}
			// 再添加新的分类
			for _, category := range in.Categories {
				newProductCategory := model.CategoryModel{
					ProductId: oldProduct.Id,
					Category:  category,
				}
				if err := tx.Create(&newProductCategory).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return &product.UpdateProductResp{Success: false}, status.Error(500, err.Error())
	}

	return &product.UpdateProductResp{Success: true}, nil
}
