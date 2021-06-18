package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./views/build/*.html")        // html
	r.Static("/static/", "./views/build/static/") // react
	r.GET("/", index)                             // handler

	r.Run(":8080")
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
