package v1

import (
	"TransProxy/model/request"
	"TransProxy/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InsertItem(c *gin.Context)  {
	var params request.Basic
	_ = c.ShouldBindJSON(&params)

	fmt.Println("controller: ",params)

	response.OkWithMessage("insert item successfully.", c)
}
