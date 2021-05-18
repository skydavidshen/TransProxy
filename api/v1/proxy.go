package v1

import (
	"TransProxy/model/response"
	"github.com/gin-gonic/gin"
)

func InsertItem(c *gin.Context)  {
	response.OkWithMessage("insert item successfully.", c)
}
