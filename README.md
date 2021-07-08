# TransProxy
## 是什么？ 
一个翻译代理服务，集成了google，bing，百度等平台，方便集成IP代理库，避免翻译平台的屏蔽。
## 开发语言
Golang
## 技术栈
Gin框架 + RabbitMq + Gorm + Nacos配置中心 + docker容器化部署(支持k8s编排，rancher)
## 基础库
Viper：配置文件处理服务:支持热修改
Zap：日志服务
Gorm：数据库服务
redis：缓存服务
rabbitMQ：消息中间件
validator：请求验证库
