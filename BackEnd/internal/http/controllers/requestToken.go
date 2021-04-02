package controllers

import (
	"airflight/internal/utils"
	"regexp"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var Token *jwt.Token

type Claims struct {
	jwt.StandardClaims
	Special string `json:"spc,omitempty"`
}

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
	Token = jwt.New(jwt.SigningMethodHS256)

	claims := Token.Claims.(jwt.MapClaims)
	claims["name"] = "Cortney Knorr"
	claims["admin"] = true

	t, err := Token.SignedString([]byte(utils.Config.Server.SecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"authentificate": t})

}

func ifToken(c *fiber.Ctx) bool {
	reBearer := regexp.MustCompile("(?i)^Bearer ")
	ts := c.Get("Authorization")
	if !reBearer.MatchString(ts) {
		c.Status(403).SendString("no bearer")
		return false
	}

	// _, err := jwt.ParseWithClaims(ts[len("Bearer "):], &Claims{}, func(t *jwt.Token) (interface{}, error) {
	// 	if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v",
	// 			t.Header["alg"])
	// 	}
	// 	return ts, nil
	// })
	// if err != nil {
	// 	return false
	// }
	return true
}
