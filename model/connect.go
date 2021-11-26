package model

import (
	"database/sql"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Db *sql.DB

type config struct {
	dbConnection dbConnection `yaml:"dbConnection"`
	initUsers    []initUser   `yaml:"initUsers"`
}

type dbConnection struct {
	dbDriver string `yaml:"dbDriver"`
	dsn      string `yaml:"dsn"`
}

type initUser struct {
	email    string `yaml:"email"`
	name     string `yaml:"name"`
	password string `yaml:"password"`
}

func init() {
	buf, err := ioutil.ReadFile("config/dbConnect.yaml")
	if err != nil {
		log.Println("Cannot init database")
		log.Println(err)
	}

	var c config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		log.Println("Cannot unmarshal yaml file")
		log.Println(err)
	}
	log.Println(c)

	Db, err = sql.Open("postgres", "host=postgres user=pguser password=password dbname=workbookstock sslmode=disable")
	if err != nil {
		log.Println("Cannot connect to database")
		panic("Cannot to connect database")
	}
}
