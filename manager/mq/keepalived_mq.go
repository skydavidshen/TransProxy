package mq

import (
	"TransProxy/manager"
	"github.com/streadway/amqp"
	"log"
	"time"
)

const monitorNotifyCloseTime = 3 //单位 s
const renewConnTime = 3 //单位 s

func MonitorConn(conn *amqp.Connection) {
	log.Println("MonitorConn start...")
	receiver := make(chan *amqp.Error)
	notifyClose := conn.NotifyClose(receiver)
	for {
		flag := false
		select {
		case e:= <-notifyClose:
			log.Printf("amqp.Connection communication error: %s", e.Error())
			RenewConn("notifyClose")
			flag = true
		default:
			log.Println("MonitorConn loop..., IsClosed: ", conn.IsClosed())
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
	}
}

func MonitorChannel(ch *amqp.Channel) {
	closeChan := make(chan *amqp.Error)
	notifyClose := ch.NotifyClose(closeChan)
	select {
	case e:= <-notifyClose:
		log.Printf("Channel communication error: %s", e.Error())
	}
}