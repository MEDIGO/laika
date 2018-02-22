package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
)

type FeatureCreated struct {
	Name string `json:"name"`
}

func (e *FeatureCreated) Validate(s *State) (error, error) {
	if e.Name == "" {
		return errors.New("Name must not be empty"), nil
	}

	if s.getFeatureByName(e.Name) != nil {
		return errors.New("Name is already in use"), nil
	}

	return nil, nil
}

func (e *FeatureCreated) Update(s *State, t time.Time) *State {
	state := *s
	state.Features = append(state.Features, Feature{
		Name:      e.Name,
		CreatedAt: t,
	})
	return &state
}

func (e *FeatureCreated) PrePersist(*State) (Event, error) {
	return e, nil
}

func (*FeatureCreated) Notify(*State, notifier.Notifier) {
}
