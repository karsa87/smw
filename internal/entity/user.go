// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

const (
	GENDER_MALE   = "male"
	GENDER_FEMALE = "female"
)

// User -.
type User struct {
	ID       int
	Name     string
	Gender   string
	Address  string
	Email    string
	Password string
}

func (User) TableName() string {
	return "user"
}

func (r User) GetGenderLabel() *string {
	var result string
	if r.Gender == GENDER_MALE {
		result = "Laki-Laki"
	} else if r.Gender == GENDER_FEMALE {
		result = "Perempuan"
	} else {
		return nil
	}

	return &result
}
