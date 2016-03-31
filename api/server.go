package api

import (
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"github.com/MEDIGO/laika/store"
)

func NewServer(store store.Store, stats *statsd.Client) *standard.Server {
	e := echo.New()

	e.Use(LogRequestMiddleware())
	e.Use(InstrumentMiddleware(stats))
	e.Use(ResponseEncoderMiddleware())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(func(user, password string) bool {
		return user == os.Getenv("LAIKA_AUTH_USERNAME") && password == os.Getenv("LAIKA_AUTH_PASSWORD")
	}))

	health := NewHealthResource(store, stats)
	features := NewFeatureResource(store, stats)
	environments := NewEnvironmentResource(store, stats)

	e.Get("/api/health", echo.HandlerFunc(health.Get))

	e.Get("/api/features/:name", echo.HandlerFunc(features.Get))
	e.Get("/api/features", echo.HandlerFunc(features.List))
	e.Post("/api/features", echo.HandlerFunc(features.Create))
	e.Patch("/api/features/:name", echo.HandlerFunc(features.Update))

	e.Get("/api/environments/:name", echo.HandlerFunc(environments.Get))
	e.Get("/api/environments", echo.HandlerFunc(environments.List))
	e.Post("/api/environments", echo.HandlerFunc(environments.Create))
	e.Patch("/api/environments/:name", echo.HandlerFunc(environments.Update))

	e.Static("/", "public")

	server := standard.NewFromConfig(engine.Config{})
	server.SetHandler(e)

	return server
}
