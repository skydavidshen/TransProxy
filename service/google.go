package service

import (
	"TransProxy/manager"
	"TransProxy/model/request"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

const mqKey = "google"
const exchange = "trans-items"
const contentType = "application/json"

type Google struct{}

func (g *Google) InsertItem(item request.Item) error {
	ch, _ := manager.TP_MQ_RABBIT.Channel()
	body, _ := json.Marshal(item)
	err := ch.Publish(
		exchange,
		mqKey,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  contentType,
			Body:         body,
		})
	if err != nil {
		log.Printf("amqp publish msg fail, err: %s", err)
		return err
	}

	return nil
}
