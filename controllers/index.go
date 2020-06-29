package controllers

import "github.com/gofiber/fiber"

type IBasicController interface {
	GetMany(c *fiber.Ctx)
	Create(c *fiber.Ctx)
	Get(c *fiber.Ctx)
	Update(c *fiber.Ctx)
	Delete(c *fiber.Ctx)
}
