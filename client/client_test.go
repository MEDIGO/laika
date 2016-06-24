package client

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MEDIGO/laika/api"
)

func TestClientIsEnabled(t *testing.T) {
	server := api.NewTestServer(t)
	defer server.Close()

	feature := api.CreateFeatureStatus(t)

	client, err := NewClient(Config{
		Addr:        server.URL,
		Username:    "root",
		Password:    "root",
		Environment: "test",
	})
	require.NoError(t, err)

	status := client.IsEnabled(feature, false)
	require.True(t, status)
}

func TestClientIsEnabledUnknown(t *testing.T) {
	server := api.NewTestServer(t)
	defer server.Close()

	client, err := NewClient(Config{
		Addr:        server.URL,
		Username:    "root",
		Password:    "root",
		Environment: "test",
	})
	require.NoError(t, err)

	status := client.IsEnabled("awesome_unknown_feature", true)
	require.True(t, status)
}
