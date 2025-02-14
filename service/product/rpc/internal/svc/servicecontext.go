package svc

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gomall/common/init_db"
	"gomall/service/product/model"
	"gomall/service/product/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDB.AutoMigrate(&model.ProductModel{})
	mysqlDB.AutoMigrate(&model.CategoryModel{})
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
		DB:     mysqlDB,
		RDB:    rdb,
	}
}
