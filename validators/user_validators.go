package validators

import (
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CreateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u CreateUser) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.FirstName, v.Length(2, 32)),
		v.Field(&u.LastName, v.Length(2, 32)),
		v.Field(&u.Username, v.Required, v.Length(2, 32)),
		v.Field(&u.Email, v.Required, is.Email),
		v.Field(&u.Password, v.Required, v.Length(6, 32)),
	)
}

type UpdateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

func (u UpdateUser) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.FirstName, v.Length(2, 32)),
		v.Field(&u.LastName, v.Length(2, 32)),
		v.Field(&u.Username, v.Length(2, 32)),
	)
}
