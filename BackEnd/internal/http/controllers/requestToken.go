package controllers

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetTokenLogin(app fiber.Router) {
	app.Post("/", getToken)

}

func getToken(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	if user != "Cortney" || pass != "Knorr" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	/// FwZa3PjGN1hrD1mk/71Pxj/pAXdh5fC4bxV2eSu00OPKxfFw0WSCPEujaP4pSsVxw9SD+1Y5pvxFnffoeLPSxJyY0HPrrKGOvBRwnfwLBa51HMPS5C/DCj6WQodpyHCiEWfNUZmJZ0lLfBWP+cPQJ5L4I1MiyjYdU3N5X+HNhgkYbcPSzJNAOdW+FeXi8SdvBLIcOqGWuWO3uffKFlBH9I0AjiSxYeAywidZZ2yzMdBMGYKLr2eDaQ7NdblF5aCRh+EFs7U+24414RFhKVNGmYMYvGsTKDJy4gg7wooB8gp3rftG3iseproRQ0tOhA/j8t9mci4vxefmkWWwXy119Q==

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true

	t, err := token.SignedString([]byte("FwZa3PjGN1hrD1mk/71Pxj/pAXdh5fC4bxV2eSu00OPKxfFw0WSCPEujaP4pSsVxw9SD+1Y5pvxFnffoeLPSxJyY0HPrrKGOvBRwnfwLBa51HMPS5C/DCj6WQodpyHCiEWfNUZmJZ0lLfBWP+cPQJ5L4I1MiyjYdU3N5X+HNhgkYbcPSzJNAOdW+FeXi8SdvBLIcOqGWuWO3uffKFlBH9I0AjiSxYeAywidZZ2yzMdBMGYKLr2eDaQ7NdblF5aCRh+EFs7U+24414RFhKVNGmYMYvGsTKDJy4gg7wooB8gp3rftG3iseproRQ0tOhA/j8t9mci4vxefmkWWwXy119Q=="))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}

func if_token(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return name
}
