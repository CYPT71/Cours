package controllers

import (
	// "gitrest/internal/domain"
	// "airflight/internal/domain"

	"airflight/internal/sql_request"

	"github.com/gofiber/fiber/v2"
)

func EmployeesBootstrap(app fiber.Router) {
	app.Get("/", employeesGetlist)
	app.Get("/categories", employeesGetByCategories)
	app.Post("/", employeePost)

	app.Patch("/", departuresUpdate)

	app.Delete("/", departuresDelete)

}

func employeesGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"value":   sql_request.GetEmployees("", ""),
		"message": "Hello from the other side",
	})
	return nil
}

func employeesGetByCategories(c *fiber.Ctx) error {
	c.JSON(&fiber.Map{
		"success": true,
		"value": map[string]interface{}{
			"Pilotes":      sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote`)"),
			"cabin crew":   sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `cabincrew`)"),
			"ground staff": sql_request.GetEmployees("", "`id`  NOT in (SELECT `staff_id` FROM `pilote`) AND `id` NOT IN (SELECT `staff_id` FROM `cabincrew`)"),
		},
		"message": "personnel by categores",
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
	sql_request.AddDevices(device.Capacity, device.Model_type)
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
	var device updateEmployees
	c.BodyParser(&device)

	sql_request.UpdateEmployees(device.Column, device.Value, device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Update Employees",
	})
	return nil
}

func employeesDelete(c *fiber.Ctx) error {

	var device updateEmployees
	c.BodyParser(&device)

	sql_request.DeleteEmployees(device.Condition)
	c.JSON(&fiber.Map{
		"success": true,
		"message": "Delete Employees",
	})
	return nil
}
