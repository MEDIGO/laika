package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetHealth(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	found := new(Health)
	err := client.get("/api/health", found)
	require.NoError(t, err)
}
