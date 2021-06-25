package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/controllers"
	"github.com/ygjken/workbook-stock/middlewares"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./views/build/*.html")        // html
	router.Static("/static/", "./views/build/static/") // react

	router.GET("/", controllers.Index) // homeページに飛ぶ

	store := cookie.NewStore([]byte("_secret"))
	router.Use(sessions.Sessions("_session", store))
	router.GET("/login", controllers.GetLogin) // loginページに飛ぶ
	router.POST("/auth", controllers.Auth)

	user := router.Group("/user")
	user.Use(middlewares.SessionCheck)
	{
		user.GET("/main", controllers.TestMain) // cookicのテスト
	}

	router.Run(":8080")
}
