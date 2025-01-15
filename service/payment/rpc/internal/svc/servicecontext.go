package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/common/init_db"
	"gomall/service/mq/rpc/types/mq"
	"gomall/service/order/rpc/orderservice"
	"gomall/service/payment/model"
	"gomall/service/payment/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	OrderRpc orderservice.OrderService
	MqRpc    mq.MqClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDB.AutoMigrate(&model.PaymentModel{}, &model.PaymentLogModel{})
	return &ServiceContext{
		Config:   c,
		DB:       mysqlDB,
		OrderRpc: orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRPC)),
		MqRpc:    mq.NewMqClient(zrpc.MustNewClient(c.Mq).Conn()),
	}
}
