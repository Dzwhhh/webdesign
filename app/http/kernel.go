package http

import (
	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/gin"
)

func NewHttpEngine(container framework.Container) (*gin.Engine, error) {
	// 设置为Debug，启动后输出调试信息
	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.SetContainer(container)

	// 路由
	Routes(r)
	return r, nil
}
