package proxyip

import "com.pippishen/trans-proxy/model/business"

type Handler interface {
	GetProxy() []business.ProxyIP
}
