package casbin_init

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"gomall/common/convert"
	"gorm.io/gorm"
	"log"
)

// InitCasbin 初始化 Casbin 的 Enforcer，可从配置文件加载模型和策略，下面是示例
func InitCasbin(db *gorm.DB) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("创建Casbin adapter失败，:%v", err)
	}
	//fmt.Println(os.Getwd())
	enforcer, err := casbin.NewEnforcer("../casbin_init/rbac_model.conf", adapter)
	if err != nil {
		log.Fatalf("创建 Casbin enforcer 失败: %v", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Fatalf("加载策略失败： %v", err)
	}

	return enforcer
}

//////////////////////////
// RocketMQ Watcher 相关实现
//////////////////////////

// RocketMQWatcher 实现了 Casbin 的 Watcher 接口，用 RocketMQ 作为消息中间件。
// 当策略变更时，会调用 Update() 方法发送通知消息；其它实例通过订阅相同主题收到消息后触发回调。
type RocketMQWatcher struct {
	producer rocketmq.Producer
	consumer rocketmq.PushConsumer
	topic    string
	// 回调函数，当接收到更新消息时调用。典型场景下调用 enforcer.LoadPolicy() 重新加载策略。
	callback func(string)
}

// NewRocketMQWatcher 创建一个 RocketMQWatcher 实例
func NewRocketMQWatcher(nameServers []string, topic, groupName string) (*RocketMQWatcher, error) {
	var servers []string
	for _, server := range nameServers {
		ipAddr, err := convert.ResolveNameServer(server)
		if err == nil {
			servers = append(servers, ipAddr)
		} else {
			logx.Error(fmt.Sprintf("resolve name server %s error", server))
		}
	}
	logx.Infof("nameServers: %v", servers)
	// 创建 Producer
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(nameServers),
		producer.WithGroupName(groupName), // 此处 groupName 也可单独为 producer 指定不同的名称
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RocketMQ producer: %v", err)
	}
	err = p.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start RocketMQ producer: %v", err)
	}

	// 创建 Consumer
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(servers),
		consumer.WithGroupName(groupName),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RocketMQ consumer: %v", err)
	}

	watcher := &RocketMQWatcher{
		producer: p,
		consumer: c,
		topic:    topic,
	}

	// 订阅指定主题，收到消息后调用 messageHandler 处理
	err = c.Subscribe(topic, consumer.MessageSelector{}, watcher.messageHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to RocketMQ topic: %v", err)
	}
	err = c.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start RocketMQ consumer: %v", err)
	}

	return watcher, nil
}

// messageHandler 处理收到的 RocketMQ 消息，调用注册的回调函数
func (w *RocketMQWatcher) messageHandler(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		logx.Infof("RocketMQWatcher 收到消息: %s", string(msg.Body))
		if w.callback != nil {
			logx.Info("即将调用回调函数")
			w.callback(string(msg.Body))
		} else {
			logx.Info("回调函数未设置")
		}
	}
	return consumer.ConsumeSuccess, nil
}

// SetUpdateCallback 设置更新回调函数（当收到消息时触发）
func (w *RocketMQWatcher) SetUpdateCallback(callback func(string)) error {
	w.callback = callback
	return nil
}

// Update 发送更新通知到 RocketMQ 主题，通知其它实例重新加载策略
func (w *RocketMQWatcher) Update() error {
	msg := &primitive.Message{
		Topic: w.topic,
		Body:  []byte("casbin_policy_updated"),
	}
	// 发送消息（同步方式）
	res, err := w.producer.SendSync(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("failed to send update message: %v", err)
	}
	log.Printf("RocketMQWatcher 发送更新消息成功, result: %v", res)
	return nil
}

// Close 关闭 Producer 和 Consumer（应用退出时调用）
func (w *RocketMQWatcher) Close() {
	if err := w.producer.Shutdown(); err != nil {
		log.Printf("关闭 RocketMQ producer 失败: %v", err)
	}
	if err := w.consumer.Shutdown(); err != nil {
		log.Printf("关闭 RocketMQ consumer 失败: %v", err)
	}
}

//////////////////////////
// 集成 Watcher 到 Casbin Enforcer 中
//////////////////////////

// InitCasbinWithWatcher 初始化 Casbin Enforcer，并集成基于 RocketMQ 的 Watcher
func InitCasbinWithWatcher(db *gorm.DB, nameServers []string) (*casbin.Enforcer, *RocketMQWatcher) {
	// 创建 Casbin Adapter（以 Gorm Adapter 为例）
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("创建 Casbin adapter 失败: %v", err)
	}

	// 使用配置文件加载模型（rbac_model.conf 路径根据实际情况调整）
	enforcer, err := casbin.NewEnforcer("/usr/src/code/service/auth/casbin_init/rbac_model.conf", adapter)
	if err != nil {
		log.Fatalf("创建 Casbin enforcer 失败: %v", err)
	}

	// 加载当前策略
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("加载策略失败： %v", err)
	}

	// 初始化 RocketMQ Watcher
	//nameServer := []string{"127.0.0.1:9876"} // RocketMQ NameServer 地址
	var servers []string
	for _, server := range nameServers {
		ipAddr, err := convert.ResolveNameServer(server)
		if err == nil {
			servers = append(servers, ipAddr)
		} else {
			logx.Error(fmt.Sprintf("resolve name server %s error", server))
		}
	}
	logx.Infof("nameServers: %v", servers)
	topic := "CasbinPolicyUpdate"     // 自定义主题名称
	groupName := "CasbinWatcherGroup" // 消费者组名称（生产者也可共用此组名）
	watcher, err := NewRocketMQWatcher(servers, topic, groupName)
	if err != nil {
		log.Fatalf("创建 RocketMQ Watcher 失败: %v", err)
	}

	// 设置更新回调：当收到更新通知时，重新加载 Casbin 策略
	watcher.SetUpdateCallback(func(message string) {
		log.Printf("Watcher 收到更新通知: %s", message)
		if err := enforcer.LoadPolicy(); err != nil {
			log.Printf("重新加载策略失败: %v", err)
		} else {
			log.Printf("Casbin 策略已重新加载")
		}
	})

	// 将 Watcher 注册到 Enforcer 中
	enforcer.SetWatcher(watcher)
	return enforcer, watcher
}
