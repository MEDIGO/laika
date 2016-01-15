package model

import (
  "time"
)

type EnvironmentFeatureUpdate struct {
  Id              int64         `json:"id"                         meddler:"id,pk"`
  Date            *time.Time    `json:"date,omitempty"             meddler:"date"`
  Enabled         *bool         `json:"enabled,omitempty"          meddler:"enabled"`
  EnvironmentID   *int64        `json:"environment_id,omitempty"   meddler:"environment_id"`
  FeatureID       *int64        `json:"feature_id,omitempty"       meddler:"feature_id"`

}

func NewEnvironmentFeatureUpdate(featureID *int64, environmentID *int64, enabled *bool) *EnvironmentFeatureUpdate {
  environmentFeatureUpdate := new (EnvironmentFeatureUpdate)

  environmentFeatureUpdate.FeatureID = featureID
  environmentFeatureUpdate.EnvironmentID = environmentID
  environmentFeatureUpdate.Enabled = enabled;

  t := time.Now()
  environmentFeatureUpdate.Date = &t

  return environmentFeatureUpdate
}
