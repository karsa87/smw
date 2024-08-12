package repo

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/pkg/hashing"
	"gorm.io/gorm"
)

// UserRepo -.
type UserRepo struct {
	*gorm.DB
}

// New -.
func NewGormUser(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// GetHistory -.
func (r *UserRepo) GetUser(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	if err := r.DB.Table("user").Find(&users).Error; err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

// Create User
func (r *UserRepo) Create(ctx context.Context, user entity.User) error {
	password, err := hashing.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = password

	if err := r.DB.Table("user").Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// Find User
func (r *UserRepo) FindUser(ctx context.Context, id int) (entity.User, error) {
	user := entity.User{}

	if err := r.DB.Table("user").First(&user, id).Error; err != nil && err != gorm.ErrRecordNotFound {
		return entity.User{}, err
	}

	return user, nil
}

// Find User
func (r *UserRepo) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}

	if err := r.DB.First(&user, "email = ?", email).Error; err != nil && err != gorm.ErrRecordNotFound {
		return entity.User{}, err
	}

	return user, nil
}

// Find User
func (r *UserRepo) FindUserByPassword(ctx context.Context, password string) (entity.User, error) {
	user := entity.User{}

	if err := r.DB.First(&user, "password = ?", password).Error; err != nil && err != gorm.ErrRecordNotFound {
		return entity.User{}, err
	}

	return user, nil
}

// Update User
func (r *UserRepo) UpdateUserByModel(ctx context.Context, user entity.User, userInput entity.User) (entity.User, error) {
	user.Name = userInput.Name
	user.Gender = userInput.Gender
	user.Address = userInput.Address
	user.Email = userInput.Email
	password, err := hashing.HashPassword(userInput.Password)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = password

	// Check if any error exists
	if err = r.DB.Table("user").Save(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return entity.User{}, err
	}

	return user, nil
}

// Delete User
func (r *UserRepo) DeleteUserByModel(ctx context.Context, user entity.User) error {
	if err := r.DB.Table("user").Delete(&user, user.ID).Error; err != nil {
		return err
	}

	return nil
}
