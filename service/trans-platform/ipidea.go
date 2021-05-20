package trans_platform

import (
	"TransProxy/lib/proxyip"
	"TransProxy/utils"
	url2 "net/url"
)

const url = "http://tiqu.linksocket.com:81/abroad?num=1&type=2&lb=1&sb=0&flow=1&regions=&port=1&n=0"

type IpIdea struct {}

func (i *IpIdea) ProxyUrl() *url2.URL {
	proxyIps := proxyip.NewIpIdea(url).GetProxy()
	proxyIp := proxyIps[0]
	urlProxy := utils.BuildIpIdeaUrl(proxyIp.IP, proxyIp.Port)
	return urlProxy
}
