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

	if _, err = s.GetUserByUsername("awesome_username"); err != nil {
		user := store.User{
			Username:     "awesome_username",
			PasswordHash: "awesome_password",
		}
		err = s.CreateUser(&user)
		require.NoError(t, err)
	}

	env, err := s.GetEnvironmentByName("test")
	if err != nil {
		environment := store.Environment{
			Name: store.String("test"),
		}
		err = s.CreateEnvironment(&environment)
		require.NoError(t, err)

		env = &environment
	}

	feat, err := s.GetFeatureByName("awesome_feature")
	if err != nil {
		feature := store.Feature{
			Name: store.String("awesome_feature"),
		}
		err = s.CreateFeature(&feature)
		require.NoError(t, err)

		feat = &feature
	}

	if _, err = s.GetFeatureStatus(feat.Id, env.Id); err != nil {
		featureStatus := store.FeatureStatus{
			Enabled:       store.Bool(true),
			FeatureId:     store.Int(feat.Id),
			EnvironmentId: store.Int(env.Id),
		}
		err = s.CreateFeatureStatus(&featureStatus)
		require.NoError(t, err)
	}

	server, err := NewServer(ServerConfig{
		Store:        s,
		RootUsername: "root",
		RootPassword: "root",
	})
	require.NoError(t, err)

	return httptest.NewServer(server)
}
