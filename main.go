package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./views/*.html") // html
	r.Static("/src", "./views/src")  // react
	r.GET("/", index)                // handler

	r.Run(":8080")
}

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
