package server

import (
	"github.com/gofiber/fiber/v2"
	"golang/domain/product"
	"log"
)

func instanceRoutes(app *fiber.App) {
	product.Routes(app.Group("/product"))
}

func Init() fiber.App {
	app := *fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})

	instanceRoutes(&app)

	log.Fatal(app.Listen(":8080"))

	return app
}
