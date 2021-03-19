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

func departuresSetFligth(c *fiber.Ctx) error {

	sql.AddDepartures(c.Query("values"))
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Set Fligth",
	})
	return nil
}
