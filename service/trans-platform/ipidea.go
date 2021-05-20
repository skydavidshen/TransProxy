package trans_platform

import (
	"TransProxy/lib/googletranslate"
	"TransProxy/lib/googletranslate/method"
	"TransProxy/lib/proxyip"
	"TransProxy/utils"
)

const url = "http://tiqu.linksocket.com:81/abroad?num=1&type=2&lb=1&sb=0&flow=1&regions=&port=1&n=0"

type IpIdea struct {}

func (i *IpIdea) Translate(to, text string) (string, error) {
	proxyIps := proxyip.NewIpIdea(url).GetProxy()
	proxyIp := proxyIps[0]
	urlProxy := utils.BuildIpIdeaUrl(proxyIp.IP, proxyIp.Port)
	translate := googletranslate.TranslationParams {
		From: "auto",
		To:   to,
		Method: method.NewProxy(urlProxy),
	}
	result, err := translate.Translate(text)
	if err != nil {
		return "", err
	}

	return result, nil
}
