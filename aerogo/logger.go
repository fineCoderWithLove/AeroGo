package aerogo

import (
	"log"
	"time"
)

// 自定义的日志中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("code: [%d] route:%s  cost:%v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
