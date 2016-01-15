package api

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/store"
	"github.com/labstack/echo"
	"strconv"
)

type EnvironmentHistoryResource struct {
	store store.Store
	stats *statsd.Client
}

func NewEnvironmentHistoryResource(store store.Store, stats *statsd.Client) *EnvironmentHistoryResource {
	return &EnvironmentHistoryResource{store, stats}
}

func (r *EnvironmentHistoryResource) Get(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return BadRequest(err)
	}

	environment, err := r.store.GetEnvironmentHistoryById(id)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	return OK(environment)
}

func (r *EnvironmentHistoryResource) List(c *echo.Context) error {
	environmentHistory, err := r.store.ListEnvironmentHistory(nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return InternalServerError(err)
	}
	return OK(environmentHistory)
}

func (r *EnvironmentHistoryResource) Create(c *echo.Context) error {
	in := new(model.EnvironmentHistory)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	environmentHistory := model.NewEnvironmentHistory(*in.CreatedAt, *in.Enabled, *in.FeatureId, *in.Name, *in.Timestamp)

	if err := environmentHistory.Validate(); err != nil {
		return BadRequest(err)
	}

	if err := r.store.CreateEnvironmentHistory(environmentHistory); err != nil {
		return InternalServerError(err)
	}

	return Created(environmentHistory)
}

func (r *EnvironmentHistoryResource) Update(c *echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return BadRequest(err)
	}

	environmentHistory, err := r.store.GetEnvironmentHistoryById(id)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(err)
		} else {
			return InternalServerError(err)
		}
	}

	in := new(model.EnvironmentHistory)
	if err := c.Bind(&in); err != nil {
		return BadRequest(err)
	}

	if in.Name != nil {
		environmentHistory.Name = in.Name
	}

	if in.CreatedAt != nil {
		environmentHistory.CreatedAt = in.CreatedAt
	}

	if in.Enabled != nil {
		environmentHistory.Enabled = in.Enabled
	}

	if in.FeatureId != nil {
		environmentHistory.FeatureId = in.FeatureId
	}

	if in.Timestamp != nil {
		environmentHistory.Timestamp = in.Timestamp
	}

	if err := environmentHistory.Validate(); err != nil {
		return BadRequest(err)
	}

	if err := r.store.UpdateEnvironmentHistory(environmentHistory); err != nil {
		return InternalServerError(err)
	}

	return OK(environmentHistory)
}
