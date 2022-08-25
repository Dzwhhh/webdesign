package main

import (
	"context"
	"time"

	"github.com/demian/webdesign/framework"
)

func LoginController(ctx *framework.Context) error {
	ctx.Json(200, "login success")
	return nil
}

func TimeoutController(ctx *framework.Context) error {
	durationCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	finish := make(chan struct{})

	// 执行具体的业务逻辑
	go func() {
		time.Sleep(12 * time.Second)
		finish <- struct{}{}
	}()

	select {
	case <-durationCtx.Done():
		ctx.Json(500, "time out")
	case <-finish:
		ctx.Json(200, "finish")
	}
	return nil
}

func SubjectFinishController(ctx *framework.Context) error {
	ctx.Json(200, "subject finish")
	return nil
}

func SubjectStartController(ctx *framework.Context) error {
	ctx.Json(200, "subject start")
	return nil
}
