package middlewares

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, err := ctx.Cookie("uuid")
		if err != nil {
			log.Println("middleware/LoginCheck Error:", err)
			ctx.Redirect(http.StatusSeeOther, "/")
			ctx.Abort()
		}

		session := sessions.Default(ctx)
		logined := session.Get("logined_uuid_str")
		if logined == nil {
			ctx.Redirect(http.StatusSeeOther, "/")
			ctx.Abort()
		} else {
			loginedstr, err := dproxy.New(logined).String()
			if err != nil {
				log.Println("middleware/LoginCheck Error:", err)
				ctx.Redirect(http.StatusSeeOther, "/")
				ctx.Abort()
			}

			r := regexp.MustCompile(uuid)
			if !r.MatchString(loginedstr) {
				ctx.Redirect(http.StatusSeeOther, "/")
			} else {
				ctx.Next()
			}
		}

	}
}
