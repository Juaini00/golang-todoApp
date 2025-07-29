package entity

import (
	"context"
	"time"
)

type UserDetail struct {
	ID           uint
	AccessToken  string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type UserDetailRepository interface {
	SaveOrUpdate(ctx context.Context, user *UserDetail) error
	FindByToken(ctx context.Context, token string) (*UserDetail, error)
}
