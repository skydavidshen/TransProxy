// 优雅地关机或重启
// 上线代码之前「重新build项目程序」，然后执行 kill -1 {pid}，重新访问web服务，会获取最新代码功能，实现「无缝重启」
// 执行 kill -9 {pid} 会杀死进程

package main

import (
	"com.pippishen/trans-proxy/manager"
	"com.pippishen/trans-proxy/manager/cache"
	"com.pippishen/trans-proxy/manager/config"
	"com.pippishen/trans-proxy/manager/db"
	"com.pippishen/trans-proxy/manager/log"
	"com.pippishen/trans-proxy/manager/mq"
	"com.pippishen/trans-proxy/manager/server"
	"com.pippishen/trans-proxy/utils"
	"fmt"
)

func main() {
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
	rabbitMqVHost := manager.TP_CONFIG.Get("mq.rabbitmq.default-vhost").(string)
	manager.TP_MQ_RABBIT = mq.Amqp(rabbitMqVHost)

	//release db
	if manager.TP_DB != nil {
		//main函数结束之前关闭资源
		defer db.Close()

		//初始化表和数据
		db.InitDB()
	}

	//release cache
	if manager.TP_CACHE_REDIS != nil {
		defer cache.Close()
	}

	//release mq
	if manager.TP_MQ_RABBIT != nil {
		defer mq.Close()
	}

	fmt.Printf("\nServer conf: %v\n", manager.TP_SERVER_CONFIG)

	//Run web server with endless
	server.Run()
}