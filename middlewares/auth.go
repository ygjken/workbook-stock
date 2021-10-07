package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionInfo struct {
	UserId interface{}
}

func LoginCheck(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var logininfo SessionInfo

		session := sessions.Default(ctx)
		logininfo.UserId = session.Get("UserId")

		// セッションがない場合、ログインフォームをだす
		if logininfo.UserId == nil {
			log.Println("ログインしていません")
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			ctx.Abort()
		} else {
			ctx.Set("UserId", logininfo.UserId)
			ctx.Next()
		}
		log.Println("ログインチェック終わり")
	}
}
