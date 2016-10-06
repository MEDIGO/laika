package models

import "time"

type Environment struct {
	ID        int64     `json:"id"         meddler:"id,pk"`
	CreatedAt time.Time `json:"created_at" meddler:"created_at"`
	Name      string    `json:"name"       meddler:"name"`
}
