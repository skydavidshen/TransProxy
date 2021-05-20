package trans_platform

import (
	"TransProxy/lib/googletranslate"
	"TransProxy/lib/googletranslate/method"
	"TransProxy/utils"
)

const userName = "sp9e3fd0b2"
const password = "12345678"
const proxyHost = "gate.smartproxy.com:7000"

type SmartProxy struct {}

func (s *SmartProxy) Translate(to, text string) (string, error) {
	urlProxy := utils.BuildSmartProxyUrl(userName, password, proxyHost)

	translate := googletranslate.TranslationParams {
		From:   "auto",
		To:     to,
		Method: method.NewProxy(urlProxy),
	}
	return translate.Translate(text)
}
