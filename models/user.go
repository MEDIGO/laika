package models

import "time"

type User struct {
	ID           int64      `json:"id"         meddler:"id,pk"`
	CreatedAt    time.Time  `json:"created_at" meddler:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" meddler:"updated_at"`
	Username     string     `json:"username"   meddler:"username"`
	Password     string     `json:"password"   meddler:"-"`
	PasswordHash string     `json:"-"          meddler:"password_hash"`
}
