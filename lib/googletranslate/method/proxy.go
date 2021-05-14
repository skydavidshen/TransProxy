package method

import (
	"com.pippishen/trans-proxy/model/business"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Proxy struct {
	ProxyConf business.ProxyConf
}

func (p *Proxy) Content(resourceUrl string) ([]byte, error) {
	urli := url.URL{}
	urlproxy, _ := urli.Parse(fmt.Sprintf("http://%s:%d", p.ProxyConf.IP, p.ProxyConf.Port))
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	rqt, err := http.NewRequest("GET", resourceUrl, nil)
	if err != nil {
		panic(err)
		return []byte{}, err
	}
	response, err := client.Do(rqt)
	if err != nil {
		panic(err)
		return []byte{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
		return []byte{}, err
	}

	return body, nil
}