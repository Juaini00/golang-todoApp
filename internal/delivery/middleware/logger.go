package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()
		status := c.Writer.Status()
		latency := time.Since(start)

		log.Printf("%s %s | %d | %v", method, path, status, latency)
	}
}
