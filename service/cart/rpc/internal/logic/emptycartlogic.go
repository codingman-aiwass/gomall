package logic

import (
	"context"
	"gomall/service/cart/model"
	"google.golang.org/grpc/status"

	"gomall/service/cart/rpc/internal/svc"
	"gomall/service/cart/rpc/types/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmptyCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyCartLogic {
	return &EmptyCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmptyCartLogic) EmptyCart(in *cart.EmptyCartReq) (*cart.EmptyCartResp, error) {
	// 根据传入的userID，删除cartmodel表下，所有的cartitem
	err := l.svcCtx.DB.First(&model.CartModel{}, "user_id = ?", in.UserId).Error
	if err != nil {
		return nil, status.Error(100, "cart is empty, no need to empty again.")
	}
	err = l.svcCtx.DB.Delete(&model.CartModel{}, "user_id = ?", in.UserId).Error

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &cart.EmptyCartResp{}, nil
}
