package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
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
	var us Users
	err := e.DB.Select(&us, `SELECT * from users`)
	spew.Dump(err)
	spew.Dump(us)
}

func (e *Experiment) GetMethod() {
	u := &User{}
	err := e.DB.Get(u, `SELECT * from users limit 1`)
	spew.Dump(err)
	spew.Dump(us)
}
