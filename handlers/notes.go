package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello World"})
}
