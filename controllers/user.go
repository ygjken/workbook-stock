package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/crypto"
	"github.com/ygjken/workbook-stock/data"
)

// TODO: セッションとクッキーに対応できるように書き換え
func UserSignUp(ctx *gin.Context) {
	println("post/signup")
	username := ctx.PostForm("username")
	email := ctx.PostForm("emailaddress")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("passwordconfirmation")

	if password != passwordConf {
		println("Error: password and passwordConf not match")
		ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
		return
	}

	db := data.DummyDB()
	if err := db.SaveUser(username, email, password); err != nil {
		println("Error: " + err.Error())
	} else {
		println("Signup success!!")
		println("  username: " + username)
		println("  email: " + email)
		println("  password: " + password)
	}

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

func UserLogIn(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// ログインできるかどうかをチェック
	db := data.DummyDB()
	user, err := db.GetUser(username, password)
	if err != nil {
		log.Printf("Error: " + err.Error())
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// セッションとクッキーをセット
	uuid := crypto.SecureRandomBase64()
	session := sessions.Default(ctx)
	ctx.SetCookie("uuid", uuid, 3600, "/", "localhost", true, true) // jsからクッキーは利用できない
	session.Set("uuid", uuid)
	session.Save()

	log.Printf("Authentication Success!!")
	log.Printf("  username: " + user.Username)
	log.Printf("  email: " + user.Email)
	log.Printf("  password: " + user.Password)
	log.Println("  setted session: ", session.Get("uuid"))
	user.Authenticate()

	ctx.Redirect(http.StatusSeeOther, "/")
}
