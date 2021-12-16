package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	mdl "github.com/ygjken/workbook-stock/model"
)

func Threads(ctx *gin.Context) {
	errorMsg := ""

	threads, err := mdl.GetThreads()
	if err != nil {
		m := "Cannot get threads. "
		log.Println(m, err)
		errorMsg = m
	}

	ctx.JSON(http.StatusOK, gin.H{
		"thread": threads,
		"Error":  errorMsg,
	})
}

func CreateThread(ctx *gin.Context) {

	tmp := ctx.Value("session")
	s := tmp.(mdl.Session)
	u, err := s.GetUser()
	if err != nil {
		log.Println("controllers/CreateThread Error:", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": "ログインを行う必要があります。",
		})
		return
	}

	topic := ctx.PostForm("topic")
	_, err = u.CreateThread(topic)
	if err != nil {
		log.Println("controllers/CreateThread Error:", err)
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"Error": "スレッドを作成できませんでした。再度お試しください。",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Error": "スレッドの作成に成功しました",
	})
}
