package routes

import (
	"fiber-rest-api/controllers"
	"fiber-rest-api/models"
	"fiber-rest-api/services"
	"github.com/gofiber/fiber"
)

func userRoutes(router *fiber.Group) {
	userController := controllers.NewUserController(services.NewUserService(models.GetDB()))

	router.Get(
		"/",
		userController.GetMany,
	)

	router.Post(
		"/",
		userController.Create,
	)

	router.Get(
		"/:id",
		userController.Get,
	)

	router.Patch(
		"/:id",
		userController.Update,
	)

	router.Delete(
		"/:id",
		userController.Delete,
	)
}
