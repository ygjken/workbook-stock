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
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Set("logined", "no")
			ctx.Abort()
		}

		s := mdl.Session{Uuid: uuid}
		ok, _ := s.Check()
		if !ok {
			log.Println("middleware/LoginCheck Error:", err)
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Set("logined", "no")
			ctx.Abort()
		}

		ctx.Set("logined", "yes")
		ctx.Next()
	}
}
