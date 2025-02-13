package svc

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gomall/common/init_db"
	"gomall/service/auth/casbin_init"
	"gomall/service/auth/model"
	"gomall/service/auth/rpc/internal/config"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	RDB            *redis.Client
	MySQLDB        *gorm.DB
	CasbinEnforcer *casbin.Enforcer
	WhiteList      []string
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 从etcd读取密钥信息
	// 从环境变量或配置中获取 etcd 连接信息
	endpoints := []string{os.Getenv("ETCD_ENDPOINTS")}
	username := os.Getenv("ETCD_USERNAME")
	password := os.Getenv("ETCD_PASSWORD")
	secretKey := os.Getenv("ETCD_SECRET_KEY")
	logx.Infof("endpoints: %v,username: %s, password: %s, secretKey: %s", endpoints, username, password, secretKey)
	// 如果没有使用 TLS，则 tlsConfig 为 nil
	// var tlsConfig *tls.Config
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
		Username:    username,
		Password:    password,
		//TLS:         tlsConfig,
	})
	if err != nil {
		logx.Errorf("failed to create etcd client: %v", err)
		panic(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Get(ctx, secretKey)
	cancel()
	if err != nil {
		logx.Errorf("failed to get secret from etcd: %v", err)
		panic(err)
	}
	if len(resp.Kvs) == 0 {
		logx.Errorf("secret not found for key: %s", secretKey)
		panic(err)
	}
	secret := string(resp.Kvs[0].Value)
	logx.Infof("Fetched secret: %s", secret)
	keys := strings.Split(secret, "|")
	c.AuthConfig.AccessSecret, c.AuthConfig.RefreshSecret = keys[0], keys[1]

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RedisConfig.Host,
		Password: c.RedisConfig.Password,
		DB:       c.RedisConfig.DB,
		PoolSize: c.RedisConfig.PoolSize,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	mysqlDB := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDB.AutoMigrate(&model.WhiteListModel{})
	whiteListModels := make([]*model.WhiteListModel, 0)
	var whiteList []string
	err = mysqlDB.Find(&whiteListModels).Error
	if err != nil {
		logx.Errorf("load white list models err:%v", err)
	} else {
		for _, single_model := range whiteListModels {
			whiteList = append(whiteList, single_model.Path)
		}
	}

	return &ServiceContext{
		Config:         c,
		RDB:            rdb,
		MySQLDB:        mysqlDB,
		WhiteList:      whiteList,
		CasbinEnforcer: casbin_init.InitCasbin(mysqlDB),
	}
}
