package db

import (
	"arnov17/echo-test/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	// username:password@protocol(address:port)/dbname

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error")
	}

	err = db.Ping()

	if err != nil {
		panic("DSN invalid")
	}
}

func CreateCont() *sql.DB {
	return db
}
