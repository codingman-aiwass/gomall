package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/service/auth/rpc/authservice"
	"gomall/service/checkout/api/internal/config"
	"gomall/service/checkout/api/internal/middleware"
	"gomall/service/checkout/rpc/checkoutservice"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	AuthRpc        authservice.AuthService
	CheckoutRpc    checkoutservice.CheckoutService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(zrpc.MustNewClient(c.AuthRPC)).Handle,

		AuthRpc:     authservice.NewAuthService(zrpc.MustNewClient(c.AuthRPC)),
		CheckoutRpc: checkoutservice.NewCheckoutService(zrpc.MustNewClient(c.CheckoutRPC)),
	}
}
