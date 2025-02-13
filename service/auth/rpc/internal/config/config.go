package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	AuthConfig struct {
		AccessSecret string
		AccessExpire int64

		RefreshSecret string
		RefreshExpire int64
	}

	RedisConfig struct {
		Host     string
		Password string
		DB       int
		PoolSize int
	}

	Mysql struct {
		DataSource string
	}
	RocketMQ struct {
		NameServers []string
	}
}
