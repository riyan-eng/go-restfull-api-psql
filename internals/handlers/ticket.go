package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/go-restfull-api-psql/database"
	"github.com/riyan-eng/go-restfull-api-psql/internals/models"
)

func CreateTicket(c *fiber.Ctx) error {
	var db = database.DB
	ticket := new(models.Ticket)

	err := c.BodyParser(&ticket)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	// communicate to database
	ticket.ID = uuid.NewString()
	db.Create(ticket)

	return c.JSON(fiber.Map{
		"data":    ticket,
		"message": "ok",
	})
}

// tiket siapa
// acara apa
// deskripsinya
// kapan
// jumlah tiketnya berapa
// harganya berapa

func OrderTicket(c *fiber.Ctx) error {
	var ticketId = c.Params("ticketId")
	var db = database.DB
	var order = new(models.Order)

	// chect ticket available
	avail, err := CheckAvailable(ticketId)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "fail",
		})
	}

	switch avail.Status {
	case "available":
		err = c.BodyParser(&order)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data":    err.Error(),
				"message": "fail",
			})
		}
		order.ID = uuid.NewString()
		order.Ticket = ticketId
		db.Create(order)
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"InsertedId": order.ID,
			},
			"message": "ok",
		})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "unavailable",
			"message": "fail",
		})
	}
}

type Available struct {
	Amount int
	Status string
}

func CheckAvailable(ticketId string) (Available, error) {
	var db = database.DB
	query := fmt.Sprintf(`
	select coalesce(count(o.ticket), 0) as amount,
	case when coalesce(count(o.ticket), 0) >= (select t.quantity from public.tickets t where t.id='%v')
	then 'unavailable' else 'available' end as status
	from public.orders o where o.ticket = '%v'
	`, ticketId, ticketId)
	var available Available
	err := db.Raw(query).Scan(&available).Error
	if err != nil {
		return available, err
	} else {
		return available, nil
	}
}

func ValidateTicketOrder(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	var order = new(models.Order)
	var db = database.DB
	query := fmt.Sprintf(`
	select * from public.orders o 
	join public.tickets t on t.id=o.ticket 
	where o.id='%v' and now() >= t.start_time and now() <= t.end_time 
	`, orderId)
	result := db.Raw(query).Scan(&order)
	switch result.RowsAffected {
	case 1:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": fiber.Map{
				"name":   order.Name,
				"ticket": order.TicketID,
			},
			"message": "ok",
		})
	default:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data":    "time doesn't match",
			"message": "fail",
		})
	}
}
