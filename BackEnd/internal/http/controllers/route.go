package controllers

import (
	"airfilgth/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func RouteBootstrap(app fiber.Router) {
	app.Get("/", routeGetlist)

	app.Post("/", routePos)

	app.Patch("/", routeUpdate)

	app.Delete("/", routeDelete)

}

func routeGetlist(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql.GetRoute(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

func routePos(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
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

	sql.UpdateTickets(updateRoute.Column, updateRoute.Value, updateRoute.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set ticket",
	})
	return nil
}

func routeDelete(c *fiber.Ctx) error {

	var device updateRoute
	c.BodyParser(&device)

	sql.DeleteTickets(updateRoute.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set passenger",
	})
	return nil
}
