package model

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=postgres user=pguser password=password dbname=workbookstock sslmode=disable")
	if err != nil {
		log.Println("Cannot connect to database")
		panic("Cannot to connect database")
	}
}
