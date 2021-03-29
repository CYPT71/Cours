package controllers

import (
	// "gitrest/internal/domain"
	// "airflight/internal/domain"

	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func TicketsBootstrap(app fiber.Router) {
	app.Get("/", ticketsGetlist)

	app.Get("/total", ticketGetTotal)

	app.Post("/", ticketsPos)

	app.Patch("/", ticketUpdate)

	app.Delete("/", ticketDelete)
}

func ticketGetTotal(c *fiber.Ctx) error {
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"Total":   sql_request.TotalSales(),
			"message": "Total sales",
		})
	}
	return nil
}

func ticketsGetlist(c *fiber.Ctx) error {
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetTickets(c.Query("selector"), c.Query("filter")),
			"message": "List of tickets",
		})
	}
	return nil
}

func ticketsPos(c *fiber.Ctx) error {
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Ticket added",
		})
	}
	return nil
}

type updateTicket struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func ticketUpdate(c *fiber.Ctx) error {
	var device updateTicket
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
			"message": "Update ticket",
		})
	}
	return nil
}

func ticketDelete(c *fiber.Ctx) error {

	var device updateTicket
	c.BodyParser(&device)

	sql_request.DeleteTickets(device.Condition)
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Delete ticket",
		})
	}
	return nil
}
