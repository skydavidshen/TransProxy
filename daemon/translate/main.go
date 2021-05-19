package main

import (
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/service"
	"fmt"
	goLog "log"
)

const queue = "trans-item-1"

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

	readItems(3)
}

func readItems(goCount int) {
	ch, _ := manager.TP_MQ_RABBIT.Channel()
	messages, err := ch.Consume(
		queue,
		"",
		false,
		true,
		false,
		false,
		nil)
	if err != nil {
		fmt.Printf("consume get msg error: %s", err)
		return
	}

	forever := make(chan bool)
	println()
	for i:=0; i < goCount; i++ {
		go func(i int) {
			for d := range messages { // messages 是一个channel,从中取东西
				goLog.Printf("User-%d: Received a message: %s\n", i, string(d.Body))

				//手动ack
				//_ = d.Ack(false)  // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
			}
		}(i)
	}
	<- forever
}
