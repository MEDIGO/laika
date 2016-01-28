package model

import (
	"time"

	valid "github.com/asaskevich/govalidator"
)

type Feature struct {
	Id        int64      `json:"id"                         meddler:"id,pk"            valid:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"       valid:"-"`
	Name      *string    `json:"name,omitempty"             meddler:"name"             valid:"alphanum,required"`
}

func NewFeature(name string) *Feature {
	feature := new(Feature)

	feature.Name = &name

	return feature
}

func (f *Feature) Validate() error {
	_, err := valid.ValidateStruct(f)
	return err
}
