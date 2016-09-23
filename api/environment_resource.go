package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/laika/store"
)

type Environment store.Environment

type EnvironmentResource struct {
	store store.Store
	stats *statsd.Client
}

func NewEnvironmentResource(store store.Store, stats *statsd.Client) *EnvironmentResource {
	return &EnvironmentResource{store, stats}
}

func (r *EnvironmentResource) Get(c echo.Context) error {
	name := c.Param("name")

	environment, err := r.store.GetEnvironmentByName(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c)
		}

		return InternalServerError(c, err)
	}

	return OK(c, environment)
}

func (r *EnvironmentResource) List(c echo.Context) error {
	environments, err := r.store.ListEnvironments()
	if err != nil {
		return InternalServerError(c, err)
	}
	return OK(c, environments)
}

func (r *EnvironmentResource) Create(c echo.Context) error {
	in := new(Environment)
	if err := c.Bind(&in); err != nil {
		return BadRequest(c, "Payload must be a valid JSON object")
	}

	if in.Name == nil {
		return Invalid(c, "Name is required")
	}

	environment := &store.Environment{
		Name: store.String(*in.Name),
	}

	if err := r.store.CreateEnvironment(environment); err != nil {
		return InternalServerError(c, err)
	}

	return Created(c, environment)
}

func (r *EnvironmentResource) Update(c echo.Context) error {
	name := c.Param("name")

	environment, err := r.store.GetEnvironmentByName(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c)
		}
		return InternalServerError(c, err)
	}

	in := new(store.Environment)
	if err := c.Bind(&in); err != nil {
		return BadRequest(c, "Payload must be a valid JSON object")
	}

	if in.Name != nil {
		environment.Name = in.Name
	}

	if err := r.store.UpdateEnvironment(environment); err != nil {
		return InternalServerError(c, err)
	}

	return OK(c, environment)
}
