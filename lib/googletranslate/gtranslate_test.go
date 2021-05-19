package googletranslate

import (
	"TransProxy/lib/googletranslate/method"
	"TransProxy/lib/proxyip"
	"TransProxy/utils"
	"fmt"
	"testing"
)

func TestIpIdea(t *testing.T) {
	urlStr := "http://tiqu.linksocket.com:81/abroad?num=1&type=2&lb=1&sb=0&flow=1&regions=&port=1&n=0"
	proxyIps := proxyip.NewIpIdea(urlStr).GetProxy()

	for _, proxyIp := range proxyIps {
		urlProxy := utils.BuildIpIdeaUrl(proxyIp.IP, proxyIp.Port)

		translate := TranslationParams {
			From:   "auto",
			To:     "en",
			Method: method.NewProxy(urlProxy),
		}
		result, err := translate.Translate("你好")
		if err != nil {
			panic(err)
		}

		fmt.Printf("result: %v", result)
	}
}

func TestSmartProxy(t *testing.T) {
	urlProxy := utils.BuildSmartProxyUrl("sp9e3fd0b2", "12345678", "gate.smartproxy.com:7000")

	translate := TranslationParams {
		From:   "auto",
		To:     "en",
		Method: method.NewProxy(urlProxy),
	}
	result, err := translate.Translate("你好")
	if err != nil {
		panic(err)
	}

	fmt.Printf("result: %v", result)
}
