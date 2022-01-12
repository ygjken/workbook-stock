package model

import (
	"time"

	"github.com/ygjken/workbook-stock/crypto"
)

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (p *Post) GetCreateAt() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (p *Post) GetUser() (u User) {
	u = User{}
	Db.QueryRow("SELECT id, uuid, user_name, email, created_at FROM users WHERE id = $1", p.UserId).
		Scan(&u.Id, &u.Uuid, &u.UserName, &u.Email, &u.CreatedAt)
	return
}

func (u *User) CreatePost(t Thread, body string) (p Post, err error) {
	s := "INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(s)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(crypto.LongSecureRandomBase64(), body, u.Id, t.Id, time.Now()).Scan(&p.Id, &p.Uuid, &p.Body, &p.UserId, &p.ThreadId, &p.CreatedAt)
	return
}
