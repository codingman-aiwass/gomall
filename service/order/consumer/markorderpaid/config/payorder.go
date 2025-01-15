package config

import (
	"context"
	"encoding/json"
	"gomall/service/order/model"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type MarkOrderPaidConsumer struct {
	ctx    context.Context
	config Config
	logx.Logger
	db *gorm.DB
}

func NewMarkOrderPaidConsumer(db *gorm.DB, ctx context.Context, config Config) *MarkOrderPaidConsumer {
	return &MarkOrderPaidConsumer{
		ctx:    ctx,
		config: config,
		Logger: logx.WithContext(ctx),
		db:     db,
	}
}

func (c *MarkOrderPaidConsumer) Start() error {
	pushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("mark_order_paid_group"),
		consumer.WithNameServer(c.config.Mq.NameServers),
	)
	if err != nil {
		return err
	}

	err = pushConsumer.Subscribe("mark_order_paid", consumer.MessageSelector{}, c.handleMessage)
	if err != nil {
		return err
	}

	return pushConsumer.Start()
}

func (c *MarkOrderPaidConsumer) handleMessage(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		var payload struct {
			OrderId uint32 `json:"order_id"`
		}

		if err := json.Unmarshal(msg.Body, &payload); err != nil {
			logx.Errorf("解析消息失败: %v", err)
			continue
		}

		// 查询订单状态
		var order model.OrderModel
		if err := c.db.First(&order, payload.OrderId).Error; err != nil {
			logx.Errorf("查询订单失败: %v", err)
			continue
		}

		// 如果订单状态仍为未支付，则完成订单
		if order.Status == 0 {
			if err := c.db.Model(&order).Update("status", 1).Error; err != nil {
				logx.Errorf("订单标记已付款失败: %v", err)
				continue
			}
			logx.Infof("订单 %d 标记为已付款", payload.OrderId)
		}
	}
	return consumer.ConsumeSuccess, nil
}
