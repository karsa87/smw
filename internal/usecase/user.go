package usecase

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepo
}

// New -.
func NewUserUsecase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

// User get user list from table user
func (uc *UserUseCase) User(ctx context.Context) ([]entity.User, error) {
	users, err := uc.repo.GetUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - User - s.repo.GetUser: %w", err)
	}

	return users, nil
}

// Create user to table user
func (uc *UserUseCase) Create(ctx context.Context, user entity.User) (entity.User, error) {
	err := uc.repo.Create(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - Create - s.repo.Create: %w", err)
	}

	return user, nil
}

// Update user to table user
func (uc *UserUseCase) FindUser(ctx context.Context, id int) (entity.User, error) {
	user, err := uc.repo.FindUser(ctx, id)
	if err != nil {
		return user, fmt.Errorf("UserUseCase - FindUser - s.repo.FindUser: %w", err)
	}

	return user, nil
}

// Update user to table user
func (uc *UserUseCase) FindUserByPassword(ctx context.Context, password string) (entity.User, error) {
	user, err := uc.repo.FindUserByPassword(ctx, password)
	if err != nil {
		return user, fmt.Errorf("UserUseCase - FindUser - s.repo.FindUserByPassword: %w", err)
	}

	return user, nil
}

// Update user to table user
func (uc *UserUseCase) Update(ctx context.Context, id int, userInput entity.User) (entity.User, error) {
	user, err := uc.FindUser(ctx, id)
	if err != nil {
		return user, fmt.Errorf("UserUseCase - Update - s.repo.Update: %w", err)
	}

	resultUpdate, err := uc.repo.UpdateUserByModel(ctx, user, userInput)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - Update - s.repo.Update: %w", err)
	}

	return resultUpdate, nil
}

// Delete user from table user
func (uc *UserUseCase) Delete(ctx context.Context, id int) error {
	user, err := uc.FindUser(ctx, id)
	if err != nil {
		return fmt.Errorf("UserUseCase - Delete - s.repo.Delete: %w", err)
	}

	err = uc.repo.DeleteUserByModel(ctx, user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Delete - s.repo.Delete: %w", err)
	}

	return nil
}
