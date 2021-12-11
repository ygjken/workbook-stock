package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/crypto"
	mdl "github.com/ygjken/workbook-stock/model"
)

// TODO: セッションとクッキーに対応できるように書き換え
func UserSignUp(ctx *gin.Context) {
	var dummyUser mdl.User
	var err error

	dummyUser.UserName = "tester"
	dummyUser.Password, err = crypto.PasswordEncrypt("admintest")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": "Password Encrypt Error",
		})
	}
}

// ログインの処理を行う
func UserLogIn(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user, err := mdl.GetUserByUserName(username)

	// ユーザが存在するかどうか
	if err != nil {
		log.Printf("UserLogin Error: " + err.Error())
		ctx.HTML(http.StatusFound, "login.html", gin.H{
			"Error": "ユーザが見つかりませんでした",
		})
		return
	}

	// パスワードが正しいかどうか
	if err = crypto.CompareHashAndPassword(user.Password, password); err != nil {
		log.Println("UserLogin Error: " + err.Error())
		ctx.HTML(http.StatusFound, "login.html", gin.H{
			"Error": "パスワードが正しくありませんでした",
		})
		return
	}

	// セッションとクッキーをセット
	session, err := user.CreateSession()
	if err != nil {
		log.Println("UserLogin Error: " + err.Error())
		ctx.HTML(http.StatusFound, "login.html", gin.H{
			"Error": "現在ログインすることができません",
		})
		return
	}
	ctx.SetCookie("uuid", session.Uuid, 3600, "/", "localhost", true, true) // jsからクッキーは利用できない
	ctx.Set("logined", "yes")
	ctx.Redirect(http.StatusSeeOther, "/")
}

// ログアウト処理を行う
func UserLogOut(ctx *gin.Context) {
	uuid, err := ctx.Cookie("uuid")
	if err != nil {
		log.Println("controllers/UserLogOut Debug:", err)
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	s := mdl.Session{Uuid: uuid}
	err = s.DeleteByUUID()
	if err != nil {
		log.Println("controllers/UserLogOut Debug:", err)
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	ctx.Set("logined", "no")
	ctx.Redirect(http.StatusSeeOther, "/")
}

// 指定のキーが配列内に存在しているかどうか
func isContains(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
