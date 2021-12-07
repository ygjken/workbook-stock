package model

import (
	"time"

	"github.com/ygjken/workbook-stock/crypto"
)

type Session struct {
	Id        int
	Uuid      string
	UserName  string
	UserId    int
	CreatedAt time.Time
}

// 発生したセッションをデータベース内に永久保存
func (user *User) CreateSession() (s Session, err error) {
	statement := "INSERT INTO sessions (uuid, user_name, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, user_name, user_id, created_at"
	stmt, err := Db.Prepare(statement) // 複数SQL文を実行できるように待機する
	if err != nil {
		return
	}
	defer stmt.Close()

	// 今のUserがセッションを持っていない場合に実行される
	err = stmt.QueryRow(crypto.LongSecureRandomBase64(), user.UserName, user.Id, time.Now()).Scan(&s.Id, &s.Uuid, &s.UserName, &s.UserId, &s.CreatedAt)
	return
}

// セッションが有効かどうかをチャック
func (s *Session) Check() (valid bool, err error) {
	Db.QueryRow("SELECT id, uuid, user_name, user_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).Scan(&s.Id, &s.Uuid, &s.UserName, &s.UserId, &s.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if s.Id != 0 {
		valid = true
	}

	return
}

// セッションを削除する
func (s *Session) DeleteByUUID() (err error) {
	statement := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Uuid)
	return
}
