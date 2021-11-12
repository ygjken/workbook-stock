package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/controllers"
)

func TestUserLogIn(t *testing.T) {
	type userinfo struct {
		username string
		password string
	}

	tests := []struct {
		name string
		info userinfo
	}{
		// TODO: Add test cases.
		{
			name: "logout_test",
			info: userinfo{username: "tester", password: "admintest"},
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
			_, router := gin.CreateTestContext(resp)
			router.POST("/login_user", controllers.UserLogIn)

			// set request into gin.context
			req, _ := http.NewRequest(
				http.MethodPost,
				"/login_user",
				reqBody,
			)

			req.Header.Set("Context-Type", "application/x-www-form-urlencoded")
			req.PostForm = values

			router.ServeHTTP(resp, req)
			log.Print()

		})
	}
}
