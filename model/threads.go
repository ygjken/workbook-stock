package model

import (
	"time"

	"github.com/ygjken/workbook-stock/crypto"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (t *Thread) GetCreateAt() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (t *Thread) GetUser() (u User) {
	u = User{}
	Db.QueryRow("SELECT id, uuid, user_name, email, created_at FROM users WHERE id = $1", t.UserId).
		Scan(&u.Id, &u.Uuid, &u.UserName, &u.Email, &u.CreatedAt)
	return
}

func (t *Thread) GetPosts() (p []Post) {
	rows, err := Db.Query("select id, uuid, body, user_id, thread_id, created_at from posts where thread_id = $1", t.Id)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
		if err != nil {
			return
		}

		p = append(p, post)
	}
	rows.Close()
	return
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

func (u *User) CreateThread(topic string) (t Thread, err error) {
	s := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err := Db.Prepare(s)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(crypto.LongSecureRandomBase64(), topic, u.Id, time.Now()).Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	return
}

func (u *User) CreatePost(t Thread, body string) (p Post, err error) {
	s := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(s)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(crypto.LongSecureRandomBase64(), body, u.Id, t.Id, time.Now()).Scan(&p.Id, &p.Uuid, &p.Body, &p.UserId, &p.ThreadId, &p.CreatedAt)
	return
}

func GetThreads() (t []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at desc")
	if err != nil {
		return
	}
	for rows.Next() {
		c := Thread{}
		err = rows.Scan(&c.Id, &c.Uuid, &c.Topic, &c.UserId, &c.CreatedAt)
		if err != nil {
			return
		}
		t = append(t, c)
	}
	rows.Close()
	return
}

func GetThreadByUUID(uuid string) (t Thread, err error) {
	t = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	return
}
