package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	Salt string
	Mq   zrpc.RpcClientConf
}
