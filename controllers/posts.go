package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mdl "github.com/ygjken/workbook-stock/model"
)

func CreatePost(ctx *gin.Context) {
	sessionVal := ctx.Value("session")
	s := sessionVal.(mdl.Session)
	u, err := s.GetUser()
	if err != nil {
		log.Println("controllers/CreateThread Error:", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Msg": "ログインを行う必要があります。",
		})
	}

	body := ctx.PostForm("body")
	uuid := ctx.PostForm("uuid") //hidden タグで渡してもらう必要がある
	if err != nil {
		log.Println("controllers/CreatePost Error:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Msg": "Somethings went worng!",
		})
	}

	t, err := mdl.GetThreadByUUID(uuid)
	if err != nil {
		log.Println("controllers/CreatePost Error:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Msg": "Somethings went worng!",
		})
	}

	_, err = u.CreatePost(t, body)
	if err != nil {
		log.Println("controllers/CreatePost Error:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Msg": "Somethings went worng!",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Location": "/threads/read?id=" + t.Uuid,
		"Msg":      "ポストしました。",
	})

}
