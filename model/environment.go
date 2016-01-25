package model

import (
	valid "github.com/asaskevich/govalidator"
	"time"
)

type Environment struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Enabled   *bool      `json:"enabled,omitempty"          meddler:"enabled"`
	FeatureId *int64     `json:"feature_id,omitempty"       meddler:"feature_id"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

func NewEnvironment(name string, time time.Time, enabled bool, featureId int64) *Environment {
	environment := new(Environment)

	environment.Name = &name
	environment.Enabled = &enabled
	environment.CreatedAt = &time
	environment.FeatureId = &featureId

	return environment
}

func (e *Environment) Validate() error {
	_, err := valid.ValidateStruct(e)
	return err
}
