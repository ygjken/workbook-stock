// TODO: このファイルはデプロイ時に削除を行う
package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DebugSetCookie(ctx *gin.Context) {
	ctx.SetCookie("debug-cookie", "debug-text", 3600, "/", "localhost", true, true) // jsからクッキーは利用できない
	ctx.JSON(http.StatusOK, gin.H{"text": "status_ok"})
}

func DebugReadCookie(ctx *gin.Context) {
	debugCookie, _ := ctx.Cookie("debug-cookie")
	log.Println("ctl.DebugReadCookie(): ", debugCookie)
	ctx.JSON(http.StatusOK, gin.H{})
}
