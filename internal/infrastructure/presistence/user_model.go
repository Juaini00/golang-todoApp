package presistence

import (
	"gorm.io/gorm"
	"time"
	"todo_app/internal/domain/entity"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:100;not null"`
	Username string `gorm:"uniqueIndex; not null"`
	Password string `gorm:"not null"`
}

func (u *User) ToEntity() *entity.User {
	if u == nil {
		return nil
	}

	return &entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: func() *time.Time {
			if u.DeletedAt.Valid {
				return &u.DeletedAt.Time
			}
			return nil
		}(),
	}
}

func UserModelFromEntity(u *entity.User) *User {
	if u == nil {
		return nil
	}

	return &User{
		Model: gorm.Model{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: gorm.DeletedAt{
				Time: func() time.Time {
					if u.DeletedAt != nil {
						return *u.DeletedAt
					}
					return time.Time{}
				}(),
				Valid: u.DeletedAt != nil,
			},
		},
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
	}
}
