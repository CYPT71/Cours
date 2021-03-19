package controllers

import (
	// "gitrest/internal/domain"
	"airfilgth/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func DeparturesBootstrap(app fiber.Router) {
	app.Get("/", departuresGetlist)

	app.Post("/", departuresSetFligth)

}

func departuresGetlist(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql.GetDepartures(c.Query("specific"), c.Query("filter")),
		"message": "Hello from the other side",
	})
	return nil
}

type departus struct {
	Id_flight   int    `json:"id_flight"`
	Date        string `json:"date"`
	Pilote      int    `json:"pilote"`
	Copilote    int    `json:"copilote"`
	Aircrew     string `json:"aircrew"`
	Free_places int    `json:"free_paces"`
	Occupied    int    `json:"occupied"`
	Ticket_id   int    `json:"ticket_id"`
}

func departuresSetFligth(c *fiber.Ctx) error {
	var device deviceStruc
	c.BodyParser(&device)
	// sql.AddDepartures()
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set Fligth",
	})
	return nil
}
