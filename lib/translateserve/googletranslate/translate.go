package googletranslate

import (
	methodHandler "TransProxy/lib/translateserve/googletranslate/method"
	"time"
)

var GoogleHost = "google.com"

// TranslationParams is a util struct to pass as parameter to indicate how to Translate
type TranslationParams struct {
	From       string
	To         string
	Tries      int
	Delay      time.Duration
	GoogleHost string
	Method     methodHandler.Method
}

func (t *TranslationParams) Translate(text string) (string, error) {
	if t.GoogleHost == "" {
		GoogleHost = "google.com"
	} else {
		GoogleHost = t.GoogleHost
	}
	translated, err := translate(text, t.From, t.To, true, t.Method)
	if err != nil {
		return "", err
	}
	return translated, nil
}
