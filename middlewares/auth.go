package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LoginCheck(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
