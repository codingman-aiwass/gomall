package logic

import (
	"context"
	"gomall/service/auth/rpc/types/auth"
	"gomall/service/user/api/internal/svc"
	"gomall/service/user/api/internal/types"
	"gomall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	deliverResp, err := l.svcCtx.AuthRpc.DeliverTokenByRPC(l.ctx, &auth.DeliverTokenReq{
		UserId:      res.UserId,
		ServiceName: l.svcCtx.Config.Name,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccessToken:  deliverResp.Token,
		RefreshToken: deliverResp.RefreshToken,
	}, nil
}
