package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	user := ctx.PostForm("user")
	pass := ctx.PostForm("pass")
	fmt.Println(ctx)

	// DEBUG: Formから受診した内容をそのまま返す
	ctx.JSON(http.StatusOK, gin.H{"user": user, "pass": pass})
}
