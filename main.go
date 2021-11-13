package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ctl "github.com/ygjken/workbook-stock/controllers"
	mid "github.com/ygjken/workbook-stock/middlewares"
)

func router() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("_secret"))
	store.Options(sessions.Options{MaxAge: 3600})
	router.Use(sessions.Sessions("_session", store))

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

	// router().Run(":8080")

	// make request
	values := url.Values{}
	values.Add("username", "tester")
	values.Add("password", "admintest")
	reqBody := strings.NewReader(values.Encode())

	// response
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)

	// set request into gin.context
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/user_login",
		reqBody,
	)

	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	ctl.CreateUserAction(c)
	log.Print()

}
