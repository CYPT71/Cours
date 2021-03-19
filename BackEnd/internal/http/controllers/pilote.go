package controllers

import (
	"airfilgth/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func PiloteBootstrap(app fiber.Router) {
	app.Get("/", piloteGetlist)

	app.Post("/", pilotePos)

	app.Patch("/", piloteUpdate)

	app.Delete("/", piloteDelete)

}

func piloteGetlist(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql.GetPilote(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

func pilotePos(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
	return nil
}

type updatePilote struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func piloteUpdate(c *fiber.Ctx) error {
	var device updatePilote
	c.BodyParser(&device)

	sql.UpdateTickets(device.Column, device.Value, device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set ticket",
	})
	return nil
}

func piloteDelete(c *fiber.Ctx) error {

	var device updatePilote
	c.BodyParser(&device)

	sql.DeleteTickets(device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set passenger",
	})
	return nil
}
