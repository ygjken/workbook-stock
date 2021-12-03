package model

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/ygjken/workbook-stock/crypto"
	"gopkg.in/yaml.v2"
)

var Db *sql.DB

type Config struct {
	DbConnection DbConnection `yaml:"dbConnection"`
	InitUsers    []InitUser   `yaml:"initUsers"`
}

type DbConnection struct {
	DbDriver string `yaml:"dbDriver"`
	Dsn      string `yaml:"dsn"`
}

type InitUser struct {
	Email    string `yaml:"email"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

func init() {
	buf, err := ioutil.ReadFile("config/dbConnect.yaml")
	if err != nil {
		log.Println("Cannot init database")
		log.Println(err)
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		log.Println("Cannot unmarshal yaml file")
		log.Println(err)
	}
	log.Println(c)

	// connect to DB
	Db, err = sql.Open(c.DbConnection.DbDriver, c.DbConnection.Dsn)
	if err != nil {
		log.Println("Cannot connect to database")
		panic("Cannot to connect database")
	}

	// insert test user
	encodedPw, err := crypto.PasswordEncrypt(c.InitUsers[0].Password)
	if err != nil {
		log.Println("Cannot set init user: ", err)
	}
	u := User{
		Email:    c.InitUsers[0].Email,
		UserName: c.InitUsers[0].Name,
		Password: encodedPw,
	}

	if err = u.Create(); err != nil {
		log.Println("Cannot set init user: ", err)
	}
}
