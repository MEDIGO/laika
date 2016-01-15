package api

import (
	"strconv"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/store"
)

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

	return OK(feature)
}

func (r *FeatureResource) List(c *echo.Context) error {
	features, err := r.store.ListFeatures(nil, nil, nil)
	if err != nil {
		return InternalServerError(err)
	}
	return OK(features)
}

func (r *FeatureResource) Create(c *echo.Context) error {
	in := new(model.Feature)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	feature := model.NewFeature(*in.Name)

	if err := feature.Validate(); err != nil {
		return BadRequest(err)
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

	in := new(model.Feature)
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
