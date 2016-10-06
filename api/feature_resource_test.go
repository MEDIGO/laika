package api

import (
	"testing"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestCreateFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := "awesome_feature" + store.Token()

	found := new(models.Feature)
	err := client.post("/api/features", &models.Feature{
		Name: name,
	}, found)
	require.NoError(t, err)

	require.NotEqual(t, 0, found.ID)
	require.Equal(t, name, found.Name)
}

func TestGetFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := "awesome_feature" + store.Token()

	err := client.post("/api/features", &models.Feature{
		Name: name,
	}, nil)
	require.NoError(t, err)

	found := new(models.Feature)
	err = client.get("/api/features/"+name, found)
	require.NoError(t, err)

	require.Equal(t, name, found.Name)
}

func TestUpdateFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := "awesome_feature" + store.Token()

	err := client.post("/api/features", &models.Feature{
		Name: name,
	}, nil)
	require.NoError(t, err)

	newName := "not_awesome_feature" + store.Token()

	found := new(models.Feature)
	err = client.patch("/api/features/"+name, &models.Feature{
		Name: newName,
	}, found)
	require.NoError(t, err)

	require.Equal(t, newName, found.Name)
}

func TestListFeatures(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	for i := 0; i < 10; i++ {
		err := client.post("/api/features", &models.Feature{
			Name: "awesome_feature" + store.Token(),
		}, nil)
		require.NoError(t, err)
	}

	found := []models.Feature{}
	err := client.get("/api/features", &found)
	require.NoError(t, err)

	require.NotEqual(t, len(found), 0)
}
