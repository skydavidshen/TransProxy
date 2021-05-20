package translator

import (
	"TransProxy/model/business"
	"TransProxy/model/request"
	transPlatform "TransProxy/service/trans-platform"
)

type Bing struct {
	PlatformHandler transPlatform.Handler
}

func (b Bing) InsertItem(item request.Item) error {
	panic("implement me")
}

func (b Bing) Translate(item request.Item) (business.TranslateItem, error) {
	panic("implement me")
}


