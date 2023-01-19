package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/riyan-eng/go-restfull-api-psql/config"
	"github.com/riyan-eng/go-restfull-api-psql/database"
	"github.com/riyan-eng/go-restfull-api-psql/router"
)

func init() {
	config.LoadEnv()
	database.ConnectDb()
}

func main() {
	// start fiber app
	app := fiber.New()

	// midleware
	app.Use(recover.New())
	// app.Use(logger.New())
	app.Use(cors.New())

	// route
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("help me!!!")
		return err
	})

	router.SetupRoutes(app)

	// listen app
	app.Listen(":3000")
}
