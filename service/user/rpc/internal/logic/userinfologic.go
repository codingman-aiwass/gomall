package logic

import (
	"context"
	"errors"
	"gomall/service/user/model"
	"gorm.io/gorm"

	"gomall/service/user/rpc/internal/svc"
	"gomall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	var userModel model.UserModel
	err := l.svcCtx.DB.Take(&userModel, "id = ?", in.UserId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, errors.New("查询用户信息失败")
	}
	return &user.UserInfoResp{
		UserId: userModel.Id,
		Email:  userModel.Email,
	}, nil
}
