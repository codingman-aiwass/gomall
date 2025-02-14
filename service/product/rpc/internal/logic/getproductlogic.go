package logic

import (
	"context"
	"errors"
	"fmt"
	"gomall/service/product/model"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strconv"
	"time"

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
	ctx := context.Background()
	hashKey := fmt.Sprintf("product:%d", in.Id)
	// 先尝试从 Redis Hash 中获取商品详情
	productData, err := l.svcCtx.RDB.HGetAll(ctx, hashKey).Result()
	var prod *product.Product
	if err == nil && len(productData) > 0 {
		id, _ := strconv.ParseUint(productData["id"], 10, 32)
		price, _ := strconv.ParseFloat(productData["price"], 32)
		prod = &product.Product{
			Id:          uint32(id),
			Name:        productData["name"],
			Description: productData["description"],
			Picture:     productData["picture"],
			Price:       float32(price),
		}
	} else {
		// 若 Redis 中未命中，则从数据库中查询
		var dbProduct model.ProductModel
		if err := l.svcCtx.DB.First(&dbProduct, in.Id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Error(100, "product not found")
			}
			return nil, status.Error(500, err.Error())
		}
		prod = &product.Product{
			Id:          dbProduct.Id,
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Picture:     dbProduct.Picture,
			Price:       float32(dbProduct.Price),
		}
		// 缓存查询结果到 Redis Hash，设置过期时间
		l.svcCtx.RDB.HMSet(ctx, hashKey, map[string]interface{}{
			"id":          prod.Id,
			"name":        prod.Name,
			"description": prod.Description,
			"picture":     prod.Picture,
			"price":       prod.Price,
		})
		l.svcCtx.RDB.Expire(ctx, hashKey, time.Hour)
	}
	// 查询商品所属类别（这里暂时未加入 Redis 缓存，可根据需求扩展）
	var categories []model.CategoryModel
	if err := l.svcCtx.DB.Where("product_id = ?", prod.Id).Find(&categories).Error; err != nil {
		return nil, status.Error(500, err.Error())
	}
	for _, cat := range categories {
		prod.Categories = append(prod.Categories, cat.Category)
	}
	//FindProduct := model.ProductModel{}
	//result := l.svcCtx.DB.First(&FindProduct, in.Id)
	//if result.Error != nil {
	//	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//		return nil, status.Error(100, "product not found")
	//	}
	//	return nil, status.Error(500, result.Error.Error())
	//}
	//var categories []model.CategoryModel
	//err := l.svcCtx.DB.Where("product_id = ?", FindProduct.Id).Find(&categories).Error
	//if err != nil {
	//	return nil, status.Error(500, err.Error())
	//}
	//
	//newProduct := product.Product{
	//	Id:          FindProduct.Id,
	//	Name:        FindProduct.Name,
	//	Description: FindProduct.Description,
	//	Picture:     FindProduct.Picture,
	//	Price:       float32(FindProduct.Price),
	//}
	//for _, category := range categories {
	//	newProduct.Categories = append(newProduct.Categories, category.Category)
	//}

	//return &product.GetProductResp{Product: &newProduct}, nil
	return &product.GetProductResp{Product: prod}, nil
}
