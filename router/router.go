package router

import (
	v1 "TransProxy/api/v1"
	"TransProxy/enum"
	"TransProxy/manager"
	"TransProxy/model/response"
	"TransProxy/router/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	//ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong..")
	})

	//test
	router.GET("/test", func(ctx *gin.Context) {
		response.OkWithDetailed(manager.TP_SERVER_CONFIG.System.Oss, "test successfully.", ctx)
	})

	//insert item
	googleProxyRouter := router.Group("google")
	{
		useAuthBasic(googleProxyRouter)
		googleProxyRouter.POST("insert-item", v1.InsertItem)
	}

	fmt.Println("Router register success.")
	return router
}

func useAuthBasic(proxyRouter *gin.RouterGroup) {
	if manager.TP_SERVER_CONFIG.System.Env != enum.Env_Dev {
		proxyRouter.Use(middleware.AuthBasic())
	}
}
