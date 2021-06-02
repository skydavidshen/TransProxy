package mq

import (
	"TransProxy/config"
	"TransProxy/manager"
	"github.com/mitchellh/mapstructure"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
)

func InitMQ() {
	conn := manager.TP_MQ_RABBIT
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Build Exchange and Queue fail.", zap.Any("error", err))
	}
	defer ch.Close()

	var exchangeMap map[string]string
	_ = mapstructure.Decode(manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange, &exchangeMap)
	for _, exchange := range exchangeMap {
		_ = GenExchange(ch, exchange)
	}

	_ = GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.TransItem, nil)
	_ = GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.InsertTransItem,
		map[string]interface{}{
			"x-dead-letter-exchange":    manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange.DeadInsertTransItems,
			"x-dead-letter-routing-key": "dead",
		})
	err = GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.DeadInsertTransItem, nil)
	if err != nil {
		manager.TP_LOG.Error("Build Exchange and Queue fail.", zap.Any("error", err))
		log.Println("Build Exchange and Queue fail.", zap.Any("error", err))
	}
}

func GenExchange(ch *amqp.Channel, name string) error {
	err := ch.ExchangeDeclare(
		name,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func GenQueue(ch *amqp.Channel, Item config.MqRabbitOptionQueueItem, args amqp.Table) error {
	q, _ := ch.QueueDeclare(
		Item.Name,
		true,
		false,
		false, //exclusive为true: 连接关闭时会被删除，所以一般设为false
		false,
		args,
	)

	var err error
	for _, bind := range Item.Binds {
		err = ch.QueueBind(
			q.Name,
			bind.Key,
			bind.Exchange,
			false,
			nil,
		)
	}
	return err
}
