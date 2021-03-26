// package main

// import (
// 	"time"

// 	"github.com/gofiber/fiber/v2"

// 	jwt "github.com/form3tech-oss/jwt-go"
// 	jwtware "github.com/gofiber/jwt/v2"
// )

// func main() {
// 	app := fiber.New()

// 	// Login route
// 	app.Post("/login", login)

// 	// Unauthenticated route
// 	app.Get("/", accessible)

// 	// JWT Middleware
// 	app.Use(jwtware.New(jwtware.Config{
// 		SigningKey: []byte("secret"),
// 	}))

// 	// Restricted Routes
// 	app.Get("/restricted", restricted)

// 	app.Listen(":3000")
// }

// func login(c *fiber.Ctx) error {
// 	user := c.FormValue("user")
// 	pass := c.FormValue("pass")

// 	// Throws Unauthorized error
// 	if user != "john" || pass != "doe" {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	// Create token
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	// Set claims
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["name"] = "John Doe"
// 	claims["admin"] = true
// 	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// 	// Generate encoded token and send it as response.

// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusInternalServerError)
// 	}

// 	return c.JSON(fiber.Map{"token": t})
// }

// func accessible(c *fiber.Ctx) error {
// 	return c.SendString("Accessible")
// }

// func restricted(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)

// 	return c.SendString("Welcome " + name)
// }
