package main

import "github.com/demian/webdesign/framework"

func main() {
	// 获取handler
	core := framework.NewCore()

	// 注册路由
	registerRoute(core)

	// 监听请求
	core.Listen("8080")
}
