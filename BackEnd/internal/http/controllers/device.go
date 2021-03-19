package controllers

import (
	// "gitrest/internal/domain"
	"airfilgth/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func DevicesBootstrap(app fiber.Router) {
	app.Get("/", devicesGetlist)
	app.Post("/", devicesPost)

}

func devicesGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql.GetDevices(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

type deviceStruc struct {
	Capacity   int    `json:"capacity"`
	Model_type string `json:"model_type"`
}

func devicesPost(c *fiber.Ctx) error {
	var device deviceStruc
	c.BodyParser(&device)
	sql.AddDevices(device.Capacity, device.Model_type)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "You added " + device.Model_type,
	})

	return nil

}
