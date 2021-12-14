package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/crypto"
	mdl "github.com/ygjken/workbook-stock/model"
)

type formInfo struct {
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
		form formInfo
		want wantedResponse
	}{
		// TODO: Add test cases.
		{
			name: "correct login",
			form: formInfo{username: "tester", password: "admintest"},
			want: wantedResponse{code: http.StatusOK, location: "/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// パスワード暗号化
			pw, err := crypto.PasswordEncrypt(tt.form.password)
			if err != nil {
				t.Errorf("cannot encrypt password")
			}
			// モック宣言
			var mock sqlmock.Sqlmock
			mdl.Db, mock, _ = sqlmock.New()
			// defer mdl.Db.Close()

			// モックの反応を定義
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, uuid, user_name, email, password, created_at FROM users WHERE user_name=$1`)).
				WithArgs(tt.form.username).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "uuid", "user_name", "email", "password", "created_at"}).
						AddRow(1, crypto.LongSecureRandomBase64(), tt.form.username, "tester@admin.com", pw, time.Now()))

			mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO sessions (uuid, user_name, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, user_name, user_id, created_at`))

			mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO sessions (uuid, user_name, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, user_name, user_id, created_at`)).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "uuid", "user_name", "user_id", "created_at"}).
						AddRow(1, crypto.LongSecureRandomBase64(), tt.form.username, 1, time.Now()))
			// モックの反応テスト
			// u, err := mdl.GetUserByUserName(tt.form.username)

			// make request
			values := url.Values{}
			values.Add("username", tt.form.username)
			values.Add("password", tt.form.password)
			reqBody := strings.NewReader(values.Encode())

			// response
			resp := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(resp)

			// make request
			ctx.Request, _ = http.NewRequest(
				http.MethodPost,
				"/user_login",
				reqBody,
			)
			ctx.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// r.ServeHTTP(resp, req)
			UserLogIn(ctx)

			// check response
			if resp.Code != tt.want.code {
				t.Errorf("Login failed with \"%s\": Status Code not match.", tt.name)
			}
			if resp.HeaderMap.Get("Location") != tt.want.location {
				t.Errorf("Login failed with \"%s\": Location not match.", tt.name)
			}
			if resp.HeaderMap.Get("Set-Cookie") == "" {
				t.Errorf("Login failed with \"%s\": Cookie not be setted.", tt.name)
			}
		})
	}
}
