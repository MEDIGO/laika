package model

import (
  "time"
)

type Feature struct {
  Id              int64         `json:"id"                         meddler:"id,pk"`
  FeatureKey      *string       `json:"feature_key,omitempty"      meddler:"feature_key"`
  CreationDate    *time.Time    `json:"creation_date,omitempty"    meddler:"creation_date"`
}

func NewFeature(featureKey *string) *Feature {
  feature := new (Feature)

  feature.FeatureKey = featureKey

  t := time.Now()
  feature.CreationDate = &t

  return feature
}
