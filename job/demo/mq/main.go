package main

import (
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/service"
	"github.com/mitchellh/mapstructure"
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
	if manager.TP_MQ_RABBIT != nil {
		defer mq.Close()
	}

	conn := manager.TP_MQ_RABBIT
	ch, _ := conn.Channel()
	defer conn.Close()
	defer ch.Close()

	var exchangeMap map[string]string
	_ = mapstructure.Decode(manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange, &exchangeMap)
	for _, exchange := range exchangeMap {
		_ = mq.GenExchange(ch, exchange)
	}

	//_ = mq.GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.TransItem)
	//_ = mq.GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.InsertTransItem)
	//err := mq.GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.DeadInsertTransItem)
	//
	//if err != nil {
	//	manager.TP_LOG.Error("Build Exchange and Queue fail.", zap.Any("error", err))
	//}
	manager.TP_LOG.Info("Create Exchange and Queue successfully.")

}
