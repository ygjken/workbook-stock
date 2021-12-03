package controllers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
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
		log.Printf("Error: " + err.Error())
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

	ctx.Redirect(http.StatusSeeOther, "/")
}

// TODO: 作成途中
// ログアウト処理を行う
func UserLogOut(ctx *gin.Context) {
	uuid, err := ctx.Cookie("uuid")
	if err != nil {
		log.Println("controllers/UserLogout Error:", err)
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	session := sessions.Default(ctx)
	logined := session.Get("logined_uuid_str")
	if logined != nil {
		log.Println("controllers/UserLogout Error:", err)
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	loginedstr, err := dproxy.New(logined).String()
	if err != nil {
		log.Println("controllers/UserLogout Error:", err)
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	r := regexp.MustCompile(uuid)
	if !r.MatchString(loginedstr) {
		log.Println("controllers/UserLogout Error: Can't find uuid of the logined user in session")
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	loginedstr = r.ReplaceAllString(loginedstr, "")
	session.Set("logined_uuid_str", loginedstr)
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
