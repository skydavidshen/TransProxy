package main

import (
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/service"
)

func main() {
	// init manager
	service.InitManager()

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
}
