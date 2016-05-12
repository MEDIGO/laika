package api

import (
	"errors"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/labstack/echo"

	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
)

type Feature struct {
	Id        int64           `json:"id"`
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	Name      *string         `json:"name,omitempty"`
	Status    map[string]bool `json:"status,omitempty"`
}

func (f *Feature) Validate() error {
	if f.Name == nil {
		return errors.New("missing name")
	}
	return nil
}

type FeatureResource struct {
	store    store.Store
	stats    *statsd.Client
	notifier notifier.Notifier
}

func NewFeatureResource(store store.Store, stats *statsd.Client, notifier notifier.Notifier) *FeatureResource {
	return &FeatureResource{store, stats, notifier}
}

func (r *FeatureResource) Get(c echo.Context) error {
	name := c.Param("name")

	feature, err := r.store.GetFeatureByName(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	featureStatus, err := r.store.ListFeatureStatus(&feature.Id, nil)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	environments, err := r.store.ListEnvironments()
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
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
		Status:    featureStatusMap,
	}

	return OK(c, apiFeature)
}

func (r *FeatureResource) List(c echo.Context) error {
	features, err := r.store.ListFeatures()
	if err != nil {
		return InternalServerError(c, err)
	}

	environments, err := r.store.ListEnvironments()
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	featureList := make([]*Feature, len(features))
	featureIndex := make(map[int64]*Feature, len(features))
	environmentNames := make(map[int64]string, len(environments))

	featureStatus, err := r.store.ListFeatureStatus(nil, nil)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	for i, feature := range features {
		apiFeature := Feature{
			Id:        feature.Id,
			CreatedAt: feature.CreatedAt,
			Name:      feature.Name,
			Status:    make(map[string]bool),
		}

		for _, environment := range environments {
			apiFeature.Status[*environment.Name] = false
			environmentNames[environment.Id] = *environment.Name
		}

		featureList[i] = &apiFeature
		featureIndex[feature.Id] = &apiFeature
	}

	for _, status := range featureStatus {
		featureIndex[*status.FeatureId].Status[environmentNames[*status.EnvironmentId]] = *status.Enabled
	}

	return OK(c, featureList)
}

func (r *FeatureResource) Create(c echo.Context) error {
	in := new(Feature)
	if err := c.Bind(&in); err != nil {
		return BadRequest(c, err)
	}

	feature, err := r.store.GetFeatureByName(*in.Name)
	if err != nil {
		if err == store.ErrNoRows {
			if err := in.Validate(); err != nil {
				return BadRequest(c, err)
			}

			feature = &store.Feature{
				Name: store.String(*in.Name),
			}

			if err := r.store.CreateFeature(feature); err != nil {
				return InternalServerError(c, err)
			}

			return Created(c, feature)
		}
		return InternalServerError(c, err)
	}

	return Conflict(c, errors.New("Feature already exists"))
}

func (r *FeatureResource) Update(c echo.Context) error {
	name := c.Param("name")

	feature, err := r.store.GetFeatureByName(name)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	in := new(Feature)
	if err := c.Bind(&in); err != nil {
		return BadRequest(c, err)
	}

	if in.Name != nil {
		feature.Name = in.Name
	}

	environments, err := r.store.ListEnvironments()
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	featureStatus, err := r.store.ListFeatureStatus(&feature.Id, nil)
	if err != nil {
		if err == store.ErrNoRows {
			return NotFound(c, err)
		}
		return InternalServerError(c, err)
	}

	for _, environment := range environments {
		var status *store.FeatureStatus
		for _, s := range featureStatus {
			if *s.EnvironmentId == environment.Id {
				status = s
				break
			}
		}

		if status != nil {
			if *status.Enabled != in.Status[*environment.Name] {
				status.Enabled = store.Bool(in.Status[*environment.Name])

				if err := r.store.UpdateFeatureStatus(status); err != nil {
					return InternalServerError(c, err)
				}

				if err := r.notifier.NotifyStatusChange(*feature.Name, in.Status[*environment.Name], *environment.Name); err != nil {
					return InternalServerError(c, err)
				}
			}
		} else {
			status = &store.FeatureStatus{
				CreatedAt:     store.Time(time.Now()),
				Enabled:       store.Bool(in.Status[*environment.Name]),
				FeatureId:     store.Int(feature.Id),
				EnvironmentId: store.Int(environment.Id),
			}

			if err := r.store.CreateFeatureStatus(status); err != nil {
				return InternalServerError(c, err)
			}

			if err := r.notifier.NotifyStatusChange(*feature.Name, in.Status[*environment.Name], *environment.Name); err != nil {
				return InternalServerError(c, err)
			}
		}
	}

	if err := r.store.UpdateFeature(feature); err != nil {
		return InternalServerError(c, err)
	}

	return OK(c, feature)
}
