package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong...")
	})
	fmt.Println("router register success.")
	return router
}
