package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gomall/common/jwtx"
	"gomall/service/auth/rpc/internal/svc"
	"gomall/service/auth/rpc/types/auth"
	"strconv"
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
	// 先去数据库检查该值是否存在，不存在说明无效
	var err error
	var hashKey string
	if in.TokenType == "access" {
		hashKey = fmt.Sprintf("accessToken:%s", in.Token)
	} else if in.TokenType == "refresh" {
		hashKey = fmt.Sprintf("refreshToken:%s", in.Token)
	}
	result := l.svcCtx.RDB.Get(l.ctx, hashKey)
	err = result.Err()
	if err != nil {
		if err != redis.Nil {
			// redis 数据库故障，视为无缓存模式，进行token的检查
			logx.Errorf("Redis error:%s", err.Error())
		} else if err == redis.Nil {
			// redis 未找到该值，说明已被删除
			logx.Errorf("Redis error:%s", err.Error())
			logx.Infof("redis did not find hashKey:%s", hashKey)
			return &auth.VerifyResp{Res: false}, err
		}
	} else {
		// redis找到token了，直接返回true
		logx.Info("token is valid")
		uid, _ := strconv.ParseUint(result.Val(), 10, 32)
		return &auth.VerifyResp{Res: true, UserId: uint32(uid)}, nil
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
		return &auth.VerifyResp{Res: false}, err
	}
	logx.Infof("expire time:%d", exp)
	// 记录该token
	if in.TokenType == "access" {
		l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("accessToken:%s", in.Token), 0, time.Duration(exp)*time.Second)
	} else if in.TokenType == "refresh" {
		l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("refreshToken:%s", in.Token), 0, time.Duration(exp)*time.Second)
	}
	return &auth.VerifyResp{Res: true, Exp: exp, UserId: uid}, nil
}
