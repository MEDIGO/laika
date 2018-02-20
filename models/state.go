package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"name"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

type Environment struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Feature struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type EnvFeature struct {
	EnvID     string
	FeatureID string
}

type State struct {
	Users        []User
	Environments []Environment
	Features     []Feature
	Enabled      map[EnvFeature]bool
}

func NewState() *State {
	return &State{
		Environments: []Environment{},
		Features:     []Feature{},
		Users:        []User{},
		Enabled:      map[EnvFeature]bool{},
	}

}

func (s *State) getFeatureByID(id string) *Feature {
	for _, feature := range s.Features {
		if feature.ID == id {
			return &feature
		}
	}

	return nil
}

func (s *State) getEnvByID(id string) *Environment {
	for _, env := range s.Environments {
		if env.ID == id {
			return &env
		}
	}

	return nil
}
