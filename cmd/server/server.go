package main

import (
	"github.com/ezenicofogar/go-learn-blog/internal/api"
	"github.com/ezenicofogar/go-learn-blog/internal/website"
	"github.com/gofiber/fiber/v2"
)

func main() {
	core := fiber.New(fiber.Config{})

	core.Mount("/api/", api.CreateFiber())
	core.Mount("/", website.CreateFiber())

	core.Listen(":8080")
}
