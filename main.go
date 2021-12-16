package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	ctl "github.com/ygjken/workbook-stock/controllers"
	mdl "github.com/ygjken/workbook-stock/model"
)

func router() *gin.Engine {
	router := gin.Default()
	mdl.InitDb()

	// 取り除く予定
	store, err := postgres.NewStore(mdl.Db, []byte("secret"))
	if err != nil {
		log.Println("Cannot to use store for cookie in postgres")
		panic(err)
	}
	router.Use(sessions.Sessions("othersession", store))

	router.LoadHTMLGlob("./views/build/*.html")
	router.Static("/static/", "./views/build/static/")

	// GETメソッド
	router.GET("/", ctl.Index)
	router.GET("/login", ctl.Login)

	api := router.Group("/api")
	api.Use()
	{
		router.POST("/login", ctl.UserLogIn)
	}

	// user := router.Group("/u")
	// user.Use(mid.LoginCheck()) // ユーザー認証が必要となるグループ
	// {
	// 	user.GET("/testmain", ctl.TestMain)
	// }

	return router
}

func main() {

	router().Run(":8080")

}
