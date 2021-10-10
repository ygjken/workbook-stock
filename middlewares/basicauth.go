package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Accounts *gin.Accounts
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func init() {
	Accounts = &gin.Accounts{
		"foo":   "foo",
		"user1": "user11",
	}
}

func Secrets(ctx *gin.Context) {

	user := ctx.MustGet(gin.AuthUserKey).(string)
	secret, ok := secrets[user]
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}
