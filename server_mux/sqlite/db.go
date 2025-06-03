package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type sqliter interface {
	connect() error
	execute(query string, args ...interface{}) (sql.Result, error)
	init()
}

type sqlite struct {
	db *sql.DB
}

func (db *sqlite) connect() error {
	fmt.Println("Connecting to sqlite")
	sqlDB, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return err
	}

	db.db = sqlDB
	return nil
}

func (lite *sqlite) execute(query string, args ...interface{}) (sql.Result, error) {
	err := lite.connect()
	if err != nil {
		panic(err)
		return nil, err
	}

	result, err := lite.db.Exec(query, args...)
	if err != nil {
		panic(err)
		return nil, err
	}

	return result, nil
}

func (lite *sqlite) init() {
	err := lite.connect()
	if err != nil {
		panic(err)
	}

	_, err = lite.db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
		id TEXT PRIMARY KEY,
		email TEXT NOT NULL,
		password TEXT
	)`)

	if err != nil {
		panic(err)
	}
}
