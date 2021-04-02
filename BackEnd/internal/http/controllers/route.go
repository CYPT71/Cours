package controllers

import (
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func RouteBootstrap(app fiber.Router) {
	app.Get("/", routeGetlist)

	app.Post("/", routePos)

	app.Patch("/", routeUpdate)

	app.Delete("/", routeDelete)

}

func routeGetlist(c *fiber.Ctx) error {

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetRoute(c.Query("specific"), c.Query("filter")),
			"message": "List of travels",
		})
	}
	return nil
}

func DeservedCities(c *fiber.Ctx) error {

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.GetRoute("origin", ""),
			"message": "desserved cities",
		})
	}
	return nil
}

func routePos(c *fiber.Ctx) error {

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Route added",
		})
	}
	return nil
}

type updateRoute struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func routeUpdate(c *fiber.Ctx) error {
	var device updateRoute
	c.BodyParser(&device)

	sql_request.UpdateTickets(device.Column, device.Value, device.Condition)

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Update ticket",
		})
	}
	return nil
}

func routeDelete(c *fiber.Ctx) error {

	var device updateRoute
	c.BodyParser(&device)

	sql_request.DeleteTickets(device.Condition)

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Delete passenger",
		})
	}
	return nil
}
