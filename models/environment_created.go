package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
)

type EnvironmentCreated struct {
	Name string `json:"name"`
}

func (e *EnvironmentCreated) Validate(s *State) (error, error) {
	if e.Name == "" {
		return errors.New("Name must not be empty"), nil
	}

	if s.getEnvByName(e.Name) != nil {
		return errors.New("Name is already in use"), nil
	}

	return nil, nil
}

func (e *EnvironmentCreated) Update(s *State, t time.Time) *State {
	state := *s
	state.Environments = append(state.Environments, Environment{
		Name:      e.Name,
		CreatedAt: t,
	})
	return &state
}

func (e *EnvironmentCreated) PrePersist(*State) (Event, error) {
	return e, nil
}

func (*EnvironmentCreated) Notify(*State, notifier.Notifier) {
}
