package model

import (
	"time"

	"github.com/ygjken/workbook-stock/crypto"
)

type User struct {
	Id        int
	Uuid      string
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func GetUserByUserName(username string) (u User, err error) {
	statement := "SELECT INTO users (id, uuid, user_name, email, password, created_at) WHERE user_name=$1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&u.Id, &u.Uuid, &u.UserName, &u.Email, &u.Password, &u.CreatedAt)
	return
}

func (u *User) Create() (err error) {
	s := "INSERT INTO users (uuid, user_name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(s)
	if err != nil {
		return
	}

	err = stmt.QueryRow(crypto.LongSecureRandomBase64(), u.UserName, u.Email, u.Password, time.Now()).Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	return
}
