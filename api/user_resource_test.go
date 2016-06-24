package api

import (
	"testing"

	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	username := store.Token()

	found := new(User)
	err := client.post("/api/users", &User{
		Username: username,
		Password: "awesome_password",
	}, found)
	require.NoError(t, err)

	require.NotEqual(t, 0, found.ID)
	require.Equal(t, username, found.Username)
	require.Equal(t, "", found.Password)
}

func TestGetUserByUsername(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	found := new(User)
	err := client.get("/api/users/awesome_username", found)
	require.NoError(t, err)
	require.Equal(t, "awesome_username", found.Username)
}
