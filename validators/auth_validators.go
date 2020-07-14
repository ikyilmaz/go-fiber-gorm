package validators

import (
	"errors"
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SignUp struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (s *SignUp) Validate() error {
	return v.ValidateStruct(s,
		v.Field(&s.FirstName, v.Length(2, 32)),
		v.Field(&s.LastName, v.Length(2, 32)),
		v.Field(&s.Username, v.Required, v.Length(2, 32)),
		v.Field(&s.Email, v.Required, is.Email),
		v.Field(&s.Password, v.Required, v.By(func(val interface{}) error {
			if v, ok := val.(string); !ok || v != s.PasswordConfirm {
				return errors.New("must equal with passwordConfirm")
			}

			return nil
		})),
		v.Field(&s.PasswordConfirm, v.Required),
	)
}

type SignIn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *SignIn) Validate() error {
	err := v.ValidateStruct(s,
		v.Field(&s.Username, v.When(s.Email == "", v.Required), v.Length(2, 32)),
		v.Field(&s.Email, v.When(s.Username == "", v.Required), is.Email),
		v.Field(&s.Password, v.Required),
	)

	if s.Email != "" && s.Username != "" {
		return errors.New("provide only one field: email, username")
	}

	return err
}
