package api

import (
	"net/url"
	"time"

	"github.com/MEDIGO/laika/models"
	"github.com/labstack/echo"
)

func GetFeature(c echo.Context) error {
	name, err := url.QueryUnescape(c.Param("name"))
	if err != nil {
		return BadRequest(c, "Bad feature name")
	}

	state := getState(c)
	for _, feature := range state.Features {
		if feature.Name == name {
			return OK(c, *getFeature(&feature, state))
		}
	}

	return NotFound(c)
}

func ListFeatures(c echo.Context) error {
	state := getState(c)
	status := []featureResource{}
	for _, feature := range state.Features {
		status = append(status, *getFeature(&feature, state))
	}
	return OK(c, status)
}

func getFeature(feature *models.Feature, s *models.State) *featureResource {
	f := featureResource{
		Feature:         *feature,
		Status:          map[string]bool{},
		FeatureStatuses: []featureStatus{},
	}
	for _, env := range s.Environments {
		status, ok := s.Enabled[models.EnvFeature{
			Env:     env.Name,
			Feature: feature.Name,
		}]
		toggled := ok && status.Enabled
		f.Status[env.Name] = toggled
		f.FeatureStatuses = append(f.FeatureStatuses, featureStatus{
			Name:      env.Name,
			Status:    toggled,
			ToggledAt: status.ToggledAt,
		})
	}

	return &f
}

func GetFeatureStatus(c echo.Context) error {
	name_param, err := url.QueryUnescape(c.Param("name"))
	if err != nil {
		return BadRequest(c, "Bad feature name")
	}

	env_param, err := url.QueryUnescape(c.Param("env"))
	if err != nil {
		return BadRequest(c, "Bad env name")
	}

	state := getState(c)
	for _, environment := range s.Environments {
		status, ok := s.Enabled[models.EnvFeature{
			Env:     environment.Name,
			Feature: feature.Name,
		}]
		toggled := ok && status.Enabled
		if(env_param == environment.Name && name_param == feature.Name) {
			return OK(c, toggled)
		}
	}
	return OK(c, false)
}

type featureResource struct {
	models.Feature
	Status          map[string]bool `json:"status"`
	FeatureStatuses []featureStatus `json:"feature_status"`
}

type featureStatus struct {
	Name      string     `json:"name"`
	Status    bool       `json:"status"`
	ToggledAt *time.Time `json:"toggled_at,omitempty"`
}
