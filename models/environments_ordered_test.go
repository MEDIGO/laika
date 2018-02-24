package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEnvironmentsOrdered(t *testing.T) {
	var time time.Time
	s := NewState()

	s = requireValid(t, s, &EnvironmentCreated{Name: "e1"}).Update(s, time)
	s = requireValid(t, s, &EnvironmentCreated{Name: "e2"}).Update(s, time)

	// successful reordering
	s = requireValid(t, s, &EnvironmentsOrdered{Order: []string{"e2", "e1"}}).Update(s, time)

	require.Len(t, s.Environments, 2)
	require.Equal(t, "e2", s.Environments[0].Name)
	require.Equal(t, "e1", s.Environments[1].Name)

	// various errors
	requireInvalid(t, s, &EnvironmentsOrdered{Order: []string{"e1"}})
	requireInvalid(t, s, &EnvironmentsOrdered{Order: []string{"e1", "e2", "e3"}})
	requireInvalid(t, s, &EnvironmentsOrdered{Order: []string{"e1", "e1"}})

	// check graceful error handling
	s = (&EnvironmentsOrdered{Order: []string{"e3"}}).Update(s, time)
	require.Len(t, s.Environments, 2)
	require.Equal(t, "e2", s.Environments[0].Name)
	require.Equal(t, "e1", s.Environments[1].Name)

	s = (&EnvironmentsOrdered{Order: []string{"e1", "e1", "e2"}}).Update(s, time)
	require.Len(t, s.Environments, 2)
	require.Equal(t, "e2", s.Environments[0].Name)
	require.Equal(t, "e1", s.Environments[1].Name)

	s = (&EnvironmentsOrdered{Order: []string{"e1"}}).Update(s, time)
	require.Len(t, s.Environments, 2)
	require.Equal(t, "e1", s.Environments[0].Name)
	require.Equal(t, "e2", s.Environments[1].Name)
}
