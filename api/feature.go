package api

import (
	"strconv"
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

func NewFeature(id int64, createdAt time.Time, name string, status map[string]bool) *Feature {
	feature := new(Feature)

	feature.Id = id
	feature.CreatedAt = &createdAt
	feature.Name = &name
	feature.Status = &status

	return feature
}

type FeatureResource struct {
	store store.Store
	stats *statsd.Client
}

func NewFeatureResource(store store.Store, stats *statsd.Client) *FeatureResource {
	return &FeatureResource{store, stats}
}

func (r *FeatureResource) Get(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return BadRequest(err)
	}

	feature, err := r.store.GetFeatureById(id)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	featureStatus, err := r.store.ListFeaturesStatus(&id, nil)
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

	var m map[string]bool
	m = make(map[string]bool)

	for i := 0; i < len(environments); i++ {
		for j := 0; j < len(featureStatus); j++ {
			if *featureStatus[j].EnvironmentId == environments[i].Id {
				m[*environments[i].Name] = *featureStatus[j].Enabled
				break
			}
		}
	}

	apiFeature := &Feature{
		Id:        feature.Id,
		CreatedAt: feature.CreatedAt,
		Name:      feature.Name,
		Status:    &m,
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
	in := new(store.Feature)
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return BadRequest(err)
	}

	feature, err := r.store.GetFeatureById(id)
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

	if err := feature.Validate(); err != nil {
		return BadRequest(err)
	}

	if err := r.store.UpdateFeature(feature); err != nil {
		return InternalServerError(err)
	}

	return OK(feature)
}
