package controllers

import (
	"fiber-rest-api/services"
	"github.com/gofiber/fiber"
)

type UserController struct{ userService *services.UserService }

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) GetMany(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented)
}

func (u *UserController) Create(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented)
}

func (u *UserController) Get(c *fiber.Ctx) {
	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Hello world!"})
}

func (u *UserController) Update(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented)
}

func (u *UserController) Delete(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented)
}
