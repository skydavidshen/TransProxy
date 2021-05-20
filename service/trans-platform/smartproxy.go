package trans_platform

import (
	"TransProxy/utils"
	url2 "net/url"
)

const userName = "sp9e3fd0b2"
const password = "12345678"
const proxyHost = "gate.smartproxy.com:7000"

type SmartProxy struct {}

func (s SmartProxy) ProxyUrl() *url2.URL {
	urlProxy := utils.BuildSmartProxyUrl(userName, password, proxyHost)
	return urlProxy
}