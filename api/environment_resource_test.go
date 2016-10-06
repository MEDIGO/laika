package api

import (
	"testing"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestCreateEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := "prod" + store.Token()

	found := new(models.Environment)
	err := client.post("/api/environments", &models.Feature{
		Name: name,
	}, found)
	require.NoError(t, err)

	require.NotEqual(t, 0, found.ID)
	require.Equal(t, name, found.Name)
}

func TestGetEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := "prod" + store.Token()

	err := client.post("/api/environments", &models.Feature{
		Name: name,
	}, nil)
	require.NoError(t, err)

	found := new(models.Environment)
	err = client.get("/api/environments/"+name, found)
	require.NoError(t, err)

	require.Equal(t, name, found.Name)
}

func TestUpdateEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := "prod" + store.Token()

	err := client.post("/api/environments", &models.Feature{
		Name: name,
	}, nil)
	require.NoError(t, err)

	newName := "not_prod" + store.Token()

	found := new(models.Environment)
	err = client.patch("/api/environments/"+name, &models.Feature{
		Name: newName,
	}, found)
	require.NoError(t, err)

	require.Equal(t, newName, found.Name)
}

func TestListEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	for i := 0; i < 10; i++ {
		err := client.post("/api/environments", &models.Feature{
			Name: "prod_" + store.Token(),
		}, nil)
		require.NoError(t, err)
	}

	found := []models.Environment{}
	err := client.get("/api/environments", &found)
	require.NoError(t, err)

	// ten created plus one that is always present in the test server
	require.NotEqual(t, len(found), 0)
}
