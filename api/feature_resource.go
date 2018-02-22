package api

import (
	"net/url"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
	log "github.com/Sirupsen/logrus"
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

	feature, err := r.store.GetFeatureByName(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c)
		}
		return InternalServerError(c, err)
	}

	return OK(c, feature)
}

func (r *FeatureResource) List(c echo.Context) error {
	features, err := r.store.ListFeatures()
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c)
		}
		return InternalServerError(c, err)
	}

	return OK(c, features)
}

func (r *FeatureResource) Create(c echo.Context) error {
	input := struct {
		Name string `json:"name"`
	}{}

	if err := c.Bind(&input); err != nil {
		return BadRequest(c, "Payload must be a valid JSON object")
	}

	if input.Name == "" {
		return Invalid(c, "Name is required")
	}

	found, err := r.store.GetFeatureByName(input.Name)
	if err != nil && err != store.ErrNoRows {
		return InternalServerError(c, err)
	}

	if found != nil {
		return Conflict(c, "Feature already exists")
	}

	feature := &models.Feature{
		Name: input.Name,
	}

	if err := r.store.CreateFeature(feature); err != nil {
		return InternalServerError(c, err)
	}

	return OK(c, feature)
}

func (r *FeatureResource) Update(c echo.Context) error {
	name := c.Param("name")

	feature, err := r.store.GetFeatureByName(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c)
		}
		return InternalServerError(c, err)
	}

	input := struct {
		Name   string          `json:"name"`
		Status map[string]bool `json:"status"`
	}{}

	if err := c.Bind(&input); err != nil {
		return BadRequest(c, "Payload must be a valid JSON object")
	}

	if input.Name != "" {
		feature.Name = input.Name
	}

	// keep the previous status so we can notify changes
	prevStats := make(map[string]bool)
	for name, enabled := range feature.Status {
		prevStats[name] = enabled
	}

	if input.Status != nil {
		feature.Status = input.Status
	}

	if err := r.store.UpdateFeature(feature); err != nil {
		return InternalServerError(c, err)
	}

	for envName, enabled := range feature.Status {
		if prevStats[envName] != enabled {
			go func(featureName string, enabled bool, envName string) {
				if err := r.notifier.NotifyStatusChange(featureName, enabled, envName); err != nil {
					log.Error("failed to notify feature status change: ", err)
				}
			}(feature.Name, enabled, envName)
		}
	}

	return OK(c, feature)
}
