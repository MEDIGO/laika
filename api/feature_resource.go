package api

import (
	"net/url"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
	"github.com/labstack/echo"
)

type FeatureResource struct {
	store    store.Store
	stats    *statsd.Client
	notifier notifier.Notifier
}

func NewFeatureResource(store store.Store, stats *statsd.Client, notifier notifier.Notifier) *FeatureResource {
	return &FeatureResource{store, stats, notifier}
}

func (r *FeatureResource) Get(c echo.Context) error {
	name, err := url.QueryUnescape(c.Param("name"))
	if err != nil {
		return BadRequest(c, "Bad feature name")
	}

	state, err := r.store.State()
	if err != nil {
		return InternalServerError(c, err)
	}

	for _, feature := range state.Features {
		if feature.Name == name {
			return OK(c, *getFeatureStatus(&feature, state))
		}
	}

	return NotFound(c)
}

func (r *FeatureResource) List(c echo.Context) error {
	state, err := r.store.State()
	if err != nil {
		return InternalServerError(c, err)
	}

	status := []featureStatus{}
	for _, feature := range state.Features {
		status = append(status, *getFeatureStatus(&feature, state))
	}
	return OK(c, status)
}

func getFeatureStatus(feature *models.Feature, s *models.State) *featureStatus {
	fs := featureStatus{
		Feature: *feature,
		Status:  map[string]bool{},
	}
	for _, env := range s.Environments {
		enabled, ok := s.Enabled[models.EnvFeature{
			Env:     env.Name,
			Feature: feature.Name,
		}]
		fs.Status[env.Name] = ok && enabled
	}

	return &fs
}

type featureStatus struct {
	models.Feature
	Status map[string]bool `json:"status"`
}
