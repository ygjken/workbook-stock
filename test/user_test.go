package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	ctl "github.com/ygjken/workbook-stock/controllers"
)

type userinfo struct {
	username string
	password string
}

type wantedResponse struct {
	code     int
	location string
}

func TestUserLogIn(t *testing.T) {

	tests := []struct {
		name string
		info userinfo
		want wantedResponse
	}{
		{
			name: "correct login",
			info: userinfo{username: "tester", password: "admintest"},
			want: wantedResponse{
				code:     http.StatusSeeOther,
				location: "/",
			},
		},
		{
			name: "uncorrect login",
			info: userinfo{username: "s", password: "f"},
			want: wantedResponse{
				code:     http.StatusFound,
				location: "/login",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// make request
			values := url.Values{}
			values.Add("username", tt.info.username)
			values.Add("password", tt.info.password)
			reqBody := strings.NewReader(values.Encode())

			// response
			resp := httptest.NewRecorder()
			_, r := gin.CreateTestContext(resp)

			// session and cookie
			s := cookie.NewStore([]byte("_secret"))
			s.Options(sessions.Options{MaxAge: 3600})
			r.Use(sessions.Sessions("_session", s))

			// set handler function
			r.POST("/user_login", func(c *gin.Context) {
				// session := sessions.Default(c)
				// log.Println("before login: ", session.Get("logined_uuid_str"))
				ctl.UserLogIn(c)
				// log.Println("after login: ", session.Get("logined_uuid_str"))
			})

			// make request
			req, _ := http.NewRequest(
				http.MethodPost,
				"/user_login",
				reqBody,
			)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			r.ServeHTTP(resp, req)

			// check response
			// log.Println("response code:", resp.Code)
			if resp.Code != tt.want.code {
				t.Errorf("Login failed with %s.", tt.name)
			}

			if resp.HeaderMap.Get("Location") != tt.want.location {
				t.Errorf("Login failed with %s.", tt.name)
			}
		})
	}
}
