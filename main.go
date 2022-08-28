package main

import (

	// "net/http"

	"github.com/demian/webdesign/app/console"
	"github.com/demian/webdesign/app/http"
	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/provider/app"
	"github.com/demian/webdesign/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewWdContainer()

	// 绑定APP服务
	container.Bind(&app.WdAppProvider{})

	// 初始化HTTP引擎并绑定到服务容器
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.KernelServiceProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)

	//
	// // 创建engine
	// core := gin.New()

	// // 绑定具体的服务
	// core.Bind(&demo.DemoServiceProvider{})
	// core.Bind(&echo.EchoServiceProvider{})

	// // 配置全局中间件
	// core.Use(gin.Recovery())
	// core.Use(middleware.Cost())

	// // 注册路由
	// registerRoute(core)

	// 监听请求
	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: core,
	// }
	// go func() {
	// 	server.ListenAndServe()
	// }()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// <-quit // 阻塞 等待退出信号

	// if err := server.Shutdown(context.Background()); err != nil {
	// 	log.Fatal("Server Shutdown Error:", err)
	// }
}
