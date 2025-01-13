package svc

import (
	"gomall/common/init_db"
	"gomall/service/product/model"
	"gomall/service/product/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDB.AutoMigrate(&model.ProductModel{})
	mysqlDB.AutoMigrate(&model.CategoryModel{})
	return &ServiceContext{
		Config: c,
		DB:     mysqlDB,
	}
}
