package mq

import (
	"bytes"
	"com.pippishen/trans-proxy/manager"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func Amqp(vHost string) *amqp.Connection {
	url := getDsn(vHost)
	conn, err := amqp.Dial(url)
	if err != nil {
		manager.TP_LOG.Error("amqp rabbit mq connect failed, err:",
			zap.String("url", url),
			zap.String("err", err.Error()))
		return nil
	}
	return conn
}

// vHost: rabbitMq virtual host
func getDsn(vHost string) string {
	config := manager.TP_CONFIG.Get("mq.rabbitmq").(map[string]interface{})

	var buffer bytes.Buffer
	buffer.WriteString("amqp://")
	buffer.WriteString(config["username"].(string))
	buffer.WriteString(":")
	buffer.WriteString(config["password"].(string))
	buffer.WriteString("@")
	buffer.WriteString(config["addr"].(string))
	buffer.WriteString("/")
	buffer.WriteString(vHost)
	return buffer.String()
}

func Close()  {
	err := manager.TP_MQ_RABBIT.Close()
	if err != nil {
		manager.TP_LOG.Error("amqp rabbit mq close failed, err:",
			zap.String("err", err.Error()))
	}
}