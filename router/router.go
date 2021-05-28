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

	// hello world
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to the translation service.")
	})

	// ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong..")
	})

	// test
	router.GET("/test", func(ctx *gin.Context) {
		response.OkWithDetailed(manager.TP_SERVER_CONFIG.System.Oss, "test successfully.", ctx)
	})


	// 华丽分界线 ================ Business API =================================

	googleProxyRouter := router.Group("google")
	{
		useAuthBasic(googleProxyRouter)
		// 异步请求翻译
		googleProxyRouter.POST("async-translate", v1.AsyncTranslate)
		// 同步请求翻译
		googleProxyRouter.POST("translate", v1.Translate)
	}

	fmt.Println("Router register success.")
	return router
}

func useAuthBasic(proxyRouter *gin.RouterGroup) {
	if manager.TP_SERVER_CONFIG.System.Env != enum.Env_Dev {
		proxyRouter.Use(middleware.AuthBasic())
	}
}
