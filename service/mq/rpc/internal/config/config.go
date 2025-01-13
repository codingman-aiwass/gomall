package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	RocketMQ struct {
		NameServers []string
	}
}
