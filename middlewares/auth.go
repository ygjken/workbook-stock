package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mdl "github.com/ygjken/workbook-stock/model"
)

func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, err := ctx.Cookie("uuid")
		if err != nil {
			log.Println("middleware/LoginCheck Error:", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Location": "/login",
				"Msg":      "ログインを行なってください。",
			})
			ctx.Abort()
		}

		s := mdl.Session{Uuid: uuid}
		ok, _ := s.Check()
		if !ok {
			log.Println("middleware/LoginCheck Error:", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Location": "/login",
				"Msg":      "再度ログインを行なってください。",
			})
			ctx.Abort()
		}

		ctx.Set("session", s)
		ctx.Next()
	}
}
