package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/riyan-eng/go-restfull-api-psql/internals/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	routes.SetupNoteRoutes(api)
	routes.SetupTicketRoutes(api)
}
