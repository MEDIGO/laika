package store

import "time"

type User struct {
	ID           int64      `meddler:"id,pk"`
	Username     string     `meddler:"username"`
	PasswordHash string     `meddler:"password_hash"`
	CreatedAt    time.Time  `meddler:"created_at"`
	UpdatedAt    *time.Time `meddler:"updated_at"`
}

type Feature struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

type FeatureStatus struct {
	Id            int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt     *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Enabled       *bool      `json:"enabled,omitempty"          meddler:"enabled"`
	FeatureId     *int64     `json:"feature_id,omitempty"       meddler:"feature_id"`
	EnvironmentId *int64     `json:"environment_id,omitempty"   meddler:"environment_id"`
}

type FeatureStatusHistory struct {
	Id              int64      `json:"id"                            meddler:"id,pk"`
	CreatedAt       *time.Time `json:"created_at,omitempty"          meddler:"created_at"`
	Enabled         *bool      `json:"enabled,omitempty"             meddler:"enabled"`
	FeatureId       *int64     `json:"feature_id,omitempty"          meddler:"feature_id"`
	EnvironmentId   *int64     `json:"environment_id,omitempty"      meddler:"environment_id"`
	FeatureStatusId *int64     `json:"feature_status_id,omitempty"   meddler:"feature_status_id"`
}

type Environment struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}
