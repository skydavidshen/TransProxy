package mq

import (
	"TransProxy/config"
	"TransProxy/manager"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
)

func InitMQ() {
	ch := GenChannel()
	doInit(ch)
}

func doInit(ch *amqp.Channel) {
	// 由于会主动close channel，所以，不需要监听channel的MonitorChannel
	defer ch.Close()

	go MonitorChannel(ch, func(data interface{}) {
		val, ok := data.(*amqp.Error)
		if ok && val != nil {
			// 真实错误
			log.Printf("Case amqp.Error - MonitorChannel communication error: %s", val.Error())
			manager.TP_LOG.Error("Case amqp.Error - MonitorChannel communication error: ", zap.Error(val))

			panic(fmt.Sprintf("Case amqp.Error - MonitorChannel communication error: %s", val.Error()))
		} else {
			// close channel导致的错误，这是正常的，此时data == nil
			log.Printf("Case default - MonitorChannel communication message: %v", data)
		}
	})
	
	log.Println("doInit start...")
	var errors []error
	var err error
	var exchangeMap map[string]string
	_ = mapstructure.Decode(manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange, &exchangeMap)
	for _, exchange := range exchangeMap {
		err = GenExchange(ch, exchange)
		if err != nil {
			errors = append(errors, err)
		}
	}
	log.Println("errors: ", errors)

	err = GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.TransItem, map[string]interface{}{
		"x-dead-letter-exchange":    manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange.DeadTransItems,
		"x-dead-letter-routing-key": "dead",
	})
	if err != nil {
		errors = append(errors, err)
	}
	log.Println("Queue.TransItem errors: ", errors, ch)

	err = GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.InsertTransItem,
		map[string]interface{}{
			"x-dead-letter-exchange":    manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange.DeadInsertTransItems,
			"x-dead-letter-routing-key": "dead",
		})
	if err != nil {
		errors = append(errors, err)
	}

	err = GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.DeadTransItem, nil)
	if err != nil {
		errors = append(errors, err)
	}

	err = GenQueue(ch, manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Queue.DeadInsertTransItem, nil)
	if err != nil {
		errors = append(errors, err)
	}

	log.Println("errors: ", errors)

	if len(errors) > 0 {
		manager.TP_LOG.Error("Build GenQueue fail.", zap.Any("errors", errors))
		log.Println("Build GenQueue fail.", zap.Any("errors", errors))
		log.Println("InitMQ IsClosed: ", manager.TP_MQ_RABBIT.IsClosed())
	}
	log.Println("doInit done.")
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
