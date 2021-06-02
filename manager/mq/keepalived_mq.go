package mq

import (
	"TransProxy/manager"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"time"
)

// 该文件使用到了多种golang技术实现
// 1. 函数作为函数参数传递
// 2. goroutine 协程使用
// 3. channel使用，并且使用select方法进行调度
// 4. rabbitMQ 链接重试机制
// 5. 使用了很多递归和无限循环，递归和无限循环的控制是关键点

const monitorNotifyCloseTime = 3 //单位 s
const renewConnTime = 3          //单位 s

// NotifyCloseChannelCallBack
// 函数编程：Channel close notify call back function
// data 如果有必要的话，可以封装成结构体
type NotifyCloseChannelCallBack func(data interface{})

func MonitorConn(conn *amqp.Connection) {
	log.Println("MonitorConn start...")
	receiver := make(chan *amqp.Error)
	notifyClose := conn.NotifyClose(receiver)
	for {
		flag := false
		select {
		case e := <-notifyClose:
			log.Printf("amqp.Connection communication error: %s", e.Error())
			RenewConn("notifyClose")
			flag = true
		//default:
		//	log.Println("MonitorConn loop..., IsClosed: ", conn.IsClosed())
		}
		if flag {
			break
		}
		time.Sleep(time.Second * monitorNotifyCloseTime)
	}
	log.Println("MonitorConn done, IsClosed: ", conn.IsClosed())
}

func RenewConn(source string) {
	log.Println("source: ", source)
	log.Println("RenewConn..., IsClosed: ", manager.TP_MQ_RABBIT.IsClosed())
	rabbitMqVHost := manager.TP_SERVER_CONFIG.MQ.RabbitMQ.DefaultVhost
	tpMq, errMq := Amqp(rabbitMqVHost)
	if errMq == nil {
		manager.TP_MQ_RABBIT = tpMq
		log.Println("create new manager.TP_MQ_RABBIT success: ", manager.TP_MQ_RABBIT)
		// 监听rabbitmq connection, 死后重启, 协程执行
		go MonitorConn(manager.TP_MQ_RABBIT)
	} else {
		log.Println("create new manager.TP_MQ_RABBIT error: ", errMq)
		time.Sleep(time.Second * renewConnTime)
		log.Println("RenewConn sleep 3 second.")
		RenewConn(source)
		return
	}
	log.Println("RenewConn success...")
}

func GenChannel() *amqp.Channel {
	conn := manager.TP_MQ_RABBIT
	ch, err := conn.Channel()
	if err != nil {
		log.Println("conn.Channel fail.", zap.Any("error", err))
		time.Sleep(time.Second * 3)
		log.Println("GenChannel wait 3 seconds...")
		return GenChannel()
	}
	log.Println("GenChannel success...")
	return ch
}

// MonitorChannel 用于业务中处理，而非genChannel中
// 另外在业务处理中，如果有必要主动ch.close(), 那么不要使用MonitorChannel()进行监听channel的状态，否则会造成死循环抛出panic
// 例如：doInit()中，执行完doInit后会close channel，则不需要进行监听
// 为了让业务端更加自由的处理MonitorChannel的处理业务，可以在业务端自定义callback方法
func MonitorChannel(ch *amqp.Channel, handle NotifyCloseChannelCallBack) {
	receiver := make(chan *amqp.Error)
	notifyClose := ch.NotifyClose(receiver)
	for {
		flag := false
		select {
		case e := <-notifyClose:
			if e != nil {
				manager.TP_LOG.Error("MonitorChannel communication error:", zap.Any("error", e.Error()))
			} else {
				log.Println("MonitorChannel communication error: ", e)
			}
			handle(e)
			flag = true
		//default:
		//	log.Println("MonitorChannel loop...")
		}
		if flag {
			break
		}
	}
}
