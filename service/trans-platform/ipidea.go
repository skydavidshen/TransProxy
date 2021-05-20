package trans_platform

import (
	"TransProxy/lib/proxyip"
	"TransProxy/model/business"
	"TransProxy/utils"
	url2 "net/url"
)

const url = "http://tiqu.linksocket.com:81/abroad?num=1&type=2&lb=1&sb=0&flow=1&regions=&port=1&n=0"

type IpIdea struct {}

func (i *IpIdea) ProxyUrl() *url2.URL {
	proxyIp := getProxyIp()
	urlProxy := utils.BuildIpIdeaUrl(proxyIp.IP, proxyIp.Port)
	return urlProxy
}

// 获取proxy IP，可以通过其他形式获取(并不一定需要实时通过网络URL请求)，例如：从IP池缓存，hashmap本地缓存等
func getProxyIp() business.ProxyIP {
	proxyIps := proxyip.NewIpIdea(url).GetProxy()
	return proxyIps[0]
}