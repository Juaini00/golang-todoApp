package entity

import (
	"context"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id uint) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

type UserUsecase interface {
	Register(ctx context.Context, name, username, password string) (*User, error)
	Login(ctx context.Context, username string, password string) (*User, string, error)
}
