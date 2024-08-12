package request

type UserStore struct {
	Name     string `json:"name" binding:"required"`
	Gender   string `json:"gender" binding:"oneof=male female"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetJsonFieldName will return json tag name
func (req *UserStore) GetJsonFieldName(field string) string {
	return map[string]string{
		"Name":   "name",
		"Gender": "gender",
	}[field]
}

// ErrMessages contains map of error messages
func (req *UserStore) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"name": {
			"required": "Name is required",
		},
		"gender": {
			"oneof": "Gender not allowed value",
		},
	}
}
