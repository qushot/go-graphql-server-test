package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	c *sql.DB
}

func NewClient() (*Client, error) {
	// エラーハンドリングは一旦適当（エラーが発生するとシステムが止まる）

	db, err := sql.Open("mysql", "api:apiapi@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatalf("sql.Open error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping error: %v", err)
	}

	return &Client{
		c: db,
	}, nil
}
