package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
)

type EnvironmentDeleted struct {
	Name string `json:"name"`
}

func (e *EnvironmentDeleted) Validate(s *State) (error, error) {
	if s.getEnvByName(e.Name) == nil {
		return errors.New("Bad environment"), nil
	}

	return nil, nil
}

func (e *EnvironmentDeleted) Update(s *State, t time.Time) *State {
	state := *s
	state.Environments = []Environment{}

	for _, env := range s.Environments {
		if env.Name != e.Name {
			state.Environments = append(state.Environments, env)
		}
	}

	for _, feature := range s.Features {
		delete(state.Enabled, EnvFeature{e.Name, feature.Name})
	}

	return &state
}

func (e *EnvironmentDeleted) PrePersist(*State) (Event, error) {
	return e, nil
}

func (*EnvironmentDeleted) Notify(*State, notifier.Notifier) {
}
