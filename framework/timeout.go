package framework

import (
	"context"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	// 使用函数回调
	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 初始化超时context
		durationCtx, cancel := context.WithTimeout(c, d)
		defer cancel()

		go func() {
			defer func() {
				if err := recover(); err != nil {
					panicChan <- err
				}
			}()
			// 执行具体的业务逻辑
			fun(c)
		}()

		select {
		case p := <-panicChan:
			c.Json(500, map[string]string{"server error": p.(string)})
		case <-durationCtx.Done():
			c.Json(500, "time out")
		case <-finish:
			c.Json(200, "finish")
		}
		return nil
	}
}
