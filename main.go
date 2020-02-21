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
	db, err := newDB()
	if err != nil {
		return err
	}
	defer db.Close()

	selectMethod(db)
	getMethod(db)

	return nil
}

func newDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root@tcp(db:3306)/test_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Users []*User

func selectMethod(db *sqlx.DB) {
	fmt.Println("start: select")
	var us Users
	err := db.Select(&us, `SELECT * from users`)
	if err == sql.ErrNoRows {
		// これは表示されない
		fmt.Println("err: err no row")
	}
	fmt.Println("done: select")
}

func getMethod(db *sqlx.DB) {
	fmt.Println("start: get")
	u := &User{}
	err := db.Get(u, `SELECT * from users limit 1`)
	if err == sql.ErrNoRows {
		// これは表示される
		fmt.Println("err: err no row")
	}
	fmt.Println("done: get")
}
