package controllers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"github.com/ygjken/workbook-stock/crypto"
	"github.com/ygjken/workbook-stock/model"
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

	// ログインできるかどうかをチェック
	db := model.DummyDB()
	user, err := db.GetUser(username, password)
	if err != nil {
		log.Printf("Error: " + err.Error())
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	// セッションとクッキーをセット
	uuid := crypto.SecureRandomBase64()
	session := sessions.Default(ctx)
	ctx.SetCookie("uuid", uuid, 3600, "/", "localhost", true, true) // jsからクッキーは利用できない

	// セッションの制御
	uuids := session.Get("logined_uuid_str")
	if uuids == nil {
		session.Set("logined_uuid_str", uuid)
	} else {
		if uuidstr, err := dproxy.New(uuids).String(); err == nil {
			uuids = uuidstr + uuid
			session.Set("logined_uuid_str", uuids)
		}
	}
	session.Save()

	// DEBUG:
	// log.Printf("Authentication Success!!")
	// log.Printf("  username: " + user.Username)
	// log.Printf("  email: " + user.Email)
	// log.Printf("  password: " + user.Password)
	// log.Println("  setted session: ", session.Get("uuid"))
	user.Authenticate()

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
