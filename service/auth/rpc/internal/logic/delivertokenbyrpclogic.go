package logic

import (
	"context"
	"gomall/common/jwtx"
	"time"

	"gomall/service/auth/rpc/internal/svc"
	"gomall/service/auth/rpc/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeliverTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliverTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliverTokenByRPCLogic {
	return &DeliverTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeliverTokenByRPCLogic) DeliverTokenByRPC(in *auth.DeliverTokenReq) (*auth.DeliveryResp, error) {
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.AuthConfig.AccessSecret, time.Now().Unix(), l.svcCtx.Config.AuthConfig.AccessExpire, int64(in.UserId))
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwtx.GetToken(l.svcCtx.Config.AuthConfig.RefreshSecret, time.Now().Unix(), l.svcCtx.Config.AuthConfig.RefreshExpire, int64(in.UserId))
	if err != nil {
		return nil, err
	}
	return &auth.DeliveryResp{Token: accessToken, RefreshToken: refreshToken}, nil
}
