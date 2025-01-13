package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/service/auth/rpc/authservice"
	"gomall/service/product/api/internal/config"
	"gomall/service/product/api/internal/middleware"
	"gomall/service/product/rpc/productcatalogservice"
)

type ServiceContext struct {
	Config         config.Config
	ProductRpc     productcatalogservice.ProductCatalogService
	AuthRpc        authservice.AuthService
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ProductRpc:     productcatalogservice.NewProductCatalogService(zrpc.MustNewClient(c.ProductRPC)),
		AuthRpc:        authservice.NewAuthService(zrpc.MustNewClient(c.AuthRPC)),
		AuthMiddleware: middleware.NewAuthMiddleware(zrpc.MustNewClient(c.AuthRPC)).Handle,
	}
}
