package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"todo_app/internal/domain/entity"
	"todo_app/internal/infrastructure/presistence"
)

var _ entity.TodoRepository = (*TodoRepositoryImpl)(nil)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) entity.TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) Save(ctx context.Context, todo *entity.Todo) error {
	model := presistence.TodoModelFromEntity(todo)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return fmt.Errorf("error saving todo: %w", err)
	}
	*todo = *model.ToEntity()
	return nil
}

func (r *TodoRepositoryImpl) FindByUserID(ctx context.Context, userID uint) ([]entity.Todo, error) {
	var todos []presistence.Todo
	if err := r.db.WithContext(ctx).Find(&todos, "user_id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("error finding todos: %w", err)
	}
	var result []entity.Todo
	for _, t := range todos {
		result = append(result, *t.ToEntity())
	}
	return result, nil
}
