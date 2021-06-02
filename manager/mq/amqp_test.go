package mq

import (
	"TransProxy/manager"
	TPTesting "TransProxy/manager/testing"
	"bytes"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/smartystreets/goconvey/convey"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"os"
	"testing"
	"time"
)

const url = "amqp://transproxy:transproxy@172.100.200.5:30608/"
const vHost = "transproxy"
const queue = "trans-item-1"

func TestConn(t *testing.T) {
	conn, _ := amqp.Dial(url + vHost)
	defer conn.Close()
	convey.Convey("connect rabbitmq", t, func() {
		convey.So(conn, convey.ShouldNotBeNil)
	})
}

func TestMap(t *testing.T) {
	var exchangeMap map[string]string
	_ = mapstructure.Decode(manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange, &exchangeMap)
	fmt.Println(exchangeMap)
}

func TestExchangeDeclare(t *testing.T) {
	conn, _ := amqp.Dial(url + vHost)
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	err = ch.ExchangeDeclare(
		"test.trans-items",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	q, _ := ch.QueueDeclare(
		"test.trans-item-1",
		true,
		false,
		false, //exclusive为true: 连接关闭时会被删除，所以一般设为false
		false,
		nil,
	)

	err = ch.QueueBind(
		q.Name,
		"google",
		"test.trans-items",
		false,
		nil,
	)
	err = ch.QueueBind(
		q.Name,
		"bing",
		"test.trans-items",
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("create exchange and queue successfully.")
}

func TestSetItem(t *testing.T) {
	for i := 0; i < 1000; i++ {
		setItem(i)
	}
}

// 消息确认：发布消息消息确认机制
func setItem(i int) {
	conn, _ := amqp.Dial(url + vHost)
	ch, err := conn.Channel()
	defer ch.Close()
	defer conn.Close()
	if err != nil {
		fmt.Printf("amqp create channel fail, err: %s", err)
	}

	body := fmt.Sprintf("david 333 student soft-engineer: %d", i)
	_ = ch.Confirm(false)
	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1)) // 处理确认逻辑
	defer confirmOne(confirms, body)                              // 处理方法

	err = ch.Publish(
		"acktest",
		"acktest",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	log.Printf("amqp publish msg fail, err: %s", err)
	if err != nil {
		log.Printf("amqp publish msg fail, err: %s", err)
	}
	log.Printf("david 333 student soft-engineer: %d\n", i)

	time.Sleep(time.Second * 3)
}

// 消息确认
func confirmOne(confirms <-chan amqp.Confirmation, body string) {
	if confirmed := <-confirms; confirmed.Ack {
		fmt.Printf("confirmed delivery with ack confirmed: %d\n", confirmed.DeliveryTag)
		manager.TP_LOG.Info("confirmed delivery with ack confirmed\n", zap.Uint64("tag", confirmed.DeliveryTag),
			zap.String("body", body))
	} else {
		fmt.Printf("confirmed delivery no: %d\n", confirmed.DeliveryTag)
		manager.TP_LOG.Info("confirmed delivery no\n", zap.Uint64("tag", confirmed.DeliveryTag),
			zap.String("body", body))
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
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Printf("consume get msg error: %s", err)
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs { // msgs 是一个channel,从中取东西
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte(".")) // 统计d.Body中的"."的个数
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second) // 有几个点就sleep几秒
			log.Printf("Done")

			//手动ack
			d.Ack(false) // 手动ACK，如果不ACK的话，那么无法保证这个消息被处理，可能它已经丢失了（比如消息队列挂了）
		}
	}()
	<-forever
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	TPTesting.NEW().InitConfig()
	os.Exit(m.Run())
}
