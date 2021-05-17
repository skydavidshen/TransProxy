package proxyip

import (
	"com.pippishen/trans-proxy/model/business"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewIpIdea(url string) *IpIdea {
	i := &IpIdea{
		url,
	}
	return i
}

type IpIdea struct {
	Url string
}

func (i IpIdea) GetProxy() []business.ProxyIP {
	client := &http.Client{}
	rqt, err := http.NewRequest("GET", i.Url, nil)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(rqt)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var proxyJson business.IPIdeaResponse
	if err := json.Unmarshal([]byte(body), &proxyJson); err == nil {
		return proxyJson.IPIdeaRespData
	} else {
		panic(err)
	}
}

