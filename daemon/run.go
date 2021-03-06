package main

import (
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/service"
	"TransProxy/service/daemon"
	"log"
	"os"
)

// 配置: 需要注册的daemon实现类
var daemons = []daemon.Handler{
	new(daemon.Translate),
	new(daemon.CallInsertTrans),
}

// daemon脚本，主脚本一直会阻塞，业务goroutine会根据自己实际情况独立coding
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
		defer mq.Close()
	}

	// 执行业务daemon task
	for _, item := range daemons {
		go item.DoTask()
	}

	log.Println("Daemon script is running, and it is blocking...")
	// 阻塞脚本, 等待业务代码执行完成后退出
	block := make(chan bool)
	<-block
}
