package middlewares

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func InitMiddleWares(app *fiber.App) {
	app.Use(middleware.Logger())
	app.Use(middleware.RequestID("X-Request-Id"))
}
