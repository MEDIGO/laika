package api

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MEDIGO/laika/store"
	"github.com/stretchr/testify/require"
)

// NewTestServer creates a new initialised Laika httptest.Server. The server
// root credentials are "root" as username and password. It contains an
// environment named "test" with an enabled featured named "test_feature",
// and a user whose username is "user" and password is "password".
func NewTestServer(t *testing.T) *httptest.Server {
	s, err := store.NewStore(
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

	user := store.User{
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

// CreateFeatureStatus creates an environment, a feature, and a freature status. Returns the name of the feature.
func CreateFeatureStatus(t *testing.T) string {
	s, err := store.NewStore(
		os.Getenv("LAIKA_MYSQL_USERNAME"),
		os.Getenv("LAIKA_MYSQL_PASSWORD"),
		os.Getenv("LAIKA_MYSQL_HOST"),
		os.Getenv("LAIKA_MYSQL_PORT"),
		os.Getenv("LAIKA_MYSQL_DBNAME"),
	)
	require.NoError(t, err)

	env, err := s.GetEnvironmentByName("test")
	if err != nil {
		environment := store.Environment{
			Name: store.String("test"),
		}
		err = s.CreateEnvironment(&environment)
		require.NoError(t, err)
		env = &environment
	}

	feature := store.Feature{
		Name: store.String("test_feature" + store.Token()),
	}

	err = s.CreateFeature(&feature)
	require.NoError(t, err)

	status := store.FeatureStatus{
		Enabled:       store.Bool(true),
		FeatureId:     store.Int(feature.Id),
		EnvironmentId: store.Int(env.Id),
	}

	err = s.CreateFeatureStatus(&status)
	require.NoError(t, err)

	return *feature.Name
}
