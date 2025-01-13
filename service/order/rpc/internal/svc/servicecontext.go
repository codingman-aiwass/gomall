package svc

import (
	"gomall/common/init_db"
	"gomall/service/mq/rpc/types/mq"
	"gomall/service/order/model"
	"gomall/service/order/rpc/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	MqRpc  mq.MqClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDB.AutoMigrate(&model.AddressModel{}, &model.OrderModel{}, &model.OrderItemModel{})

	return &ServiceContext{
		Config: c,
		DB:     mysqlDB,
		MqRpc:  mq.NewMqClient(zrpc.MustNewClient(c.Mq).Conn()),
	}
}
