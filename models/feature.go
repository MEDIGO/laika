package models

import "time"

type Feature struct {
	ID        int64           `json:"id"         meddler:"id,pk"`
	CreatedAt time.Time       `json:"created_at" meddler:"created_at"`
	Name      string          `json:"name"       meddler:"name"`
	Status    map[string]bool `json:"status"     meddler:"-"`
}
