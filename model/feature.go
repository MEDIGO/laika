package model

import (
	valid "github.com/asaskevich/govalidator"
	"time"
)

type Feature struct {
	Id        int64      `json:"id"                        meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

type Features []Feature

func NewFeature(name string, time time.Time) *Feature {
	feature := new(Feature)

	feature.Name = &name
	feature.CreatedAt = &time

	return feature
}

func (f *Feature) Validate() error {
	_, err := valid.ValidateStruct(f)
	return err
}
