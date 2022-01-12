package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/crypto"
	mdl "github.com/ygjken/workbook-stock/model"
)

// ログインの処理を行う
func UserLogIn(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user, err := mdl.GetUserByUserName(username)

	// ユーザが存在するかどうか
	if err != nil {
		log.Printf("UserLogin Error: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"Msg": "ユーザ名またはパスワードが間違っています",
		})
		return
	}

	// パスワードが正しいかどうか
	if err = crypto.CompareHashAndPassword(user.Password, password); err != nil {
		log.Println("UserLogin Error: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"Msg": "ユーザ名またはパスワードが間違っています",
		})
		return
	}

	// セッションとクッキーをセット
	session, err := user.CreateSession()
	if err != nil {
		log.Println("UserLogin Error: " + err.Error())
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"Msg": "エラーが発生しました。再度ログインしてください。",
		})
		return
	}
	ctx.SetCookie("uuid", session.Uuid, 3600, "/", "localhost", true, true) // jsからクッキーは利用できない
	ctx.Set("logined", "yes")

	ctx.JSON(http.StatusMovedPermanently, gin.H{
		"Location": "/",
	})

}

// ログアウト処理を行う
func UserLogOut(ctx *gin.Context) {
	uuid, err := ctx.Cookie("uuid")
	if err != nil {
		log.Println("controllers/UserLogOut Debug:", err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"Msg": "ログインしていません",
		})
		return
	}

	s := mdl.Session{Uuid: uuid}
	err = s.DeleteByUUID()
	if err != nil {
		log.Println("controllers/UserLogOut Debug:", err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"Msg": "エラーが発生しました。再度ログインしてください。",
		})
		return
	}

	ctx.Set("logined", "no")
	ctx.JSON(http.StatusMovedPermanently, gin.H{
		"Location": "/",
	})
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
