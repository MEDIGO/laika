package model

import (
  "time"
)

type EnvironmentFeature struct {
  Id              int64         `json:"id"                         meddler:"id,pk"`
  CreationDate    *time.Time    `json:"creation_date,omitempty"    meddler:"creation_date"`
  Enabled         *bool         `json:"enabled,omitempty"          meddler:"enabled"`
  EnvironmentID   *int64        `json:"environment_id,omitempty"   meddler:"environment_id"`
  FeatureID       *int64        `json:"feature_id,omitempty"       meddler:"feature_id"`
}

func NewEnvironmentFeature(featureID *int64, environmentID *int64) *EnvironmentFeature {
  environmentFeature := new (EnvironmentFeature)

  environmentFeature.FeatureID = featureID
  environmentFeature.EnvironmentID = environmentID

  e := false
  environmentFeature.Enabled = &e;

  t := time.Now()
  environmentFeature.CreationDate = &t

  return environmentFeature
}
