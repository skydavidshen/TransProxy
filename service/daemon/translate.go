package daemon

import (
	"TransProxy/enum"
	"TransProxy/manager"
	"TransProxy/manager/mq"
	"TransProxy/model/business"
	"TransProxy/model/request"
	translatorHandler "TransProxy/service/translator"
	"TransProxy/utils"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync/atomic"
	"time"
)

type Translate struct{}

var translateRetryCount = 0

// 原子性操作，多协程安全
var getTranslationTryCount int64 = 0

// 原子性操作，多协程安全
var doInsertTransItemTryCount int64 = 0

func (t Translate) DoTask() {
	fmt.Println("Do task: Translate...")

	// 使用多少个协程消费待翻译队列Items
	var goroutineCount = manager.TP_SERVER_CONFIG.Handler.TransItemGoroutineCount
	readItems(goroutineCount)
}

func readItems(goCount int) {
	// defer ch.Close() 不能在该方法执行，因为，所有的处理消息队列都是通过goroutine(协程)处理的，所以，readItems方法很快会结束
	// 如果readItems结束之后就ch.Close()，那么，协程处理的业务，就会出问题，channel不存在。

	log.Println("readItems start ...", "retry: ", translateRetryCount)
	translateRetryCount++

	ch := mq.GenChannel()
	go mq.MonitorChannel(ch, func(data interface{}) {
		// close channel导致的错误，这是正常的，此时data == nil
		log.Printf("MonitorChannel communication message: %v", data)
		if translateRetryCount > enum.MonitorRabbitMqRetryMaxCount {
			panic(fmt.Sprintf("MonitorChannel communication error: %v", data))
		} else {
			readItems(goCount)
		}
	})

	var transItemQueue = manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.TransItem
	messages, err := ch.Consume(
		transItemQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Printf("consume get msg error: %s", err)
		return
	}

	chInsert, _ := manager.TP_MQ_RABBIT.Channel()
	for i := 0; i < goCount; i++ {
		go func(i int) {
			manager.TP_LOG.Info(fmt.Sprintf("Goroutine-%d start running ... ", i))
			log.Println(fmt.Sprintf("Goroutine-%d start running ... ", i))

			for msg := range messages { // messages 是一个channel,从中取东西
				manager.TP_LOG.Info(fmt.Sprintf("Goroutine-%d: Received a message: %s", i, string(msg.Body)))
				log.Println(fmt.Sprintf("Goroutine-%d: Received a message: %s", i, string(msg.Body)))

				var item request.Item
				parseErr := json.Unmarshal(msg.Body, &item)
				if parseErr != nil {
					manager.TP_LOG.Info(fmt.Sprintf("parse json error: %v", parseErr))
					_ = msg.Nack(false, true)
					continue
				}

				transItem, transErr := getTranslation(item)
				if transErr != nil {
					_ = msg.Nack(false, false)
					manager.TP_LOG.Info(fmt.Sprintf("Translate item error: %v", transErr))
					continue
				}

				manager.TP_LOG.Info(fmt.Sprintf("transItem: %v", transItem))
				insertErr := doInsertTransItem(chInsert, transItem)
				if insertErr != nil {
					fmt.Println("Insert trans item error: ", insertErr)
					_ = msg.Nack(false, false)

					manager.TP_LOG.Info(fmt.Sprintf("Insert trans item error: %v", insertErr))
					continue
				}

				//手动ack
				errAck := msg.Ack(false) // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
				log.Println("ack success", string(msg.Body), "errAck: ", errAck)
			}
			log.Println("goroutine-readItems done ..., i: ", i)
		}(i)
	}
	log.Println("readItems done ...")
}

func getTranslation(item request.Item) (business.TranslateItem, error) {
	var err error
	for getTranslationTryCount < 5 {
		// 原子性操作变量
		atomic.AddInt64(&getTranslationTryCount, 1)
		transItem, transErr := translatorHandler.TranslateFromItem(item)
		err = transErr
		if transErr != nil {
			// sleep一秒再跑
			time.Sleep(time.Second * 1)
			return getTranslation(item)
		}
		getTranslationTryCount = 0
		return transItem, nil
	}
	getTranslationTryCount = 0
	return business.TranslateItem{}, fmt.Errorf("transErr: %v", err)
}

func doInsertTransItem(ch *amqp.Channel, item business.TranslateItem) error {
	var err error
	for doInsertTransItemTryCount < 5 {
		// 原子性操作变量
		atomic.AddInt64(&doInsertTransItemTryCount, 1)
		insertErr := insertTransItem(ch, item)
		err = insertErr
		if insertErr != nil {
			// sleep一秒再跑
			time.Sleep(time.Second * 1)
			return doInsertTransItem(ch, item)
		}
		doInsertTransItemTryCount = 0
		return nil
	}
	doInsertTransItemTryCount = 0
	return fmt.Errorf("transErr: %v", err)
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
			Expiration: manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Expiration,
			MessageId:  utils.GenUUID(),
		})
	if err != nil {
		manager.TP_LOG.Info(fmt.Sprintf("amqp publish msg fail, err: %s", err))
		return err
	}
	return nil
}
