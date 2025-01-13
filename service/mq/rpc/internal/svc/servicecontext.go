package svc

import (
	"gomall/service/mq/rpc/internal/config"

	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type ServiceContext struct {
	Config   config.Config
	Producer rocketmq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(c.RocketMQ.NameServers),
		producer.WithGroupName("mq_producer_group"),
		producer.WithRetry(2),
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
		producer.WithCreateTopicKey("order_timeout"),
	)
	if err != nil {
		panic(err)
	}

	var startErr error
	for i := 0; i < 3; i++ {
		startErr = p.Start()
		if startErr == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if startErr != nil {
		panic(startErr)
	}

	return &ServiceContext{
		Config:   c,
		Producer: p,
	}
}
