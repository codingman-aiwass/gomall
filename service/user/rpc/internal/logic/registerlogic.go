package logic

import (
	"context"
	"errors"
	"gomall/common/cryptx"
	"gomall/service/user/model"
	"gorm.io/gorm"

	"gomall/service/user/rpc/internal/svc"
	"gomall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	var userModel model.UserModel
	err := l.svcCtx.DB.Take(&userModel, "email = ?", in.Email).Error
	if err == nil {
		return nil, errors.New("用户已存在")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := model.UserModel{
			Email:    in.Email,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		err = l.svcCtx.DB.Create(&newUser).Error
		if err != nil {
			return nil, errors.New("用户注册失败")
		}
		id := newUser.Id
		return &user.RegisterResp{
			UserId: id,
		}, nil
	}

	return nil, errors.New("用户注册失败")
}
