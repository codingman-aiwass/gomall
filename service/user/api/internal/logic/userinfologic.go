package logic

import (
	"context"
	"gomall/service/user/rpc/types/user"

	"gomall/service/user/api/internal/svc"
	"gomall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	//uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	uid := l.ctx.Value("userId").(int64)
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		UserId: uint32(uid),
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		Id: res.UserId,
	}, nil
}
