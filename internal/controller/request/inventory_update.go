package request

type InventoryUpdate struct {
	Name        string  `json:"name" binding:"required"`
	UserID      int     `json:"user_id" binding:"required"`
	Stock       int32   `json:"stock" binding:"number"`
	Price       float64 `json:"price" binding:"number"`
	Description string  `json:"description"`
}

// GetJsonFieldName will return json tag name
func (req *InventoryUpdate) GetJsonFieldName(field string) string {
	return map[string]string{
		"Name":        "name",
		"UserID":      "user_id",
		"Stock":       "stock",
		"Price":       "price",
		"Description": "description",
	}[field]
}

// ErrMessages contains map of error messages
func (req *InventoryUpdate) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"name": {
			"required": "Name is required",
		},
		"user_id": {
			"required": "User ID is required",
		},
		"stock": {
			"number": "Stock must a number",
		},
		"price": {
			"number": "Price must a number",
		},
	}
}
