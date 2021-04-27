package middleware

import (
	"com.pippishen/trans-proxy/model/request"
	"github.com/gin-gonic/gin"
)

func AuthBasic() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params request.Basic
		_ = c.ShouldBindJSON(&params)
		c.Next()
	}
}
