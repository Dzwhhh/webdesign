package main

import (
	"context"
	"fmt"
	"time"

	"github.com/demian/webdesign/framework"
)

func LoginController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("Login Success")
	return nil
}

func TimeoutController(ctx *framework.Context) error {
	d, _ := ctx.ParamInt("duration", 10)
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
		ctx.SetStatus(500).Json("Time Out")
	case <-finish:
		ctx.SetOkStatus().Json("Finish")
	}
	return nil
}

func SubjectFinishController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("Subject Finish")
	return nil
}

func SubjectStartController(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("Subject Start")
	return nil
}
