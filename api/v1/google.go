package v1

import (
	"TransProxy/manager"
	"TransProxy/model/request"
	"TransProxy/model/response"
	"TransProxy/service/translator"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"log"
)

var googleService translator.Google

// AsyncTranslate 异步获取翻译信息
func AsyncTranslate(c *gin.Context) {
	var basic request.Basic
	_ = c.ShouldBindJSON(&basic)

	var item request.Item
	_ = mapstructure.Decode(basic.Data, &item)
	errItem := manager.TP_VALIDATE.Struct(item)
	if errItem != nil {
		response.FailWithMessage("Request data invalid.", c)
		log.Println(errItem)
		return
	}

	err := googleService.InsertItem(item)
	if err != nil {
		response.FailWithMessage("Failed to insert item.", c)
		log.Println(err)
		return
	}
	response.OkWithDetailed(item, "Asynchronous translation succeeded.", c)
}

// Translate 同步获取翻译信息
func Translate(c *gin.Context) {
	var basic request.Basic
	_ = c.ShouldBindJSON(&basic)

	var item request.Item
	_ = mapstructure.Decode(basic.Data, &item)
	transItem, err := translator.TranslateFromItem(item)

	if err != nil {
		response.FailWithMessage("Failed to translate item.", c)
		log.Println(err)
		return
	}
	response.OkWithDetailed(transItem, "Translation succeeded.", c)
}
