package repository

import (
	"log"

	"github.com/qushot/go-graphql-server-test/graph/model"
)

type User struct {
	*Client
}

func NewUser(c *Client) *User {
	return &User{
		Client: c,
	}
}

func (u *User) SelectById(userId string) (*model.User, error) {
	log.Printf("START: repository#User#SelectById")

	stmt, err := u.c.Prepare("SELECT id, name FROM user WHERE id = ? LOCK IN SHARE MODE")
	if err != nil {
		log.Fatalf("'user' u.c.Prepare error: %v", err)
	}
	defer stmt.Close()

	var m model.User
	if err := stmt.QueryRow(userId).Scan(&m.ID, &m.Name); err != nil {
		log.Fatalf("'user' stmt.QueryRow error: %v", err)
	}

	log.Printf("m: %+v", m)
	log.Printf("DONE: repository#User#SelectById")

	return &m, nil
}
