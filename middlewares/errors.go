package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.ByType(gin.ErrorTypePublic).Last()
		fmt.Println(err)
		// if err != nil {
		// 	log.Print(err.Err)

		// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		// 		"Msg": err.Error(),
		// 	})
		// }
	}
}
