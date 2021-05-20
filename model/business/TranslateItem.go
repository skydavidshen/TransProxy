package business

import "TransProxy/enum"

type TranslateItem struct {
	UUID     string            `json:"uuid" mapstructure:"uuid" validate:"required"`
	Platform enum.PlatformType `json:"platform" mapstructure:"platform" validate:"required"`
	To       string            `json:"to" mapstructure:"to" validate:"required"`
	Text     string            `json:"text" mapstructure:"text" validate:"required"`
	Source   enum.SourceType   `json:"source" mapstructure:"source" validate:"required"`
	LangItem []LangItem        `json:"lang_item" mapstructure:"lang_item" validate:"required"`
}

type LangItem struct {
	Lang string `json:"lang" mapstructure:"lang"`
	Text string `json:"text" mapstructure:"text"`
}
