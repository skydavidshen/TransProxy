package daemon

import (
	"TransProxy/enum"
	"TransProxy/manager"
	"TransProxy/model/business"
	"TransProxy/model/request"
	"TransProxy/utils"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 业务消息如果消费失败，会尝试{deadNum}次，如果依然处理失败，放入死信队列

var retryMsg chan amqp.Delivery

// 需要重试的消息队列池子
var retryPool map[string]int

// 进入死信队列的阈值：尝试{deadNum}次
const deadNum = 5

type CallInsertTrans struct {}

func (c CallInsertTrans) DoTask() {
	fmt.Println("Do task: CallInsertTrans...")

	// 给变量分配内存，初始化
	retryPool = make(map[string]int)
	retryMsg = make(chan amqp.Delivery)

	// 使用多少个协程消费待翻译队列Items
	var goroutineCount = manager.TP_SERVER_CONFIG.Handler.CallInsertTransItemGoroutineCount
	callInsertTransItem(goroutineCount)

	// 接收retry message，并检查retry pool
	go receiveDelivery()

	// 定时打印「需要重试的消息队列池子」
	go func() {
		for {
			log.Println("retryPool: " ,retryPool)
			time.Sleep(time.Second * 10)
		}
	}()
}

func receiveDelivery() {
	for msg := range retryMsg {
		if _, ok := retryPool[msg.MessageId]; !ok {
			retryPool[msg.MessageId] = 1
		} else {
			retryPool[msg.MessageId]++
		}

		var err error
		if retryPool[msg.MessageId] >= deadNum {
			err = msg.Nack(false, false)
			if err == nil {
				delete(retryPool, msg.MessageId)
			}
		} else {
			err = msg.Nack(false, true)
		}
		// log err
		if err != nil {
			manager.TP_LOG.Error("receiveDelivery and Nack message fail",
				zap.String("err", err.Error()),
			)
		}
	}
}

func callInsertTransItem(goCount int) {
	var transItemQueue = manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.InsertTransItem
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
				resp, err := SendItem(item)
				manager.TP_LOG.Info(fmt.Sprintf(string(resp)))
				if err != nil {
					manager.TP_LOG.Error("call insert trans item error",
						zap.String("err", err.Error()),
					)
					continue
				}

				check, errCheck := CheckResp(resp)
				manager.TP_LOG.Info(fmt.Sprintf("check: %v ", check))
				if errCheck != nil {
					// 塞入channel
					retryMsg <- msg

					manager.TP_LOG.Error("call insert trans item response error",
						zap.String("err", errCheck.Error()),
					)

					time.Sleep(time.Second * 3)
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

func CheckResp(resp []byte) (bool, error) {
	var callInsertTransResp request.CallInsertTransResp
	err := json.Unmarshal(resp, &callInsertTransResp)
	if err != nil {
		return false, err
	}

	if callInsertTransResp.Code == 0 {
		return true, nil
	} else {
		return false, fmt.Errorf("%s", callInsertTransResp.Msg)
	}
}


// 发送transItem 到 对应的source方，数据安全通过：对称加密
// 加密方式：token = md5(md5(data) + privateKey + timestamp)

func SendItem(item business.TranslateItem) ([]byte, error) {
	var url string
	var privateKey string

	switch item.Source {
	case enum.Source_VAFFLE:
		url = manager.TP_SERVER_CONFIG.ThirdParty.ThirdPartyVaffle.InsertTransItem
		privateKey = manager.TP_SERVER_CONFIG.ThirdParty.ThirdPartyVaffle.PrivateKey
	default:
		url = manager.TP_SERVER_CONFIG.ThirdParty.ThirdPartyVaffle.InsertTransItem
		privateKey = manager.TP_SERVER_CONFIG.ThirdParty.ThirdPartyVaffle.PrivateKey
	}

	itemJson, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	itemJsonStr := string(itemJson)
	timeStamp := time.Now().Unix()

	// token算法: 对称hash加密
	// token = md5(md5(bodyStr) + privateKey + timestamp)
	preStr := fmt.Sprintf("%s%s%s", utils.GetMD5Hash(itemJsonStr), privateKey,
		strconv.Itoa(int(timeStamp)))
	genToken := utils.GetMD5Hash(preStr)

	data := map[string]interface{} {
		"token": genToken,
		"timestamp": strconv.FormatInt(timeStamp,10),
		"data": item,
	}

	resp, _ := utils.DoRequest(data, url, http.MethodPost)
	return resp, nil
}