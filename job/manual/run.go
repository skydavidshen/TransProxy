package main

import (
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/service"
	"TransProxy/service/job/manual"
	"fmt"
	"os"
)

// 手动执行脚本：需要手动进行执行，处理需要人工介入的业务
// 面向接口编程 manual.Handler
// 执行脚本，必须在项目根目录下执行如下语句：
// go run job/manual/run.go dead-insert-trans-items

var jobMap = map[string]manual.Handler{
	"dead-insert-trans-items": new(manual.DeadInsertTransItems),
}

func main() {
	args := os.Args
	job := args[1]
	println(job)
	if job == "" {
		fmt.Println("请输入需要执行的job名称.")
	}

	jobObj, ok := jobMap[job]
	if !ok  {
		fmt.Println("请输入正确的job名称，可选项：dead-insert-trans-items")
	}

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

	jobObj.DoTask()
}
