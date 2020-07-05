package validators

import (
	"fiber-rest-api/utils"
	"github.com/faceair/jio"
	"github.com/gofiber/fiber"
)

func ValidateCreateUserForm(c *fiber.Ctx) {
	body := []byte(c.Body())

	_, err := jio.ValidateJSON(&body, jio.Object().Keys(jio.K{
		"firstName": jio.String().Optional().Min(2).Max(32),
		"lastName":  jio.String().Optional().Min(2).Max(32),
		"username":  jio.String().Required().Min(2).Max(32),
		"email":     jio.String().Required().Max(128),
		"password":  jio.String().Required().Min(8).Max(32),
	}).Required())

	if err != nil {
		c.Next(utils.NewAPIError(400, err.Error()))
		return
	}

	c.Next()
}

func ValidateUpdateUserForm(c *fiber.Ctx) {
	body := []byte(c.Body())

	_, err := jio.ValidateJSON(&body, jio.Object().Keys(jio.K{
		"firstName": jio.String().Optional().Min(2).Max(32),
		"lastName":  jio.String().Optional().Min(2).Max(32),
		"username":  jio.String().Optional().Min(2).Max(32),
	}).Required())

	if err != nil {
		c.Next(utils.NewAPIError(400, err.Error()))
		return
	}

	c.Next()
}
