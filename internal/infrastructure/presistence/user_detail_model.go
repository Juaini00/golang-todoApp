package presistence

import (
	"gorm.io/gorm"
	"time"
	"todo_app/internal/domain/entity"
)

type UserDetail struct {
	gorm.Model
	AccessToken  string `gorm:"not null"`
	RefreshToken string
}

func (u *UserDetail) ToEntity() *entity.UserDetail {
	if u == nil {
		return nil
	}

	return &entity.UserDetail{
		ID:           u.ID,
		RefreshToken: u.RefreshToken,
		AccessToken:  u.AccessToken,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		DeletedAt: func() *time.Time {
			if u.DeletedAt.Valid {
				return &u.DeletedAt.Time
			}
			return nil
		}(),
	}
}

func UserDetailModelFromEntity(u *entity.UserDetail) *UserDetail {
	if u == nil {
		return nil
	}

	return &UserDetail{
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
		RefreshToken: u.RefreshToken,
		AccessToken:  u.AccessToken,
	}
}
