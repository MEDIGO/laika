package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
	log "github.com/Sirupsen/logrus"
)

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
	state.Enabled[EnvFeature{e.Environment, e.Feature}] = Status{
		Enabled:   e.Status,
		ToggledAt: &t,
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
