package translator

import (
	"TransProxy/lib/translateserve/googletranslate"
	methodHandler "TransProxy/lib/translateserve/googletranslate/method"
	"TransProxy/manager"
	"TransProxy/model/business"
	"TransProxy/model/request"
	transPlatform "TransProxy/service/trans-platform"
	"encoding/json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"strings"
)

const mqKey = "google"
const exchange = "trans-items"

type Google struct{
	platformHandler transPlatform.Handler
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

func (g *Google) Translate(item request.Item) (business.TranslateItem, error) {
	toArr := strings.Split(item.To, ",")

	var transItem business.TranslateItem
	transItem.UUID = item.UUID
	transItem.Source = item.Source
	transItem.Platform = item.Platform
	transItem.To = item.To
	transItem.Text = item.Text

	for _, to := range toArr {
		var langItem business.LangItem
		urlProxy := g.platformHandler.ProxyUrl()
		translate := googletranslate.TranslationParams{
			From:   "auto",
			To:     to,
			Method: methodHandler.NewProxy(urlProxy),
		}
		transText, err := translate.Translate(item.Text)
		if err != nil {
			manager.TP_LOG.Error("Translate fail",
				zap.String("err", err.Error()),
				zap.Any("item", item),
			)
			return business.TranslateItem{}, err
		}
		langItem.Lang = to
		langItem.Text = transText
	}
	return transItem, nil
}