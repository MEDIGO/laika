package client

import "time"

// Feature represents a Feature.
type Feature struct {
	ID        int64           `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	Name      string          `json:"name"`
	Status    map[string]bool `json:"status"`
}

// Error represents an API error.
type Error struct {
	Message string `json:"message"`
}
