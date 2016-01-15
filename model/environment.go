package model

import (
	"time"

	valid "github.com/asaskevich/govalidator"
)

type Environment struct {
	Id        int64      `json:"id"                         meddler:"id,pk"            valid:"required"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"       valid:"-"`
	Enabled   *bool      `json:"enabled,omitempty"          meddler:"enabled"          valid:"required"`
	FeatureId *int64     `json:"feature_id,omitempty"       meddler:"feature_id"       valid:"required"`
	Name      *string    `json:"name,omitempty"             meddler:"name"             valid:"alphanum,required"`
}

func NewEnvironment(name string, enabled bool, featureId int64) *Environment {
	environment := new(Environment)

	environment.Name = &name
	environment.Enabled = &enabled
	environment.FeatureId = &featureId

	return environment
}

func (e *Environment) Validate() error {
	_, err := valid.ValidateStruct(e)
	return err
}
