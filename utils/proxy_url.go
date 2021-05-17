package utils

import (
	"fmt"
	"net/url"
)

func BuildIpIdeaUrl(ip string, port int) *url.URL {
	urlObj := url.URL{}
	urlValue, _ := urlObj.Parse(fmt.Sprintf("http://%s:%d", ip, port))
	return urlValue
}

func BuildSmartProxyUrl(username string, password string, proxyHost string) *url.URL {
	urlValue := &url.URL{
		Scheme: "http",
		User:   url.UserPassword(username, password),
		Host:   proxyHost,
	}
	return urlValue
}