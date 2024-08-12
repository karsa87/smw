package usecase

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// InventoryUseCase -.
type InventoryUseCase struct {
	repo InventoryRepo
}

// New -.
func NewInventoryUsecase(r InventoryRepo) *InventoryUseCase {
	return &InventoryUseCase{
		repo: r,
	}
}

// Inventory get inventory list from table inventory
func (uc *InventoryUseCase) Inventory(ctx context.Context) ([]entity.Inventory, error) {
	inventories, err := uc.repo.GetInventory(ctx)
	if err != nil {
		return nil, fmt.Errorf("InventoryUseCase - Inventory - s.repo.GetInventory: %w", err)
	}

	return inventories, nil
}

// Create inventory to table inventory
func (uc *InventoryUseCase) Create(ctx context.Context, inventory entity.Inventory) (entity.Inventory, error) {
	err := uc.repo.Create(ctx, inventory)
	if err != nil {
		return entity.Inventory{}, fmt.Errorf("InventoryUseCase - Create - s.repo.Create: %w", err)
	}

	return inventory, nil
}

// Update inventory to table inventory
func (uc *InventoryUseCase) FindInventory(ctx context.Context, id int) (entity.Inventory, error) {
	inventory, err := uc.repo.FindInventory(ctx, id)
	if err != nil {
		return inventory, fmt.Errorf("InventoryUseCase - FindInventory - s.repo.FindInventory: %w", err)
	}

	return inventory, nil
}

// Update inventory to table inventory
func (uc *InventoryUseCase) Update(ctx context.Context, id int, inventoryInput entity.Inventory) (entity.Inventory, error) {
	inventory, err := uc.FindInventory(ctx, id)
	if err != nil {
		return inventory, fmt.Errorf("InventoryUseCase - Update - s.repo.Update: %w", err)
	}

	resultUpdate, err := uc.repo.UpdateInventoryByModel(ctx, inventory, inventoryInput)
	if err != nil {
		return entity.Inventory{}, fmt.Errorf("InventoryUseCase - Update - s.repo.Update: %w", err)
	}

	return resultUpdate, nil
}

// Delete inventory from table inventory
func (uc *InventoryUseCase) Delete(ctx context.Context, id int) error {
	inventory, err := uc.FindInventory(ctx, id)
	if err != nil {
		return fmt.Errorf("InventoryUseCase - Delete - s.repo.Delete: %w", err)
	}

	err = uc.repo.DeleteInventoryByModel(ctx, inventory)
	if err != nil {
		return fmt.Errorf("InventoryUseCase - Delete - s.repo.Delete: %w", err)
	}

	return nil
}
