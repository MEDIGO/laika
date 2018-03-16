package api

import (
	"net/url"
	"time"

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
			return OK(c, *getFeature(&feature, state))
		}
	}

	return NotFound(c)
}

func (r *FeatureResource) List(c echo.Context) error {
	state, err := r.store.State()
	if err != nil {
		return InternalServerError(c, err)
	}

	status := []featureResource{}
	for _, feature := range state.Features {
		status = append(status, *getFeature(&feature, state))
	}
	return OK(c, status)
}

func getFeature(feature *models.Feature, s *models.State) *featureResource {
	f := featureResource{
		Feature:         *feature,
		Status:          map[string]bool{},
		FeatureStatuses: []featureStatus{},
	}
	for _, env := range s.Environments {
		status, ok := s.Enabled[models.EnvFeature{
			Env:     env.Name,
			Feature: feature.Name,
		}]
		toggled := ok && status.Enabled
		f.Status[env.Name] = toggled
		f.FeatureStatuses = append(f.FeatureStatuses, featureStatus{
			Name:      env.Name,
			Status:    toggled,
			ToggledAt: status.ToggledAt,
		})
	}

	return &f
}

type featureResource struct {
	models.Feature
	Status          map[string]bool `json:"status"`
	FeatureStatuses []featureStatus `json:"feature_status"`
}

type featureStatus struct {
	Name      string     `json:"name"`
	Status    bool       `json:"status"`
	ToggledAt *time.Time `json:"toggled_at,omitempty"`
}
