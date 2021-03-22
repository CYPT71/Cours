package controllers

import (
	// "gitrest/internal/domain"
	// "airfilgth/internal/domain"

	"airfilgth/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func FligthsBootstrap(app fiber.Router) {
	app.Get("/", fligthsGetlist)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)

}

func fligthsGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
	return nil
}

type UpdateFligth struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func fligthUpdate(c *fiber.Ctx) error {
	var device UpdateFligth
	c.BodyParser(&device)

	sql_request.UpdateFligth(device.Column, device.Value, device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set Fligth",
	})
	return nil
}

func fligthDelete(c *fiber.Ctx) error {

	var device UpdateFligth
	c.BodyParser(&device)

	sql_request.DeleteFligth(device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set Fligth",
	})
	return nil
}
