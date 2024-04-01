package storage

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
	Website  string `json:"website"`
	Bio      string `json:"bio"`
}

type Post struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
	Like     string `json:"likes"`
	Dislikes string `json:"dislikes"`
	Views    string `json:"views"`
	Category string `json:"category"`
	OwnerId  string `json:"owner_id"`
}

type Comment struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	PostId  string `json:"post_id"`
	OwnerId string `json:"owner_id"`
}

// User info validation
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Name, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[A-Z][a-z]*$"))),
	)
}

type Message struct {
	Message string `json:"message"`
}
