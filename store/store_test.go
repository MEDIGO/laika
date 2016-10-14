package store

import (
	"os"
	"testing"

	"github.com/MEDIGO/laika/models"
	"github.com/stretchr/testify/require"
)

func testStoreUsers(t *testing.T, store Store) {
	// it should fail to get an user that doesn't exist
	found, err := store.GetUserByUsername("testy")
	require.Equal(t, ErrNoRows, err)

	// it should create an user
	user := &models.User{
		Username: "testy",
		Password: "qwerty",
	}

	err = store.CreateUser(user)
	require.NoError(t, err)
	require.NotEqual(t, user.ID, 0)
	require.Empty(t, user.Password)
	require.NotEmpty(t, user.PasswordHash)

	// it should get an user by its username
	found, err = store.GetUserByUsername("testy")
	require.NoError(t, err)
	require.Equal(t, user.ID, found.ID)
}

func testStoreEnvironments(t *testing.T, store Store) {
	// it should failt to get an environment that doesn't exist
	found, err := store.GetEnvironmentByName("staging")
	require.Equal(t, ErrNoRows, err)

	// it should get an empty list when there is no environments
	list, err := store.ListEnvironments()
	require.NoError(t, err)
	require.Empty(t, list)

	// it should create an environment
	env := &models.Environment{
		Name: "staging",
	}

	err = store.CreateEnvironment(env)
	require.NoError(t, err)
	require.NotEqual(t, env.ID, 0)
	require.Equal(t, "staging", env.Name)

	// it should get an environment
	found, err = store.GetEnvironmentByName("staging")
	require.NoError(t, err)
	require.Equal(t, env.ID, found.ID)

	// it should update an environment
	env.Name = "not_staging"

	err = store.UpdateEnvironment(env)
	require.NoError(t, err)

	found, err = store.GetEnvironmentByName("not_staging")
	require.NoError(t, err)

	// it should list all the environments
	list, err = store.ListEnvironments()
	require.NoError(t, err)
	require.Len(t, list, 1)
}

func testStoreFeatures(t *testing.T, store Store) {
	// it should failt to get a feature that doesn't exist
	found, err := store.GetFeatureByName("awesome_feature")
	require.Equal(t, ErrNoRows, err)

	// it should get an empty list when there is no features
	list, err := store.ListFeatures()
	require.NoError(t, err)
	require.Empty(t, list)

	// it should create a feature
	feature := &models.Feature{
		Name: "awesome_feature",
	}

	err = store.CreateFeature(feature)
	require.NoError(t, err)
	require.NotEqual(t, feature.ID, 0)
	require.Equal(t, "awesome_feature", feature.Name)
	require.Empty(t, feature.Status)

	// it should update a feature
	feature.Name = "not_awesome_feature"

	err = store.UpdateFeature(feature)
	require.NoError(t, err)

	// it should list all features
	list, err = store.ListFeatures()
	require.NoError(t, err)
	require.Len(t, list, 1)

	// it should get a feature with the status of all environments
	err = store.CreateEnvironment(&models.Environment{
		Name: "dev",
	})
	require.NoError(t, err)

	err = store.CreateEnvironment(&models.Environment{
		Name: "staging",
	})
	require.NoError(t, err)

	err = store.CreateEnvironment(&models.Environment{
		Name: "prod",
	})
	require.NoError(t, err)

	found, err = store.GetFeatureByName("not_awesome_feature")
	require.NoError(t, err)
	require.Len(t, found.Status, 3)
	require.False(t, found.Status["dev"])
	require.False(t, found.Status["staging"])
	require.False(t, found.Status["prod"])

	// it should update the status of a feature in an environment
	found.Status["dev"] = true

	err = store.UpdateFeature(found)
	require.NoError(t, err)

	found, err = store.GetFeatureByName("not_awesome_feature")
	require.NoError(t, err)
	require.True(t, found.Status["dev"])
	require.False(t, found.Status["staging"])
	require.False(t, found.Status["prod"])

	found.Status["staging"] = true

	// update a second time to execute a different code path when the environment status
	// already exist
	err = store.UpdateFeature(found)
	require.NoError(t, err)

	found, err = store.GetFeatureByName("not_awesome_feature")
	require.NoError(t, err)
	require.True(t, found.Status["dev"])
	require.True(t, found.Status["staging"])
	require.False(t, found.Status["prod"])

	// it should list all features with the status on all environments
	list, err = store.ListFeatures()
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, found.ID, list[0].ID)
	require.True(t, list[0].Status["dev"])
	require.True(t, list[0].Status["staging"])
	require.False(t, list[0].Status["prod"])
}

func getenv(name, val string) string {
	if found := os.Getenv(name); found != "" {
		return found
	}
	return val
}
