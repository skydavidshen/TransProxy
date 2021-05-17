package proxyip

import (
	"fmt"
	"testing"
)

func TestGetProxy(t *testing.T)  {
	url := "http://tiqu.linksocket.com:81/abroad?num=1&type=2&lb=1&sb=0&flow=1&regions=&port=1&n=0"

	i := NewIpIdea(url)
	fmt.Println(i.GetProxy())
}
