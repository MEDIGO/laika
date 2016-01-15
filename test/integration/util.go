package integration

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/feature-flag/api"
	"github.com/MEDIGO/feature-flag/client"
	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/store"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"os"
)

type FeatureFlagSuite struct {
	suite.Suite
	client client.Client
	server *httptest.Server
}

func (s *FeatureFlagSuite) SetupTest() {

	store, err := store.NewStore(
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	)
	if err != nil {
		panic(err)
	}

	stats, err := statsd.New(os.Getenv("STATSD_HOST") + ":" + os.Getenv("STATSD_PORT"))
	if err != nil {
		panic(err)
	}

	s.server = httptest.NewServer(api.NewServer(store, stats))

	s.client, err = client.NewClient(s.server.URL)
	if err != nil {
		panic(err)
	}
}

func (s *FeatureFlagSuite) CreateRandFeature(userID int64) (*model.Feature, error) {
	input := &model.Feature{}
	return s.client.FeatureCreate(input)
}

func (s *FeatureFlagSuite) TearDownSuite() {
	s.server.Close()
}
