package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/demian/webdesign/framework"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		// 初始化超时context
		durationCtx, cancel := context.WithTimeout(c, d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			c.Json(500, "server error")
			fmt.Println(p.(string))
		case <-durationCtx.Done():
			c.Json(500, "time out")
		case <-finish:
			fmt.Println("finish")
		}
		return nil
	}
}
