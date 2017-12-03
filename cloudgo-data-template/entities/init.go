package entities

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

var mydb *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
