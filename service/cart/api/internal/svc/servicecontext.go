package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/service/auth/rpc/authservice"
	"gomall/service/cart/api/internal/config"
	"gomall/service/cart/api/internal/middleware"
	"gomall/service/cart/rpc/cartservice"
)

type ServiceContext struct {
	Config         config.Config
	CartRpc        cartservice.CartService
	AuthRpc        authservice.AuthService
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CartRpc:        cartservice.NewCartService(zrpc.MustNewClient(c.CartRPC)),
		AuthRpc:        authservice.NewAuthService(zrpc.MustNewClient(c.AuthRPC)),
		AuthMiddleware: middleware.NewAuthMiddleware(zrpc.MustNewClient(c.AuthRPC)).Handle,
	}
}
