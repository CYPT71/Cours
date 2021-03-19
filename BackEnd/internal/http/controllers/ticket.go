package controllers

import (
	// "gitrest/internal/domain"
	// "airfilgth/internal/domain"

	"airfilgth/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func TicketsBootstrap(app fiber.Router) {
	app.Get("/", ticketsGetlist)

	app.Put("/", ticketsPos)

	app.Patch("/", ticketUpdate)

	app.Delete("/", ticketDelete)
}

func ticketsGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
	return nil
}

func ticketsPos(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
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

	sql.UpdateTickets(device.Column, device.Value, device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set ticket",
	})
	return nil
}

func ticketDelete(c *fiber.Ctx) error {

	var device updateTicket
	c.BodyParser(&device)

	sql.DeleteTickets(device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set passenger",
	})
	return nil
}
