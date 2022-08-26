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
	"github.com/demian/webdesign/provider/demo"
	"github.com/demian/webdesign/provider/echo"
)

func main() {
	// 创建engine
	core := gin.New()

	// 绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})
	core.Bind(&echo.EchoServiceProvider{})

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
