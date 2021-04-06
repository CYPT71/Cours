package controllers

import (
	"airflight/internal/utils"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var Token *jwt.Token
var TokenString string

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
	TokenString = t
	is_logged := "isLogged : 1"
	file, err := os.Create("security.txt")

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer file.Close()

	file.WriteString(is_logged)

	// log.Print(c.Request().Header)

	return c.JSON(fiber.Map{"authentificate": t})

}

func ifNotToken(c *fiber.Ctx) bool {
	reBearer := regexp.MustCompile("(?i)^Bearer ")
	ts := c.Get("Authorization")

	if !reBearer.MatchString(ts) {

		c.Status(403).SendString("no bearer")
		return true
	}

	reader, _ := ioutil.ReadFile("./security.txt")

	reponse := string(reader)
	if strings.Split(strings.Split(reponse, "\n")[0], ":")[1] == " 1" {
		return false
	}

	bearerToken := strings.Split(ts, " ")
	isToken, err := jwt.Parse(bearerToken[1], func(t *jwt.Token) (interface{}, error) {

		return []byte(utils.Config.Server.SecretKey), nil
	})

	if err != nil {
		log.Print(err)
		return true
	}
	log.Print(TokenString)
	return isToken.Raw != TokenString
}
