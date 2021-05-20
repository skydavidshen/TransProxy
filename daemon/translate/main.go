package main

import (
	"TransProxy/enum"
	"TransProxy/manager"
	"TransProxy/manager/cache"
	"TransProxy/manager/db"
	"TransProxy/manager/mq"
	"TransProxy/model/business"
	"TransProxy/model/request"
	"TransProxy/service"
	transPlatform "TransProxy/service/trans-platform"
	translatorHandler "TransProxy/service/translator"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	goLog "log"
)

const transItemQueue = "trans-item-1"
const insertExchange = "insert-trans-items"

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
		transItemQueue,
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
			for msg := range messages { // messages 是一个channel,从中取东西
				goLog.Printf("User-%d: Received a message: %s\n", i, string(msg.Body))
				var item request.Item
				parseErr := json.Unmarshal(msg.Body, &item)
				if parseErr != nil {
					goLog.Printf("parse json error: %v\n", parseErr)
					continue
				}

				var translator translatorHandler.Handler
				// 使用哪个翻译平台可以根据实际情况替换
				var platform transPlatform.Handler = transPlatform.SmartProxy{}
				switch item.Platform {
				case enum.Platform_Google:
					translator = &translatorHandler.Google{PlatformHandler: platform}
				case enum.Platform_Bing:
					translator = &translatorHandler.Bing{PlatformHandler: platform}
				}
				transItem, _ := translator.Translate(item)
				
				goLog.Printf("transItem: %v\n", transItem)

				insertErr := insertTransItem(transItem)
				if insertErr != nil {
					goLog.Printf("Insert trans item error: %v\n", insertErr)
					continue
				}

				//手动ack
				_ = msg.Ack(false) // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
			}
		}(i)
	}
	<- forever
}

func insertTransItem(item business.TranslateItem) error {
	ch, _ := manager.TP_MQ_RABBIT.Channel()
	body, _ := json.Marshal(item)
	err := ch.Publish(
		insertExchange,
		string(item.Platform),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  translatorHandler.ContentType,
			Body:         body,
		})
	if err != nil {
		goLog.Printf("amqp publish msg fail, err: %s", err)
		return err
	}
	return nil
}
