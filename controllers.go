package main

import (
	"context"
	"fmt"
	"time"

	"github.com/demian/webdesign/framework/gin"
	"github.com/demian/webdesign/provider/demo"
	"github.com/demian/webdesign/provider/echo"
)

func LoginController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("Login Success")
}

func TimeoutController(ctx *gin.Context) {
	d, _ := ctx.DefaultParamInt("duration", 10)
	fmt.Println("duration:", d)
	durationCtx, cancel := context.WithTimeout(ctx, time.Duration(d)*time.Second)
	defer cancel()
	finish := make(chan struct{})

	// 执行具体的业务逻辑
	go func() {
		time.Sleep(8 * time.Second)
		finish <- struct{}{}
	}()

	select {
	case <-durationCtx.Done():
		ctx.ISetStatus(500).IJson("Time Out")
	case <-finish:
		ctx.ISetOkStatus().IJson("Finish")
	}
}

func SubjectFinishController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("Subject Finish")
}

func SubjectStartController(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("Subject Start")
}

func UseServiceController(ctx *gin.Context) {
	serviceName, _ := ctx.DefaultParamString("name", "")

	if serviceName == "demo" {
		instance := ctx.MustMake(demo.Key).(demo.Service)
		groot := instance.GetGroot()
		ctx.ISetOkStatus().IJson(groot)
	} else {
		ctx.ISetStatus(403).IJson("Service Not Found")
		return
	}
}

func EchoServiceController(ctx *gin.Context) {
	type info struct {
		Msg  string `json:"msg"`
		Name string `json:"name"`
	}
	var inf info
	err := ctx.BindJSON(&inf)

	if err != nil {
		ctx.ISetStatus(501).IJson("Parse Request Failed")
	}

	msg, name := inf.Msg, inf.Name
	serviceName, _ := ctx.DefaultParamString("echo", "")

	if serviceName == "echo" {
		instance, err := ctx.MakeNew(echo.Key, []interface{}{msg, name})
		if err == nil {
			service := instance.(echo.Service)
			result := service.Echo()
			ctx.ISetOkStatus().IJson(result)
			return
		}
	}
	ctx.ISetStatus(403).IJson("Service Not Found")
}
