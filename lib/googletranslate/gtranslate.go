package googletranslate

import (
	"com.pippishen/trans-proxy/lib/googletranslate/method"
	"time"
)

var GoogleHost = "google.com"

// TranslationParams is a util struct to pass as parameter to indicate how to translate
type TranslationParams struct {
	From       string
	To         string
	Tries      int
	Delay      time.Duration
	GoogleHost string
	method     method.Method
}

func (t *TranslationParams) translate(text string) (string, error) {
	if t.GoogleHost == "" {
		GoogleHost = "google.com"
	} else {
		GoogleHost = t.GoogleHost
	}
	translated, err := translate(text, t.From, t.To, true, t.method)
	if err != nil {
		return "", err
	}
	return translated, nil
}
