package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
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

	// セッションの制御
	logined := session.Get("logined_uuid")
	if logined == nil {
		a := []string{uuid}
		if ary, err := json.Marshal(a); err == nil {
			session.Set("logined_uuid", string(ary))
			session.Save()
		}

	} else {
		var ary []string
		if str, err := dproxy.New(logined).String(); err == nil {
			log.Println(err)
			json.Unmarshal([]byte(str), &ary)
		}
		logined = append(ary, uuid)
		if byt, err := json.Marshal(logined); err == nil {
			session.Set("logined_uuid", string(byt))
			session.Save()
		}
	}

	log.Printf("Authentication Success!!")
	log.Printf("  username: " + user.Username)
	log.Printf("  email: " + user.Email)
	log.Printf("  password: " + user.Password)
	log.Println("  setted session: ", session.Get("uuid"))
	user.Authenticate()

	ctx.Redirect(http.StatusSeeOther, "/")
}

// TODO: 作成途中
func UserLogout(ctx *gin.Context) {
	var ary []string
	session := sessions.Default(ctx)
	logined := session.Get("logined_uuid")
	str, err := dproxy.New(logined).String()
	if err != nil { // loginedが存在しない場合もこの例外処理が走る
		log.Println("UserLogout():", err)
		ctx.Redirect(http.StatusUnauthorized, "/")
		return
	}
	if err = json.Unmarshal([]byte(str), &ary); err != nil {
		log.Println("UserLogout():", err)
		ctx.Redirect(http.StatusUnauthorized, "/")
		return
	}

	if uuid, err := ctx.Cookie("uuid"); err == nil {
		if isContains(uuid, ary) {

		}
	}

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
