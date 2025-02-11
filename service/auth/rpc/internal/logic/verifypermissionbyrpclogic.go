package logic

import (
	"context"
	"errors"
	"fmt"
	"gomall/service/auth/rpc/internal/svc"
	"gomall/service/auth/rpc/types/auth"
	"gomall/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPermissionByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyPermissionByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPermissionByRPCLogic {
	return &VerifyPermissionByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyPermissionByRPCLogic) VerifyPermissionByRPC(in *auth.VerifyPermissionReq) (*auth.VerifyPermissionResp, error) {
	// 需要根据uid去user_roles中获取用户角色id,再到role_models中获取角色名
	userRole := &model.UserRole{}
	err := l.svcCtx.MySQLDB.First(userRole, "user_id = ?", in.UserId).Error
	if err != nil {
		logx.Errorf("get userrole err:%v", err)
		return nil, err
	}

	roleModel := &model.RoleModel{}
	err = l.svcCtx.MySQLDB.First(roleModel, "id = ?", userRole.RoleID).Error
	if err != nil {
		logx.Errorf("get role info err:%v", err)
		return nil, err
	}

	// 进行权限校验
	ok, err := l.svcCtx.CasbinEnforcer.Enforce(roleModel.Name, in.Resource, in.Method)
	if err != nil {
		logx.Errorf("casbin enforce err:%v", err)
		return nil, err
	}
	if !ok {
		return nil, errors.New(fmt.Sprintf("用户ID: %v 无权限通过 %v 访问 %v", userRole.UserID, in.Method, in.Resource))
	}

	return &auth.VerifyPermissionResp{}, nil
}
