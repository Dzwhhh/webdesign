package main

import "github.com/demian/webdesign/framework"

func registerRoute(core *framework.Core) {
	// 静态路由
	core.Get("/user/login", LoginController)

	// 共同前缀
	groupCore := core.Group("/subject")
	groupCore.Get("/finish", SubjectFinishController)
	groupCore.Post("/start", SubjectStartController)

	// 动态路由
	core.Get("/timeout/:seconds", TimeoutController)
}
