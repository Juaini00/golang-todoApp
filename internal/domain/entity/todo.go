package entity

import (
	"context"
	"time"
)

type Todo struct {
	ID        uint
	UserID    uint
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type TodoRepository interface {
	Save(ctx context.Context, todo *Todo) error
	FindByUserID(ctx context.Context, userID uint) ([]Todo, error)
}

type TodoUsecase interface {
	CreateTodo(ctx context.Context, userID uint, title string) (*Todo, error)
	ListTodo(ctx context.Context, userID uint) ([]Todo, error)
}
