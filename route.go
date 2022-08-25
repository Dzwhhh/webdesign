package main

import (
	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/middleware"
)

func registerRoute(core *framework.Core) {
	// 静态路由
	core.Get("/user/login", middleware.Test1(), LoginController)

	// 共同前缀
	groupCore := core.Group("/subject")
	groupCore.Get("/finish", middleware.Test2(), SubjectFinishController)
	groupCore.Post("/start", SubjectStartController)

	// 动态路由
	core.Get("/timeout/:seconds", TimeoutController)
}
