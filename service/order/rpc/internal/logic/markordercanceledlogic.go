package logic

import (
	"context"
	"gomall/service/order/model"
	"google.golang.org/grpc/status"

	"gomall/service/order/rpc/internal/svc"
	"gomall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkOrderCanceledLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkOrderCanceledLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkOrderCanceledLogic {
	return &MarkOrderCanceledLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkOrderCanceledLogic) MarkOrderCanceled(in *order.MarkOrderCanceledReq) (*order.MarkOrderCanceledResp, error) {
	err := l.svcCtx.DB.Model(&model.OrderModel{}).Where("user_id = ? and id = ?", in.UserId, in.OrderId).Update("status", 2).Error
	if err != nil {
		return nil, status.Error(100, "更新订单状态失败")
	}

	return &order.MarkOrderCanceledResp{}, nil
}
