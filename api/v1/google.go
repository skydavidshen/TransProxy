package v1

import (
	"TransProxy/manager"
	"TransProxy/model/request"
	"TransProxy/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InsertItem(c *gin.Context)  {
	var item request.Item
	_ = c.ShouldBindJSON(&item)

	errItem := manager.TP_VALIDATE.Struct(item)
	if errItem != nil {
		response.FailWithMessage("Request data invalid.", c)
		manager.TP_LOG.Error("Request data invalid.",
			zap.String("err", errItem.Error()),
		)
		c.Abort()
		return
	}

	fmt.Println("controller: ", item)

	response.OkWithMessage("insert item successfully.", c)
}
