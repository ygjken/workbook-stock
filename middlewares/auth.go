package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionCheck(ctx *gin.Context) {
	var UserId interface{}

	session := sessions.Default(ctx)
	UserId = session.Get("UserId")

	// if don't have session, redircet to /login
	if UserId == nil {
		log.Println("don't login")
		ctx.Redirect(http.StatusMovedPermanently, "/login")
		ctx.Abort()
	} else {
		ctx.Set("UserId", UserId)
		ctx.Next()
	}

}
