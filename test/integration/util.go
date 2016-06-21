package integration

import (
	"net/http/httptest"
	"os"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/laika/api"
	"github.com/MEDIGO/laika/client"
	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
)

type FeatureFlagSuite struct {
	suite.Suite
	client client.Client
	server *httptest.Server
}

func (s *FeatureFlagSuite) SetupTest() {
	store, err := store.NewStore(
		os.Getenv("LAIKA_MYSQL_USERNAME"),
		os.Getenv("LAIKA_MYSQL_PASSWORD"),
		os.Getenv("LAIKA_MYSQL_HOST"),
		os.Getenv("LAIKA_MYSQL_PORT"),
		os.Getenv("LAIKA_MYSQL_DBNAME"),
	)
	if err != nil {
		require.NoError(s.T(), err)
	}

	if err := store.Ping(); err != nil {
		require.NoError(s.T(), err)
	}

	if err := store.Reset(); err != nil {
		require.NoError(s.T(), err)
	}

	notifier := notifier.NewSlackNotifier(os.Getenv("SLACK_TOKEN"), os.Getenv("SLACK_CHANNEL"))
	s.server = httptest.NewServer(api.NewServer(store, nil, notifier))

	s.client, err = client.NewClient(s.server.URL)
	if err != nil {
		require.NoError(s.T(), err)
	}
}

func (s *FeatureFlagSuite) CreateRandFeature(userID int64) (*api.Feature, error) {
	input := &api.Feature{}
	return s.client.FeatureCreate(input)
}

func (s *FeatureFlagSuite) TearDownSuite() {
	s.server.Close()
}
