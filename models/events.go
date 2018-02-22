package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
)

type Event interface {
	// Validate validates the event data against given (immutable) state.
	Validate(*State) (error, error)
	// PrePersist can return a modified event just before persisting.
	PrePersist(*State) (Event, error)
	// Update returns the new state with the event's effect applied.
	Update(*State, time.Time) *State
	// Notify can call a notifier about the event.
	Notify(*State, notifier.Notifier)
}

var types = map[string](func() Event){
	"environment_created": func() Event { return &EnvironmentCreated{} },
	"feature_created":     func() Event { return &FeatureCreated{} },
	"feature_toggled":     func() Event { return &FeatureToggled{} },
	"user_created":        func() Event { return &UserCreated{} },
}

func EventForType(eventType string) (Event, error) {
	f, ok := types[eventType]
	if !ok {
		return nil, errors.New("unknown event type")
	}

	return f(), nil
}
