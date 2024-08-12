// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Usecase -.
	Auth interface {
		Login(context.Context, string, string) (entity.User, error)
	}
	User interface {
		User(context.Context) ([]entity.User, error)
		Create(context.Context, entity.User) (entity.User, error)
		FindUser(context.Context, int) (entity.User, error)
		Update(context.Context, int, entity.User) (entity.User, error)
		Delete(context.Context, int) error
		FindUserByPassword(context.Context, string) (entity.User, error)
	}
	Inventory interface {
		Inventory(context.Context) ([]entity.Inventory, error)
		Create(context.Context, entity.Inventory) (entity.Inventory, error)
		FindInventory(context.Context, int) (entity.Inventory, error)
		Update(context.Context, int, entity.Inventory) (entity.Inventory, error)
		Delete(context.Context, int) error
	}

	// Repo -.
	UserRepo interface {
		GetUser(context.Context) ([]entity.User, error)
		Create(context.Context, entity.User) error
		FindUser(context.Context, int) (entity.User, error)
		FindUserByEmail(context.Context, string) (entity.User, error)
		FindUserByPassword(context.Context, string) (entity.User, error)
		UpdateUserByModel(context.Context, entity.User, entity.User) (entity.User, error)
		DeleteUserByModel(context.Context, entity.User) error
	}
	InventoryRepo interface {
		GetInventory(context.Context) ([]entity.Inventory, error)
		Create(context.Context, entity.Inventory) error
		FindInventory(context.Context, int) (entity.Inventory, error)
		UpdateInventoryByModel(context.Context, entity.Inventory, entity.Inventory) (entity.Inventory, error)
		DeleteInventoryByModel(context.Context, entity.Inventory) error
	}
)
