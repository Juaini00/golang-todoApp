package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"todo_app/internal/domain/entity"
	"todo_app/internal/infrastructure/presistence"
)

var _ entity.UserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []presistence.User

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("error finding users: %w", err)
	}

	var entities []entity.User
	for _, u := range users {
		entities = append(entities, *u.ToEntity())
	}

	return entities, nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var user presistence.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	return user.ToEntity(), nil
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user presistence.User
	if err := r.db.WithContext(ctx).First(&user, "username = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with username %s not found", username)
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	return user.ToEntity(), nil
}

func (r *UserRepositoryImpl) Save(ctx context.Context, user *entity.User) error {
	userModel := presistence.UserModelFromEntity(user)
	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	*user = *userModel.ToEntity()
	return nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	userModel := presistence.UserModelFromEntity(user)
	if err := r.db.WithContext(ctx).Save(userModel).Error; err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	*user = *userModel.ToEntity()
	return nil
}
