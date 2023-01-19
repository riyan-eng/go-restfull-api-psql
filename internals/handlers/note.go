package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/go-restfull-api-psql/database"
	"github.com/riyan-eng/go-restfull-api-psql/internals/helpers"
	"github.com/riyan-eng/go-restfull-api-psql/internals/models"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []models.Note

	// find all notes in db
	db.Raw("SELECT * FROM public.notes").Scan(&notes)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    notes,
		"message": "ok",
	})
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(models.Note)

	err := c.BodyParser(&note)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	errVal := helpers.ValidateNote(*note)
	if errVal != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    errVal,
			"message": "fail",
		})
	}

	note.ID = uuid.New().String()
	db.Create(note)

	// parsing body to var note
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"InsertId": note.ID,
		},
		"message": "ok",
	})
}
