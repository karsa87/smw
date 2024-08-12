package request

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// GetJsonFieldName will return json tag name
func (req *Login) GetJsonFieldName(field string) string {
	return map[string]string{
		"Email":    "email",
		"Password": "password",
	}[field]
}

// ErrMessages contains map of error messages
func (req *Login) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"email": {
			"required": "Email is required",
		},
		"password": {
			"required": "Password is required",
		},
	}
}
