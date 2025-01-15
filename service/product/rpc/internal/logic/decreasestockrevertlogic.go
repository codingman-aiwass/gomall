package logic

import (
	"context"

	"gomall/service/product/rpc/internal/svc"
	"gomall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecreaseStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecreaseStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecreaseStockRevertLogic {
	return &DecreaseStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecreaseStockRevertLogic) DecreaseStockRevert(in *product.DecreaseStockReq) (*product.DecreaseStockResp, error) {
	// todo: add your logic here and delete this line

	return &product.DecreaseStockResp{}, nil
}
