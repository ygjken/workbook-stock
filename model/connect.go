package model

import (
	"database/sql"
	"log"

	"github.com/ygjken/workbook-stock/crypto"
)

var Db *sql.DB

func InitDb() {

	var err error

	// connect to DB
	Db, err = sql.Open("postgres", "host=postgres user=pguser password=password dbname=workbookstock sslmode=disable")
	if err != nil {
		log.Println("Cannot connect to database")
		panic("Cannot to connect database")
	}

	// insert test user
	encodedPw, err := crypto.PasswordEncrypt("admintest")
	if err != nil {
		log.Println("Cannot set init user: ", err)
	}
	u := User{
		Email:    "tester@admin.com",
		UserName: "tester",
		Password: encodedPw,
	}
	if err = u.Create(); err != nil {
		log.Println("Cannot set init user: ", err)
	}

}
