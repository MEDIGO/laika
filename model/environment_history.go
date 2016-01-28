package model

import (
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

func NewEnvironmentHistory(createdAt time.Time, enabled bool, featureId int64, name string) *EnvironmentHistory {
	environmentHistory := new(EnvironmentHistory)

	environmentHistory.CreatedAt = &createdAt
	environmentHistory.Enabled = &enabled
	environmentHistory.FeatureId = &featureId
	environmentHistory.Name = &name

	return environmentHistory
}

func (e *EnvironmentHistory) Validate() error {
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
