package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/qushot/go-graphql-server-test/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoRepository *repository.Todo
}
