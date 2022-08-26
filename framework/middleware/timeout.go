package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/demian/webdesign/framework/gin"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			c.ISetStatus(500).IJson("Internal Error")
			fmt.Println(p.(string))
		case <-durationCtx.Done():
			c.ISetStatus(500).IJson("Time Out")
		case <-finish:
			fmt.Println("finish")
		}
	}
}
