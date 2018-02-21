package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

type EventMeta struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

var types = map[string](func() Event){
	"environment_created": func() Event { return &EnvironmentCreated{} },
	"feature_created":     func() Event { return &FeatureCreated{} },
	"feature_toggled":     func() Event { return &FeatureToggled{} },
	"user_created":        func() Event { return &UserCreated{} },
}

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

func (*FeatureCreated) Notify(*State, notifier.Notifier) {}

type FeatureToggled struct {
	Feature     string `json:"feature"`
	Environment string `json:"environment"`
	Status      bool   `json:"status"`
}

func (e *FeatureToggled) Validate(s *State) (error, error) {
	if s.getEnvByName(e.Environment) == nil {
		return errors.New("Bad environment"), nil
	}

	if s.getFeatureByName(e.Feature) == nil {
		return errors.New("Bad Feature"), nil
	}

	return nil, nil
}

func (e *FeatureToggled) Update(s *State, t time.Time) *State {
	state := *s

	if e.Status {
		state.Enabled[EnvFeature{e.Environment, e.Feature}] = true
	} else {
		delete(state.Enabled, EnvFeature{e.Environment, e.Feature})
	}

	return &state
}

func (e *FeatureToggled) Notify(s *State, n notifier.Notifier) {
	go func() {
		if err := n.NotifyStatusChange(e.Feature, e.Status, e.Environment); err != nil {
			log.Error("failed to notify feature status change: ", err)
		}
	}()
}

func (e *FeatureToggled) PrePersist(*State) (Event, error) {
	return e, nil
}

type UserCreated struct {
	Username     string  `json:"username"`
	Password     *string `json:"password,omitempty"`
	PasswordHash string  `json:"password_hash"`
}

func (e *UserCreated) Validate(s *State) (error, error) {
	if e.Username == "" {
		return errors.New("Username must not be empty"), nil
	}

	if s.getUserByName(e.Username) != nil {
		return errors.New("Name is already in use"), nil
	}

	if e.Password == nil && e.PasswordHash == "" ||
		e.Password != nil && e.PasswordHash != "" {
		return errors.New("Exactly one of either password or password hash is required"), nil
	}

	return nil, nil
}

func (e *UserCreated) Update(s *State, t time.Time) *State {
	state := *s

	state.Users = append(state.Users, User{
		Username:     e.Username,
		PasswordHash: e.PasswordHash,
	})

	return &state
}

func (e *UserCreated) PrePersist(s *State) (Event, error) {
	if e.Password == nil {
		// already hashed
		return e, nil
	}

	event := *e

	hash, err := bcrypt.GenerateFromPassword([]byte(*e.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// discard plain text
	event.Password = nil
	event.PasswordHash = string(hash)

	return &event, nil
}

func (e *UserCreated) Notify(*State, notifier.Notifier) {
}

func EventForType(eventType string) (Event, error) {
	f, ok := types[eventType]
	if !ok {
		return nil, errors.New("unknown event type")
	}

	return f(), nil
}
