package controllers

import (
	"fiber-rest-api/forms"
	"fiber-rest-api/services"
	"fiber-rest-api/utils"
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
	createUserForm := new(forms.CreateUser)

	if err := c.BodyParser(createUserForm); err != nil {
		c.Next(err)
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
	c.Status(fiber.StatusNotImplemented).JSON(utils.NewAPIError(fiber.StatusNotImplemented, "Not Implemented"))
}

func (u *UserController) Delete(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented).JSON(utils.NewAPIError(fiber.StatusNotImplemented, "Not Implemented"))
}
