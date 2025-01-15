package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/common/init_db"
	"gomall/service/checkout/rpc/internal/config"
	"gomall/service/mq/rpc/types/mq"
	"gomall/service/order/rpc/orderservice"
	"gomall/service/payment/rpc/paymentservice"
	"gomall/service/product/rpc/productcatalogservice"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	PaymentRpc     paymentservice.PaymentService
	OrderRpc       orderservice.OrderService
	ProductRpc     productcatalogservice.ProductCatalogService
	DTMServer      string
	ProductService string
	OrderService   string
	PaymentService string
	MqRpc          mq.MqClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)

	return &ServiceContext{
		Config:         c,
		DB:             mysqlDB,
		PaymentRpc:     paymentservice.NewPaymentService(zrpc.MustNewClient(c.PaymentRpc)),
		OrderRpc:       orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc:     productcatalogservice.NewProductCatalogService(zrpc.MustNewClient(c.ProductRpc)),
		DTMServer:      c.Dtm.Server,
		ProductService: c.ProductService,
		OrderService:   c.OrderService,
		PaymentService: c.PaymentService,
		MqRpc:          mq.NewMqClient(zrpc.MustNewClient(c.Mq).Conn()),
	}
}
