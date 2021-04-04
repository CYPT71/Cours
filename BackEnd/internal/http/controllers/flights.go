package controllers

import (
	// "gitrest/internal/domain"
	// "airflight/internal/domain"

	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func FligthsBootstrap(app fiber.Router) {
	app.Get("/", flightsGetlist)

	app.Get("/occupancyRate", occupancyRate)

	app.Get("/:city", cityGetFLight)

	app.Patch("/", departuresUpdate)

	app.Delete("/:name", departuresDelete)

}

func cityGetFLight(c *fiber.Ctx) error {

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.GetFlightByCity(c.Params("city")),
			"message": "List of cities served by a flight",
		})
	}

	return nil
}

func occupancyRate(c *fiber.Ctx) error {

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success":        true,
			"Occupancy Rate": sql_request.OccupancyRate(),
			"message":        "Most profitable destinations",
		})
	}
	return nil
}

func flightsGetlist(c *fiber.Ctx) error {

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetFlight(c.Query("selector"), c.Query("filter")),
			"message": "List of flights",
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

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Update Flight",
		})
	}
	return nil
}

func flightDelete(c *fiber.Ctx) error {

	sql_request.DeleteFlight(c.Params("name"))

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Delete Flight",
		})
	}
	return nil
}
