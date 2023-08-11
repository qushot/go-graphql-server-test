package repository

import (
	"fmt"
	"log"
	"math/rand"

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
	log.Printf("count: %v", len(ts))

	return ts, nil
}

func (t *Todo) Create(input model.NewTodo) (*model.Todo, error) {
	stmt, err := t.c.Prepare("INSERT INTO todo (id, text, user_id) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("'todo' t.c.Prepare error: %v", err)
	}
	defer stmt.Close()

	id := fmt.Sprintf("T%09d", rand.Intn(999999999))

	if _, err := stmt.Exec(id, input.Text, input.UserID); err != nil {
		log.Fatalf("'todo' stmt.Exec error: %v", err)
	}

	m := model.Todo{
		ID:     id,
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}

	log.Printf("DONE: repository#Todo#Create")
	log.Printf("m: %+v", m)

	return &m, nil
}
