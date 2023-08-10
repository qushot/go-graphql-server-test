package repository

import (
	"log"

	"github.com/qushot/go-graphql-server-test/graph/model"
)

type Todo struct {
	*Client
}

func NewTodo(c *Client) *Todo {
	return &Todo{
		Client: c,
	}
}

func (t *Todo) GetAll() ([]model.Todo, error) {
	rows, err := t.c.Query("SELECT id, text, done, user_id FROM todo LOCK IN SHARE MODE")
	if err != nil {
		log.Fatalf("'todo' t.c.Query error: %v", err)
	}
	defer rows.Close()

	var ts []model.Todo

	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Text, &t.Done, &t.UserID); err != nil {
			log.Fatalf("'todo' rows.Scan error: %v", err)
		}
		ts = append(ts, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("'todo' rows.Err() error: %v", err)
	}

	log.Printf("DONE: repository#Todo#GetAll")

	return ts, nil
}
