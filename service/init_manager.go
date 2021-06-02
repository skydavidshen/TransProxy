package service

import (
	"TransProxy/enum"
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/config"
	"TransProxy/manager/db"
	"TransProxy/manager/log"
	"TransProxy/manager/mq"
	"TransProxy/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

func InitManager(args []string) {
	manager.TP_ENV = enum.Env_Dev
	if len(os.Args) > 1 {
		// 设置当前执行环境
		switch args[1] {
		case enum.Env_Dev:
			manager.TP_ENV = enum.Env_Dev
		case enum.Env_Prod:
			manager.TP_ENV = enum.Env_Prod
		default:
			manager.TP_ENV = enum.Env_Dev
		}
	}

	//项目根目录
	manager.TP_ROOT_DIR = utils.GetRootDir()
	//配置文件处理服务:支持热修改
	tpConf, errConf := config.Viper()
	if errConf == nil {
		manager.TP_CONFIG = tpConf
	} else {
		panic(fmt.Errorf("manager.TP_CONFIG init error: %v", errConf))
	}

	//日志服务
	manager.TP_LOG = log.Zap()
	if manager.TP_LOG == nil {
		panic(fmt.Sprint("manager.TP_LOG init error."))
	}

	//数据库服务
	tpDB, errDb := db.Gorm()
	if errDb == nil {
		manager.TP_DB = tpDB
	} else {
		panic(fmt.Errorf("manager.TP_DB init error: %v", errDb))
	}

	//缓存 - redis服务
	tpRedis, errRedis := cache.Redis()
	if errRedis == nil {
		manager.TP_CACHE_REDIS = tpRedis
	} else {
		panic(fmt.Errorf("TP_CACHE_REDIS init error: %v", errRedis))
	}

	//消息中间件 - rabbitMQ服务
	rabbitMqVHost := manager.TP_SERVER_CONFIG.MQ.RabbitMQ.DefaultVhost
	tpMq, errMq := mq.Amqp(rabbitMqVHost)
	if errMq == nil {
		manager.TP_MQ_RABBIT = tpMq
		if manager.TP_MQ_RABBIT.IsClosed() {
			panic(fmt.Errorf("TP_MQ_RABBIT connection is closed: %v", errMq))
		}
	} else {
		panic(fmt.Errorf("TP_MQ_RABBIT init error: %v", errMq))
	}
	// 监听rabbitmq connection, 死后重启, 协程执行
	go mq.MonitorConn(manager.TP_MQ_RABBIT)

	//请求验证库
	manager.TP_VALIDATE = validator.New()
}
