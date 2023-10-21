package product

import (
	"github.com/gofiber/fiber/v2"
	"golang/common/http"
)

func Routes(r fiber.Router) {

	r.Get("/", func(c *fiber.Ctx) error {

		products, err := UseCases().findAll()

		if err != nil {
			return c.Status(404).JSON(http.ReturnError(err.Error()))
		}

		return c.Status(200).JSON(http.ReturnSuccess(&products))
	})

	r.Get("/:id", func(c *fiber.Ctx) error {

		product, err := UseCases().findById(c.Params("id"))

		if err != nil {
			return c.Status(404).JSON(http.ReturnError(err.Error()))
		}

		return c.Status(200).JSON(http.ReturnSuccess(&product))
	})

	r.Post("/", func(c *fiber.Ctx) error {

		product, err := UseCases().create(c.Body())

		if err != nil {
			return c.Status(400).JSON(http.ReturnError(err.Error()))
		}

		return c.Status(201).JSON(http.ReturnSuccess(&product))
	})

	r.Put("/", func(c *fiber.Ctx) error {

		product, err := UseCases().update(c.Body())

		if err != nil {
			return c.Status(400).JSON(http.ReturnError(err.Error()))
		}

		return c.Status(200).JSON(http.ReturnSuccess(&product))
	})

}
