package api

import (
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
)

func NewServer(store store.Store, stats *statsd.Client, notifier notifier.Notifier) *standard.Server {
	e := echo.New()

	basicAuthMiddleware := middleware.BasicAuth(func(user, password string) bool {
		return user == os.Getenv("LAIKA_AUTH_USERNAME") && password == os.Getenv("LAIKA_AUTH_PASSWORD")
	})

	e.Use(LogRequestMiddleware())
	e.Use(InstrumentMiddleware(stats))
	e.Use(ResponseEncoderMiddleware())
	e.Use(middleware.Recover())

	health := NewHealthResource(store, stats)
	features := NewFeatureResource(store, stats, notifier)
	environments := NewEnvironmentResource(store, stats)

	e.Get("/api/health", echo.HandlerFunc(health.Get))

	e.Get("/api/features/:name", echo.HandlerFunc(features.Get), basicAuthMiddleware)
	e.Get("/api/features", echo.HandlerFunc(features.List), basicAuthMiddleware)
	e.Post("/api/features", echo.HandlerFunc(features.Create), basicAuthMiddleware)
	e.Patch("/api/features/:name", echo.HandlerFunc(features.Update), basicAuthMiddleware)

	e.Get("/api/environments/:name", echo.HandlerFunc(environments.Get), basicAuthMiddleware)
	e.Get("/api/environments", echo.HandlerFunc(environments.List), basicAuthMiddleware)
	e.Post("/api/environments", echo.HandlerFunc(environments.Create), basicAuthMiddleware)
	e.Patch("/api/environments/:name", echo.HandlerFunc(environments.Update), basicAuthMiddleware)

	e.Static("/", "public")

	server := standard.NewFromConfig(engine.Config{})
	server.SetHandler(e)

	return server
}
