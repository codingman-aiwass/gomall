package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gomall/common/jwtx"
	redis2 "gomall/common/redis"

	"gomall/service/auth/rpc/internal/svc"
	"gomall/service/auth/rpc/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenByRPCLogic {
	return &VerifyTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// VerifyTokenByRPC 验证accessToken是否有效
func (l *VerifyTokenByRPCLogic) VerifyTokenByRPC(in *auth.VerifyTokenReq) (*auth.VerifyResp, error) {
	// 先去数据库检查该值是否存在，存在说明无效
	err := l.svcCtx.RDB.Get(l.ctx, in.Token).Err()
	if err != nil {
		if err != redis.Nil {
			return nil, redis2.RedisError
		} else {
			logx.Info("token is valid")
			err = nil
		}
	} else {
		logx.Info("token is invalid")
		return &auth.VerifyResp{Res: false, UserId: 0}, nil
	}

	var uid uint32

	var exp int64
	if in.TokenType == "access" {
		uid, exp, err = jwtx.ValidateToken(l.svcCtx.Config.AuthConfig.AccessSecret, in.Token)
	} else if in.TokenType == "refresh" {
		uid, exp, err = jwtx.ValidateToken(l.svcCtx.Config.AuthConfig.RefreshSecret, in.Token)
	}
	if err != nil {
		return nil, err
	}
	if err != nil {
		return &auth.VerifyResp{Res: false, UserId: 0}, err
	}

	return &auth.VerifyResp{Res: true, UserId: uid, Exp: exp}, nil
}
