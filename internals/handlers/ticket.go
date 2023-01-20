package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateTicket(c *fiber.Ctx) error {
	// var db = database.DB

	return c.JSON(fiber.Map{
		"data":    1,
		"message": "ok",
	})
}
