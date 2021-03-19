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
	//router.SetupRoutes(app, conf.Repository.Name)
	// Group
	//repo := app.Group("/")

	controllers.DeparturesBootstrap(app.Group("/departures"))
	controllers.DevicesBootstrap(app.Group("/devices"))
	controllers.DevicesBootstrap(app.Group("/employees"))
	controllers.DevicesBootstrap(app.Group("/fligths"))
	controllers.DevicesBootstrap(app.Group("/passengers"))
	controllers.DevicesBootstrap(app.Group("/tickets"))
	controllers.DevicesBootstrap(app.Group("/route"))
	controllers.DevicesBootstrap(app.Group("/pilote"))
	controllers.DevicesBootstrap(app.Group("/cabincrews"))

	// Commits
	//repo.Get("/branches", controllers.Branches)

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
