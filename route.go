package main

import (
	"github.com/demian/webdesign/framework/gin"
	"github.com/demian/webdesign/framework/middleware"
)

func registerRoute(core *gin.Engine) {
	// 静态路由
	core.GET("/user/login", middleware.Test1(), LoginController)

	// 共同前缀
	groupCore := core.Group("/subject")
	groupCore.GET("/finish", middleware.Test2(), SubjectFinishController)
	groupCore.POST("/start", SubjectStartController)

	// 动态路由
	core.GET("/timeout/:duration", TimeoutController)
}
