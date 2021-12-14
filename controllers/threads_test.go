package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/ygjken/workbook-stock/crypto"
	mdl "github.com/ygjken/workbook-stock/model"
)

type Response struct {
	Err     string       `json:"Error"`
	Threads []mdl.Thread `json:"thread"`
}

func TestThreads(t *testing.T) {
	tests := []struct {
		name    string
		threads []mdl.Thread
	}{
		{
			name: "default",
			threads: []mdl.Thread{
				{Id: 1, Uuid: crypto.SecureRandomBase64(), Topic: "test1", UserId: 1, CreatedAt: time.Now()},
				{Id: 2, Uuid: crypto.SecureRandomBase64(), Topic: "test2", UserId: 1, CreatedAt: time.Now()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mock sqlmock.Sqlmock
			mdl.Db, mock, _ = sqlmock.New()
			defer mdl.Db.Close()

			// モックの反応を定義
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at desc`)).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "uuid", "topic", "user_id", "created_at"}).
						AddRow(tt.threads[0].Id, tt.threads[0].Uuid, tt.threads[0].Topic, tt.threads[0].UserId, tt.threads[0].CreatedAt).
						AddRow(tt.threads[1].Id, tt.threads[1].Uuid, tt.threads[1].Topic, tt.threads[1].UserId, tt.threads[1].CreatedAt))

			// response
			resp := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(resp)

			// make request
			ctx.Request, _ = http.NewRequest(
				http.MethodPost,
				"/threads",
				nil,
			)

			Threads(ctx)

			// check response
			wantedJson := Response{}
			err := json.Unmarshal([]byte(resp.Body.String()), &wantedJson)
			if err != nil {
				t.Errorf("cannot unmarshal wanted response json.")
			}

			respStr, _ := json.Marshal(wantedJson.Threads)
			wantedStr, _ := json.Marshal(tt.threads)

			if string(respStr) != string(wantedStr) {
				t.Errorf("response is not correct.")
			}

		})
	}
}
