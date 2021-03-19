package controllers

import (
	// "gitrest/internal/domain"
	// "airfilgth/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func EmployeesBootstrap(app fiber.Router) {
	app.Get("/", employeesGetlist)

}

func employeesGetlist(c *fiber.Ctx) error {

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Hello from the other side",
	})
	return nil
}
