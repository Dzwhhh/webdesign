package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/demian/webdesign/framework/gin"
	"github.com/demian/webdesign/framework/middleware"
)

func main() {
	// 获取handler
	core := gin.New()

	// 配置全局中间件
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	// 注册路由
	registerRoute(core)

	// 监听请求
	server := http.Server{
		Addr:    ":8080",
		Handler: core,
	}
	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit // 阻塞 等待退出信号

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
}
