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
	 CREATE TABLE IF NOT EXISTS clients (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name VARCHAR(255), cpf VARCHAR(255), email VARCHAR(255), password VARCHAR(255));
	 CREATE TABLE IF NOT EXISTS transactions (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, account_origin_id INTEGER, account_destination_id INTEGER, value INTEGER)
	`
	_, err = DB.Exec(sqlStmt)

	if err != nil {
		panic(err)
	}

}
