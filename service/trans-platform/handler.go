package trans_platform

import url2 "net/url"

type Handler interface {
	ProxyUrl() *url2.URL
}
