package controllers

import (
	// "gitrest/internal/domain"
	// "airflight/internal/domain"

	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func FligthsBootstrap(app fiber.Router) {
	app.Get("/", flightsGetlist)

	app.Get("/:city", cityGetFLight)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)

}

func cityGetFLight(c *fiber.Ctx) error {
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.GetFlightByCity(c.Params("city")),
			"message": "Hello from the other side",
		})
	}

	return nil
}

func flightsGetlist(c *fiber.Ctx) error {
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.GetFlight(c.Query("selector"), c.Query("filter")),
			"message": "Hello from the other side",
		})
	}
	return nil
}

type UpdateFlight struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func flightUpdate(c *fiber.Ctx) error {
	var device UpdateFlight
	c.BodyParser(&device)

	sql_request.UpdateFlight(device.Column, device.Value, device.Condition)
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set Flight",
		})
	}
	return nil
}

func flightDelete(c *fiber.Ctx) error {

	var device UpdateFlight
	c.BodyParser(&device)

	sql_request.DeleteFlight(device.Condition)
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set Flight",
		})
	}
	return nil
}
