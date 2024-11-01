package middelware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)
		clientIP := c.ClientIP()
		if clientIP == "::1" {
			clientIP = "127.0.0.1"
		}

		fmt.Printf("Client IP: %s\n ,Method: %s, Path: %s, Status: %d, Latency: %v\n",
			clientIP,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
		)
	}
}
