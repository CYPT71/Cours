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

	app.Get("/:interval", soldsPer)

	app.Post("/", ticketsPos)

	app.Patch("/", ticketUpdate)

	app.Delete("/", ticketDelete)
}

func ticketGetTotal(c *fiber.Ctx) error {

	if ifToken(c) {
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

func soldsPer(c *fiber.Ctx) error {

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success":                          true,
			"Solds per" + c.Params("interval"): sql_request.SoldsPer(c.Params("interval")),
			"message":                          "Solds per" + c.Params("interval"),
		})
	}
	return nil
}

func ticketsGetlist(c *fiber.Ctx) error {

	if ifToken(c) {
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

	if ifToken(c) {
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

func ticketDelete(c *fiber.Ctx) error {

	var device updateTicket
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
			"message": "Delete ticket",
		})
	}
	return nil
}
