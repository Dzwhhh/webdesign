package http

import (
	"github.com/demian/webdesign/app/http/module/demo"
	"github.com/demian/webdesign/framework/gin"
)

func Routes(engine *gin.Engine) {
	engine.Static("/dist", "./dist/")
	demo.Register(engine)
}
