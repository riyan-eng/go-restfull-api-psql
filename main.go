package main

import "github.com/gofiber/fiber/v2"

func main() {
	// start fiber app
	app := fiber.New()

	// route
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("help me!!!")
		return err
	})

	// listen app
	app.Listen(":3000")
}
