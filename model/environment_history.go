package model

import (
	"time"

	valid "github.com/asaskevich/govalidator"
)

type EnvironmentHistory struct {
	Id        int64      `json:"id"                         meddler:"id,pk"            valid:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"       valid:"-"`
	Enabled   *bool      `json:"enabled,omitempty"          meddler:"enabled"          valid:"-"`
	FeatureId *int64     `json:"feature_id,omitempty"       meddler:"feature_id"       valid:"required"`
	Name      *string    `json:"name,omitempty"             meddler:"name"             valid:"alphanum,required"`
	Timestamp *time.Time `json:"timestamp,omitempty"        meddler:"timestamp"        valid:"-"`
}

func NewEnvironmentHistory(createdAt time.Time, enabled bool, featureId int64, name string) *EnvironmentHistory {
	environmentHistory := new(EnvironmentHistory)

	environmentHistory.CreatedAt = &createdAt
	environmentHistory.Enabled = &enabled
	environmentHistory.FeatureId = &featureId
	environmentHistory.Name = &name

	return environmentHistory
}

func (eH *EnvironmentHistory) Validate() error {
	_, err := valid.ValidateStruct(eH)
	return err
}
