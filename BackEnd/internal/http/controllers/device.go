package controllers

import (
	// "gitrest/internal/domain"
	"airfilgth/internal/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func DevicesBootstrap(app fiber.Router) {
	app.Get("/", devicesGetlist)
	app.Post("/", devicesPost)

}

func devicesGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql.GetDevice(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

type deviceStruc struct {
	capacity   int    `json:"capacity"`
	model_type string `json:"model_type"`
}

func devicesPost(c *fiber.Ctx) error {
	var device deviceStruc
	err := c.BodyParser(&device)
	log.Print(err)
	// sql.AddDevices(body.capacity, body.model_type)
	c.JSON(device)

	return nil

}
