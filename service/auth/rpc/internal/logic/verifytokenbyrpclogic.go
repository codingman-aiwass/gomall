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
	var err error
	if in.TokenType == "access" {
		err = l.svcCtx.RDB.Get(l.ctx, fmt.Sprintf("accessToken:%d:%s", in.UserId, in.Token)).Err()
	} else if in.TokenType == "refresh" {
		err = l.svcCtx.RDB.Get(l.ctx, fmt.Sprintf("refreshToken:%d:%s", in.UserId, in.Token)).Err()
	}

	if err != nil {
		//if err != redis.Nil {
		//	// redis 数据库故障，视为无缓存模式
		//	logx.Errorf("Redis error:%s", err.Error())
		//} else if err == redis.Nil {
		//	// redis 未找到该值，视为无缓存模式
		//	logx.Info("redis not found token, execute check token process later...")
		//	err = nil
		//}
		logx.Info("redis is down or not found token, execute check token process later")
	} else {
		// redis找到token了，直接返回true
		//logx.Info("token is valid")
		return &auth.VerifyResp{Res: true, UserId: in.UserId}, nil
	}
	logx.Infof("start executing token-check.")

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
	logx.Infof("expire time:%d", exp)
	// 记录该token
	if in.TokenType == "access" {
		l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("accessToken:%d:%s", in.UserId, in.Token), 0, time.Duration(exp)*time.Second)
	} else if in.TokenType == "refresh" {
		l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("refreshToken:%d:%s", in.UserId, in.Token), 0, time.Duration(exp)*time.Second)
	}
	return &auth.VerifyResp{Res: true, UserId: uid, Exp: exp}, nil
}
