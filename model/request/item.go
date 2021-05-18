package request

// UUID：唯一码
// Platform: 翻译平台，例如：google、bing、baidu
// To: 翻译成什么语言(缩写)，批量多语言翻译以逗号隔开，例如：en,fr,jp 等 language.English
// Text：翻译文本
// Source：来源平台，例如：vaffle, hg, dms

type Item struct {
	UUID     string `json:"uuid" validate:"required"`
	Platform string `json:"platform" validate:"required"`
	To       string `json:"to" validate:"required"`
	Text     string `json:"text" validate:"required"`
	Source   string `json:"source" validate:"required"`
}