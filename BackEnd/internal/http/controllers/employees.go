package controllers

import (
	// "gitrest/internal/domain"
	// "airfilgth/internal/domain"

	"airfilgth/internal/sql"

	"github.com/gofiber/fiber/v2"
)

func EmployeesBootstrap(app fiber.Router) {
	app.Get("/", employeesGetlist)

	app.Post("/", employeePost)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)

}

func employeesGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
	return nil
}

type EmployeesStruc struct {
	Capacity   int    `json:"capacity"`
	Model_type string `json:"model_type"`
}

func employeePost(c *fiber.Ctx) error {
	var device deviceStruc
	c.BodyParser(&device)
	sql.AddDevices(device.Capacity, device.Model_type)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "You added " + device.Model_type,
	})

	return nil

}

type updateEmployees struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func employeesUpdate(c *fiber.Ctx) error {
	var device updateDepart
	c.BodyParser(&device)

	sql.UpdateEmployees(updateDepart.Column, updateDepart.Value, updateDepart.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Update Employees",
	})
	return nil
}

func employeesDelete(c *fiber.Ctx) error {

	var device updateDepart
	c.BodyParser(&device)

	sql.DeleteEmployees(updateDepart.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Delete Employees",
	})
	return nil
}
