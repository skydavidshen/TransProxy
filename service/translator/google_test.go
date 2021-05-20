package translator

import (
	"TransProxy/model/request"
	trans_platform "TransProxy/service/trans-platform"
	"fmt"
	"testing"
)

func TestTranslateIpIdea(t *testing.T) {
	google := Google{platformHandler: &trans_platform.IpIdea{}}
	item := request.Item {
		UUID: "99a977b749fda07975953f52cff7e093",
		To: "en,ja",
		Platform: "google",
		Text: "武汉",
		Source: "vaffle",
	}
	transItem, _ := google.Translate(item)
	fmt.Println("transItem: ", transItem)
}

func TestTranslateSmartProxy(t *testing.T) {
	google := Google{platformHandler: &trans_platform.SmartProxy{}}
	item := request.Item {
		UUID: "99a977b749fda07975953f52cff7e093",
		To: "en,ja",
		Platform: "google",
		Text: "武汉",
		Source: "vaffle",
	}
	transItem, _ := google.Translate(item)
	fmt.Println("transItem: ", transItem)
}


