package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func InitDBArg(dbType string, username string, password string, url string) {
	db, err := sql.Open(dbType, fmt.Sprintf("%v:%v@(%v)/dbname?parseTime=true", username, password, url))

	if (err == nil) {
		db.Ping()
		return
	}

	log.Fatal(err)
}

func InitDB() {
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

	if (err == nil) {
		db.Ping()
		return
	}

	log.Fatal(err)
}

