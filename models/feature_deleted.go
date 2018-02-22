package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
)

type FeatureDeleted struct {
	Name string `json:"name"`
}

func (e *FeatureDeleted) Validate(s *State) (error, error) {
	if s.getFeatureByName(e.Name) == nil {
		return errors.New("Bad feature"), nil
	}

	return nil, nil
}

func (e *FeatureDeleted) Update(s *State, t time.Time) *State {
	state := *s
	state.Features = []Feature{}

	for _, feature := range s.Features {
		if feature.Name != e.Name {
			state.Features = append(state.Features, feature)
		}
	}

	for _, env := range s.Environments {
		delete(state.Enabled, EnvFeature{env.Name, e.Name})
	}

	return &state
}

func (e *FeatureDeleted) PrePersist(*State) (Event, error) {
	return e, nil
}

func (*FeatureDeleted) Notify(*State, notifier.Notifier) {
}
