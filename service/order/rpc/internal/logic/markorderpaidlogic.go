package logic

import (
	"context"
	"gomall/service/order/model"
	"google.golang.org/grpc/status"

	"gomall/service/order/rpc/internal/svc"
	"gomall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkOrderPaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkOrderPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkOrderPaidLogic {
	return &MarkOrderPaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkOrderPaidLogic) MarkOrderPaid(in *order.MarkOrderPaidReq) (*order.MarkOrderPaidResp, error) {
	err := l.svcCtx.DB.Model(&model.OrderModel{}).Where("user_id = ? and id = ?", in.UserId, in.OrderId).Update("status", 1).Error
	if err != nil {
		return nil, status.Error(100, "更新订单状态失败")
	}

	return &order.MarkOrderPaidResp{}, nil
}
