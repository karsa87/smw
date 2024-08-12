package response

import "github.com/evrone/go-clean-template/internal/entity"

type InventoryResponse struct {
	Inventory []Inventory `json:"inventory"`
}

type Inventory struct {
	ID          int     `json:"id"`
	User        User    `json:"user"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int32   `json:"stock"`
	Price       float64 `json:"price"`
}

func (r InventoryResponse) Make(inventorys []entity.Inventory) InventoryResponse {
	responses := make([]Inventory, 0, len(inventorys))
	for _, inventory := range inventorys {
		responses = append(responses, Inventory{
			ID:          inventory.ID,
			Name:        inventory.Name,
			Description: inventory.Description,
			Stock:       inventory.Stock,
			Price:       inventory.Price,
			User: User{
				ID:      inventory.User.ID,
				Name:    inventory.User.Name,
				Gender:  inventory.User.GetGenderLabel(),
				Address: inventory.User.Address,
			},
		})
	}

	r.Inventory = responses

	return r
}
