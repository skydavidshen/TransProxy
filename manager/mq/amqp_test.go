package mq

import (
	"bytes"
	TPTesting "TransProxy/manager/testing"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/streadway/amqp"
	"log"
	"os"
	"testing"
	"time"
)

const url   = "amqp://transproxy:transproxy@172.100.200.5:30608/"
const vHost = "transproxy"
const queue = "trans-item-1"

func TestConn(t *testing.T) {
	conn, _ := amqp.Dial(url + vHost)
	defer conn.Close()
	convey.Convey("connect rabbitmq", t, func() {
		convey.So(conn, convey.ShouldNotBeNil)
	})
}

func TestSetItem(t *testing.T) {
	conn, _ := amqp.Dial(url + vHost)
	ch, err := conn.Channel()
	defer ch.Close()
	defer conn.Close()
	if err != nil {
		fmt.Printf("amqp create channel fail, err: %s", err)
	}

	body := "david 31 student soft-engineer"
	err = ch.Publish(
		"trans-items",
		"google",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: []byte(body),
		})
	if err != nil {
		fmt.Printf("amqp publish msg fail, err: %s", err)
	}
}

func TestGetItem(t *testing.T) {
	conn, _ := amqp.Dial(url + vHost)
	ch, _ := conn.Channel()
	defer ch.Close()
	defer conn.Close()

	msgs, err := ch.Consume(
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
	go func() {
		for d := range msgs {  // msgs 是一个channel,从中取东西
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte(".")) // 统计d.Body中的"."的个数
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second) // 有几个点就sleep几秒
			log.Printf("Done")

			//手动ack
			d.Ack(false)  // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
		}
	}()
	<- forever
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	TPTesting.NEW().InitConfig()
	os.Exit(m.Run())
}