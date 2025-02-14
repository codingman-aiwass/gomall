package logic

import (
	"context"
	"fmt"
	"gomall/common/jwtx"
	"gomall/service/auth/rpc/internal/svc"
	"gomall/service/auth/rpc/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpireTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpireTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpireTokenByRPCLogic {
	return &ExpireTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpireTokenByRPCLogic) ExpireTokenByRPC(in *auth.ExpireTokenReq) (*auth.ExpireTokenResp, error) {
	_, exp, err := jwtx.ValidateToken(l.svcCtx.Config.AuthConfig.AccessSecret, in.AccessToken)
	if err != nil {
		return nil, err
	}
	if exp > 0 {
		l.svcCtx.RDB.Del(l.ctx, fmt.Sprintf("accessToken:%s", in.AccessToken))
	}
	_, exp, err = jwtx.ValidateToken(l.svcCtx.Config.AuthConfig.RefreshSecret, in.RefreshToken)
	if err != nil {
		return nil, err
	}
	if exp > 0 {
		l.svcCtx.RDB.Del(l.ctx, fmt.Sprintf("refreshToken:%s", in.RefreshToken))
	}
	return &auth.ExpireTokenResp{Res: true}, nil
}
