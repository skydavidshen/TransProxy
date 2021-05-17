package main

import (
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

var testApi = "https://translate.google.com/translate_a/single?client=gtx&dt=at&dt=bd&dt=ex&dt=ld&dt=md&dt=qca&dt=rw&dt=rm&dt=ss&dt=t&hl=en&ie=UTF-8&kc=7&oe=UTF-8&otf=1&q=%E4%BD%A0%E5%A5%BD&sl=auto&ssel=0&tk=879267.879267&tl=en&tsel=0"

func main() {
	fmt.Println("代理测试")

	httpProxy("49.51.73.134", 17939)

	//go socks5Proxy("8.210.11.115", 15767)

	//time.Sleep(time.Hour)
}

func httpProxy(ip string, port int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), "http", "返回信息:", err)
		}
	}()
	urli := url.URL{}
	urlproxy, _ := urli.Parse(fmt.Sprintf("http://%s:%d", ip, port))
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	rqt, err := http.NewRequest("GET", testApi, nil)
	if err != nil {
		panic(err)
		return
	}
	response, err := client.Do(rqt)
	if err != nil {
		panic(err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), "http", "返回信息:", string(body))
	return
}

func socks5Proxy(ip string, port int) {
	dialer, err := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%d", ip, port), nil, proxy.Direct)
	if err != nil {
		_, _ = fmt.Println("can't connect to the proxy:", err.Error())
		os.Exit(1)
	}
	//setup a http client
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		},
	}
	httpClient.Timeout = time.Second * 10
	// set our socks5 as the dialer
	if resp, err := httpClient.Get(testApi); err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), "socks5", "返回信息:", err.Error())
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), "socks5", "返回信息:", string(body))
	}

}
