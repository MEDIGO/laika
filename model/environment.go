package model

import (
	"time"
)

type Environment struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Enabled   *bool      `json:"enabled,omitempty"          meddler:"enabled"`
	FeatureId *int64     `json:"feature_id,omitempty"       meddler:"feature_id"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

func NewEnvironment(name string, enabled bool, featureId int64) *Environment {
	environment := new(Environment)

	environment.Name = &name
	environment.Enabled = &enabled
	environment.FeatureId = &featureId

	return environment
}

func (e *Environment) Validate() error {
	if e.Enabled == nil {
		return CustomError{
			"Enabled: non zero value required;",
		}
	}
	if e.FeatureId == nil {
		return CustomError{
			"FeatureId: non zero value required;",
		}
	}
	if e.Name == nil {
		return CustomError{
			"Name: non zero value required;",
		}
	}
	return nil
}
