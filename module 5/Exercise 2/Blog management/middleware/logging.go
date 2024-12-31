package middleware
import (
	"log"
	"github.com/gin-gonic/gin"
	"time"
)
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		log.Printf("Method: %s, Path: %s, Status: %d, Duration: %v",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(), endTime.Sub(startTime))
	}
}
