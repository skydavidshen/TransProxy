package translator

import (
	"TransProxy/model/request"
	trans_platform "TransProxy/service/trans-platform"
	"testing"
)

func TestTranslate(t *testing.T) {
	google := Google{platformHandler: trans_platform.IpIdea{}}
	item := request.Item {
		UUID: "99a977b749fda07975953f52cff7e093",
		To: "en,jp",
		Platform: "google",
		Text: "武汉",
		Source: "vaffle",
	}
	_ = google.Translate(item)
}
