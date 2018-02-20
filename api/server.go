package api

import (
	"errors"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/go-healthz"
	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// ServerConfig is used to parametrize a Server.
type ServerConfig struct {
	RootUsername string
	RootPassword string
	Store        store.Store
	Stats        *statsd.Client
	Notifier     notifier.Notifier
}

// NewServer creates a new server.
func NewServer(conf ServerConfig) (*standard.Server, error) {
	if conf.RootPassword == "" {
		return nil, errors.New("missing root username")
	}

	if conf.RootPassword == "" {
		return nil, errors.New("missing root password")
	}

	if conf.Store == nil {
		return nil, errors.New("missing store")
	}

	if conf.Notifier == nil {
		conf.Notifier = notifier.NewNOOPNotifier()
	}

	e := echo.New()

	basicAuthMiddleware := AuthMiddleware(conf.RootUsername, conf.RootPassword, conf.Store)

	e.Use(TraceMiddleware())
	e.Use(LogMiddleware())
	e.Use(InstrumentMiddleware(conf.Stats))
	e.Use(middleware.Recover())

	features := NewFeatureResource(conf.Store, conf.Stats, conf.Notifier)
	environments := NewEnvironmentResource(conf.Store, conf.Stats)
	events := NewEventResource(conf.Store, conf.Stats, conf.Notifier)

	e.Get("/api/health", standard.WrapHandler(healthz.Handler()))

	e.Post("/api/events/:type", events.Create, basicAuthMiddleware)

	e.Get("/api/features/:name", features.Get, basicAuthMiddleware)
	e.Get("/api/features", features.List, basicAuthMiddleware)

	e.Get("/api/environments", environments.List, basicAuthMiddleware)

	e.Static("/static", "public")
	e.File("/*", "public/index.html")

	server := standard.WithConfig(engine.Config{})
	server.SetHandler(e)

	return server, nil
}
