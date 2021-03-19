package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/ansrivas/fiberprometheus/v2"

	"airfilgth/internal/http/controllers"
	"airfilgth/internal/utils"

	"airfilgth/internal/http/middlewares"
)

func Run() {
	// Setup Configuration
	utils.Setup()
	conf := utils.GetConfig()

	// Create fiber app
	app := fiber.New()

	// Setup routes
	controllers.DeparturesBootstrap(app.Group("/departures"))
	controllers.DevicesBootstrap(app.Group("/devices"))
	controllers.EmployeesBootstrap(app.Group("/employees"))
	controllers.FligthsBootstrap(app.Group("/fligths"))
	controllers.PassagersBootstrap(app.Group("/passengers"))
	controllers.TicketsBootstrap(app.Group("/tickets"))
	controllers.RouteBootstrap(app.Group("/route"))
	controllers.PiloteBootstrap(app.Group("/pilote"))
	controllers.DevicesBootstrap(app.Group("/cabincrews"))

	// Setup CORS/CSRF
	app.Use(middlewares.CORS())
	app.Use(middlewares.CSRF())
	// Setup Logging
	app.Use(logger.New())
	// Setup Limiter - Need to be configured before
	// app.Use(limiter.New())

	// Pprof - Profiling (Remove for production environment)
	// app.Use(pprof.New())

	// Prometheus Endpoint
	prometheus := fiberprometheus.New("gitrest")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Start Listening
	// app.ListenTLS("addr", "certFile", "keyFile") - In production - Istio Passthrough - Cert Files mounted with k8s
	if err := app.Listen(conf.Server.Address + ":" + conf.Server.Port); err != nil {
		log.Fatalf("Err: %v", err)
	}
}
