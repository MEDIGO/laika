package store

import (
	"github.com/MEDIGO/laika/models"
)

// Store describes a data store.
type Store interface {
	// Persist persists an event and returns its ID.
	Persist(eventType string, data string) (int64, error)
	// State returns the current state.
	State() (*models.State, error)
	// Migrate migrates the database schema to the latest available version.
	Migrate() error
	// Ping checks the contnectivity with the store.
	Ping() error
	// Reset removes all stored data.
	Reset() error
}
