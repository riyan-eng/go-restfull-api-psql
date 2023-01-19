package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/go-restfull-api-psql/database"
	"github.com/riyan-eng/go-restfull-api-psql/internals/models"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []models.Note

	// find all notes in db
	db.Find(&notes)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    notes,
		"message": "ok",
	})
}
