// 优雅地关机或重启
// 上线代码之前「重新build项目程序」，然后执行 kill -1 {pid}，重新访问web服务，会获取最新代码功能，实现「无缝重启」
// 执行 kill -9 {pid} 会杀死进程
// 面向接口编程，让代码扩展性更强，代码更清晰，高内聚，低耦合

package main

import (
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/manager/server"
	"TransProxy/service"
	"log"
	"os"
)

func main() {
	// init manager
	service.InitManager(os.Args)

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
	if manager.TP_MQ_RABBIT.IsClosed() == false {
		// 初始化MQ
		mq.InitMQ()
		defer mq.Close()
	}

	log.Println("Run web server with endless, server is running...")
	//Run web server with endless
	server.Run()
}