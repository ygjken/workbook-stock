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
		ctx.JSON(http.StatusOK, gin.H{
			"Status": "FAIL",
			"Error":  "スレッドを作成できませんでした",
		})
		return
	}

	topic := ctx.PostForm("topic")
	_, err = u.CreateThread(topic)
	if err != nil {
		log.Println("controllers/CreateThread Error:", err)
		ctx.JSON(http.StatusOK, gin.H{
			"Status": "FAIL",
			"Error":  "スレッドを作成できませんでした",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status": "ok",
		"Error":  "スレッドを作成に成功しました",
	})
}
