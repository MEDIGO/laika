package integration

import (
	"net/http/httptest"
	"os"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/laika/api"
	"github.com/MEDIGO/laika/client"
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
		panic(err)
	}

	stats, err := statsd.New(os.Getenv("LAIKA_STATSD_HOST") + ":" + os.Getenv("LAIKA_STATSD_PORT"))
	if err != nil {
		panic(err)
	}

	s.server = httptest.NewServer(api.NewServer(store, stats))

	s.client, err = client.NewClient(s.server.URL)
	if err != nil {
		panic(err)
	}
}

func (s *FeatureFlagSuite) CreateRandFeature(userID int64) (*api.Feature, error) {
	input := &api.Feature{}
	return s.client.FeatureCreate(input)
}

func (s *FeatureFlagSuite) TearDownSuite() {
	s.server.Close()
}
