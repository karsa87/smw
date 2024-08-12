// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Inventory -.
type Inventory struct {
	ID          int
	UserID      int
	Name        string
	Description string
	Stock       int32
	Price       float64

	User User `gorm:"foreignKey:UserID;references:ID;"`
}

func (Inventory) TableName() string {
	return "inventory"
}
