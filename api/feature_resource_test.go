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

	name := store.Token()

	err := client.post("/api/features", &Feature{
		Name: store.String(name),
	}, nil)
	require.NoError(t, err)

	found := new(Feature)
	err = client.get("/api/features/"+name, found)
	require.NoError(t, err)

	require.Equal(t, name, *found.Name)
}

func TestUpdateFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	name := store.Token()

	err := client.post("/api/features", &Feature{
		Name: store.String(name),
	}, nil)
	require.NoError(t, err)

	newName := store.Token()

	found := new(Feature)
	err = client.patch("/api/features/"+name, &Feature{
		Name: store.String(newName),
	}, found)
	require.NoError(t, err)

	require.Equal(t, newName, *found.Name)
}

func TestListFeatures(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	for i := 0; i < 10; i++ {
		err := client.post("/api/features", &Feature{
			Name: store.String(store.Token()),
		}, nil)
		require.NoError(t, err)
	}

	found := []Feature{}
	err := client.get("/api/features", &found)
	require.NoError(t, err)

	// ten created plus one that is always present in the test server
	require.NotEqual(t, len(found), 0)
}
