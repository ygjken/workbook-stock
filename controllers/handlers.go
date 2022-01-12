package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mdl "github.com/ygjken/workbook-stock/model"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Location": "/",
	})
}

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Location": "/login",
	})
}

func Threads(ctx *gin.Context) {
	threads, err := mdl.GetThreads()
	if err != nil {
		log.Println("handlers/Threads Error:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Msg": "Somethings went worng!",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Threads": threads,
	})
}
