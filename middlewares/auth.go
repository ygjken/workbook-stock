package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var err error = nil

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
