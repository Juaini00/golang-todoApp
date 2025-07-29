package usecase

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"todo_app/internal/domain/entity"
	"todo_app/pkg/utils"
)

var _ entity.UserUsecase = (*UserUsecaseImpl)(nil)

type UserUsecaseImpl struct {
	repository           entity.UserRepository
	userDetailRepository entity.UserDetailRepository
}

func NewUserUsecase(repo entity.UserRepository, userDetailRepo entity.UserDetailRepository) UserUsecaseImpl {
	return UserUsecaseImpl{repository: repo, userDetailRepository: userDetailRepo}
}

func (u *UserUsecaseImpl) Register(ctx context.Context, name, username, password string) (*entity.User, error) {

	existingUser, err := u.repository.FindByUsername(ctx, username)
	if err != nil || existingUser != nil {
		return nil, fmt.Errorf("username already exists: %s", username)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	newUser := &entity.User{
		Name:     name,
		Username: username,
		Password: string(hashedPassword),
	}

	if err := u.repository.Save(ctx, newUser); err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}

	return newUser, nil
}

func (u *UserUsecaseImpl) Login(ctx context.Context, username string, password string) (*entity.User, string, error) {

	existingUser, err := u.repository.FindByUsername(ctx, username)
	if err != nil || existingUser == nil {
		return nil, "", fmt.Errorf("username does not exists: %s", username)
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if errBcrypt != nil {
		return nil, "", fmt.Errorf("password does not match: %w", err)
	}

	payload := &entity.User{
		ID: existingUser.ID,
	}
	token, tokenErr := utils.EncryptToken(payload)
	if tokenErr != nil {
		return nil, tokenErr.Error(), nil
	}

	userDetail := &entity.UserDetail{
		AccessToken: token,
	}

	if err := u.userDetailRepository.SaveOrUpdate(ctx, userDetail); err != nil {
		return nil, "", fmt.Errorf("error saving user detail: %s", err)
	}

	return existingUser, token, nil
}
