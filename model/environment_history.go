package model

import (
	valid "github.com/asaskevich/govalidator"
	"time"
)

type EnvironmentHistory struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Enabled   *bool      `json:"enabled,omitempty"          meddler:"enabled"`
	FeatureId *int64     `json:"feature_id,omitempty"       meddler:"feature_id"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
	Timestamp *time.Time `json:"timestamp,omitempty"        meddler:"timestamp"`
}

func NewEnvironmentHistory(createdAt time.Time, enabled bool, featureId int64, name string, timestamp time.Time) *EnvironmentHistory {
	environmentHistory := new(EnvironmentHistory)

	environmentHistory.CreatedAt = &createdAt
	environmentHistory.Enabled = &enabled
	environmentHistory.FeatureId = &featureId
	environmentHistory.Name = &name
	environmentHistory.Timestamp = &timestamp

	return environmentHistory
}

func (eH *EnvironmentHistory) Validate() error {
	_, err := valid.ValidateStruct(eH)
	return err
}
