package main

import (
	"adver_cart/src/environment"

	"github.com/gofiber/fiber/v2"
)

func main() {

	environment.Init()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":10060")
}
