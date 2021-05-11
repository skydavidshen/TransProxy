package main

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
)

func main() {
	google("你好")
}

func google(kw string)  {
	text := kw
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "auto",
			To:   "en",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | ja: %s \n", text, translated)
}
