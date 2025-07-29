package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"todo_app/internal/domain/entity"
	"todo_app/internal/infrastructure/presistence"
)

var _ entity.UserDetailRepository = (*UserDetailRepositoryImpl)(nil)

type UserDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewUserDetailRepository(db *gorm.DB) entity.UserDetailRepository {
	return &UserDetailRepositoryImpl{db: db}
}

func (r *UserDetailRepositoryImpl) SaveOrUpdate(ctx context.Context, user *entity.UserDetail) error {
	db := r.db.WithContext(ctx)
	userDetailModel := presistence.UserDetailModelFromEntity(user)

	err := db.First(&userDetailModel, "access_token = ?", user.AccessToken).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

		if err := db.Create(&userDetailModel).Error; err != nil {
			return fmt.Errorf("error saving user detail: %s", err)
		}
	} else if err != nil {
		return fmt.Errorf("error saving user detail: %s", err)
	} else {
		if err := db.Save(&userDetailModel).Error; err != nil {
			return fmt.Errorf("error saving user detail: %s", err)
		}
	}

	return nil
}

func (r *UserDetailRepositoryImpl) FindByToken(ctx context.Context, token string) (*entity.UserDetail, error) {
	var userDetail presistence.UserDetail
	if err := r.db.WithContext(ctx).Find(&userDetail, "access_token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user detail with %s not found", err)
		}
		return nil, fmt.Errorf("error finding user detail: %s", err)
	}
	return userDetail.ToEntity(), nil
}
