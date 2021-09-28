package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ctl "github.com/ygjken/workbook-stock/controllers"
)

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("_secret"))
	router.Use(sessions.Sessions("_session", store))

	router.LoadHTMLGlob("./views/build/*.html")        // html
	router.Static("/static/", "./views/build/static/") // react
	router.GET("/", ctl.Index)                         // homeページに飛ぶ
	router.GET("/login", ctl.Login)

	user := router.Group("/user") // ユーザー認証が必要となるグループ
	{
		user.POST("/login", ctl.UserLogIn) // cookicのテスト
	}

	router.Run(":8080")
}
