package logic

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"

	"gomall/service/mq/rpc/internal/svc"
	"gomall/service/mq/rpc/types/mq"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SendMessage 发送普通消息
func (l *SendMessageLogic) SendMessage(in *mq.SendMessageReq) (*mq.SendMessageResp, error) {
	p := l.svcCtx.Producer
	msg := primitive.NewMessage(in.Topic, in.Payload)

	// 如果 Properties 字段非空，则添加消息属性
	if in.Properties != nil && len(in.Properties) > 0 {
		for key, value := range in.Properties {
			msg.WithProperty(key, value)
		}
	}

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
		logx.Errorf("发送普通消息失败: %v", err)
		return nil, err
	}

	return &mq.SendMessageResp{
		MessageId: res.MsgID,
	}, nil
}
