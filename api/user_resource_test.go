package api

import (
	"testing"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	username := "new_awesome_user" + store.Token()

	found := new(models.User)
	err := client.post("/api/users", &models.User{
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

	username := "new_awesome_user" + store.Token()

	err := client.post("/api/users", &models.User{
		Username: username,
		Password: "awesome_password",
	}, nil)
	require.NoError(t, err)

	found := new(models.User)
	err = client.get("/api/users/"+username, found)

	require.NoError(t, err)
	require.Equal(t, username, found.Username)
}
