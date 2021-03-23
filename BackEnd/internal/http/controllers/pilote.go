package controllers

import (
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func PiloteBootstrap(app fiber.Router) {
	app.Get("/details", piloteGetlist)
	app.Get("/", piloteGetlistDetails)
	app.Post("/", pilotePos)

	app.Patch("/", piloteUpdate)

	app.Delete("/", piloteDelete)

}

func piloteGetlist(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetPilote(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

func piloteGetlistDetails(c *fiber.Ctx) error {
	pilotes_info := sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote`)")
	c.JSON(&fiber.Map{
		"succes":  true,
		"value":   pilotes_info,
		"message": "pilotes details info",
	})

	return nil
}

func piloteGetlistRenewLissence(c *fiber.Ctx) error {
	pilotes_info := sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote` WHERE licence <= DATE_ADD(CURRENT_DATE(), INTERVAL 3 MONTH))")
	c.JSON(&fiber.Map{
		"succes":  true,
		"value":   pilotes_info,
		"message": "pilotes details info",
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

	sql_request.UpdateTickets(device.Column, device.Value, device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set ticket",
	})
	return nil
}

func piloteDelete(c *fiber.Ctx) error {

	var device updatePilote
	c.BodyParser(&device)

	sql_request.DeleteTickets(device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set passenger",
	})
	return nil
}
