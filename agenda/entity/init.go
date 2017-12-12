package entity

import (
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var createUserTable = `
CREATE TABLE IF NOT EXISTS user (
	username VARCHAR(32),
	password VARCHAR(32)
);
`
var createMeetingTable = `
CREATE TABLE IF NOT EXISTS meeting (
	title VARCHAR(32),
	sponsor VARCHAR(32),
	participants TEXT,
	stime CHAR(16)
	etime CHAR(16)
);
`


func init() {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")


}