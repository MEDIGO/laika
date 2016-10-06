package store

import (
	"database/sql"

	"github.com/MEDIGO/laika/models"
)

var ErrNoRows = sql.ErrNoRows
var ErrTxDone = sql.ErrTxDone

// Store describes a data store.
type Store interface {
	// CreateEnvironment creates an environment.
	CreateEnvironment(environment *models.Environment) error
	// GetEnvironmentByName gets an environment by name.
	GetEnvironmentByName(name string) (*models.Environment, error)
	// UpdateEnvironment updates an environment.
	UpdateEnvironment(environment *models.Environment) error
	// ListEnvironments list all the environments.
	ListEnvironments() ([]*models.Environment, error)
	// CreateFeature creates a feature.
	CreateFeature(feature *models.Feature) error
	// GetFeatureByName gets a feature by its name.
	GetFeatureByName(name string) (*models.Feature, error)
	// UpdateFeature updates a feature.
	UpdateFeature(feature *models.Feature) error
	// ListFeatures list all features.
	ListFeatures() ([]*models.Feature, error)
	// CreateUser creates an User.
	CreateUser(user *models.User) error
	//GetUserByUsername gets an User by its username.
	GetUserByUsername(username string) (*models.User, error)
	// Migrate migrates the database schema to the latest available version.
	Migrate() error
	// Ping checks the contnectivity with the store.
	Ping() error
	// Reset removes all stored data.
	Reset() error
}
