package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/store"
	"github.com/labstack/echo"
	"strconv"
)

type EnvironmentResource struct {
	store store.Store
	stats *statsd.Client
}

func NewEnvironmentResource(store store.Store, stats *statsd.Client) *EnvironmentResource {
	return &EnvironmentResource{store, stats}
}

func (r *EnvironmentResource) Get(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return BadRequest(err)
	}

	environment, err := r.store.GetEnvironmentById(id)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	return OK(environment)
}

func (r *EnvironmentResource) List(c *echo.Context) error {
	environments, err := r.store.ListEnvironments(nil, nil, nil, nil, nil)
	if err != nil {
		return InternalServerError(err)
	}
	return OK(environments)
}

func (r *EnvironmentResource) Create(c *echo.Context) error {
	in := new(model.Environment)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	environment := model.NewEnvironment(*in.Name, *in.CreatedAt, *in.Enabled, *in.FeatureId)

	if err := environment.Validate(); err != nil {
		return BadRequest(err)
	}

	if err := r.store.CreateEnvironment(environment); err != nil {
		return InternalServerError(err)
	}

	return Created(environment)
}

func (r *EnvironmentResource) Update(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return BadRequest(err)
	}

	environment, err := r.store.GetEnvironmentById(id)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	in := new(model.Environment)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	if in.Name != nil {
		environment.Name = in.Name
	}

	if in.CreatedAt != nil {
		environment.CreatedAt = in.CreatedAt
	}

	if in.Enabled != nil {
		environment.Enabled = in.Enabled
	}

	if in.FeatureId != nil {
		environment.FeatureId = in.FeatureId
	}

	if err := environment.Validate(); err != nil {
		return BadRequest(err)
	}

	if err := r.store.UpdateEnvironment(environment); err != nil {
		return InternalServerError(err)
	}

	return OK(environment)
}
