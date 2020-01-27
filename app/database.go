package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root@tcp(db:3306)/test_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
