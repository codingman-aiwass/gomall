package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	PaymentRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	Dtm        struct {
		Server string
	}
	ProductService string
	OrderService   string
	PaymentService string
}
