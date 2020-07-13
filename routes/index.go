package routes

import "github.com/gofiber/fiber"

func InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	userRoutes(v1.Group("/users"))
	authRoutes(v1.Group("/auth"))
}
