package services

import (
	"fiber-rest-api/lib"
	"fiber-rest-api/models"
	"fiber-rest-api/responses"
	"fiber-rest-api/utils"
	"fiber-rest-api/validators"
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

	users := new([]responses.UserPublic)

	if err = db.Model(new(models.UserModel)).Clauses(
		clause.Limit{Limit: limit, Offset: offset},
		//lib.LikeQuery("username", c.Query("username")),
		//lib.LikeQuery("first_name", c.Query("firstName")),
		//lib.LikeQuery("last_name", c.Query("lastName")),
	).Find(&users).Error; err != nil {
		return nil, utils.NewDBError(err)
	}

	if len(*users) == 0 {
		return nil, utils.NotFound()
	}

	return users, nil

}

func (u *UserService) CreateOneUser(body *validators.CreateUser) (*responses.UserCreatedPublic, error) {
	db := u.db

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

	return &responses.UserCreatedPublic{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *UserService) GetOneUserByID(id int) (*responses.UserPublic, error) {
	db := u.db

	user := new(responses.UserPublic)

	if err := db.Model(new(models.UserModel)).Clauses(clause.Eq{Column: "id", Value: id}).Find(user).Error; err != nil {
		return nil, utils.NewDBError(err)
	}

	if user.ID == 0 {
		return nil, utils.NotFound()
	}

	return user, nil
}

func (u *UserService) UpdateOneUserByID(id int, body *validators.UpdateUser) (*responses.UserPublic, error) {
	db := u.db

	user := new(responses.UserPublic)

	err := db.Transaction(func(tx *gorm.DB) error {

		// Check if user exists or not
		if err := tx.
			Model(new(models.UserModel)).
			Clauses(clause.Eq{Column: "id", Value: id}).
			Find(user).Error; err != nil {
			return utils.NewDBError(err)
		}

		// if not exists then return not found error
		if user.ID == 0 {
			return utils.NotFound()
		}

		// update user with the given fields. gorm will handle with zero values
		if err := tx.Clauses(clause.Eq{Column: "id", Value: id}).Updates(&models.UserModel{
			FirstName: body.FirstName,
			LastName:  body.LastName,
			Username:  body.Username,
		}).Error; err != nil {
			return utils.NewDBError(err)
		}

		// then find updated user
		if err := tx.Clauses(clause.Eq{Column: "id", Value: id}).Find(user).Error; err != nil {
			return utils.NewDBError(err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) DeleteOneUserByID(id int) error {
	db := u.db

	err := db.Transaction(func(tx *gorm.DB) error {
		userModel := new(models.UserModel)

		// check if user exists or not
		if err := tx.Clauses(clause.Eq{Column: "id", Value: id}).Find(userModel).Error; err != nil {
			return utils.NewDBError(err)
		}

		// if not exists then throw not found error
		if userModel.ID == 0 {
			return utils.NotFound()
		}

		// delete
		if err := tx.Clauses(clause.Eq{Column: "id", Value: id}).Delete(new(models.UserModel)).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
