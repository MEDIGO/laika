package model

import (
	"time"
)

type Feature struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

func NewFeature(name string) *Feature {
	feature := new(Feature)

	feature.Name = &name

	return feature
}

func (f *Feature) Validate() error {
	if f.Name == nil {
		return CustomError{
			"Name: non zero value required;",
		}
	}
	return nil
}
