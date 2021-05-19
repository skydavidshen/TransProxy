package v1

import (
	"TransProxy/manager"
	"TransProxy/model/request"
	"TransProxy/model/response"
	"TransProxy/service"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"log"
)

var googleService service.Google

func InsertItem(c *gin.Context)  {
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
		response.OkWithMessage("Failed to insert item.", c)
		log.Println(err)
		return
	}
	response.OkWithMessage("Insert item successfully.", c)
}
