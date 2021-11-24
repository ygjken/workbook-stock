package model

import (
	"time"
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
