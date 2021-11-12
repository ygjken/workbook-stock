package controllers

import (
	"log"
	"net/http"
	"regexp"

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

// ログインの処理を行う
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

	log.Printf("Authentication Success!!")
	log.Printf("  username: " + user.Username)
	log.Printf("  email: " + user.Email)
	log.Printf("  password: " + user.Password)
	log.Println("  setted session: ", session.Get("uuid"))
	user.Authenticate()

	ctx.Redirect(http.StatusSeeOther, "/")
}

// TODO: 作成途中
// ログアウト処理を行う
func UserLogout(ctx *gin.Context) {
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

// testをデバックする用のtestハンドラー
func CreateUserAction(c *gin.Context) {
	var err error

	un := c.PostForm("username")
	pw := c.PostForm("password")

	log.Printf("un:%s pw:%s", un, pw)

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{"message": "success", "name": un, "pass": pw})
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})

	}
}
