package service

import (
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/config"
	"TransProxy/manager/db"
	"TransProxy/manager/log"
	"TransProxy/manager/mq"
	"TransProxy/utils"
	"github.com/go-playground/validator/v10"
)

func InitManager() {
	//项目根目录
	manager.TP_ROOT_DIR = utils.GetRootDir()
	//配置文件处理服务:支持热修改
	manager.TP_CONFIG = config.Viper()
	//日志服务
	manager.TP_LOG = log.Zap()
	//数据库服务
	manager.TP_DB = db.Gorm()
	//缓存 - redis服务
	manager.TP_CACHE_REDIS = cache.Redis()
	//消息中间件 - rabbitMQ服务
	rabbitMqVHost := manager.TP_SERVER_CONFIG.MQ.RabbitMQ.DefaultVhost
	manager.TP_MQ_RABBIT = mq.Amqp(rabbitMqVHost)
	//请求验证库
	manager.TP_VALIDATE = validator.New()
}
