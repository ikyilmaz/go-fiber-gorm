package main

import (
	"fiber-rest-api/lib"
	"fiber-rest-api/middlewares"
	"fiber-rest-api/models"
	"fiber-rest-api/routes"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	lib.CheckErr(godotenv.Load(".env"), "while loading environment variables")
}

func main() {
	app := fiber.New()

	models.InitDB()

	middlewares.InitMiddleWares(app)
	routes.InitRoutes(app)

	err := app.Listen(os.Getenv("PORT"))
	lib.CheckErr(err, "listening")
}
