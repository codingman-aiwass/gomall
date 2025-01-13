package logic

import (
	"context"
	"errors"
	"gomall/common/cryptx"
	"gomall/service/user/model"
	"google.golang.org/grpc/status"

	"gomall/service/user/rpc/internal/svc"
	"gomall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	var userModel model.UserModel
	err := l.svcCtx.DB.Take(&userModel, "email = ?", in.Email).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != userModel.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user.LoginResp{UserId: userModel.Id}, nil
}
