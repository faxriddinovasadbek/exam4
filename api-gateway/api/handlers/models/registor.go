package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)


type UserResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	Username     string `json:"user_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	// CreatedAt    string `json:"created_at"`
	// UpdatedAt    string `json:"udpated_at"`
}

func (rm *RegisterUser) Validate() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Email, validation.Required, is.Email),
		validation.Field(
			&rm.Password,
			validation.Required,
			validation.Length(8, 30),
			validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]")),
		),
	)
}



// type RegisterModel struct {
// 	FullName string `json:"full_name"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
