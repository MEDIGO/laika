package api

import (
	"errors"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/go-healthz"
	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
	"github.com/labstack/echo"
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
func NewServer(conf ServerConfig) (*echo.Echo, error) {
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

	events := NewEventResource(conf.Store, conf.Stats, conf.Notifier)

	e.GET("/api/health", echo.WrapHandler(healthz.Handler()))

	e.GET("/api/feature/:name/status?env=:env", GetFeatureStatus)

	api := e.Group("/api", StateMiddleware(conf.Store), basicAuthMiddleware)
	api.POST("/events/:type", events.Create)

	api.GET("/features/:name", GetFeature)
	api.GET("/features", ListFeatures)

	api.GET("/environments", ListEnvironments)
	api.GET("/*", func(c echo.Context) error { return NotFound(c) })

	e.Static("/assets", "dashboard/public/assets")
	e.File("/*", "dashboard/public/index.html")

	return e, nil
}
