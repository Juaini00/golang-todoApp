package presistence

import (
	"gorm.io/gorm"
	"time"
	"todo_app/internal/domain/entity"
)

type Todo struct {
	gorm.Model
	UserID    uint
	Title     string
	Completed bool
}

func (t *Todo) ToEntity() *entity.Todo {
	if t == nil {
		return nil
	}
	return &entity.Todo{
		ID:        t.ID,
		UserID:    t.UserID,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		DeletedAt: func() *time.Time {
			if t.DeletedAt.Valid {
				return &t.DeletedAt.Time
			}
			return nil
		}(),
	}
}

func TodoModelFromEntity(t *entity.Todo) *Todo {
	if t == nil {
		return nil
	}
	return &Todo{
		Model: gorm.Model{
			ID:        t.ID,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			DeletedAt: gorm.DeletedAt{
				Time: func() time.Time {
					if t.DeletedAt != nil {
						return *t.DeletedAt
					}
					return time.Time{}
				}(),
				Valid: t.DeletedAt != nil,
			},
		},
		UserID:    t.UserID,
		Title:     t.Title,
		Completed: t.Completed,
	}
}
