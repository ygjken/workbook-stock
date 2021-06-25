package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
	// gin.H{}はテンプレートエンジンに埋め込むためのもの
}

func GetLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func TestMain(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "main.html", gin.H{})
}
