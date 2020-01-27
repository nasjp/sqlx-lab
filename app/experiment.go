package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type Experiment struct {
	DB *sqlx.DB
}

func NewExperiment(db *sqlx.DB) *Experiment {
	return &Experiment{
		DB: db,
	}
}

func (e *Experiment) SelectMethod() {
	fmt.Println("start: select")
	var us Users
	err := e.DB.Select(&us, `SELECT * from users`)
	if err == sql.ErrNoRows {
		// これは表示されない
		fmt.Println("err: err no row")
	}
	fmt.Println("done: select")
}

func (e *Experiment) GetMethod() {
	fmt.Println("start: get")
	u := &User{}
	err := e.DB.Get(u, `SELECT * from users limit 1`)
	if err == sql.ErrNoRows {
		// これは表示される
		fmt.Println("err: err no row")
	}
	fmt.Println("done: get")
}

func (e *Experiment) XError() error {
	fmt.Println("start: get")
	u := &User{}
	err := e.DB.Get(u, `SELECT * from users limit 1`)
	if err == sql.ErrNoRows {
		return xerrors.Errorf("XError: %w", err)
	}
	if err != nil {
		return xerrors.Errorf("XError: %w", err)
	}
	return nil
}
