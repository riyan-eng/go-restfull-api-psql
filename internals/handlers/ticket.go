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
		"data": fiber.Map{
			"InsertedId": ticket.ID,
		},
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

var BodyJsonValidate struct {
	OrderID string `json:"order_id"`
}

var DataOrder struct {
	ID         string
	Name       string
	Email      string
	WA         string
	Ticket     string
	TicketName string
}

func ValidateTicketOrder(c *fiber.Ctx) error {
	ticketId := c.Params("ticketId")
	var order = DataOrder
	var db = database.DB
	var bodyJsonValdate = BodyJsonValidate

	c.BodyParser(&bodyJsonValdate)
	// fmt.Println(bodyJsonValdate)

	query := fmt.Sprintf(`
	select o.id, o.name, o.email, o.wa, o.ticket, t.ticket_name
	from public.orders o 
	join public.tickets t on t.id=o.ticket 
	where o.id='%v' and now() >= t.start_time and now() <= t.end_time 
	`, bodyJsonValdate.OrderID)
	result := db.Raw(query).Scan(&order)

	if ticketId != order.Ticket {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    "Your ticket cannot be used at this event.",
			"message": "fail",
		})
	}

	switch result.RowsAffected {
	case 1:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": fiber.Map{
				"order":  order.ID,
				"name":   order.Name,
				"email":  order.Email,
				"ticket": order.TicketName,
			},
			"message": "ok",
		})
	default:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"data":    "time doesn't match or ticket doesn't exist",
			"message": "fail",
		})
	}
}
