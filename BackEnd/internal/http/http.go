package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	"github.com/ansrivas/fiberprometheus/v2"

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
	app := fiber.New()

	// Setup routes"
	controllers.DeparturesBootstrap(app.Group("/departures"))
	controllers.DevicesBootstrap(app.Group("/devices"))
	controllers.EmployeesBootstrap(app.Group("/employees"))
	controllers.FligthsBootstrap(app.Group("/flights"))
	controllers.PassagersBootstrap(app.Group("/passengers"))
	controllers.TicketsBootstrap(app.Group("/tickets"))
	controllers.RouteBootstrap(app.Group("/route"))
	controllers.PiloteBootstrap(app.Group("/pilote"))
	controllers.CabincrewBootstrap(app.Group("/cabincrew"))

	// Setup CORS/CSRF
	app.Use(middlewares.CORS())
	app.Use(middlewares.CSRF())
	// Setup Logging
	app.Use(logger.New())
	// Setup Limiter - Need to be configured before
	// app.Use(limiter.New())
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("FwZa3PjGN1hrD1mk/71Pxj/pAXdh5fC4bxV2eSu00OPKxfFw0WSCPEujaP4pSsVxw9SD+1Y5pvxFnffoeLPSxJyY0HPrrKGOvBRwnfwLBa51HMPS5C/DCj6WQodpyHCiEWfNUZmJZ0lLfBWP+cPQJ5L4I1MiyjYdU3N5X+HNhgkYbcPSzJNAOdW+FeXi8SdvBLIcOqGWuWO3uffKFlBH9I0AjiSxYeAywidZZ2yzMdBMGYKLr2eDaQ7NdblF5aCRh+EFs7U+24414RFhKVNGmYMYvGsTKDJy4gg7wooB8gp3rftG3iseproRQ0tOhA/j8t9mci4vxefmkWWwXy119Q=="),
	}))
	// Pprof - Profiling (Remove for production environment)
	app.Use(pprof.New())

	// Prometheus Endpoint
	prometheus := fiberprometheus.New("airflight")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Start Listening
	// app.ListenTLS("addr", "certFile", "keyFile") - In production - Istio Passthrough - Cert Files mounted with k8s
	if err := app.Listen(conf.Server.Address + ":" + conf.Server.Port); err != nil {
		log.Fatalf("Err: %v", err)
	}
}
