package main

import (
	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/middleware"
)

func main() {
	// 获取handler
	core := framework.NewCore()

	// 配置全局中间件
	core.Use(middleware.Recovery())

	// 注册路由
	registerRoute(core)

	// 监听请求
	core.Listen("8080")
}
