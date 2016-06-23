package client

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MEDIGO/laika/api"
)

func TestClientToggle(t *testing.T) {
	server := api.NewTestServer(t)
	defer server.Close()

	client, err := NewClient(Config{
		Addr:        server.URL,
		Username:    "root",
		Password:    "root",
		Environment: "test",
	})
	require.NoError(t, err)

	status := client.IsEnabled("test_feature", false)
	require.True(t, status)
}

func TestClientToggleUnknown(t *testing.T) {
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
