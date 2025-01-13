package logic

import (
	"context"
	"gomall/service/auth/rpc/types/auth"
	"gomall/service/user/api/internal/svc"
	"gomall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutRequest) (resp *types.LogoutResponse, err error) {
	expireTokenResp, err := l.svcCtx.AuthRpc.ExpireTokenByRPC(l.ctx, &auth.ExpireTokenReq{AccessToken: req.AccessToken, RefreshToken: req.RefreshToken})
	if err != nil {
		return nil, err
	}

	return &types.LogoutResponse{
		Res: expireTokenResp.Res,
	}, nil
}
