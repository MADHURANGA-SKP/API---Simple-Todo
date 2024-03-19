// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteTodo(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, userName string) (Account, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTodo(ctx context.Context, id int64) (Todo, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Account, error)
	ListTodo(ctx context.Context, arg ListTodoParams) ([]Todo, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error)
}

var _ Querier = (*Queries)(nil)
