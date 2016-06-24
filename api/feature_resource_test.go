package api

import (
	"testing"

	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestCreateFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := store.Token()

	found := new(Feature)
	err := client.post("/api/features", &Feature{
		Name: store.String(name),
	}, found)
	require.NoError(t, err)

	require.NotEqual(t, 0, found.Id)
	require.Equal(t, name, *found.Name)
}

func TestGetFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	found := new(Feature)
	err := client.get("/api/features/awesome_feature", found)
	require.NoError(t, err)

	require.Equal(t, "awesome_feature", *found.Name)
}

func TestUpdateFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	newName := store.Token()

	found := new(Feature)
	err := client.patch("/api/features/awesome_feature", &Feature{
		Name: store.String(newName),
	}, found)
	require.NoError(t, err)

	require.Equal(t, newName, *found.Name)
}

func TestListFeatures(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	found := []Feature{}
	err := client.get("/api/features", &found)
	require.NoError(t, err)

	require.NotEqual(t, len(found), 0)
}
