package daemon

import (
	"TransProxy/enum"
	"TransProxy/manager"
	"TransProxy/model/business"
	"TransProxy/model/request"
	translatorHandler "TransProxy/service/translator"
	"TransProxy/utils"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type Translate struct{}

func (t Translate) DoTask() {
	fmt.Println("Do task: Translate...")

	// 使用多少个协程消费待翻译队列Items
	var goroutineCount = manager.TP_SERVER_CONFIG.Handler.TransItemGoroutineCount
	readItems(goroutineCount)
}

func readItems(goCount int) {
	var transItemQueue = manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.TransItem
	ch, _ := manager.TP_MQ_RABBIT.Channel()
	messages, err := ch.Consume(
		transItemQueue.Name,
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

	println()

	chInsert, _ := manager.TP_MQ_RABBIT.Channel()
	for i := 0; i < goCount; i++ {
		go func(i int) {
			manager.TP_LOG.Info(fmt.Sprintf("Goroutine-%d start running ... ", i))

			for msg := range messages { // messages 是一个channel,从中取东西
				manager.TP_LOG.Info(fmt.Sprintf("Goroutine-%d: Received a message: %s", i, string(msg.Body)))
				var item request.Item
				parseErr := json.Unmarshal(msg.Body, &item)
				if parseErr != nil {
					manager.TP_LOG.Info(fmt.Sprintf("parse json error: %v", parseErr))
					_ = msg.Nack(false, true)
					continue
				}

				transItem, transErr := translatorHandler.TranslateFromItem(item)
				if transErr != nil {
					_ = msg.Nack(false, true)

					manager.TP_LOG.Info(fmt.Sprintf("Translate item error: %v", transErr))
					continue
				}

				manager.TP_LOG.Info(fmt.Sprintf("transItem: %v", transItem))
				insertErr := insertTransItem(chInsert, transItem)
				if insertErr != nil {
					fmt.Println("Insert trans item error: ", insertErr)
					_ = msg.Nack(false, true)

					manager.TP_LOG.Info(fmt.Sprintf("Insert trans item error: %v", insertErr))
					continue
				}

				//手动ack
				_ = msg.Ack(false) // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
			}
		}(i)
	}
}

func insertTransItem(ch *amqp.Channel, item business.TranslateItem) error {
	var insertExchange = manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange.InsertTransItems
	body, _ := json.Marshal(item)
	err := ch.Publish(
		insertExchange,
		string(item.Platform),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  enum.ContentType_Json,
			Body:         body,
			// Expiration 单位为 ms，1000ms = 1s
			// 设置了expired可以防止程序本身故障导致重试次数计算不准，就算重试机制失效，通过消息超时也可以将超时消息塞入「死信队列」
			Expiration:   manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Expiration,
			MessageId: utils.GenUUID(),
		})
	if err != nil {
		manager.TP_LOG.Info(fmt.Sprintf("amqp publish msg fail, err: %s", err))
		return err
	}
	return nil
}
