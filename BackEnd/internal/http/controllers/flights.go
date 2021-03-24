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

	filter := "`id` IN SELECT `id_departures` FROM flight(`id_route` IN (SELECT id FROM `route` WHERE `origin` = " + c.Params("city") + " or `arival` = " + c.Params("city") + ")"

	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetFlight("", filter),
		"message": "Hello from the other side",
	})

	return nil
}

func flightsGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetFlight(c.Query("selector"), c.Query("filter")),
		"message": "Hello from the other side",
	})
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
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set Flight",
	})
	return nil
}

func flightDelete(c *fiber.Ctx) error {

	var device UpdateFlight
	c.BodyParser(&device)

	sql_request.DeleteFlight(device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set Flight",
	})
	return nil
}
