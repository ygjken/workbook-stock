package main

import (
	"log"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	ctl "github.com/ygjken/workbook-stock/controllers"
	mdl "github.com/ygjken/workbook-stock/model"
)

func router() *gin.Engine {
	router := gin.Default()
	mdl.InitDb()

	// TODO:取り除く予定
	store, err := postgres.NewStore(mdl.Db, []byte("secret"))
	if err != nil {
		log.Println("Cannot to use store for cookie in postgres")
		panic(err)
	}
	router.Use(sessions.Sessions("othersession", store))

	// Reactルーティング
	router.Use(static.Serve("/", static.LocalFile("./views/build/", true)))
	folderPath := "./views/build/"
	router.NoRoute(func(ctx *gin.Context) {
		_, file := path.Split(ctx.Request.RequestURI) // ディレクトリ名とファイル名を分ける
		ext := filepath.Ext(file)                     // 拡張子取得

		log.Println(file)
		log.Println(ext)

		//ディレクトリアクセス（ファイル名がない）かパスクエリ（拡張子がない）
		if file == "" || ext == "" {
			ctx.File(folderPath + "/index.html")
		} else {
			ctx.File(folderPath + ctx.Request.RequestURI)
		}
	})

	// CORSの設定
	router.Use(cors.New(getCorsConfig()))

	//// router.LoadHTMLGlob("./views/build/*.html")
	router.Static("/static/", "./views/build/static/")

	// GETメソッド
	router.GET("/", ctl.Index)
	router.GET("/login", ctl.Login)
	router.GET("/threads", ctl.Threads)

	// debug用メソッド
	router.GET("/debug-set-cookie", ctl.DebugSetCookie)
	router.GET("/debug-read-cookie", ctl.DebugReadCookie)

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

func getCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Credentials", "Content-Type", "Set-Cookie"},
		AllowCredentials: true,      // cookieやTLSなどの資格情報を含むリクエストを承認
		MaxAge:           time.Hour, // preflightリクエストがキャッシュされる時間
	}
}
