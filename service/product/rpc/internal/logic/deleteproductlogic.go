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

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductReq) (*product.DeleteProductResp, error) {
	oldProduct := model.ProductModel{}
	result := l.svcCtx.DB.First(&oldProduct, in.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(100, "product not found")
		}
		return nil, status.Error(500, result.Error.Error())
	}
	// 需要同时删除product_models中的记录和category_models中的记录
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&oldProduct).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.CategoryModel{}, "product_id = ?", in.Id).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return &product.DeleteProductResp{Success: false}, status.Error(500, err.Error())
	}

	return &product.DeleteProductResp{Success: true}, nil
}
