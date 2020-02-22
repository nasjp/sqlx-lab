package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	db, err := sqlx.Open("mysql", "root@tcp(db:3306)/test_db")
	if err != nil {
		return err
	}
	defer db.Close()

	selectStructs(db)
	selectInts(db)
	selectIntsWithScannable(db)
	getStruct(db)

	return nil
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Users []*User

type id uint

func (id) Scan(src interface{}) error {
	return nil
}

func selectStructs(db *sqlx.DB) {
	var us Users
	err := db.Select(&us, `SELECT * from users`)
	fmt.Printf("- db.Select(&us, `SELECT * from users`)\n  err == sql.ErrNoRows => %t\n", err == sql.ErrNoRows)
}

func selectInts(db *sqlx.DB) {
	var ids []uint
	err := db.Select(&ids, `SELECT id from users`)
	fmt.Printf("- db.Select(&ids, `SELECT id from users`)\n  err == sql.ErrNoRows => %t\n", err == sql.ErrNoRows)
}

func selectIntsWithScannable(db *sqlx.DB) {
	var ids []id
	err := db.Select(&ids, `SELECT id from users`)
	fmt.Printf("- db.Select(&ids, `SELECT id from users`)\n  err == sql.ErrNoRows => %t\n", err == sql.ErrNoRows)
}

func getStruct(db *sqlx.DB) {
	u := &User{}
	err := db.Get(u, `SELECT * from users limit 1`)
	fmt.Printf("- db.Get(u, `SELECT * from users limit 1`)\n  err == sql.ErrNoRows => %t\n", err == sql.ErrNoRows)
}
