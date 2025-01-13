package logic

import (
	"context"
	"time"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zeromicro/go-zero/core/logx"

	"gomall/service/mq/rpc/internal/svc"
	"gomall/service/mq/rpc/types/mq"
)

type SendDelayMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendDelayMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendDelayMessageLogic {
	return &SendDelayMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendDelayMessageLogic) SendDelayMessage(in *mq.SendDelayMessageReq) (*mq.SendDelayMessageResp, error) {
	p := l.svcCtx.Producer
	msg := primitive.NewMessage(in.Topic, in.Payload)
	msg.WithDelayTimeLevel(l.getDelayTimeLevel(in.DelaySeconds))

	// 添加重试逻辑
	var res *primitive.SendResult
	var err error
	for i := 0; i < 3; i++ {
		res, err = p.SendSync(l.ctx, msg)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if err != nil {
		logx.Errorf("发送延时消息失败: %v", err)
		return nil, err
	}

	return &mq.SendDelayMessageResp{
		MessageId: res.MsgID,
	}, nil
}

// RocketMQ的延时级别: 1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
func (l *SendDelayMessageLogic) getDelayTimeLevel(delaySeconds int64) int {
	switch {
	case delaySeconds <= 1:
		return 1
	case delaySeconds <= 5:
		return 2
	case delaySeconds <= 10:
		return 3
	case delaySeconds <= 30:
		return 4
	case delaySeconds <= 60:
		return 5
	case delaySeconds <= 120:
		return 6
	// ... 可以根据需要添加更多级别
	default:
		return 16 // 30分钟
	}
}
