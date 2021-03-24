package controllers

import (
	"airflight/internal/sql_request"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PiloteBootstrap(app fiber.Router) {
	app.Get("/details", piloteGetlist)
	app.Get("/", piloteGetlistDetails)

	app.Get("/", piloteArrivalByCapitain)

	app.Get("/flightHours", piloteGetAmong)

	app.Post("/", pilotePos)

	app.Patch("/", piloteUpdate)

	app.Delete("/:name", piloteDelete)

}

func piloteGetAmong(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetPiloteAmong(),
		"message": "Hello from the other side",
	})
	return nil
}

func piloteArrivalByCapitain(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetPiloteDestination(c.Params("name")),
		"message": "Hello from the other side",
	})
	return nil
}

func piloteGetlist(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetPilote(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

func piloteGetlistDetails(c *fiber.Ctx) error {
	pilotes_info := sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote`)")
	c.JSON(&fiber.Map{
		"succes":  true,
		"value":   pilotes_info,
		"message": "pilotes details info",
	})

	return nil
}

func piloteGetlistRenewLissence(c *fiber.Ctx) error {
	pilotes_info := sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote` WHERE licence <= DATE_ADD(CURRENT_DATE(), INTERVAL 3 MONTH))")
	c.JSON(&fiber.Map{
		"succes":  true,
		"value":   pilotes_info,
		"message": "pilotes details info",
	})

	return nil
}

func pilotePos(c *fiber.Ctx) error {

	type addPilote struct {
		Licence  string `json:"licence"`
		Among    string `json:"among"`
		Staff_id int    `json:"staff_id"`
	}
	var pilote addPilote
	c.BodyParser(&pilote)

	layout := "2006-01-02 15:04:05"

	licence, _ := time.Parse(pilote.Licence, layout)

	among, _ := time.Parse(pilote.Among, layout)

	sql_request.AddPilote(licence, among, pilote.Staff_id)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
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
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set ticket",
	})
	return nil
}

func piloteDelete(c *fiber.Ctx) error {

	sql_request.DeleteTickets(c.Params("name"))
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set passenger",
	})
	return nil
}
