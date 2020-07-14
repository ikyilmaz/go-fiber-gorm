package controllers

import (
	"fiber-rest-api/services"
	"fiber-rest-api/utils"
	"fiber-rest-api/validators"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(s *services.AuthService) *AuthController { return &AuthController{s} }

func (a *AuthController) SignUp(c *fiber.Ctx) {
	signUp := new(validators.SignUp)

	if err := utils.ValidateBody(signUp, c); err != nil {
		c.Next(utils.BadRequest(err.Error()))
		return
	}

	userCreatedPublic, err := a.authService.SignUp(signUp)

	if err != nil {
		c.Next(err)
	}

	token, err := a.signToken(userCreatedPublic.ID)

	if err != nil {
		c.Next(err)
		return
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add((time.Hour * 24) * 90),
		HTTPOnly: true,
	})

	c.JSON(utils.Created(fiber.Map{
		"user":  userCreatedPublic,
		"token": token,
	}))
}

func (a *AuthController) SignIn(c *fiber.Ctx) {
	body := new(validators.SignIn)

	if err := utils.ValidateBody(body, c); err != nil {
		c.Next(utils.BadRequest(err.Error()))
		return
	}

	userPrivate, password, err := a.authService.SignIn(body)

	if err != nil {
		c.Next(err)
		return
	}

	if userPrivate.ID == 0 {
		c.Next(utils.Unauthorized())
		return
	}

	fmt.Printf("%v %v\n", password, body.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password)); err != nil {
		c.Next(utils.Unauthorized())
		return
	}

	token, err := a.signToken(userPrivate.ID)

	if err != nil {
		c.Next(err)
		return
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add((time.Hour * 24) * 90),
		HTTPOnly: true,
	})

	c.JSON(utils.OK(fiber.Map{
		"user":  userPrivate,
		"token": token,
	}))
}

func (a *AuthController) SignOut(c *fiber.Ctx) {
	c.ClearCookie("jwt")
	c.Status(fiber.StatusOK)
}

func (a *AuthController) signToken(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add((time.Hour * 24) * 90).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
