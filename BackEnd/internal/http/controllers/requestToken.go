package controllers

import (
	"airflight/internal/utils"
	"log"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var Token *jwt.Token

func GetTokenLogin(app fiber.Router) {
	app.Post("/", getToken)

}

func getToken(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	log.Print(user, pass)

	if user != "Cortney" || pass != "Knorr" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// Create token
	Token = jwt.New(jwt.SigningMethodHS256)

	claims := Token.Claims.(jwt.MapClaims)
	claims["name"] = "Cortney Knorr"
	claims["admin"] = true

	t, err := Token.SignedString([]byte(utils.Config.Server.SecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}

func ifToken(c *fiber.Ctx) string {

	user := c.Locals("user")

	log.Print(user)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	return "name"
}
