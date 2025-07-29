package usecase

import (
	"context"
	"fmt"
	"todo_app/internal/domain/entity"
)

var _ entity.TodoUsecase = (*TodoUsecaseImpl)(nil)

type TodoUsecaseImpl struct {
	repo entity.TodoRepository
}

func NewTodoUsecase(repo entity.TodoRepository) TodoUsecaseImpl {
	return TodoUsecaseImpl{repo: repo}
}

func (u *TodoUsecaseImpl) CreateTodo(ctx context.Context, userID uint, title string) (*entity.Todo, error) {
	todo := &entity.Todo{
		UserID: userID,
		Title:  title,
	}
	if err := u.repo.Save(ctx, todo); err != nil {
		return nil, fmt.Errorf("error creating todo: %w", err)
	}
	return todo, nil
}

func (u *TodoUsecaseImpl) ListTodo(ctx context.Context, userID uint) ([]entity.Todo, error) {
	todos, err := u.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting todos: %w", err)
	}
	return todos, nil
}
