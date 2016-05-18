package api

import (
	"fmt"
	"testing"

	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

func TestCreateEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	found := new(Environment)
	err := client.post("/api/environments", &Feature{
		Name: store.String("prod"),
	}, found)
	require.NoError(t, err)

	require.NotEqual(t, 0, found.Id)
	require.Equal(t, "prod", *found.Name)
}

func TestGetEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	err := client.post("/api/environments", &Feature{
		Name: store.String("prod"),
	}, nil)
	require.NoError(t, err)

	found := new(Environment)
	err = client.get("/api/environments/prod", found)
	require.NoError(t, err)

	require.Equal(t, "prod", *found.Name)
}

func TestUpdateEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	err := client.post("/api/environments", &Feature{
		Name: store.String("prod"),
	}, nil)
	require.NoError(t, err)

	found := new(Environment)
	err = client.patch("/api/environments/prod", &Feature{
		Name: store.String("not_prod"),
	}, found)
	require.NoError(t, err)

	require.Equal(t, "not_prod", *found.Name)
}

func TestListEnvironment(t *testing.T) {
	client := NewTestClient(t, "root", "root")
	defer client.Close()

	for i := 0; i < 10; i++ {
		err := client.post("/api/environments", &Feature{
			Name: store.String(fmt.Sprintf("prod_%d", i)),
		}, nil)
		require.NoError(t, err)
	}

	found := []Environment{}
	err := client.get("/api/environments", &found)
	require.NoError(t, err)

	// ten created plus one that is always present in the test server
	require.Len(t, found, 11)
}
