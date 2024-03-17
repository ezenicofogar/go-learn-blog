package api

import (
	"github.com/gofiber/fiber/v2"
)

/*
	Website Fiber App

Constructor of website.
*/
func CreateFiber() *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Status": "Online",
		})
	})

	return app
}
