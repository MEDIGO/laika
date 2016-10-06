package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store"
)

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
	input := struct {
		Name string `json:"name"`
	}{}

	if err := c.Bind(&input); err != nil {
		return BadRequest(c, "Payload must be a valid JSON object")
	}

	if input.Name == "" {
		return Invalid(c, "Name is required")
	}

	environment := &models.Environment{
		Name: input.Name,
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

	input := struct {
		Name string `json:"name"`
	}{}

	if err := c.Bind(&input); err != nil {
		return BadRequest(c, "Payload must be a valid JSON object")
	}

	if input.Name != "" {
		environment.Name = input.Name
	}

	if err := r.store.UpdateEnvironment(environment); err != nil {
		return InternalServerError(c, err)
	}

	return OK(c, environment)
}
