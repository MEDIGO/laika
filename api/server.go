package api

import (
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/MEDIGO/laika/store"
)

func NewServer(store store.Store, stats *statsd.Client) *echo.Echo {
	server := echo.New()

	server.Use(LogRequestMiddleware())
	server.Use(InstrumentMiddleware(stats))
	server.Use(ResponseEncoderMiddleware())
	server.Use(middleware.Recover())
	server.Use(middleware.BasicAuth(func(user, password string) bool {
		if user == os.Getenv("LAIKA_AUTH_USERNAME") && password == os.Getenv("LAIKA_AUTH_PASSWORD") {
			return true
		}
		return false
	}))

	health := NewHealthResource(store, stats)
	features := NewFeatureResource(store, stats)
	environments := NewEnvironmentResource(store, stats)

	server.Get("/api/health", health.Get)

	server.Get("/api/features/:name", features.Get)
	server.Get("/api/features", features.List)
	server.Post("/api/features", features.Create)
	server.Patch("/api/features/:name", features.Update)

	server.Get("/api/environments/:name", environments.Get)
	server.Get("/api/environments", environments.List)
	server.Post("/api/environments", environments.Create)
	server.Patch("/api/environments/:name", environments.Update)

	server.ServeDir("/", "public")

	return server
}
