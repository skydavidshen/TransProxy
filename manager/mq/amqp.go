package mq

import (
	"TransProxy/manager"
	"bytes"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func Amqp(vHost string) (*amqp.Connection, error) {
	url := getDsn(vHost)
	conn, err := amqp.Dial(url)
	if err != nil {
		manager.TP_LOG.Error("amqp rabbit mq connect failed, err:",
			zap.String("url", url),
			zap.String("err", err.Error()))
		return nil, err
	}
	return conn, nil
}

// vHost: rabbitMq virtual host
func getDsn(vHost string) string {
	var buffer bytes.Buffer
	buffer.WriteString("amqp://")
	buffer.WriteString(manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Username)
	buffer.WriteString(":")
	buffer.WriteString(manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Password)
	buffer.WriteString("@")
	buffer.WriteString(manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Addr)
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
