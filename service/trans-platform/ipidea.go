package trans_platform

import (
	"TransProxy/lib/proxyip"
	"TransProxy/manager"
	"TransProxy/model/business"
	"TransProxy/utils"
	url2 "net/url"
)

type IpIdea struct {}

func (i *IpIdea) ProxyUrl() *url2.URL {
	proxyIp := getProxyIp()
	urlProxy := utils.BuildIpIdeaUrl(proxyIp.IP, proxyIp.Port)
	return urlProxy
}

// 获取proxy IP，可以通过其他形式获取(并不一定需要实时通过网络URL请求)，例如：从IP池缓存，hashmap本地缓存等
func getProxyIp() business.ProxyIP {
	var url = manager.TP_SERVER_CONFIG.TransPlatform.IpIdea.Url
	proxyIps := proxyip.NewIpIdea(url).GetProxy()
	return proxyIps[0]
}