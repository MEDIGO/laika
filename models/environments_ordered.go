package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
)

type EnvironmentsOrdered struct {
	Order []string `json:"order"`
}

func (e *EnvironmentsOrdered) Validate(s *State) (error, error) {
	err := errors.New("Order must contain every environment exactly once")

	if len(e.Order) != len(s.Environments) {
		return err, nil
	}

	envs := map[string]bool{}
	for _, env := range e.Order {
		if _, ok := envs[env]; ok {
			return err, nil
		}

		if s.getEnvByName(env) == nil {
			return err, nil
		}
		envs[env] = true
	}

	return nil, nil
}

func (e *EnvironmentsOrdered) Update(s *State, t time.Time) *State {
	state := *s
	state.Environments = s.Environments[:]

	positions := map[string]int{}
	for i, name := range e.Order {
		positions[name] = i
	}

	for src, env := range state.Environments {
		dest, ok := positions[env.Name]
		if !ok {
			continue
		}

		if dest >= len(state.Environments) {
			continue
		}

		state.Environments[dest], state.Environments[src] = state.Environments[src], state.Environments[dest]
	}

	return &state
}

func (e *EnvironmentsOrdered) PrePersist(*State) (Event, error) {
	return e, nil
}

func (*EnvironmentsOrdered) Notify(*State, notifier.Notifier) {
}
