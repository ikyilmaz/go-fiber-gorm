package services

import (
	"fiber-rest-api/models"
	"fiber-rest-api/responses"
	"fiber-rest-api/utils"
	"fiber-rest-api/validators"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthService struct{ db *gorm.DB }

func NewAuthService(db *gorm.DB) *AuthService { return &AuthService{db: db} }

func (s *AuthService) SignUp(body *validators.SignUp) (*responses.UserCreatedPrivate, error) {
	db := s.db

	user := models.UserModel{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Username:  body.Username,
		Email:     body.Email,
		Password:  body.Password,
	}

	if err := db.Model(new(models.UserModel)).Create(&user).Error; err != nil {
		return nil, utils.NewDBError(err)
	}

	return &responses.UserCreatedPrivate{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *AuthService) SignIn(body *validators.SignIn) (*responses.UserPrivate, string, error) {
	db := s.db

	var cl clause.Expression

	if body.Email != "" {
		cl = clause.Eq{Column: "email", Value: body.Email}
	}

	if body.Username != "" {
		cl = clause.Eq{Column: "username", Value: body.Username}
	}

	user := new(models.UserModel)

	if err := db.Clauses(cl).Find(user).Error; err != nil {
		return nil, "", utils.NewDBError(err)
	}

	return &responses.UserPrivate{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, user.Password, nil
}
