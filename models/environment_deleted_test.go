package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEnvironmentDeleted(t *testing.T) {
	var time time.Time
	s := NewState()

	s = requireValid(t, s, &EnvironmentCreated{Name: "e1"}).Update(s, time)
	s = requireValid(t, s, &EnvironmentCreated{Name: "e2"}).Update(s, time)
	s = requireValid(t, s, &FeatureCreated{Name: "f1"}).Update(s, time)
	s = requireValid(t, s, &FeatureCreated{Name: "f2"}).Update(s, time)
	s = requireValid(t, s, &FeatureCreated{Name: "f3"}).Update(s, time)
	s = requireValid(t, s, &FeatureToggled{Environment: "e1", Feature: "f2", Status: true}).Update(s, time)

	// successful deletion
	s = requireValid(t, s, &EnvironmentDeleted{Name: "e1"}).Update(s, time)

	require.Len(t, s.Environments, 1)
	require.Equal(t, "e2", s.Environments[0].Name)

	_, ok := s.Enabled[EnvFeature{Env: "e1", Feature: "f2"}]
	require.False(t, ok)

	// non-existing feature is invalid
	fd := &EnvironmentDeleted{Name: "e1"}
	valErr, err := fd.Validate(s)
	require.NoError(t, err)
	require.Error(t, valErr)

	// non-existing environment should not cause harm
	before := *s
	s = fd.Update(s, time)
	require.Equal(t, before, *s)
}
