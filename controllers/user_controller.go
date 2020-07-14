package controllers

import (
	"fiber-rest-api/services"
	"fiber-rest-api/utils"
	"fiber-rest-api/validators"
	"github.com/gofiber/fiber"
)

type UserController struct{ userService *services.UserService }

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) GetMany(c *fiber.Ctx) {
	users, err := u.userService.GetManyUser(c)

	if err != nil {
		c.Next(err)
		return
	}

	c.Status(fiber.StatusOK).JSON(utils.OK(users))
}

func (u *UserController) Create(c *fiber.Ctx) {
	createUserForm := new(validators.CreateUser)

	if err := utils.ValidateBody(createUserForm, c); err != nil {
		c.Next(utils.BadRequest(err.Error()))
		return
	}

	userCreatedPublic, err := u.userService.CreateOneUser(createUserForm)

	if err != nil {
		c.Next(err)
		return
	}

	c.Status(fiber.StatusCreated).JSON(utils.Created(userCreatedPublic))
}

func (u *UserController) Get(c *fiber.Ctx) {
	userResponse, err := u.userService.GetOneUserByID(c.Locals("id").(int))

	if err != nil {
		c.Next(err)
		return
	}

	c.Status(fiber.StatusOK).JSON(utils.OK(userResponse))
}

func (u *UserController) Update(c *fiber.Ctx) {
	updateUserForm := new(validators.UpdateUser)

	if err := utils.ValidateBody(updateUserForm, c); err != nil {
		c.Next(err)
		return
	}

	userResponse, err := u.userService.UpdateOneUserByID(c.Locals("id").(int), updateUserForm)

	if err != nil {
		c.Next(err)
		return
	}

	c.Status(fiber.StatusOK).JSON(utils.OK(userResponse))
}

func (u *UserController) Delete(c *fiber.Ctx) {
	err := u.userService.DeleteOneUserByID(c.Locals("id").(int))

	if err != nil {
		c.Next(err)
		return
	}

	c.Status(fiber.StatusNoContent).JSON(utils.NoContent(nil))
}
