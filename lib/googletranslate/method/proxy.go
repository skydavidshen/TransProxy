package method

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func NewProxy(url *url.URL) *Proxy {
	p := &Proxy{
		url,
	}
	return p
}

type Proxy struct {
	URL *url.URL
}

func (p *Proxy) Content(resourceUrl string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(p.URL),
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