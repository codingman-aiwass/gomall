package logic

import (
	"context"
	"fmt"
	"gomall/common/jwtx"
	"gomall/service/auth/rpc/internal/svc"
	"gomall/service/auth/rpc/types/auth"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenByRPCLogic {
	return &RefreshTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenByRPCLogic) RefreshTokenByRPC(in *auth.RefreshTokenReq) (*auth.RefreshTokenResp, error) {
	token, uid, err := jwtx.RenewToken(l.svcCtx.Config.AuthConfig.RefreshSecret, l.svcCtx.Config.AuthConfig.AccessSecret, l.svcCtx.Config.AuthConfig.AccessExpire, in.AccessToken, in.RefreshToken)
	if err != nil {
		return nil, err
	}
	l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("accessToken:%s", token), uid, time.Duration(l.svcCtx.Config.AuthConfig.AccessExpire)*time.Second)

	return &auth.RefreshTokenResp{AccessToken: token}, nil
}
