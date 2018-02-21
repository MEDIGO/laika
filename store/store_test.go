package store

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/MEDIGO/laika/models"
	"github.com/stretchr/testify/require"
)

func testStoreEvents(t *testing.T, store Store) {
	event := models.EnvironmentCreated{Name: "some-env"}
	encoded, err := json.Marshal(&event)
	require.NoError(t, err)

	_, err = store.Persist("environment_created", string(encoded))
	require.NoError(t, err)

	state, err := store.State()
	require.NoError(t, err)
	require.Len(t, state.Environments, 1)
	require.Equal(t, "some-env", state.Environments[0].Name)
}

func getenv(name, val string) string {
	if found := os.Getenv(name); found != "" {
		return found
	}
	return val
}
