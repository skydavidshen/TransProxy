package v1

import (
	"com.pippishen/trans-proxy/model/response"
	"github.com/gin-gonic/gin"
)

func InsertItem(c *gin.Context)  {
	response.OkWithMessage("insert item successfully.", c)
}
