package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	mdl "github.com/ygjken/workbook-stock/model"
)

// スレッドの作成
func CreateThread(ctx *gin.Context) {

	sessionVal := ctx.Value("session")
	s := sessionVal.(mdl.Session)
	u, err := s.GetUser()
	if err != nil {
		log.Println("controllers/CreateThread Error:", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Msg": "ログインを行う必要があります。",
		})
	}

	topic := ctx.PostForm("topic")
	_, err = u.CreateThread(topic)
	if err != nil {
		log.Println("controllers/CreateThread Error:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Msg": "Somethings went worng!",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg": "スレッドの作成に成功しました",
	})
}

// スレッドに所属するポストを読む
func ReadThreads(ctx *gin.Context) {
	// url query string の取得
	threadUuid := ctx.Query("id")

	thread, err := mdl.GetThreadByUUID(threadUuid)
	if err != nil {
		log.Println("controllers/ReadThreads Error:", err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Msg": "スレッドが見つかりませんでした。",
		})
	}

	posts, err := thread.GetPosts()
	if err != nil {
		log.Println("controllers/ReadThreads Error:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Msg": "Somethings went worng!",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Posts": posts,
	})

}
