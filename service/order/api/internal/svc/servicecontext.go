package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/service/auth/rpc/authservice"
	"gomall/service/order/api/internal/config"
	"gomall/service/order/api/internal/middleware"
	"gomall/service/order/rpc/orderservice"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	AuthRpc        authservice.AuthService
	OrderRpc       orderservice.OrderService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(zrpc.MustNewClient(c.AuthRPC)).Handle,
		AuthRpc:        authservice.NewAuthService(zrpc.MustNewClient(c.AuthRPC)),
		OrderRpc:       orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRPC)),
	}
}
