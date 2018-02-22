package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func requireValid(t *testing.T, s *State, e Event) Event {
	valErr, err := e.Validate(s)

	require.NoError(t, valErr, "event must be valid")
	require.NoError(t, err, "error during validaton")

	return e
}
