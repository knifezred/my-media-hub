package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		log.Printf("[%d] %s %s", status, c.Request.Method, path)
		if latency > 100*time.Millisecond {
			log.Printf("slow request: %s %s (%v)", c.Request.Method, path, latency)
		}
	}
}
