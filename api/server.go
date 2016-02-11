package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/MEDIGO/feature-flag/store"
)

func NewServer(store store.Store, stats *statsd.Client) *echo.Echo {
	server := echo.New()

	server.Use(LogRequestMiddleware())
	server.Use(InstrumentMiddleware(stats))
	server.Use(ResponseEncoderMiddleware())
	server.Use(middleware.Recover())

	features := NewFeatureResource(store, stats)
	environments := NewEnvironmentResource(store, stats)

	server.Get("/api/features/:id", features.Get)
	server.Get("/api/features", features.List)
	server.Post("/api/features", features.Create)
	server.Patch("/api/features/:id", features.Update)

	server.Get("/api/environments/:id", environments.Get)
	server.Get("/api/environments", environments.List)
	server.Post("/api/environments", environments.Create)
	server.Patch("/api/environments/:id", environments.Update)

	server.ServeDir("/", "public")

	return server
}
