package logic

import (
	"context"
	"gomall/service/auth/rpc/types/auth"

	"gomall/service/user/api/internal/svc"
	"gomall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAccessTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAccessTokenLogic {
	return &RefreshAccessTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAccessTokenLogic) RefreshAccessToken(req *types.RefreshAccessTokenRequest) (resp *types.RefreshAccessTokenResponse, err error) {
	refreshTokenResp, err := l.svcCtx.AuthRpc.RefreshTokenByRPC(l.ctx, &auth.RefreshTokenReq{
		RefreshToken: req.RefreshToken,
		AccessToken:  req.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	return &types.RefreshAccessTokenResponse{
		AccessToken: refreshTokenResp.AccessToken,
	}, nil
}
