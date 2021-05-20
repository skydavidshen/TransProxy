package translator

import (
	"TransProxy/lib/translateserve/googletranslate"
	methodHandler "TransProxy/lib/translateserve/googletranslate/method"
	"TransProxy/manager"
	"TransProxy/model/business"
	"TransProxy/model/request"
	trans_platform "TransProxy/service/trans-platform"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

const mqKey = "google"
const exchange = "trans-items"

type Google struct{
	platformHandler trans_platform.Handler
}

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

func (g *Google) Translate(item request.Item) business.TranslateItem {
	toArr := strings.Split(item.To, ",")
	for _, to := range toArr {
		urlProxy := g.platformHandler.ProxyUrl()
		translate := googletranslate.TranslationParams{
			From:   "auto",
			To:     to,
			Method: methodHandler.NewProxy(urlProxy),
		}
		transText, err := translate.Translate(item.Text)
		
		fmt.Println("result: ", transText)
		fmt.Println("err: ", err)
	}
	return business.TranslateItem{}
}