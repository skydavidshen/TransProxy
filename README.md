# TransProxy

## 是什么？ 
一个翻译代理服务，集成了google，bing，百度等平台，方便集成IP代理库，避免翻译平台的屏蔽。

## 对您的价值
1. 您可以当做一个翻译代理服务使用，可以根据自己的业务二次开发。
2. 您可以当做一个golang学习的项目使用，项目的分层、基础服务的搭建、以及代码规范等都是作者通过借鉴优秀项目和自己深入思考的结果，对您一定有一些启发和借鉴意义。

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

## 如果有兴趣，可以联系交流
微信：shen_da_wei
邮箱：shendawei123@gmail.com
