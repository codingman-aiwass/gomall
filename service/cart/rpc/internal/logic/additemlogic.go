package logic

import (
	"context"
	"errors"
	"gomall/service/cart/model"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"gomall/service/cart/rpc/internal/svc"
	"gomall/service/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemLogic) AddItem(in *cart.AddItemReq) (*cart.AddItemResp, error) {
	// 如果购物车中已经有该商品，则数量加n
	// 如果购物车中没有该商品，则添加该商品
	// 先查询购物车中是否有该商品
	var cartItem model.CartModel
	err := l.svcCtx.DB.First(&cartItem, "user_id =? and product_id =?", in.UserId, in.Item.ProductId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if in.Item.Quantity <= 0 {
				return nil, status.Error(100, "quantity can not be less than 0")
			}
			// 如果购物车中没有该商品，则添加该商品
			cartItem = model.CartModel{
				UserId:    in.UserId,
				ProductId: in.Item.ProductId,
				Quantity:  in.Item.Quantity,
			}
			err1 := l.svcCtx.DB.Create(&cartItem).Error
			// 数据库错误，插入失败
			if err1 != nil {
				return nil, status.Error(500, err1.Error())
			}
			return &cart.AddItemResp{}, nil
		} else {
			// 数据库错误，查询失败
			return nil, status.Error(500, err.Error())
		}
	}
	// 如果购物车中已经有该商品，则数量加n
	if cartItem.Quantity+in.Item.Quantity < 0 {
		return nil, status.Error(100, "quantity can not be negative")
	}
	cartItem.Quantity += in.Item.Quantity
	err = l.svcCtx.DB.Save(&cartItem).Error
	// 数据库错误，更新失败
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
