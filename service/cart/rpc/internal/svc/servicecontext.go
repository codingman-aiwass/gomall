package svc

import (
	"gomall/common/init_db"
	"gomall/service/cart/model"
	"gomall/service/cart/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDB.AutoMigrate(&model.CartModel{})
	return &ServiceContext{
		Config: c,
		DB:     mysqlDB,
	}
}
