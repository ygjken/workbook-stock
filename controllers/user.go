package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	email := ctx.PostForm("email")
	pass := ctx.PostForm("pass")
	if email == "user0@abc.com" && pass == "user00" {
		log.Println("user auth pass")
		UserId := "user0"
		PostLogin(ctx, UserId)
		ctx.Redirect(http.StatusMovedPermanently, "/user/main")
	} else {
		ctx.Redirect(http.StatusFound, "/login")
	}

}

func PostLogin(ctx *gin.Context, UserId string) {
	session := sessions.Default(ctx)
	session.Set("UserId", UserId)
	session.Save()
	log.Println("make session done")
}
