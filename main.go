package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/middlewares"
	"github.com/ygjken/workbook-stock/server"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./views/build/*.html")        // html
	r.Static("/static/", "./views/build/static/") // react

	r.GET("/", server.Index)          // homeページに飛ぶ
	r.GET("/login", server.Login)     // loginページに飛ぶ
	r.POST("/auth", middlewares.Auth) // userを認識し,sessionを作成

	r.Run(":8080")
}
