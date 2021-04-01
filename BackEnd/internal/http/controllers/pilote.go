package controllers

import (
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func PiloteBootstrap(app fiber.Router) {
	app.Get("/details", piloteGetlist)
	app.Get("/", piloteGetlistDetails)

	app.Get("/flightHours", piloteGetAmong)

	app.Get("/pilotArrival", pilotByArrival)

	app.Get("/:name", piloteArrivalByCapitain)

	app.Post("/", pilotePos)

	app.Patch("/", piloteUpdate)

	app.Delete("/:name", piloteDelete)
}

func pilotByArrival(c *fiber.Ctx) error {

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.CityOfThePilot(),
			"message": "Which pilots flies to their cities",
		})
	}
	return nil
}

func piloteGetAmong(c *fiber.Ctx) error {

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.GetPiloteAmong(),
			"message": "Flight hours",
		})
	}
	return nil
}

func piloteArrivalByCapitain(c *fiber.Ctx) error {

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.GetPiloteDestination(c.Params("name")),
			"message": "destinations served by a captain",
		})
	}

	return nil
}

func piloteGetlist(c *fiber.Ctx) error {

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetPilote(c.Query("specific"), c.Query("filter")),
			"message": "List of pilots",
		})
	}
	return nil
}

func piloteGetlistDetails(c *fiber.Ctx) error {
	pilotes_info := sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote`)")

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"succes":  true,
			"value":   pilotes_info,
			"message": "pilotes details info",
		})
	}
	return nil
}

func piloteGetlistRenewLissence(c *fiber.Ctx) error {
	pilotes_info := sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote` WHERE license <= DATE_ADD(CURRENT_DATE(), INTERVAL 3 MONTH))")

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"succes":  true,
			"value":   pilotes_info,
			"message": "pilotes details info",
		})
	}
	return nil
}

func pilotePos(c *fiber.Ctx) error {

	type addPilote struct {
		License  string `json:"license"`
		Among    string `json:"among"`
		Staff_id int    `json:"staff_id"`
	}
	var pilote addPilote
	c.BodyParser(&pilote)

	if pilote.License == "" || pilote.Among == "" || pilote.Staff_id == 0 {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	}
	sql_request.AddPilote(pilote.License, pilote.Among, pilote.Staff_id)

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "You added a pilot",
		})
	}
	return nil
}

func piloteUpdate(c *fiber.Ctx) error {
	type updatePilote struct {
		Column    string `json:"Column"`
		Value     string `json:"Value"`
		Condition string `json:"Condition"`
	}

	var device updatePilote
	c.BodyParser(&device)

	sql_request.UpdateTickets(device.Column, device.Value, device.Condition)

	name := ifToken(c)

	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set ticket",
		})
	}
	return nil
}

func piloteDelete(c *fiber.Ctx) error {

	sql_request.DeleteTickets(c.Params("name"))

	name := ifToken(c)

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
