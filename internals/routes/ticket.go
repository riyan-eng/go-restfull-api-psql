package routes

import (
	"github.com/gofiber/fiber/v2"
	ticketHandler "github.com/riyan-eng/go-restfull-api-psql/internals/handlers"
)

func SetupTicketRoutes(router fiber.Router) {
	ticket := router.Group("/ticket")
	ticket.Post("/", ticketHandler.CreateTicket)
	ticket.Get("/", func(c *fiber.Ctx) error { return nil })
	ticket.Get("/:noteId", func(c *fiber.Ctx) error { return nil })
	ticket.Put("/:noteId", func(c *fiber.Ctx) error { return nil })
	ticket.Delete("/:noteId", func(c *fiber.Ctx) error { return nil })

	ticket.Post("/order/:ticketId", ticketHandler.OrderTicket)
	ticket.Get("/validate/:orderId", ticketHandler.ValidateTicketOrder)
}
