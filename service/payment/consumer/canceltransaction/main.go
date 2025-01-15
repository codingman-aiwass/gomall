package main

import (
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"gomall/service/payment/consumer/canceltransaction/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 读取配置文件
	var c config.Config
	configPath := "./config/config.yml"
	conf.MustLoad(configPath, &c)

	// 初始化数据库
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}

	// 创建消费者
	ctx := context.Background()
	consumer := config.NewCancelTransactionConsumer(db, ctx, c)

	// 启动消费者
	log.Println("启动取消订单消费者...")
	if err := consumer.Start(); err != nil {
		log.Fatalf("消费者启动失败: %v", err)
	}

	// 保持服务运行
	select {}
}
