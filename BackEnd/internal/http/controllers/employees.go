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
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"List":    sql_request.GetEmployees("", ""),
			"message": "List of employees",
		})
	}
	return nil
}

func employeesGetByCategories(c *fiber.Ctx) error {
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"value": map[string]interface{}{
				"Pilotes":      sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `pilote`)"),
				"cabin crew":   sql_request.GetEmployees("", "`id` in (SELECT `staff_id` FROM `cabincrew`)"),
				"ground staff": sql_request.GetEmployees("", "`id`  NOT in (SELECT `staff_id` FROM `pilote`) AND `id` NOT IN (SELECT `staff_id` FROM `cabincrew`)"),
			},
			"message": "personnel by categores",
		})
	}

	return nil
}

type EmployeesStruc struct {
	Salary          int    `json:"salary"`
	Social_security int    `json:"social_security"`
	Name            string `json:"name"`
	First_name      string `json:"first_name"`
	Address         string `json:"address"`
}

func employeePost(c *fiber.Ctx) error {
	var employees deviceStruc
	c.BodyParser(&employees)
	sql_request.AddDevices(employees.Capacity, employees.Model_type)
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {

		c.JSON(&fiber.Map{
			"success": true,
			"message": "You added ",
		})
	}
	return nil

}

type updateEmployees struct {
	Column    string `json:"Column"`
	Value     string `json:"Value"`
	Condition string `json:"Condition"`
}

func employeesUpdate(c *fiber.Ctx) error {
	var employees updateEmployees
	c.BodyParser(&employees)

	sql_request.UpdateEmployees(employees.Column, employees.Value, employees.Condition)
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Update Employees",
		})
	}
	return nil
}

func employeesDelete(c *fiber.Ctx) error {

	var employees updateEmployees
	c.BodyParser(&employees)

	sql_request.DeleteEmployees(employees.Condition)
	name := ifToken(c)
	if name == "" {
		c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Unautorized",
		})
	} else {
		c.JSON(&fiber.Map{
			"success": true,
			"message": "Delete Employees",
		})
	}
	return nil
}
