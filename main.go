package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	ctl "github.com/ygjken/workbook-stock/controllers"
	mid "github.com/ygjken/workbook-stock/middlewares"
	mdl "github.com/ygjken/workbook-stock/model"
)

func router() *gin.Engine {
	router := gin.Default()
	mdl.InitDb()

	store, err := postgres.NewStore(mdl.Db, []byte("secret"))
	if err != nil {
		log.Println("Cannot to use store for cookie in postgres")
		panic(err)
	}
	router.Use(sessions.Sessions("othersession", store))

	router.LoadHTMLGlob("./views/build/*.html")        // html
	router.Static("/static/", "./views/build/static/") // react
	router.GET("/", ctl.Index)                         // homeページに飛ぶ
	router.GET("/login", ctl.Login)
	router.POST("/user_login", ctl.UserLogIn) // cookicのテスト

	user := router.Group("/u")
	user.Use(mid.LoginCheck()) // ユーザー認証が必要となるグループ
	{
		user.GET("/testmain", ctl.TestMain)
	}

	return router
}

func main() {

	router().Run(":8080")

}
