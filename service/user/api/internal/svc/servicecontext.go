package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/service/auth/rpc/authservice"
	"gomall/service/user/api/internal/config"
	"gomall/service/user/api/internal/middleware"
	"gomall/service/user/rpc/userservice"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        userservice.UserService
	AuthRpc        authservice.AuthService
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRpc:        userservice.NewUserService(zrpc.MustNewClient(c.UserRPC)),
		AuthRpc:        authservice.NewAuthService(zrpc.MustNewClient(c.AuthRPC)),
		AuthMiddleware: middleware.NewAuthMiddleware(zrpc.MustNewClient(c.AuthRPC)).Handle,
	}
}
