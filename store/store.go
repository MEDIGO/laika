package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var ErrNoRows = sql.ErrNoRows
var ErrTxDone = sql.ErrTxDone

type Store interface {
	CreateEnvironment(environment *Environment) error
	CreateFeature(feature *Feature) error
	CreateFeatureStatus(featureStatus *FeatureStatus) error
	CreateUser(user *User) error
	GetEnvironmentByName(name string) (*Environment, error)
	GetFeatureByName(name string) (*Feature, error)
	GetFeatureStatus(featureId int64, environmentId int64) (*FeatureStatus, error)
	GetUserByUsername(username string) (*User, error)
	ListEnvironments() ([]*Environment, error)
	ListFeatures() ([]*Feature, error)
	ListFeatureStatus(featureId *int64, environmentId *int64) ([]*FeatureStatus, error)
	ListFeatureStatusHistory(featureId *int64, environmentId *int64, featureStatusId *int64) ([]*FeatureStatusHistory, error)
	// Migrate migrates the database schema to the latest available version.
	Migrate() error
	Ping() error
	// Reset removes all stored data.
	Reset() error
	UpdateEnvironment(environment *Environment) error
	UpdateFeature(feature *Feature) error
	UpdateFeatureStatus(featureStatus *FeatureStatus) error
}
