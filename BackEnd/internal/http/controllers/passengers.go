package controllers

import (
	"airfilgth/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func PassagersBootstrap(app fiber.Router) {
	app.Get("/", passagersGetlist)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)
}

func passagersGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
	return nil
}

type UpdatePassenger struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func passengerUpdate(c *fiber.Ctx) error {
	var device UpdatePassenger
	c.BodyParser(&device)

	sql.UpdatePassenger(updateDepart.Column, updateDepart.Value, updateDepart.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set passenger",
	})
	return nil
}

func passengerDelete(c *fiber.Ctx) error {

	var device UpdatePassenger
	c.BodyParser(&device)

	sql.DeletePassenger(updateDepart.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set passenger",
	})
	return nil
}
