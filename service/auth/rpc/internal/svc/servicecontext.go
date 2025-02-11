package svc

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"gomall/common/init_db"
	"gomall/service/auth/casbin_init"
	"gomall/service/auth/model"
	"gomall/service/auth/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	RDB            *redis.Client
	MySQLDB        *gorm.DB
	CasbinEnforcer *casbin.Enforcer
	WhiteList      []string
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

	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDB.AutoMigrate(&model.WhiteListModel{})
	whiteListModels := make([]*model.WhiteListModel, 0)
	var whiteList []string
	err = mysqlDB.Find(&whiteListModels).Error
	if err != nil {
		logx.Errorf("load white list models err:%v", err)
	} else {
		for _, single_model := range whiteListModels {
			whiteList = append(whiteList, single_model.Path)
		}
	}

	return &ServiceContext{
		Config:         c,
		RDB:            rdb,
		MySQLDB:        mysqlDB,
		WhiteList:      whiteList,
		CasbinEnforcer: casbin_init.InitCasbin(mysqlDB),
	}
}
