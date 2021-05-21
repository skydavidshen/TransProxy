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
	"TransProxy/utils"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	goLog "log"
	"net/http"
	"strconv"
	"time"
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

	// 使用多少个协程消费待翻译队列Items
	var goroutineCount = manager.TP_SERVER_CONFIG.Handler.CallInsertTransItemGoroutineCount
	callInsertTransItem(goroutineCount)

	// 让main阻塞，不退出
	forever := make(chan bool)
	<-forever
}

func callInsertTransItem(goCount int) {
	var transItemQueue = manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.InsertTransItem
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

	for i := 0; i < goCount; i++ {
		go func(i int) {
			goLog.Printf("Goroutine-%d start running ... \n", i)
			for msg := range messages {
				goLog.Printf("transItem: %v\n", string(msg.Body))

				var item business.TranslateItem
				_ = json.Unmarshal(msg.Body, &item)
				resp, err := sendItem(item)
				goLog.Printf(string(resp))
				if err != nil {
					manager.TP_LOG.Error("call insert trans item error",
						zap.String("err", err.Error()),
					)
					continue
				}

				check, errCheck := checkResp(resp)
				goLog.Printf("check: %v ", check)
				if errCheck != nil {
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

func checkResp(resp []byte) (bool, error) {
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
func sendItem(item business.TranslateItem) ([]byte, error) {
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
