package controllers

import (
	// "gitrest/internal/domain"
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func DevicesBootstrap(app fiber.Router) {
	app.Get("/", devicesGetlist)
	app.Get("/time", devicesTimes)
	app.Post("/", devicesPost)

	app.Patch("/", devicesUpdate)

	app.Delete("/:name", devicesDelete)
}

func devicesGetlist(c *fiber.Ctx) error {

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetDevices(c.Query("specific"), c.Query("filter")),
			"message": "List of avaiables flights",
		})
	}
	return nil
}

func devicesTimes(c *fiber.Ctx) error {

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.DeviveHours(),
			"message": "flight hours per flight",
		})
	}
	return nil
}

type deviceStruc struct {
	Capacity   int    `json:"capacity"`
	Model_type string `json:"model_type"`
}

func devicesPost(c *fiber.Ctx) error {
	var device deviceStruc
	c.BodyParser(&device)
	sql_request.AddDevices(device.Capacity, device.Model_type)

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "You added " + device.Model_type,
		})
	}
	return nil

}

type updateDevices struct {
	Column    string `json:"column"`
	Value     string `json:"value"`
	Condition string `json:"condition"`
}

func devicesUpdate(c *fiber.Ctx) error {
	var device updateDevices
	c.BodyParser(&device)

	sql_request.UpdateDevice(device.Column, device.Value, device.Condition)

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Update Device",
		})
	}
	return nil
}

func devicesDelete(c *fiber.Ctx) error {

	sql_request.DeleteDevice(c.Params("name"))

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Delete Device",
		})
	}
	return nil
}
