package api

import (
	"github.com/gin-gonic/gin"
)

// PingContext gin handler function
func PingContext(c *gin.Context) {
	c.String(200, "pong")
}
