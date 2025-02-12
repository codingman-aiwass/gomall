一、项目概述
1. 项目名称
   字节跳动青训营tiktok_e-commence项目
2. 项目背景
   随着移动互联网的普及和消费者购物习惯的变化，社交电商呈现出蓬勃发展的趋势。抖音作为一款拥有庞大用户群体的短视频社交平台，具有巨大的电商潜力。通过搭建电商平台，抖音可以为用户提供更加丰富的购物体验，同时为商家提供新的销售渠道，实现用户、商家和平台的多赢局面。
3. 项目愿景
   希望同学们可以通过完成这个项目切实实践课程中(视频中)学到的知识点包括但是不限于Go 语言编程，常用框架、数据库、对象存储，服务治理，服务上云等内容，同时对开发工作有更多的深入了解与认识，长远讲能对大家的个人技术成长或视野有启发。
4. 项目目标
   一句话做一个抖音商城。为用户提供便捷、优质的购物环境，满足用户多样化的购物需求，打造一个具有影响力的社交电商平台，提升抖音在电商领域的市场竞争力。
5. 技术栈
- Go - Hertz   -Kitex  -Consul   - OpenTelemetry   - Gorm   -cwgo   -RedisConfig
  -Java -SpringBoot  -Consul | nacos  -Dubbo - sentinel -mybatis -RedisConfig
  使用其他语言以及其他语言对应的技术生态也可以，这里不做任何限制
  二、技术需求
  （一）注册中心集成
1. 服务注册与发现
- 该服务能够与注册中心（如 Consul、Nacos 、etcd 等）进行集成，自动注册服务数据。
  （二）身份认证
1. 登录认证
- 可以使用第三方现成的登录验证框架（CasBin、Satoken等），对请求进行身份验证
- 可配置的认证白名单，对于某些不需要认证的接口或路径，允许直接访问
- 可配置的黑名单，对于某些异常的用户，直接进行封禁处理（可选）
2. 权限认证（高级）
- 根据用户的角色和权限，对请求进行授权检查，确保只有具有相应权限的用户能够访问特定的服务或接口。
- 支持正则表达模式的权限匹配（加分项）
- 支持动态更新用户权限信息，当用户权限发生变化时，权限校验能够实时生效。
  （三）可观测要求
1. 日志记录与监控
- 对服务的运行状态和请求处理过程进行详细的日志记录，方便故障排查和性能分析。
- 提供实时监控功能，能够及时发现和解决系统中的问题。
  （四）可靠性要求（高级）
1. 容错机制
- 该服务应具备一定的容错能力，当出现部分下游服务不可用或网络故障时，能够自动切换到备用服务或进行降级处理。
- 保证下游在异常情况下，系统的整体可用性不会受太大影响，且核心服务可用。
- 服务应该具有一定的流量兜底措施，在服务流量激增时，应该给予一定的限流措施。
  三、功能需求
  认证中心
- 分发身份令牌
- 续期身份令牌（高级）
- 校验身份令牌

用户服务
- 创建用户
- 登录
- 用户登出（可选）
- 删除用户（可选）
- 更新用户（可选）
- 获取用户身份信息
  商品服务
- 创建商品（可选）
- 修改商品信息（可选）
- 删除商品（可选）
- 查询商品信息（单个商品、批量商品）
  购物车服务
- 创建购物车
- 清空购物车
- 获取购物车信息
  订单服务
- 创建订单
- 修改订单信息（可选）
- 订单定时取消（高级）
  结算
- 订单结算
  支付
- 取消支付（高级）
- 定时取消支付（高级）
- 支付
  四、考核方式
- 青训营同学需要根据第二点技术需求设计一个合理且具有一定扩展性的系统架构
- 青训营同学需要根据第三点功能需求设计出完整的库表结构
- 除可选、高级标签标记的接口之外，其余功能必做。项目主要从功能实现完整度、代码质量、服务性能与安全可靠4个维度进行考核，计算规则如下所示，最终分数为所有评分项之和
- 在完成必选需求之后，如果有余力可以选择完成高级标签、可选标签的需求获得额外加分，根据完成情况最多加20分
  评价项
  评分说明
  功能实现
  60分，服务能够正常运行，接口实现完整性，边界情况处理等
  代码质量
  10分，项目结构清晰，代码符合编码规范
  服务性能
  10分，数据表是否设置了合理的索引，处理了常见的性能问题
  安全可靠
  20分，越权等安全问题的防御和处理方式
  五、加分点（高级）
- 代码目录结构分层合理，代码扩展性和可维护性较高，有较好的技术文档，能完美体现开发者技术水平
- 有比较良好的编码规范严格按照技术编码规范进行编码，与此同时，针对业务类需求编写了相对完善的单元测试用例，单元测试框架这里不做任何限制
- 完整服务迁移上云部署，可选常规服务器部署、云托管、Fass等方式，这里不做任何限制
  【参考部署方式 1  】抖音云自托管部署 抖音云
  【参考部署方式 2 】火山引擎部署 火山引擎
  【参考部署方式 3 】常规ECS部署 火山引擎  阿里云
  【参考部署方式 4 】高级Serverless部署 阿里云Serverless 火山引擎Serverless
  六、编码要求
- 在本机搭建运行环境或在云上进行开发都可，这里不做任何限制
  侧重服务端实现，会提前定义好各个功能对应的接口（接口定义推荐使用Protobuf，但不做限制），按说明实现接口即可在客户端中看到运行效果
  服务端最基本的结构只需要服务端程序和数据库即可，服务端程序连接数据库，响应客户端请求完成对应功能。同时需要根据功能，设计合理的数据模型，并创建对应的数据表，其中日志文件等可以保存到本地，这里不做限制
  为了数据库层面的安全考虑建议，建议使用提供ACL控制的云数据库，使用本地数据库也可，这里不做限制
  数据库安装配置说明：MySQL 8.0 version +
  对其他数据库或者其他中间件有了解的同学也可以根据实际情况选择，这里不做限制
  七、编码帮助
- 如何创建一个可运行的Hertz服务
  1） 使用Vscode、Goland 创建一个项目
  [Image]
  2）在根目录下新建main.go文件 并 修改main.go为下面的代码
  package main

import (
"context"

    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
// server.Default() creates a Hertz with recovery middleware.
// If you need a pure hertz, you can use server.New()
h := server.Default()

    h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
        c.String(consts.StatusOK, "Hello hertz!")
    })

    h.Spin()
}

4）基于Hertz 框架 启动Go服务
[Image]
九、技术资料&技术视频
字节跳动基础架构服务框架团队-CloudWeGo技术社区出品的电商项目系列教程 ：项目教学
仓库链接：https://github.com/cloudwego/biz-demo/blob/main/gomall
Git操作教程 Git

goctl rpc protoc .\order\rpc\order.proto --go_out=order/rpc/types --go-grpc_out=order/rpc/types --zrpc_out=order/rpc --proto_path=.\

goctl api go --api .\cart\api\cart.api --dir .\cart\api\

auth: 9006
cart: 9004
mq: 9005
payment: 9003
order: 9002
product: 9001
user: 9000
checkout: 9007

docker exec -it rmqnamesrv bash
sh mqadmin updateTopic -n localhost:9876 -t order_timeout -c DefaultCluster
sh mqadmin updateTopic -n localhost:9876 -t transaction_timeout -c DefaultCluster
sh mqadmin updateTopic -n localhost:9876 -t mark_order_paid -c DefaultCluster
sh mqadmin updateTopic -n localhost:9876 -t mark_order_canceled -c DefaultCluster

docker run --name mysql1 -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=aiwass mysql:8.4
docker run --name redis1 -d -p 6379:6379 redis:7.4.2
docker compose up -d

**etcd开发环境**
docker run -d \
-e ALLOW_NONE_AUTHENTICATION=yes \
-p 2379:2379 \
-p 2380:2380 \
--name etcd1 bitnami/etcd

**生产环境**
docker run -d \
-e ETCD_ROOT_PASSWORD=my_secure_password \
-p 2379:2379 \
-p 2380:2380 \
--name etcd1 bitnami/etcd

rpc auth 的配置：https://juejin.cn/post/7044185614811398174
核心思想是后端在Redis中存入一个hash表，表中的key为api层需要传入的App，value为Token。
api层调用rpc服务时，检查携带值和hash表中记录是否一致。

TODO 将所有的微服务按照依赖顺序启动，每个微服务一个golang容器

导出旧docker容器中的数据库： `docker exec -i mysql1 mysqldump -u root -p go_mall > 0212backup.sql`
将拷贝出的数据库文件拷贝到目标容器中： `docker cp 0212backup.sql gomall_docker-mysql-1:/backup.sql `
手动进入新docker容器 `docker exec -it gomall_docker-mysql-1 bash  `
在新容器中创建数据库后，还原数据库 `mysql -u root -p go_mall < ./backup.sql`

redis 恢复备份 `hset rpc:auth:user userapi 6jKNZbEpYGeUMAifz10gOnmoty3TV`

grafana https://juejin.cn/post/7044509187027501063#heading-12
以 path 维度统计 api 接口的 qps查询 `sum(rate(http_server_requests_duration_ms_count{app="user-api"}[5m])) by (path)`
以 method 维度统计 rpc 接口的qps查询 `sum(rate(rpc_server_requests_duration_ms_count{app="$rpc_app"}[5m])) by (method) `
以 code 维度统计 rpc 接口的状态码 `sum(rate(rpc_server_requests_code_total{app="$rpc_app"}[5m])) by (code)`