package request

import "TransProxy/enum"

// UUID：唯一码
// Platform: 翻译平台，例如：google、bing、baidu
// To: 翻译成什么语言(缩写)，批量多语言翻译以逗号隔开，例如：en,fr,jp 等 language.English
// Text：翻译文本
// Source：来源平台，例如：vaffle, hg, dms

type Item struct {
	UUID     string            `json:"uuid" mapstructure:"uuid" validate:"required"`
	Platform enum.PlatformType `json:"platform" mapstructure:"platform" validate:"required"`
	To       string            `json:"to" mapstructure:"to" validate:"required"`
	Text     string            `json:"text" mapstructure:"text" validate:"required"`
	Source   enum.SourceType   `json:"source" mapstructure:"source" validate:"required"`
}
