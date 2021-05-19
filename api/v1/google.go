package v1

import (
	"TransProxy/manager"
	"TransProxy/model/request"
	"TransProxy/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"log"
)

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

	fmt.Println("items: ", item)

	response.OkWithMessage("insert item successfully.", c)
	fmt.Println("after response")
}
