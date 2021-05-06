package mq

import (
	TPTesting "com.pippishen/trans-proxy/manager/testing"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/streadway/amqp"
	"os"
	"testing"
)

const url   = "amqp://transproxy:transproxy@172.100.200.5:30608/"
const vhost = "transproxy"

func TestConn(t *testing.T) {
	conn, _ := amqp.Dial(url + vhost)
	defer conn.Close()
	convey.Convey("connect rabbitmq", t, func() {
		convey.So(conn, convey.ShouldNotBeNil)
	})
}

func TestSetItem(t *testing.T) {
	conn, _ := amqp.Dial(url + vhost)
	ch, err := conn.Channel()
	defer ch.Close()
	defer conn.Close()
	if err != nil {
		fmt.Printf("amqp create channel fail, err: %s", err)
	}

	body := "david 30 student soft-engineer"
	err = ch.Publish(
		"trans-items",
		"google",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(body),
		})
	if err != nil {
		fmt.Printf("amqp publish msg fail, err: %s", err)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain start")

	//初始化测试环境
	TPTesting.NEW().InitConfig()
	os.Exit(m.Run())
}