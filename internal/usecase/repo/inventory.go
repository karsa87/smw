package repo

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
	"gorm.io/gorm"
)

// InventoryRepo -.
type InventoryRepo struct {
	*gorm.DB
}

// New -.
func NewGormInventory(db *gorm.DB) *InventoryRepo {
	return &InventoryRepo{db}
}

// GetHistory -.
func (r *InventoryRepo) GetInventory(ctx context.Context) ([]entity.Inventory, error) {
	var inventories []entity.Inventory

	if err := r.DB.Table("inventory").Preload("User").Find(&inventories).Error; err != nil {
		return []entity.Inventory{}, err
	}

	return inventories, nil
}

// Create Inventory
func (r *InventoryRepo) Create(ctx context.Context, inventory entity.Inventory) error {
	if err := r.DB.Table("inventory").Create(&inventory).Error; err != nil {
		return err
	}

	return nil
}

// Find Inventory
func (r *InventoryRepo) FindInventory(ctx context.Context, id int) (entity.Inventory, error) {
	inventory := entity.Inventory{}
	if err := r.DB.Table("inventory").Preload("User").First(&inventory, id).Error; err != nil && err != gorm.ErrRecordNotFound {
		return entity.Inventory{}, err
	}

	return inventory, nil
}

// Update Inventory
func (r *InventoryRepo) UpdateInventoryByModel(ctx context.Context, inventory entity.Inventory, inventoryInput entity.Inventory) (entity.Inventory, error) {
	err := r.DB.Save(entity.Inventory{
		ID:          inventory.ID,
		Name:        inventoryInput.Name,
		Description: inventoryInput.Description,
		Stock:       inventoryInput.Stock,
		Price:       inventoryInput.Price,
		UserID:      inventoryInput.UserID,
	}).Error

	// Check if any error exists
	if err != nil && err != gorm.ErrRecordNotFound {
		return entity.Inventory{}, err
	}

	return inventory, err
}

// Delete Inventory
func (r *InventoryRepo) DeleteInventoryByModel(ctx context.Context, inventory entity.Inventory) error {
	if err := r.DB.Table("inventory").Delete(&inventory, inventory.ID).Error; err != nil {
		return err
	}

	return nil
}
