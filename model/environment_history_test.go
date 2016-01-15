package model

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EnvironmentHistoryIntegrationSuite struct {
	suite.Suite
	Name                      string
	FeatureId                 int64
	Enabled                   bool
	ExampleEnvironmentHistory *EnvironmentHistory
}

func (suite *EnvironmentHistoryIntegrationSuite) TestExample() {
	suite.Name = "featureKey"
	suite.FeatureId = 1
	suite.Enabled = true
	suite.ExampleEnvironmentHistory = NewEnvironmentHistory(time.Now(), suite.Enabled, suite.FeatureId, suite.Name, time.Now())
	require.NotNil(suite.T(), suite.ExampleEnvironmentHistory.Id)
	require.Equal(suite.T(), *suite.ExampleEnvironmentHistory.FeatureId, suite.FeatureId)
	require.Equal(suite.T(), *suite.ExampleEnvironmentHistory.Name, suite.Name)
	require.Equal(suite.T(), *suite.ExampleEnvironmentHistory.Enabled, suite.Enabled)
	require.NotNil(suite.T(), suite.ExampleEnvironmentHistory.CreatedAt)
	require.NotNil(suite.T(), suite.ExampleEnvironmentHistory.Timestamp)
}

func TestEnvironmentHistoryIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentHistoryIntegrationSuite))
}
