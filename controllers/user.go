package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/data"
)

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
	println("post/login")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	db := data.DummyDB()
	user, err := db.GetUser(username, password)
	if err != nil {
		println("Error: " + err.Error())
	} else {
		println("Authentication Success!!")
		println("  username: " + user.Username)
		println("  email: " + user.Email)
		println("  password: " + user.Password)
		user.Authenticate()
	}

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}
