package controllers

import (
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func PassagersBootstrap(app fiber.Router) {
	app.Get("/", passagersGetlist)

	app.Get("/regular", getRegularProfession)

	app.Get("/perFlight", getPassengerPerFlight)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)
}

func getPassengerPerFlight(c *fiber.Ctx) error {
	// domain.RegularPassenger()
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.ListPassengerperFlight(),
			"message": "Hello from the other side",
		})
	}
	return nil
}

func getRegularProfession(c *fiber.Ctx) error {
	// domain.RegularPassenger()
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.MostRegularProfession(),
			"message": "Hello from the other side",
		})
	}
	return nil
}

func passagersGetlist(c *fiber.Ctx) error {
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.GetPassenger("", ""),
			"message": "Hello from the other side",
		})
	}
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

	sql_request.UpdatePassenger(device.Column, device.Value, device.Condition)
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set passenger",
		})
	}
	return nil
}

func passengerDelete(c *fiber.Ctx) error {

	var device UpdatePassenger
	c.BodyParser(&device)

	sql_request.DeletePassenger(device.Condition)
	name := if_token(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set passenger",
		})
	}
	return nil
}
