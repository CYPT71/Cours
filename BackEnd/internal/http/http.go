package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	"airflight/internal/http/controllers"
	"airflight/internal/utils"

	"airflight/internal/http/middlewares"

	jwtware "github.com/gofiber/jwt/v2"
)

func Run() {
	// Setup Configuration
	utils.Setup()
	conf := utils.GetConfig()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork:       true,
		ServerHeader:  "Air Crash",
		StrictRouting: true,
		ProxyHeader:   "Sup Info AirLine",
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   e.Error(),
			})

			return nil
		},
	})

	// Setup routes
	controllers.DeparturesBootstrap(app.Group("/departures"))
	controllers.DevicesBootstrap(app.Group("/devices"))
	controllers.EmployeesBootstrap(app.Group("/employees"))
	controllers.FligthsBootstrap(app.Group("/flights"))
	controllers.PassagersBootstrap(app.Group("/passengers"))
	controllers.TicketsBootstrap(app.Group("/tickets"))
	controllers.RouteBootstrap(app.Group("/route"))
	controllers.PiloteBootstrap(app.Group("/pilote"))
	controllers.CabincrewBootstrap(app.Group("/cabincrew"))

	controllers.GetTokenLogin(app.Group("/login"))

	// Setup CORS/CSRF
	app.Use(middlewares.CORS())
	app.Use(middlewares.CSRF())
	// Setup Logging
	app.Use(logger.New())
	// Setup Limiter - Need to be configured before
	// app.Use(limiter.New())

	// Pprof - Profiling (Remove for production environment)
	app.Use(pprof.New())

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(conf.Server.SecretKey),
	}))

	// Start Listening
	// app.ListenTLS("addr", "certFile", "keyFile") - In production - Istio Passthrough - Cert Files mounted with k8s
	if err := app.Listen(conf.Server.Address + ":" + conf.Server.Port); err != nil {
		log.Fatalf("Err: %v", err)
	}
}
