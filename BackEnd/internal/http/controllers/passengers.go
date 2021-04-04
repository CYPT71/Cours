package controllers

import (
	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func PassagersBootstrap(app fiber.Router) {
	app.Get("/", passagersGetlist)

	app.Get("/regular", getRegularProfession)

	app.Get("/perFlight", getPassengerPerFlight)

	app.Get("/mostRegular", regularPassenger)

	app.Get("/byPlane/:start/:end", numbOfPassengersByPeriodByPlane)

	app.Get("/total/:start/:end", numbOfPassengersByPeriod)

	app.Patch("/", departuresUpdate)

	app.Delete("/:name", departuresDelete)
}

func getPassengerPerFlight(c *fiber.Ctx) error {
	// domain.RegularPassenger()

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success":   true,
			"Passenger": sql_request.ListPassengerperFlight(),
			"message":   "List of passengers per flight",
		})
	}
	return nil
}

func getRegularProfession(c *fiber.Ctx) error {
	// domain.RegularPassenger()

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value":   sql_request.MostRegularProfession(),
			"message": "List of regularity of professions",
		})
	}
	return nil
}

func passagersGetlist(c *fiber.Ctx) error {

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetPassenger("", ""),
			"message": "List of passengers",
		})
	}
	return nil
}

type UpdatePassenger struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func passengerUpdate(c *fiber.Ctx) error {
	var device UpdatePassenger
	c.BodyParser(&device)

	sql_request.UpdatePassenger(device.Column, device.Value, device.Condition)

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Update passenger",
		})
	}
	return nil
}

func passengerDelete(c *fiber.Ctx) error {

	var device UpdatePassenger
	c.BodyParser(&device)

	sql_request.DeletePassenger(device.Condition)

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Delete passenger",
		})
	}
	return nil
}

func regularPassenger(c *fiber.Ctx) error {

	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.MostRegularPassenger(),
			"message": "List of regular passengers (2 or more flights per month)",
		})
	}
	return nil
}

func numbOfPassengersByPeriodByPlane(c *fiber.Ctx) error {
	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.NumbOfPassengersByPeriodByPlane(c.Params("start"), c.Params("end")),
			"message": "Number of passengers transported by a plane over a given period",
		})
	}
	return nil
}

func numbOfPassengersByPeriod(c *fiber.Ctx) error {
	if ifNotToken(c) {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success":              true,
			"Number of passengers": sql_request.NumbOfPassengersByPeriod(c.Params("start"), c.Params("end"))[0]["Number of passengers"],
			"message":              "Number of passengers carried over a given period",
		})
	}
	return nil
}
