package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func main() {
	templateEngine := jet.New("./templates", ".jet")

	app := fiber.New(fiber.Config{
		Views: templateEngine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Listen(":8080")
}
