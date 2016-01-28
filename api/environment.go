package api

import (
	"strconv"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/store"
)

type EnvironmentResource struct {
	store store.Store
	stats *statsd.Client
}

func NewEnvironmentResource(store store.Store, stats *statsd.Client) *EnvironmentResource {
	return &EnvironmentResource{store, stats}
}

func (r *EnvironmentResource) Get(c *echo.Context) error {
	featureName := c.Param("feature_name")
	environmentName := c.Param("environment_name")

	feature, err := r.store.GetFeatureByName(featureName)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		}
		return InternalServerError(err)
	}

	environment, err := r.store.GetEnvironment(environmentName, feature.Id)
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

	if err := in.Validate(); err != nil {
		return BadRequest(err)
	}

	environment := model.NewEnvironment(*in.Name, *in.Enabled, *in.FeatureId)

	if err := r.store.CreateEnvironment(environment); err != nil {
		return InternalServerError(err)
	}

	if err := r.store.CreateEnvironmentHistory(environment); err != nil {
		return err
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

	if in.Enabled != nil {
		environment.Enabled = in.Enabled
	}

	if in.FeatureId != nil {
		environment.FeatureId = in.FeatureId
	}

	if in.Name != nil {
		environment.Name = in.Name
	}

	if err := environment.Validate(); err != nil {
		return BadRequest(err)
	}

	if err := r.store.UpdateEnvironment(environment); err != nil {
		return InternalServerError(err)
	}

	if err := r.store.CreateEnvironmentHistory(environment); err != nil {
		return err
	}

	return OK(environment)
}
