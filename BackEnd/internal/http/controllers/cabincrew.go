package controllers

import (
	// "gitrest/internal/domain"
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func CabincrewBootstrap(app fiber.Router) {

	app.Get("/", cabincrewGetlist)
	app.Post("/", cabincrewPost)

	app.Patch("/", cabincrewUpdate)

	app.Delete("/", cabincrewDelete)
}

func cabincrewGetlist(c *fiber.Ctx) error {

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetCabincrew(c.Query("specific"), c.Query("filter")),
			"message": "List of members of the cabincrew",
		})
	}
	return nil
}

type cabincrewtruc struct {
	Among    string `json:"among"`
	Fonction string `json:"fonction"`
	Staff_id int    `json:"staff_id"`
}

func cabincrewPost(c *fiber.Ctx) error {

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		var cabincrew cabincrewtruc
		c.BodyParser(&cabincrew)
		sql_request.AddCabincrew(cabincrew.Among, cabincrew.Fonction, cabincrew.Staff_id)
		c.JSON(&fiber.Map{
			"success": true,
			"message": "You added " + cabincrew.Fonction,
		})
	}
	return nil

}

type updatecabincrew struct {
	Column    string `json:"column"`
	Value     string `json:"value"`
	Condition string `json:"condition"`
}

func cabincrewUpdate(c *fiber.Ctx) error {
	var cabincrew updatecabincrew
	c.BodyParser(&cabincrew)

	sql_request.UpdateCabincrew(cabincrew.Column, cabincrew.Value, cabincrew.Condition)

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Update cabincrew",
		})
	}
	return nil
}

func cabincrewDelete(c *fiber.Ctx) error {

	var cabincrew updatecabincrew
	c.BodyParser(&cabincrew)

	sql_request.DeleteCabincrew(cabincrew.Condition)

	if ifToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {

		c.JSON(&fiber.Map{
			"success": true,
			"message": "Delete cabincrew",
		})

	}
	return nil
}
