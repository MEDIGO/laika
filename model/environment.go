package model

import (
  "time"
)

type Environment struct {
  Id              int64         `json:"id"                         meddler:"id,pk"`
  CreationDate    *time.Time    `json:"creation_date,omitempty"    meddler:"creation_date"`
  Enabled         *bool         `json:"enabled,omitempty"          meddler:"enabled"`
  EnvironmentKey  *string       `json:"environment_key,omitempty"  meddler:"environment_key"`
}

func NewEnvironment(environmentKey *string) *Environment {
  environment := new (Environment)

  environment.EnvironmentKey = environmentKey

  e := false
  environment.Enabled = &e;

  t := time.Now()
  environment.CreationDate = &t

  return environment
}
