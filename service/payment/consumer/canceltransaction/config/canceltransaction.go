package config

import (
	"context"
	"encoding/json"
	"gomall/service/payment/model"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CancelTransactionConsumer struct {
	ctx    context.Context
	config Config
	logx.Logger
	db *gorm.DB
}

func NewCancelTransactionConsumer(db *gorm.DB, ctx context.Context, config Config) *CancelTransactionConsumer {
	return &CancelTransactionConsumer{
		ctx:    ctx,
		config: config,
		Logger: logx.WithContext(ctx),
		db:     db,
	}
}

func (c *CancelTransactionConsumer) Start() error {
	pushConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("transaction_cancel_group"),
		consumer.WithNameServer(c.config.Mq.NameServers),
	)
	if err != nil {
		return err
	}

	err = pushConsumer.Subscribe("transaction_timeout", consumer.MessageSelector{}, c.handleMessage)
	if err != nil {
		return err
	}

	return pushConsumer.Start()
}

func (c *CancelTransactionConsumer) handleMessage(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		var payload struct {
			TransactionId uint64 `json:"transaction_id"`
			UserId        uint32 `json:"user_id"`
		}

		if err := json.Unmarshal(msg.Body, &payload); err != nil {
			logx.Errorf("解析消息失败: %v", err)
			continue
		}

		// 查询交易状态
		var transactionInfo model.PaymentModel
		if err := c.db.First(&transactionInfo, "user_id = ? and transaction_id = ?", payload.UserId, payload.TransactionId).Error; err != nil {
			logx.Errorf("查询交易失败: %v", err)
			continue
		}

		// 如果订单状态仍为未支付，则取消订单并通过事务记录日志
		if transactionInfo.Status == model.PENDING {
			paymentLog := model.PaymentLogModel{
				TransactionId: payload.TransactionId,
				Action:        model.CANCEL,
				Message:       "Transaction Canceled Due to Timeout",
				Status:        model.CANCELED,
			}
			err := c.db.Transaction(func(tx *gorm.DB) error {
				if err := tx.Save(&paymentLog).Error; err != nil {
					return err
				}
				if err := tx.Save(&transactionInfo).Error; err != nil {
					return err
				}
				return nil
			})

			if err != nil {
				logx.Errorf("取消交易失败: %v", err)
				continue
			}
			logx.Infof("交易 %d 已自动取消", payload.TransactionId)
		}
	}
	return consumer.ConsumeSuccess, nil
}
