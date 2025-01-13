package svc

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gomall/service/auth/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RedisConfig.Host,
		Password: c.RedisConfig.Password,
		DB:       c.RedisConfig.DB,
		PoolSize: c.RedisConfig.PoolSize,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		RDB:    rdb,
	}
}
