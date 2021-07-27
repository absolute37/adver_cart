package routers

import (
	"adver_cart/src/core"
	"adver_cart/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Static("/", "public")

	app.Use(cors.New(cors.ConfigDefault))

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("cc", core.NewFiberContext(c))
		return c.Next()
	})

	//Conntect database
	db := database.NewConnection()

	_ = db

	return app
}
