package model

import (
	"time"

	"github.com/ygjken/workbook-stock/crypto"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// 発生したセッションをデータベース内に永久保存
func (user *User) CreateSession() (s Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement) // 複数SQL文を実行できるように待機する
	if err != nil {
		return
	}
	defer stmt.Close()

	// 今のUserがセッションを持っていない場合に実行される
	err = stmt.QueryRow(crypto.LongSecureRandomBase64(), user.Email, user.Id, time.Now()).Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	return
}
