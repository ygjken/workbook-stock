package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		logined := session.Get("uuid")
		var err error = nil
		log.Println(logined)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
