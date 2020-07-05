package pipes

import (
	"fiber-rest-api/utils"
	"fmt"
	"github.com/gofiber/fiber"
	"strconv"
)

func ParseIntPipe(paramName ...string) fiber.Handler {
	return func(c *fiber.Ctx) {
		param := "id"

		if len(paramName) != 0 {
			param = paramName[0]
		}

		paramValue, err := strconv.Atoi(c.Params(param))

		if err != nil {
			c.Next(utils.BadRequest(fmt.Sprintf("param '%s' must be type int", param)))
			return
		}

		c.Locals(param, paramValue)
		c.Next()
	}
}
