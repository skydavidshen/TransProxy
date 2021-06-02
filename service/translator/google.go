package translator

import (
	"TransProxy/enum"
	"TransProxy/lib/translateserve/googletranslate"
	methodHandler "TransProxy/lib/translateserve/googletranslate/method"
	"TransProxy/manager"
	"TransProxy/model/business"
	"TransProxy/model/request"
	transPlatform "TransProxy/service/trans-platform"
	"TransProxy/utils"
	"encoding/json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"strings"
)

type Google struct {
	PlatformHandler transPlatform.Handler
}

func (g *Google) InsertItem(item request.Item) error {
	var exchange = manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Option.Exchange.TransItems
	ch, _ := manager.TP_MQ_RABBIT.Channel()
	defer ch.Close()

	body, _ := json.Marshal(item)
	err := ch.Publish(
		exchange,
		string(enum.Platform_Google),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  enum.ContentType_Json,
			Body:         body,
			// Expiration 单位为 ms，1000ms = 1s
			// 设置了expired可以防止程序本身故障导致重试次数计算不准，就算重试机制失效，通过消息超时也可以将超时消息塞入「死信队列」
			Expiration: manager.TP_SERVER_CONFIG.MQ.RabbitMQ.Expiration,
			MessageId: utils.GenUUID(),
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
		var transText string

		if manager.TP_SERVER_CONFIG.Switch.UseRealTranslate == false {
			transText = "this is a test translate data..."
		} else {
			urlProxy := g.PlatformHandler.ProxyUrl()
			translate := googletranslate.TranslationParams{
				From:   "auto",
				To:     to,
				Method: methodHandler.NewProxy(urlProxy),
			}

			transResult, err := translate.Translate(item.Text)
			transText = transResult
			if err != nil {
				manager.TP_LOG.Error("Translate fail",
					zap.String("err", err.Error()),
					zap.Any("item", item),
				)
				return business.TranslateItem{}, err
			}
		}
		langItem.Lang = to
		langItem.Text = transText
		transItem.LangItem = append(transItem.LangItem, langItem)
	}
	return transItem, nil
}
