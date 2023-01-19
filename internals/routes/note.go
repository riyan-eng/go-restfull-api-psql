package routes

import (
	"github.com/gofiber/fiber/v2"
	noteHandler "github.com/riyan-eng/go-restfull-api-psql/internals/handlers"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", noteHandler.CreateNote)
	note.Get("/", noteHandler.GetNotes)
	note.Get("/:noteId", func(c *fiber.Ctx) error { return nil })
	note.Put("/:noteId", func(c *fiber.Ctx) error { return nil })
	note.Delete("/:noteId", func(c *fiber.Ctx) error { return nil })
}
