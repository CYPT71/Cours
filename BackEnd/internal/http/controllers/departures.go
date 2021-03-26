package controllers

import (
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func DeparturesBootstrap(app fiber.Router) {
	app.Get("/", departuresGetlist)

	app.Post("/", departuresSetFligth)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)

}

func departuresGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetDepartures(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

func departuresGetListNow(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetDepartures("", "date == CURRENT_DATE()"),
		"message": "Departures for the date",
	})
	return nil
}

type departures struct {
	Id_flight   int    `json:"id_flight"`
	Date        string `json:"date"`
	Pilote      int    `json:"pilote"`
	Copilote    int    `json:"copilote"`
	Aircrew     string `json:"aircrew"`
	Free_places int    `json:"free_paces"`
	Occupied    int    `json:"occupied"`
}

func departuresSetFligth(c *fiber.Ctx) error {
	var device departures
	c.BodyParser(&device)

	sql_request.AddDepartures(device.Id_flight, device.Date, device.Pilote, device.Copilote, device.Aircrew, device.Free_places, device.Occupied)

	name := if_token(c)
	if name == "" {
		c.JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set Departus",
		})
	}
	return nil
}

type updatedepart struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func departuresUpdate(c *fiber.Ctx) error {
	var device updatedepart
	c.BodyParser(&device)

	sql_request.UpdateDepartures(device.Column, device.Value, device.Condition)
	name := if_token(c)
	if name == "" {
		c.JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set Departus",
		})
	}
	return nil
}

func departuresDelete(c *fiber.Ctx) error {

	var device updatedepart
	c.BodyParser(&device)

	sql_request.DeleteDepartures(device.Condition)
	name := if_token(c)
	if name == "" {
		c.JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Set Departus",
		})
	}
	return nil
}
