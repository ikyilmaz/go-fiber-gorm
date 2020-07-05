package routes

import (
	"fiber-rest-api/controllers"
	"fiber-rest-api/models"
	"fiber-rest-api/pipes"
	"fiber-rest-api/services"
	"fiber-rest-api/validators"
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
		validators.ValidateCreateUserForm,
		userController.Create,
	)

	router.Get(
		"/:id",
		pipes.ParseIntPipe(),
		userController.Get,
	)

	router.Patch(
		"/:id",
		pipes.ParseIntPipe(),
		validators.ValidateUpdateUserForm,
		userController.Update,
	)

	router.Delete(
		"/:id",
		pipes.ParseIntPipe(),
		userController.Delete,
	)
}
