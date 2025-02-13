package test

import (
	"gomall/common/init_db"
	"gomall/service/auth/casbin_init"
	"testing"
	"time"
)

func TestCasbinWatcherUpdate(t *testing.T) {
	mysqlDB := init_db.InitGorm("root:aiwass@tcp(mysql:3306)/go_mall?charset=utf8mb4&parseTime=True&loc=Local")
	_, watcher := casbin_init.InitCasbinWithWatcher(mysqlDB, []string{"namesrv:9876"})
	// 定义一个 channel，用于接收回调触发信号
	callbackCh := make(chan struct{})

	// 重写 CasbinWatcher 的回调函数：收到消息时打印日志并发送信号

	err := watcher.SetUpdateCallback(func(message string) {
		t.Logf("回调函数收到更新消息：%s", message)
		// 通知测试回调已触发
		close(callbackCh)
	})
	if err != nil {
		t.Fatalf("设置更新回调失败: %v", err)
	}

	// 调用 Update 发送更新消息到 RocketMQ
	err = watcher.Update()
	if err != nil {
		t.Fatalf("发送更新消息失败: %v", err)
	}

	// 等待回调触发，超时时间设置为 5 秒
	select {
	case <-callbackCh:
		t.Log("测试通过：回调成功触发")
	case <-time.After(5 * time.Second):
		t.Error("测试失败：5秒内未收到更新回调")
	}
}
