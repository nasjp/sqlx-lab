package main

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Users []*User
