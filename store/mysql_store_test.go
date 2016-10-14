package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMySQLStore(t *testing.T) {
	host := getenv("LAIKA_TEST_MYSQL_HOST", "")
	if host == "" {
		t.Skip("Skipping store MySQL test")
	}

	port := getenv("LAIKA_TEST_MYSQL_PORT", "3306")
	username := getenv("LAIKA_TEST_MYSQL_USERNAME", "root")
	password := getenv("LAIKA_TEST_MYSQL_PASSWORD", "")
	database := getenv("LAIKA_TEST_MYSQL_DBNAME", "test")

	store, err := NewMySQLStore(username, password, host, port, database)
	require.NoError(t, err)

	err = store.Migrate()
	require.NoError(t, err)

	require.NoError(t, store.Reset())
	testStoreUsers(t, store)

	require.NoError(t, store.Reset())
	testStoreEnvironments(t, store)

	require.NoError(t, store.Reset())
	testStoreFeatures(t, store)

	require.NoError(t, store.Reset())
}
