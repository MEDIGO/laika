package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/feature-flag/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(store store.Store, stats *statsd.Client) *echo.Echo {
	server := echo.New()

	server.Use(LogRequestMiddleware())
	server.Use(InstrumentMiddleware(stats))
	server.Use(ResponseEncoderMiddleware())
	server.Use(middleware.Recover())

	features := NewFeatureResource(store, stats)
	environments := NewEnvironmentResource(store, stats)
	environmentHistory := NewEnvironmentHistoryResource(store, stats)

	server.Get("/features", features.List)
	server.Get("/features/:id", features.Get)
	server.Post("/features", features.Create)
	server.Patch("/features/:id", features.Update)

	server.Get("/environments", environments.List)
	server.Get("/environments/:id", environments.Get)
	server.Post("/environments", environments.Create)
	server.Patch("/environments/:id", environments.Update)

	server.Get("/environment_history", environmentHistory.List)
	server.Get("/environment_history/:id", environmentHistory.Get)
	server.Post("/environment_history", environmentHistory.Create)
	server.Patch("/environment_history/:id", environmentHistory.Update)

	return server
}
