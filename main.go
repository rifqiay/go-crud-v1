package main

import (
	"example.com/go-fiber-api/api"
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := api.SetupRoute()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to go fiber API")
	})

	app.Listen(":5000")
}