package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	db, err := NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	e := NewExperiment(db)

	e.GetMethod()

	return nil
}
