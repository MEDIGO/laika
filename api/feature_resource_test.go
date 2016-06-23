package api

import (
	"fmt"
	"testing"

	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestCreateFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	found := new(Feature)
	err := client.post("/api/features", &Feature{
		Name: store.String("awesome_feature"),
	}, found)
	require.NoError(t, err)

	require.NotEqual(t, 0, found.Id)
	require.Equal(t, "awesome_feature", *found.Name)
}

func TestGetFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	err := client.post("/api/features", &Feature{
		Name: store.String("awesome_feature"),
	}, nil)
	require.NoError(t, err)

	found := new(Feature)
	err = client.get("/api/features/awesome_feature", found)
	require.NoError(t, err)

	require.Equal(t, "awesome_feature", *found.Name)
}

func TestUpdateFeature(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	err := client.post("/api/features", &Feature{
		Name: store.String("awesome_feature"),
	}, nil)
	require.NoError(t, err)

	found := new(Feature)
	err = client.patch("/api/features/awesome_feature", &Feature{
		Name: store.String("not_awesome_feature"),
	}, found)
	require.NoError(t, err)

	require.Equal(t, "not_awesome_feature", *found.Name)
}

func TestListFeatures(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	for i := 0; i < 10; i++ {
		err := client.post("/api/features", &Feature{
			Name: store.String(fmt.Sprintf("awesome_feature_%d", i)),
		}, nil)
		require.NoError(t, err)
	}

	found := []Feature{}
	err := client.get("/api/features", &found)
	require.NoError(t, err)

	// ten created plus one that is always present in the test server
	require.Len(t, found, 11)
}
