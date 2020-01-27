package main

import (
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Printf("%+v\n", err)
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

	e.SelectMethod()
	e.GetMethod()
	if err := e.XError(); err != nil {
		return xerrors.Errorf("Run: %w", err)
	}

	return nil
}
