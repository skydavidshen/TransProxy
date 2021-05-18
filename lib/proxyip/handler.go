package proxyip

import "TransProxy/model/business"

type Handler interface {
	GetProxy() []business.ProxyIP
}
