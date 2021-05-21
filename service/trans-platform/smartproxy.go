package trans_platform

import (
	"TransProxy/manager"
	"TransProxy/utils"
	url2 "net/url"
)

var userName = manager.TP_SERVER_CONFIG.TransPlatform.SmartProxy.Username
var password = manager.TP_SERVER_CONFIG.TransPlatform.SmartProxy.Password
var proxyHost = manager.TP_SERVER_CONFIG.TransPlatform.SmartProxy.ProxyHost

type SmartProxy struct {}

func (s SmartProxy) ProxyUrl() *url2.URL {
	urlProxy := utils.BuildSmartProxyUrl(userName, password, proxyHost)
	return urlProxy
}