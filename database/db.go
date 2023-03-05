package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./bank.db")
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	sqlStmt := `
	 CREATE TABLE IF NOT EXISTS accounts (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, balance INTEGER, agency VARCHAR(255), account VARCHAR(255));
	`
	_, err = DB.Exec(sqlStmt)

	if err != nil {
		panic(err)
	}

}
