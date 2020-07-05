package services

import (
	"fiber-rest-api/forms"
	"fiber-rest-api/lib"
	"fiber-rest-api/models"
	"fiber-rest-api/responses"
	"fiber-rest-api/utils"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService struct{ db *gorm.DB }

func NewUserService(db *gorm.DB) *UserService { return &UserService{db: db} }

func (u *UserService) GetManyUser(c *fiber.Ctx) (*[]responses.UserPublic, error) {
	db := u.db

	offset, limit, err := lib.Paginate(c.Query("page", "1"), c.Query("limit", "20"))

	if err != nil {
		return nil, err
	}

	var users []models.UserModel

	if err = db.Clauses(
		clause.Limit{Limit: limit, Offset: offset},
		//utils.LikeQuery("username", c.Query("username")),
		//utils.LikeQuery("first_name", c.Query("firstName")),
		//utils.LikeQuery("last_name", c.Query("lastName")),
	).Find(&users).Error; err != nil {
		return nil, utils.NewDBError(err)
	}

	if len(users) == 0 {
		return nil, utils.NotFound()
	}

	var userResponse []responses.UserPublic

	for _, user := range users {
		userResponse = append(userResponse, responses.UserPublic{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		})
	}

	return &userResponse, nil
}

func (u *UserService) CreateOneUser(createUserForm *forms.CreateUser) (*responses.UserCreatedPublic, *utils.DBError) {
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

func (u *UserService) GetOneUserByID(id int) (*responses.UserPublic, error) {
	db := u.db

	var user models.UserModel

	if err := db.Clauses(clause.Eq{Column: "id", Value: id}).Find(&user).Error; err != nil {
		return nil, utils.NewDBError(err)
	}

	if user.ID == 0 {
		return nil, utils.NotFound()
	}

	return &responses.UserPublic{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.LastName,
		CreatedAt: user.CreatedAt,
	}, nil
}
