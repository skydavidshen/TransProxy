package router

import (
	v1 "TransProxy/api/v1"
	"TransProxy/router/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	//ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong...")
	})

	//insert item
	proxyRouter := router.Group("google").Use(middleware.AuthBasic())
	{
		proxyRouter.POST("insert-item", v1.InsertItem)
	}

	fmt.Println("router register success.")
	return router
}
