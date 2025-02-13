package logic

import (
	"context"
	"gomall/service/auth/rpc/types/auth"

	"gomall/service/user/api/internal/svc"
	"gomall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyAccessTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyAccessTokenLogic {
	return &VerifyAccessTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyAccessTokenLogic) VerifyAccessToken(req *types.VerifyAccessTokenRequest) (resp *types.VerifyAccessTokenResponse, err error) {
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{
		Token:     req.AccessToken,
		TokenType: "access",
		UserId:    req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.VerifyAccessTokenResponse{
		Res: verifyResp.Res,
	}, nil
}
