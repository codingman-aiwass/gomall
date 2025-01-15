package main

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConsumerIntegration(t *testing.T) {
	// 启动生产者
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithGroupName("test_group"),
		producer.WithRetry(2),
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
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

	// 发送测试消息
	topic := "mark_order_canceled"

	payload := struct {
		OrderId uint32 `json:"order_id"`
		UserId  uint32 `json:"user_id"`
	}{
		OrderId: 73,
		UserId:  2,
	}
	payloadBytes, _ := json.Marshal(payload)
	msg := primitive.NewMessage(topic, payloadBytes)

	res, err := p.SendSync(context.Background(), msg)
	assert.NoError(t, err)
	t.Logf("消息发送成功: %v", res)
}
