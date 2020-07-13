package controllers

import (
	"fiber-rest-api/services"
	"github.com/gofiber/fiber"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (a *AuthController) SignUp(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented)
}

func (a *AuthController) SignIn(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented)
}

func (a *AuthController) SignOut(c *fiber.Ctx) {
	c.Status(fiber.StatusNotImplemented)
}
