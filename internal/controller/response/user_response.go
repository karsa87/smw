package response

import "github.com/evrone/go-clean-template/internal/entity"

type UserResponse struct {
	User []User `json:"users"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Gender  *string `json:"gender"`
	Address string  `json:"address"`
	Email   string  `json:"email"`
}

func (r UserResponse) Make(user entity.User) User {
	return User{
		ID:      user.ID,
		Name:    user.Name,
		Gender:  user.GetGenderLabel(),
		Address: user.Address,
		Email:   user.Email,
	}
}

func (r UserResponse) Makes(users []entity.User) UserResponse {
	responses := make([]User, 0, len(users))
	for _, user := range users {
		responses = append(responses, r.Make(user))
	}

	r.User = responses

	return r
}
