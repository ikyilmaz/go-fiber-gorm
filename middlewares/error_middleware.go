package middlewares

import (
	"fiber-rest-api/utils"
	"github.com/gofiber/fiber"
)

func ErrorHandler(c *fiber.Ctx, err error) {
	switch err.(type) {
	case *utils.APIError:
		handleAPIError(c, err.(*utils.APIError))
	case *utils.DBError:
		handleDBError(c, err.(*utils.DBError))
	default:
		c.JSON(utils.NewAPIError(500, "something went wrong"))
	}
}

func handleAPIError(c *fiber.Ctx, err *utils.APIError) {
	c.JSON(err)
}

// TODO
func handleDBError(c *fiber.Ctx, err *utils.DBError) {
	c.JSON(utils.NewAPIError(500, err.Error()))
}
