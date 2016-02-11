package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/feature-flag/store"
)

type Environment store.Environment

func (e *Environment) Validate() error {
	if e.Name == nil {
		return CustomError{
			"Name: non zero value required;",
		}
	}
	return nil
}

type EnvironmentResource struct {
	store store.Store
	stats *statsd.Client
}

func NewEnvironmentResource(store store.Store, stats *statsd.Client) *EnvironmentResource {
	return &EnvironmentResource{store, stats}
}

func (r *EnvironmentResource) Get(c *echo.Context) error {
	name := c.Param("name")

	environment, err := r.store.GetEnvironment(name)
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
	environments, err := r.store.ListEnvironments()
	if err != nil {
		return InternalServerError(err)
	}
	return OK(environments)
}

func (r *EnvironmentResource) Create(c *echo.Context) error {
	in := new(Environment)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	if err := in.Validate(); err != nil {
		return BadRequest(err)
	}

	environment := &store.Environment{
		Name: store.String(*in.Name),
	}

	if err := r.store.CreateEnvironment(environment); err != nil {
		return InternalServerError(err)
	}

	return Created(environment)
}

func (r *EnvironmentResource) Update(c *echo.Context) error {
	name := c.Param("name")

	environment, err := r.store.GetEnvironment(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	in := new(store.Environment)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	if in.Name != nil {
		environment.Name = in.Name
	}

	if err := r.store.UpdateEnvironment(environment); err != nil {
		return InternalServerError(err)
	}

	return OK(environment)
}
