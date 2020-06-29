package services

import (
	"fiber-rest-api/forms"
	"fiber-rest-api/models"
	"fiber-rest-api/responses"
	"fiber-rest-api/utils"
	"gorm.io/gorm"
)

type UserService struct{ db *gorm.DB }

func NewUserService(db *gorm.DB) *UserService { return &UserService{db: db} }

func (u *UserService) GetManyUser() {
	panic("implement me")
}

func (u *UserService) CreateUser(createUserForm *forms.CreateUser) (*responses.UserCreatedPublic, *utils.DBError) {
	db := u.db

	userModel := models.UserModel{
		FirstName: createUserForm.FirstName,
		LastName:  createUserForm.LastName,
		Username:  createUserForm.Username,
		Email:     createUserForm.Email,
		Password:  createUserForm.Password,
	}

	if err := db.Create(&userModel).Error; err != nil {
		return nil, utils.NewDBError(err)
	}

	return &responses.UserCreatedPublic{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Username:  userModel.Username,
		CreatedAt: userModel.CreatedAt,
	}, nil
}
