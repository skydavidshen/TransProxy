package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	resourceUrl = "https://translate.google.com/translate_a/single?client=gtx&dt=at&dt=bd&dt=ex&dt=ld&dt=md&dt=qca&dt=rw&dt=rm&dt=ss&dt=t&hl=en&ie=UTF-8&kc=7&oe=UTF-8&otf=1&q=%E4%BD%A0%E5%A5%BD&sl=auto&ssel=0&tk=879267.879267&tl=en&tsel=0"
	proxyHost   = "gate.visitxiangtan.com:7000"
	username    = "sp9e3fd0b2"
	password    = "12345678"
)

func main() {
	proxyUrl := &url.URL{
		Scheme: "http",
		User:   url.UserPassword(username, password),
		Host:   proxyHost,
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	resp, err := client.Get(resourceUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	fmt.Println(string(body))
	//var body map[string]interface{}
	//if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(body)
}