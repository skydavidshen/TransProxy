package business

type TranslateItem struct {
	UUID     string     `json:"uuid" mapstructure:"uuid" validate:"required"`
	Platform string     `json:"platform" mapstructure:"platform" validate:"required"`
	To       string     `json:"to" mapstructure:"to" validate:"required"`
	Text     string     `json:"text" mapstructure:"text" validate:"required"`
	Source   string     `json:"source" mapstructure:"source" validate:"required"`
	LangItem []LangItem `json:"lang_item" mapstructure:"lang_item" validate:"required"`
}

type LangItem struct {
	Lang string `json:"lang" mapstructure:"lang"`
	Text string `json:"text" mapstructure:"text"`
}
