package main

import (
	"context"
	"fmt"
	"time"

	"github.com/demian/webdesign/framework/gin"
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
