package server

import (
	"com.pippishen/trans-proxy/manager"
	TPRouter "com.pippishen/trans-proxy/router"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type WebServer interface {
	ListenAndServe() error
}

func initServer(addr string, router *gin.Engine) WebServer {
	s := endless.NewServer(addr, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func Run() {
	router := TPRouter.Routers()
	listenAddr := fmt.Sprintf(":%d", manager.TP_SERVER_CONFIG.System.Addr)
	s := initServer(listenAddr, router)
	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("ListenAndServe Failed err: %s \n", err))
	}
}
