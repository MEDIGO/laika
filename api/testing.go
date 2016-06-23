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

	err = s.Reset()
	require.NoError(t, err)

	environment := store.Environment{
		Name: store.String("test"),
	}
	err = s.CreateEnvironment(&environment)
	require.NoError(t, err)

	feature := store.Feature{
		Name: store.String("test_feature"),
	}
	err = s.CreateFeature(&feature)
	require.NoError(t, err)

	status := store.FeatureStatus{
		Enabled:       store.Bool(true),
		FeatureId:     store.Int(feature.Id),
		EnvironmentId: store.Int(environment.Id),
	}
	err = s.CreateFeatureStatus(&status)
	require.NoError(t, err)

	user := store.User{
		Username:     "awesome_username",
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
