package controllers

import (
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func PassagersBootstrap(app fiber.Router) {
	app.Get("/", passagersGetlist)

	app.Get("/regular", getRegularProfession)

	app.Get("/perFlight", getPassengerPerFlight)

	app.Get("/mostRegular", regularPassenger)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)
}

func getPassengerPerFlight(c *fiber.Ctx) error {
	// domain.RegularPassenger()
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success":   true,
			"Passenger": sql_request.ListPassengerperFlight(),
			"message":   "List of passengers per flight",
		})
	}
	return nil
}

func getRegularProfession(c *fiber.Ctx) error {
	// domain.RegularPassenger()
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.MostRegularProfession(),
			"message": "List of regularity of professions",
		})
	}
	return nil
}

func passagersGetlist(c *fiber.Ctx) error {
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetPassenger("", ""),
			"message": "List of passengers",
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
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Update passenger",
		})
	}
	return nil
}

func passengerDelete(c *fiber.Ctx) error {

	var device UpdatePassenger
	c.BodyParser(&device)

	sql_request.DeletePassenger(device.Condition)
	name := ifToken(c)
	if name == "" {
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

func regularPassenger(c *fiber.Ctx) error {
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.MostRegularPassenger(),
			"message": "List of regular passengers (2 or more flights per month)",
		})
	}
	return nil
}
