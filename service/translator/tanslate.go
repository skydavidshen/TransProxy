package translator

import (
	"TransProxy/enum"
	"TransProxy/model/business"
	"TransProxy/model/request"
	transPlatform "TransProxy/service/trans-platform"
)

func TranslateFromItem(item request.Item) (business.TranslateItem, error) {
	var translator Handler
	// 使用哪个翻译平台可以根据实际情况替换
	var platform transPlatform.Handler = transPlatform.SmartProxy{}
	switch item.Platform {
	case enum.Platform_Google:
		translator = &Google{PlatformHandler: platform}
	case enum.Platform_Bing:
		translator = &Bing{PlatformHandler: platform}
	}
	transItem, err := translator.Translate(item)
	return transItem, err
}
