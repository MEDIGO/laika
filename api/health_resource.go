package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/laika/store"
)

type Health struct {
	Status string `json:"status"`
}

type HealthResource struct {
	store store.Store
	stats *statsd.Client
}

func NewHealthResource(store store.Store, stats *statsd.Client) *HealthResource {
	return &HealthResource{store, stats}
}

func (r *HealthResource) Get(c *echo.Context) error {
	return OK(Health{"OK"})
}
