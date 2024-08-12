package usecase

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/pkg/hashing"
)

// UserUseCase -.
type AuthUseCase struct {
	repo UserRepo
}

// New -.
func NewAuthUsecase(r UserRepo) *AuthUseCase {
	return &AuthUseCase{
		repo: r,
	}
}

// Update user to table user
func (uc *AuthUseCase) Login(ctx context.Context, email string, password string) (entity.User, error) {
	user, err := uc.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return user, fmt.Errorf("UserUseCase - Login - s.repo.Login: %w", err)
	}

	match, err := hashing.ComparePasswords(user.Password, password)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - Login - s.repo.Login - Failed decode password: %w", err)
	}

	if !match {
		return entity.User{}, fmt.Errorf("UserUseCase - Login - s.repo.Login - Wrong Password: %w", err)
	}

	return user, nil
}
