package logic

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gomall/service/product/model"
	"google.golang.org/grpc/status"
	"strconv"
	"time"

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
	// 构造 Redis Sorted Set key，存放某个类别下的商品 ID 列表
	redisKey := fmt.Sprintf("category:%s:products", in.CategoryName)
	// 使用游标分页。如果 LastId 为 0，则表示从头开始；否则排除等于 LastId 的记录
	var minScore string
	if in.LastId == 0 {
		minScore = "-inf"
	} else {
		// 注意：前缀 "(" 表示排除等于该值的成员（即 id 必须大于 LastId）
		minScore = fmt.Sprintf("(%d", in.LastId)
	}
	// 尝试从 Redis Sorted Set 中获取商品 ID 列表，使用 ZRANGEBYSCORE 命令进行分页
	redisIDs, err := l.svcCtx.RDB.ZRangeByScore(l.ctx, redisKey, &redis.ZRangeBy{
		Min:   minScore,
		Max:   "+inf",
		Count: in.PageSize,
	}).Result()
	ids := []uint32{}
	if err != nil || len(redisIDs) == 0 {
		// 如果 Redis 中没有数据或发生错误，则从数据库中查询
		var categories []model.CategoryModel
		var err1 error
		// 先去category中根据种类查询出product_id

		// lastId 不为0且page!=1；lastId为0但是page=1；lastId不为0但是page=1
		err1 = l.svcCtx.DB.
			Where("category = ? AND product_id > ?", in.CategoryName, in.LastId).
			Order("product_id asc").
			Limit(int(in.PageSize)).Find(&categories).Error

		if err1 != nil {
			return nil, status.Error(500, err1.Error())
		}
		// 再根据product_id去product中查询出product
		for _, category := range categories {
			ids = append(ids, category.ProductId)
			// 使用类别表的product_id 作为 score 保证顺序
			l.svcCtx.RDB.ZAdd(l.ctx, redisKey, redis.Z{
				Score:  float64(category.ProductId),
				Member: fmt.Sprintf("%d", category.ProductId),
			})
		}
	} else {
		// 将 Redis 返回的字符串 ID 转换为 uint32 类型
		for _, idStr := range redisIDs {
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				ids = append(ids, uint32(id))
			}
		}
	}

	// 根据获取到的商品 ID 批量查询商品详情
	// 优先尝试从 Redis Hash 中获取，若缓存未命中则查询数据库，并将结果缓存
	var respProducts []*product.Product
	for _, pid := range ids {
		hashKey := fmt.Sprintf("product:%d", pid)
		productData, err := l.svcCtx.RDB.HGetAll(l.ctx, hashKey).Result()
		var prod *product.Product
		if err == nil && len(productData) > 0 {
			// 从Redis缓存中解析数据
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
			// 缓存未命中则从数据库中查询商品详情
			var dbProduct model.ProductModel
			if err := l.svcCtx.DB.First(&dbProduct, pid).Error; err != nil {
				// 如果商品不存在则跳过（或者根据业务需求返回错误）
				continue
			}
			prod = &product.Product{
				Id:          dbProduct.Id,
				Name:        dbProduct.Name,
				Description: dbProduct.Description,
				Picture:     dbProduct.Picture,
				Price:       float32(dbProduct.Price),
			}
			// 将查询到的商品详情写入 Redis Hash，并设置适当的过期时间
			l.svcCtx.RDB.HMSet(l.ctx, hashKey, map[string]interface{}{
				"id":          prod.Id,
				"name":        prod.Name,
				"description": prod.Description,
				"picture":     prod.Picture,
				"price":       prod.Price,
			})
			l.svcCtx.RDB.Expire(l.ctx, hashKey, time.Hour)
		}
		respProducts = append(respProducts, prod)
	}

	return &product.ListProductsResp{Products: respProducts}, nil
}
