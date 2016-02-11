package api

import (
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/feature-flag/store"
)

type Feature struct {
	Id        int64            `json:"id"`
	CreatedAt *time.Time       `json:"created_at,omitempty"`
	Name      *string          `json:"name,omitempty"`
	Status    *map[string]bool `json:"status,omitempty"`
}

func (f *Feature) Validate() error {
	if f.Name == nil {
		return CustomError{
			"Name: non zero value required;",
		}
	}
	return nil
}

type FeatureResource struct {
	store store.Store
	stats *statsd.Client
}

func NewFeatureResource(store store.Store, stats *statsd.Client) *FeatureResource {
	return &FeatureResource{store, stats}
}

func (r *FeatureResource) Get(c *echo.Context) error {
	name := c.Param("name")

	feature, err := r.store.GetFeature(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	featureStatus, err := r.store.ListFeaturesStatus(&feature.Id, nil)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	environments, err := r.store.ListEnvironments()
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	featureStatusMap := make(map[string]bool)

	for _, environment := range environments {
		featureStatusMap[*environment.Name] = false
		for _, status := range featureStatus {
			if *status.EnvironmentId == environment.Id {
				featureStatusMap[*environment.Name] = *status.Enabled
				break
			}
		}
	}

	apiFeature := &Feature{
		Id:        feature.Id,
		CreatedAt: feature.CreatedAt,
		Name:      feature.Name,
		Status:    &featureStatusMap,
	}

	return OK(apiFeature)
}

func (r *FeatureResource) List(c *echo.Context) error {
	features, err := r.store.ListFeatures()
	if err != nil {
		return InternalServerError(err)
	}
	return OK(features)
}

func (r *FeatureResource) Create(c *echo.Context) error {
	in := new(Feature)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	if err := in.Validate(); err != nil {
		return BadRequest(err)
	}

	feature := &store.Feature{
		Name: store.String(*in.Name),
	}

	if err := r.store.CreateFeature(feature); err != nil {
		return InternalServerError(err)
	}

	return Created(feature)
}

func (r *FeatureResource) Update(c *echo.Context) error {
	name := c.Param("name")

	feature, err := r.store.GetFeature(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	in := new(store.Feature)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	if in.Name != nil {
		feature.Name = in.Name
	}

	if err := r.store.UpdateFeature(feature); err != nil {
		return InternalServerError(err)
	}

	return OK(feature)
}
