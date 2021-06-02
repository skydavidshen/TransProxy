package manual

import (
	"TransProxy/manager"
	"TransProxy/model/business"
	"TransProxy/service/daemon"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type DeadInsertTransItems struct {}

func (d DeadInsertTransItems) DoTask() {
	// 使用多少个协程消费待翻译队列Items
	var goroutineCount = 10
	callDeadInsertTransItem(goroutineCount)

	fmt.Println("Manual job - DeadInsertTransItems: is running, and it is blocking...")
	b := make(chan bool)
	<- b
}

func callDeadInsertTransItem(goCount int) {
	var transItemQueue = manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.DeadInsertTransItem
	ch, _ := manager.TP_MQ_RABBIT.Channel()
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

	for i := 0; i < goCount; i++ {
		go func(i int) {
			manager.TP_LOG.Info(fmt.Sprintf("Goroutine-%d start running ... ", i))
			for msg := range messages {
				manager.TP_LOG.Info(fmt.Sprintf("transItem: %v", string(msg.Body)))

				var item business.TranslateItem
				_ = json.Unmarshal(msg.Body, &item)
				resp, err := daemon.SendItem(item)
				manager.TP_LOG.Info(fmt.Sprintf(string(resp)))
				if err != nil {
					manager.TP_LOG.Error("call insert trans item error",
						zap.String("err", err.Error()),
					)
					continue
				}

				check, errCheck := daemon.CheckResp(resp)
				manager.TP_LOG.Info(fmt.Sprintf("check: %v ", check))
				if errCheck != nil {
					// 塞入channel
					_ = msg.Nack(false, true)

					manager.TP_LOG.Error("call insert trans item response error",
						zap.String("err", errCheck.Error()),
					)
					continue
				}

				//手动ack
				if check {
					_ = msg.Ack(false) // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
				}
			}
		}(i)
	}
}

