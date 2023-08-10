package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type TestUser struct {
	ID      int
	Name    string
	Address string
}

func main() {
	fmt.Println("BEGIN")
	db, err := sql.Open("mysql", "api:apiapi@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatalf("sql.Open error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping error: %v", err)
	}

	rows, err := db.Query("SELECT * FROM testuser LOCK IN SHARE MODE")
	if err != nil {
		log.Fatalf("db.Query error: %v", err)
	}
	defer rows.Close()

	var us []TestUser

	for rows.Next() {
		var u TestUser
		if err := rows.Scan(&u.ID, &u.Name, &u.Address); err != nil {
			log.Fatalf("rows.Scan error: %v", err)
		}
		us = append(us, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("rows.Err() error: %v", err)
	}

	for _, u := range us {
		fmt.Printf("ID: %d, Name: %s, Address: %s\n", u.ID, u.Name, u.Address)
	}

	fmt.Println("DONE")
}
