package integration

import (
  "testing"
  "github.com/MEDIGO/feature-flag/model"
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"
)

type EnvironmentFeatureUpdateIntegrationSuite struct {
  suite.Suite
  Key string
  FeatureID int64
  EnvironmentID int64
  Enabled bool
  ExampleEnvironmentFeatureUpdate *model.EnvironmentFeatureUpdate
}

func (suite *EnvironmentFeatureUpdateIntegrationSuite) SetupTest() {
  suite.Key = "featureKey"
  suite.FeatureID = 1
  suite.EnvironmentID = 2
  suite.Enabled = true
  suite.ExampleEnvironmentFeatureUpdate = model.NewEnvironmentFeatureUpdate(&suite.FeatureID, &suite.EnvironmentID, &suite.Enabled)
}

func (suite *EnvironmentFeatureUpdateIntegrationSuite) TestExample() {

	require.NotNil(suite.T(), suite.ExampleEnvironmentFeatureUpdate.Id)
  require.Equal(suite.T(), *suite.ExampleEnvironmentFeatureUpdate.FeatureID, suite.FeatureID)
  require.Equal(suite.T(), *suite.ExampleEnvironmentFeatureUpdate.EnvironmentID, suite.EnvironmentID)
  require.Equal(suite.T(), *suite.ExampleEnvironmentFeatureUpdate.Enabled, suite.Enabled)
  require.NotNil(suite.T(), suite.ExampleEnvironmentFeatureUpdate.Date)

}

func TestEnvironmentFeatureUpdateIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentFeatureUpdateIntegrationSuite))
}
