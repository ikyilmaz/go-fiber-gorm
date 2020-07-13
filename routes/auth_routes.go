package routes

import (
	"fiber-rest-api/controllers"
	"fiber-rest-api/models"
	"fiber-rest-api/services"
	"github.com/gofiber/fiber"
)

func authRoutes(router *fiber.Group) {
	authController := controllers.NewAuthController(services.NewAuthService(models.GetDB()))

	router.Post(
		"/sign-up",
		authController.SignUp,
	)

	router.Post(
		"/sign-in",
		authController.SignIn,
	)

	router.Get(
		"/sign-out",
		authController.SignOut,
	)
}
