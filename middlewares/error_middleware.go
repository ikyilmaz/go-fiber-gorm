package middlewares

import (
	"fiber-rest-api/utils"
	"github.com/gofiber/fiber"
)

func ErrorHandler(c *fiber.Ctx, err error) {
	switch err.(type) {
	case *utils.APIError:
		c.JSON(err.(*utils.APIError))
	default:
		c.JSON(utils.NewAPIError(500, "something went wrong"))
	}
}
