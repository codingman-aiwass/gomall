package casbin_init

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
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
