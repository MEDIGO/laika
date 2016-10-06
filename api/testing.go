package api

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

// NewTestServer creates a new initialised Laika httptest.Server. The server
// root credentials are "root" as username and password. It contains an
// environment named "test" with an enabled featured named "test_feature",
// and a user whose username is "user" and password is "password".
func NewTestServer(t *testing.T) *httptest.Server {
	s, err := store.NewMySQLStore(
		os.Getenv("LAIKA_MYSQL_USERNAME"),
		os.Getenv("LAIKA_MYSQL_PASSWORD"),
		os.Getenv("LAIKA_MYSQL_HOST"),
		os.Getenv("LAIKA_MYSQL_PORT"),
		os.Getenv("LAIKA_MYSQL_DBNAME"),
	)
	require.NoError(t, err)

	err = s.Ping()
	require.NoError(t, err)

	err = s.Migrate()
	require.NoError(t, err)

	user := models.User{
		Username:     "test_username" + store.Token(),
		PasswordHash: "awesome_password",
	}
	err = s.CreateUser(&user)
	require.NoError(t, err)

	server, err := NewServer(ServerConfig{
		Store:        s,
		RootUsername: "root",
		RootPassword: "root",
	})
	require.NoError(t, err)

	return httptest.NewServer(server)
}

// CreateTestFeature creates a random feature to be used during testing.
func CreateTestFeature(t *testing.T) *models.Feature {
	s, err := store.NewMySQLStore(
		os.Getenv("LAIKA_MYSQL_USERNAME"),
		os.Getenv("LAIKA_MYSQL_PASSWORD"),
		os.Getenv("LAIKA_MYSQL_HOST"),
		os.Getenv("LAIKA_MYSQL_PORT"),
		os.Getenv("LAIKA_MYSQL_DBNAME"),
	)
	require.NoError(t, err)

	env, err := s.GetEnvironmentByName("test")
	if err != nil {
		env = &models.Environment{
			Name: "test",
		}
		err = s.CreateEnvironment(env)
		require.NoError(t, err)
	}

	feature := &models.Feature{
		Name: "test_feature" + store.Token(),
		Status: map[string]bool{
			"test": true,
		},
	}

	err = s.CreateFeature(feature)
	require.NoError(t, err)

	return feature
}
