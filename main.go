package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ctl "github.com/ygjken/workbook-stock/controllers"
	mid "github.com/ygjken/workbook-stock/middlewares"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("_secret"))
	store.Options(sessions.Options{MaxAge: 3600})
	router.Use(sessions.Sessions("_session", store))

	router.LoadHTMLGlob("./views/build/*.html")        // html
	router.Static("/static/", "./views/build/static/") // react
	router.GET("/", ctl.Index)                         // homeページに飛ぶ
	router.GET("/login", ctl.Login)
	router.POST("/user_login", ctl.UserLogIn) // cookicのテスト

	user := router.Group("/user")
	user.Use(mid.LoginCheck()) // ユーザー認証が必要となるグループ
	{
		user.GET("/testmain", ctl.TestMain)
	}

	router.Run(":8080")
}
