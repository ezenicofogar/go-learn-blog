package website

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

/*
	Website Fiber App

Constructor of website.
*/
func CreateFiber() *fiber.App {
	jetEngine := jet.New("./internal/website/templates", ".jet")

	app := fiber.New(fiber.Config{
		Views: jetEngine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"engineName": "JET",
		})
	})

	return app
}
